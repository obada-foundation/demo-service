package http

import (
	"context"
	"github.com/obada-protocol/demo-service/build"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/dimfeld/httptreemux/v5"
)

// ctxKey represents the type of value for the context key.
type ctxKey int


// KeyValues is how request values are stored/retrieved.
const KeyValues ctxKey = 1

// Values represent state for each request.
type Values struct {
	TraceID    string
	Now        time.Time
	StatusCode int
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

var registered = make(map[string]bool)

type App struct {
	mux      *httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	mux := httptreemux.NewContextMux()

	app := &App{
		mux:      mux,
		shutdown: shutdown,
		mw:       mw,
	}

	return app
}

func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

func (a *App) HandleDebug(method string, path string, handler Handler, mw ...Middleware) {
	a.handle(true, method, path, handler, mw...)
}

func (a *App) HandleStatic() {
	contextMux := a.mux.UsingContext()

	contextMux.Handler(http.MethodGet, "/public/*", build.AssetHandler("/public/", "build"))
}

func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
	a.handle(false, method, path, handler, mw...)
}

func (a *App) handle(debug bool, method string, path string, handler Handler, mw ...Middleware) {
	if debug {
		// Track all the handlers that are being registered so we don't have
		// the same handlers registered twice to this singleton.
		if _, exists := registered[method+path]; exists {
			return
		}
		registered[method+path] = true
	}

	// First wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// Add the application's general middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request) {
		// Start or expand a distributed trace.
		ctx := r.Context()

		// Set the context with the required values to
		// process the request.
		v := Values{
			TraceID: "test",
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, KeyValues, &v)

		// Call the wrapped handler functions.
		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
			return
		}
	}

	a.mux.Handle(method, path, h)
}
