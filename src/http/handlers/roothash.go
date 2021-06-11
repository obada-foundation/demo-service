package handlers

import (
	"bytes"
	"context"
	sdk_go "github.com/obada-foundation/sdk-go"
	app "github.com/obada-protocol/demo-service/http"
	"log"
	"net/http"
	"time"
)

type rootHashGroup struct {}

func (rh rootHashGroup) calculateRootHash(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var captureSdkLogs bytes.Buffer

	captureLog := log.New(&captureSdkLogs, "", 0)
	sdk, err := sdk_go.NewSdk(captureLog, true)

	if err != nil {
		return err
	}



	var dto sdk_go.ObitDto
	dto.SerialNumberHash = r.FormValue("serial_number_hash")
	dto.Manufacturer = r.FormValue("manufacturer")
	dto.PartNumber = r.FormValue("part_number")
	dto.OwnerDid = r.FormValue("owner_did")
	dto.ObdDid = r.FormValue("obd_did")
	modifiedAt, err := time.Parse("2021-06-25T14:02", r.FormValue("modified_at"))
	dto.ModifiedAt = modifiedAt
	dto.Status = r.FormValue("status")

	log.Printf("%v", dto)


	obit, err := sdk.NewObit(dto)
	_, err = obit.GetRootHash()


	return app.RespondJson(ctx, w, nil, http.StatusOK)
}

func (rh rootHashGroup) rootHash(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	tpl, err := app.NewTpl().ParseFS(templateFs,"templates/roothash.html", "templates/base.html")

	if err != nil {
		return err
	}

	if err = tpl.ExecuteTemplate(&html, "base",nil); err != nil {
		return err
	}

	return app.Respond(ctx, w, html.Bytes(), http.StatusOK)
}
