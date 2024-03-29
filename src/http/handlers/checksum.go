package handlers

import (
	"bytes"
	"context"
	sdkgo "github.com/obada-foundation/sdkgo"
	"github.com/obada-foundation/sdkgo/properties"
	app "github.com/obada-protocol/demo-service/http"
	"log"
	"net/http"
	"strings"
)

type checksumGroup struct{}

type ObitRequest struct {
	sdkgo.ObitDto
	Metadata       interface{}
	StructuredData interface{}
	Documents      interface{}
	ModifiedOn     int64
}

func (ch checksumGroup) calculateChecksum(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var captureSdkLogs bytes.Buffer
	var requestData ObitRequest
	var dto sdkgo.ObitDto

	captureLog := log.New(&captureSdkLogs, "", 0)
	sdk, err := sdkgo.NewSdk(captureLog, true)

	if err != nil {
		return err
	}

	app.Decode(r, &requestData)

	toKV := func(dynamicKV interface{}) []properties.KV {
		var kv []properties.KV

		dynamicKVMap := dynamicKV.(map[string]interface{})

		for key := range dynamicKVMap {

			kv = append(kv, properties.KV{
				Key: key,
				Value: dynamicKVMap[key].(string),
			})
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

	if err != nil {
		return err
	}

	dto.ModifiedOn = requestData.ModifiedOn

	obit, err := sdk.NewObit(dto)

	if err != nil {
		return err
	}

	checksum, err := obit.GetChecksum(nil)

	if err != nil {
		return err
	}

	logOutput := strings.ReplaceAll(captureSdkLogs.String(), "<|", "<p style='font-weight:bold;color:#16a085'>")
	logOutput = strings.ReplaceAll(logOutput, "|>", "</p>")

	resp := struct {
		Checksum string
		Log      string
	}{
		Checksum: checksum.GetHash(),
		Log:      logOutput,
	}

	return app.RespondJson(ctx, w, resp, http.StatusOK)
}

func (ch checksumGroup) checksum(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var html bytes.Buffer

	tpl, err := app.NewTpl().ParseFS(templateFs, "templates/checksum.html", "templates/base.html")

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
