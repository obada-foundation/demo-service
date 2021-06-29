package sdk_go

import (
	"fmt"
	"github.com/obada-foundation/sdk-go/hash"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/obada-foundation/sdk-go/properties"
)

type Sdk struct {
	logger   *log.Logger
	debug    bool
	validate *validator.Validate
}

func NewSdk(log *log.Logger, debug bool) (*Sdk, error) {
	return &Sdk{
		logger:   log,
		debug:    debug,
		validate: initializeValidator(),
	}, nil
}

func initializeValidator() *validator.Validate {
	v := validator.New()

	return v
}

// NewObit s
func (sdk *Sdk) NewObit(dto ObitDto) (Obit, error) {
	var o Obit
	o.debug = sdk.debug
	o.logger = sdk.logger

	err := sdk.validate.Struct(dto)
	if err != nil {
		return o, err
	}

	if sdk.debug {
		sdk.logger.Printf("NewObit(%v)", dto)
	}

	snProp, err := properties.NewStringProperty(dto.SerialNumberHash, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	manufacturerProp, err := properties.NewStringProperty(dto.Manufacturer, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	pnProp, err := properties.NewStringProperty(dto.PartNumber, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	obdDidProp, err := properties.NewStringProperty(dto.ObdDid, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	ownerDidProp, err := properties.NewStringProperty(dto.OwnerDid, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	statusProp, err := properties.NewStatusProperty(dto.Status, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	obitIdProp, err := properties.NewObitIdProperty(snProp, manufacturerProp, pnProp, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	modifiedAt, err := properties.NewTimeProperty(dto.ModifiedAt, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	metadataProp, err := properties.NewMapProperty(dto.Matadata, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	strctDataProp, err := properties.NewMapProperty(dto.StructureData, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	documentsProp, err := properties.NewMapProperty(dto.Documents, sdk.logger, sdk.debug)

	if err != nil {
		return o, err
	}

	o.obitId = obitIdProp
	o.serialNumberHash = snProp
	o.manufacturer = manufacturerProp
	o.partNumber = pnProp
	o.obdDid = obdDidProp
	o.ownerDid = ownerDidProp
	o.status = statusProp
	o.metadata = metadataProp
	o.structureData = strctDataProp
	o.documents = documentsProp
	o.modifiedAt = modifiedAt

	return o, nil
}

func (o Obit) GetRootHash() (hash.Hash, error) {
	var rootHash hash.Hash

	if o.debug {
		o.logger.Println("\n\nObit root hash calculation")
	}

	sum := o.obitId.GetHash().GetDec() +
		o.serialNumberHash.GetHash().GetDec() +
		o.manufacturer.GetHash().GetDec() +
		o.partNumber.GetHash().GetDec() +
		o.ownerDid.GetHash().GetDec() +
		o.obdDid.GetHash().GetDec() +
		o.metadata.GetHash().GetDec() +
		o.structureData.GetHash().GetDec() +
		o.documents.GetHash().GetDec() +
		o.modifiedAt.GetHash().GetDec() +
		o.status.GetHash().GetDec()

	if o.debug {
		o.logger.Println(fmt.Sprintf(
			"(%d + %d + %d + %d + %d + %d + %d + %d + %d + %d + %d) -> %d -> Dec2Hex(%d) -> %s",
			o.obitId.GetHash().GetDec(),
			o.serialNumberHash.GetHash().GetDec(),
			o.manufacturer.GetHash().GetDec(),
			o.partNumber.GetHash().GetDec(),
			o.ownerDid.GetHash().GetDec(),
			o.obdDid.GetHash().GetDec(),
			o.metadata.GetHash().GetDec(),
			o.structureData.GetHash().GetDec(),
			o.documents.GetHash().GetDec(),
			o.modifiedAt.GetHash().GetDec(),
			o.status.GetHash().GetDec(),
			sum,
			sum,
			fmt.Sprintf("%x", sum),
		))
	}

	rootHash, err := hash.NewHash(fmt.Sprintf("%x", sum), o.logger, o.debug)

	if err != nil {
		return rootHash, err
	}

	if o.debug {
		o.logger.Printf("RootHash(%v) -> %q", o, rootHash.GetHash())
	}

	return rootHash, nil
}

// NewObitId c
func (sdk *Sdk) NewObitId(dto ObitIdDto) (properties.ObitId, error) {
	var obitId properties.ObitId

	if sdk.debug {
		sdk.logger.Printf("NewObitId(%q, %q, %q)", dto.SerialNumberHash, dto.Manufacturer, dto.PartNumber)
	}

	err := sdk.validate.Struct(dto)
	if err != nil {
		return obitId, err
	}

	snProp, err := properties.NewStringProperty(dto.SerialNumberHash, sdk.logger, sdk.debug)

	if err != nil {
		return obitId, err
	}

	mnProp, err := properties.NewStringProperty(dto.Manufacturer, sdk.logger, sdk.debug)

	if err != nil {
		return obitId, err
	}

	pnProp, err := properties.NewStringProperty(dto.PartNumber, sdk.logger, sdk.debug)

	if err != nil {
		return obitId, err
	}

	obitId, err = properties.NewObitIdProperty(snProp, mnProp, pnProp, sdk.logger, sdk.debug)

	if err != nil {
		return obitId, err
	}

	return obitId, nil
}