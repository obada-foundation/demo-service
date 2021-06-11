package handlers

import (
	"bytes"
	"context"
	app "github.com/obada-protocol/demo-service/http"
	"net/http"
)

func home(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	tpl, err := app.NewTpl().ParseFS(templateFs, "templates/home.html", "templates/base.html")

	if err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(&html, "base",nil); err != nil {
		return err
	}

	return app.Respond(ctx, w, html.Bytes(), http.StatusOK)
}
