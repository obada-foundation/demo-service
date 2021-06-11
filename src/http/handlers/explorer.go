package handlers

import (
	"bytes"
	"context"
	app "github.com/obada-protocol/demo-service/http"
	"net/http"
)

type explorerGroup struct {

}

func (eg explorerGroup) overview(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	tpl, err := app.NewTpl().ParseFS(templateFs,"templates/explorer.html", "templates/base.html")

	if err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(&html, "base",nil); err != nil {
		return err
	}

	return app.Respond(ctx, w, html.Bytes(), http.StatusOK)
}

func (eg explorerGroup) obit(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	data := struct {
		ObitId string
	}{
		ObitId: app.Param(r, "obitId"),
	}

	tpl, err := app.NewTpl().ParseFS(templateFs,"templates/obit.html", "templates/base.html")

	if err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(&html, "base",data); err != nil {
		return err
	}

	return app.Respond(ctx, w, html.Bytes(), http.StatusOK)
}

