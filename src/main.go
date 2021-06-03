package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	obadaSdk "github.com/obada-foundation/sdk-go"
)

// build is the git version of this program. It is set using build flags in the makefile.
var build = "develop"

func main() {
	log := log.New(os.Stdout, "OBADA-DEMO :", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if err := run(log); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		showForm(w, r)
	} else {
		submitForm(w, r)
	}
}

func showForm(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/form.gtpl")
	t.Execute(w, nil)
}

func submitForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	var captureSdkLogs bytes.Buffer

	captureLog := log.New(&captureSdkLogs, "", 0)

	sdk, err := obadaSdk.NewSdk(captureLog, true)

	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("Something went wrong!"))
	}

	status := r.FormValue("status")

	if status == "FUNCTIONAL" {

	}

	var dto obadaSdk.ObitDto
	dto.SerialNumberHash = r.FormValue("serial_number_hash")
	dto.Manufacturer = r.FormValue("manufacturer")
	dto.PartNumber = r.FormValue("part_number")
	dto.OwnerDid = r.FormValue("owner_did")
	dto.ObdDid = r.FormValue("obd_did")

	modifiedAt, err := time.Parse("2021-06-25T14:02", r.FormValue("modified_at"))

	dto.ModifiedAt = modifiedAt
	dto.Status = status


	obit, err := sdk.NewObit(dto)
	_, err = obit.GetRootHash()

	if err != nil {
		log.Println(err.Error())
		w.Write([]byte("Something went wrong!"))
	}

	logs := strings.ReplaceAll(captureSdkLogs.String(), "\n", "<br \\>")

	w.Write([]byte(logs))
}

func run(log *log.Logger) error {
	http.HandleFunc("/", handleRequest) // setting router rule

	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		return err
	}

	return nil
}