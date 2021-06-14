package mid

import (
	"context"
	app "github.com/obada-protocol/demo-service/http"
	"log"
	"net/http"
)

func Logger(log *log.Logger) app.Middleware {
	m := func(handler app.Handler) app.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			err := handler(ctx, w, r)
			return err
		}

		return h
	}

	return m
}
