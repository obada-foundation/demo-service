package handlers

import (
	"context"
	"embed"
	app "github.com/obada-protocol/demo-service/http"
	"github.com/obada-protocol/demo-service/system/mid"
	"log"
	"net/http"
	"os"
)

//go:embed templates
var templateFs embed.FS

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

func Server(build string, shutdown chan os.Signal, log *log.Logger, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	app := app.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Panics())

	app.Handle(http.MethodGet, "/", home)

	ch := checksumGroup{}

	app.Handle(http.MethodGet, "/checksum", ch.checksum)
	app.Handle(http.MethodGet, "/checksum/:obitId", ch.checksum)
	app.Handle(http.MethodPost, "/checksum", ch.calculateChecksum)

	eg := explorerGroup{}

	app.Handle(http.MethodGet, "/explorer", eg.overview)
	app.Handle(http.MethodGet, "/obit/:obitId", eg.obit)

	ng := nodesGroup{}

	app.Handle(http.MethodGet, "/nodes", ng.all)

	app.HandleStatic()

	// Accept CORS 'OPTIONS' preflight requests if config has been provided.
	// Don't forget to apply the CORS middleware to the routes that need it.
	// Example Config: `conf:"default:https://MY_DOMAIN.COM"`
	if opts.corsOrigin != "" {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			return nil
		}
		app.Handle(http.MethodOptions, "/*", h)
	}

	return app
}
