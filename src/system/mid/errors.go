package mid

import (
	"context"
	"log"
	"net/http"

	app "github.com/obada-protocol/demo-service/http"
)

func Errors(log *log.Logger) app.Middleware {
	m := func(handler app.Handler) app.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			_, ok := ctx.Value(app.KeyValues).(*app.Values)
			if !ok {
				return app.NewShutdownError("web value missing from context")
			}

			if err := handler(ctx, w, r); err != nil {
				log.Println("ERROR", err)

				er := "Internal error"

				status := http.StatusInternalServerError

				// Respond with the error back to the client.
				if err := app.RespondJson(ctx, w, er, status); err != nil {
					return err
				}

				// If we receive the shutdown err we need to return it
				// back to the base handler to shutdown the service.
				if ok := app.IsShutdown(err); ok {
					return err
				}
			}

			// The error has been handled so we can stop propagating it.
			return nil
		}

		return h
	}

	return m
}
