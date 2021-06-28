package sdk_go

import (
	"github.com/obada-foundation/sdk-go/properties"
	"log"
	"time"
)

type Status string

const StatusFunctional Status = "FUNCTIONAL"
const StatusNonFunctional Status = "NON_FUNCTIONAL"
const StatusDisposed Status = "DISPOSED"
const StatusStolen Status = "STOLEN"
const DisabledByOwner Status = "DISABLED_BY_OWNER"

type ObitIdDto struct {
	SerialNumberHash string `validate:"required"`
	Manufacturer     string `validate:"required"`
	PartNumber       string `validate:"required"`
}

type ObitDto struct {
	ObitIdDto
	OwnerDid      string `validate:"required"`
	ObdDid        string
	Matadata      map[string]string
	StructureData map[string]string
	Documents     map[string]string
	ModifiedAt    time.Time
	Status        string
}

type Obit struct {
	obitId           properties.ObitId
	serialNumberHash properties.StringProperty
	manufacturer     properties.StringProperty
	partNumber       properties.StringProperty
	ownerDid         properties.StringProperty
	obdDid           properties.StringProperty
	metadata         properties.KvProperty
	structureData    properties.KvProperty
	documents        properties.KvProperty
	modifiedAt       properties.TimeProperty
	status           properties.StatusProperty
	debug bool
	logger *log.Logger
}
