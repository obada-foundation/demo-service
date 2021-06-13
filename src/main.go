package main

import (
	"context"
	"expvar"
	"fmt"
	"github.com/obada-protocol/demo-service/http/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ardanlabs/conf"
	"github.com/pkg/errors"
)

// build is the git version of this program.
var build = "develop"

func main() {
	log := log.New(os.Stdout, "OBADA-WEB :", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run(log *log.Logger) error {
	var cfg struct {
		conf.Version
		Web struct {
			Host            string        `conf:"default:0.0.0.0:8080"`
			ReadTimeout     time.Duration `conf:"default:3s"`
			WriteTimeout    time.Duration `conf:"default:3s"`
			ShutdownTimeout time.Duration `conf:"default:5s"`
		}
	}

	cfg.Version.SVN = build
	cfg.Version.Desc = "(c) OBADA Foundation 2021"

	if err := conf.Parse(os.Args[1:], "OBADA-WEB", &cfg); err != nil {
		switch err {
		case conf.ErrHelpWanted:
			usage, err := conf.Usage("OBADA-WEB", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		case conf.ErrVersionWanted:
			version, err := conf.VersionString("OBADA-WEB", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config version")
			}
			fmt.Println(version)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	expvar.NewString("build").Set(build)

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}

	log.Println(out)

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	httpHandlers := handlers.Server(build, shutdown, log)

	server := http.Server{
		Addr:         cfg.Web.Host,
		Handler:      httpHandlers,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
	}

	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		log.Println("Http server started")
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case sig := <-shutdown:
		log.Println("shutdown started", sig)
		defer log.Println("shutdown completed")

		// Give outstanding requests a deadline for completion.
		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		// Asking listener to shutdown and shed load.
		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
