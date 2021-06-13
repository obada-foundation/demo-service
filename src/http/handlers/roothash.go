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

type rootHashGroup struct{}

type ObitRequest struct {
	sdk_go.ObitDto
	Metadata       interface{}
	StructuredData interface{}
	Documents      interface{}
	ModifiedAt     string
}

func (rh rootHashGroup) calculateRootHash(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var captureSdkLogs bytes.Buffer
	var requestData ObitRequest
	var dto sdk_go.ObitDto

	captureLog := log.New(&captureSdkLogs, "", 0)
	sdk, err := sdk_go.NewSdk(captureLog, true)

	if err != nil {
		return err
	}

	app.Decode(r, &requestData)

	toKV := func(dynamicKV interface{}) map[string]string {
		kv := make(map[string]string)

		dynamicKVMap := dynamicKV.(map[string]interface{})

		for key := range dynamicKVMap {
			kv[key] = dynamicKVMap[key].(string)
		}

		return kv
	}

	dto.SerialNumberHash = requestData.SerialNumberHash
	dto.Manufacturer = requestData.Manufacturer
	dto.PartNumber = requestData.PartNumber
	dto.ObdDid = requestData.ObdDid
	dto.OwnerDid = requestData.OwnerDid
	dto.Status = requestData.Status
	dto.Matadata = toKV(requestData.Metadata)
	dto.StructuredData = toKV(requestData.StructuredData)
	dto.Documents = toKV(requestData.Documents)

	modifiedAt, err := time.Parse(time.RFC3339, requestData.ModifiedAt)

	if err != nil {
		return err
	}

	dto.ModifiedAt = modifiedAt

	obit, err := sdk.NewObit(dto)
	rootHash, err := obit.GetRootHash()

	resp := struct {
		RootHash string
		Log      string
	}{
		RootHash: rootHash.GetHash(),
		Log:      captureSdkLogs.String(),
	}

	return app.RespondJson(ctx, w, resp, http.StatusOK)
}

func (rh rootHashGroup) rootHash(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	tpl, err := app.NewTpl().ParseFS(templateFs, "templates/roothash.html", "templates/base.html")

	if err != nil {
		return err
	}

	data := struct {
		ObitId string
	}{
		ObitId: app.Param(r, "obitId"),
	}

	if err = tpl.ExecuteTemplate(&html, "base", data); err != nil {
		return err
	}

	return app.Respond(ctx, w, html.Bytes(), http.StatusOK)
}
