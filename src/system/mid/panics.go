package mid

import (
	"context"
	app "github.com/obada-protocol/demo-service/http"
	"github.com/pkg/errors"
	"net/http"
)

func Panics() app.Middleware {
	m := func(handler app.Handler) app.Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					err = errors.Errorf("PANIC: %v", rec)
				}
			}()

			// Call the next handler and set its return value in the err variable.
			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
