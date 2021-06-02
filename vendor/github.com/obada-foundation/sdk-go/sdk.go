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
		sdk.Debug(fmt.Sprintf("NewObit(%v)", dto))
	}

	serialNumberProp, err := sdk.createStringProperty(dto.SerialNumberHash)

	if err != nil {
		return o, err
	}

	manufacturerProp, err := sdk.createStringProperty(dto.Manufacturer)

	if err != nil {
		return o, err
	}

	pnProp, err := sdk.createStringProperty(dto.PartNumber)

	if err != nil {
		return o, err
	}

	obdDidProp, err := sdk.createStringProperty(dto.ObdDid)

	if err != nil {
		return o, err
	}

	ownerDidProp, err := sdk.createStringProperty(dto.OwnerDid)

	if err != nil {
		return o, err
	}

	statusProp, err := sdk.createStringProperty(string(dto.Status))

	if err != nil {
		return o, err
	}

	obitId, err := properties.NewObitId(serialNumberProp, manufacturerProp, pnProp)

	if err != nil {
		return o, err
	}

	sdk.debugObitId(serialNumberProp.GetHash(), manufacturerProp.GetHash(), pnProp.GetHash(),obitId)

	o.obitId = obitId
	o.serialNumberHash = serialNumberProp
	o.manufacturer = manufacturerProp
	o.partNumber = pnProp
	o.obdDid = obdDidProp
	o.ownerDid = ownerDidProp
	o.status = statusProp
	//o.modifiedAt =

	return o, nil
}

func (o Obit) GetRootHash() (hash.Hash, error) {
	var rootHash hash.Hash

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

	rootHash, err := hash.NewHash(fmt.Sprintf("%x", sum))

	if err != nil {
		return rootHash, err
	}

	return rootHash, nil
}

// NewObitId c
func (sdk *Sdk) NewObitId(dto ObitIdDto) (properties.ObitId, error) {
	var obitId properties.ObitId

	if sdk.debug {
		sdk.Debug(fmt.Sprintf("NewObitId(%q, %q, %q)", dto.SerialNumberHash, dto.Manufacturer, dto.PartNumber))
	}

	err := sdk.validate.Struct(dto)
	if err != nil {
		return obitId, err
	}

	snProp, err := sdk.createStringProperty(dto.SerialNumberHash)

	if err != nil {
		return obitId, err
	}

	mnProp, err := sdk.createStringProperty(dto.Manufacturer)

	if err != nil {
		return obitId, err
	}

	pnProp, err := sdk.createStringProperty(dto.PartNumber)

	if err != nil {
		return obitId, err
	}

	obitId, err = properties.NewObitId(snProp, mnProp, pnProp)

	if err != nil {
		return obitId, err
	}

	sdk.debugObitId(snProp.GetHash(), mnProp.GetHash(), pnProp.GetHash(),obitId)

	return obitId, nil
}

func (sdk *Sdk) Debug(message string) {
	if sdk.debug {
		if sdk.logger != nil {
			sdk.logger.Println(message)
		} else {
			log.Println(message)
		}
	}
}

func (sdk *Sdk) debugObitId(snPropHash hash.Hash, mnPropHash hash.Hash, pnPropHash hash.Hash, obitId properties.ObitId) {
	if sdk.debug {
		sum := snPropHash.GetDec() + mnPropHash.GetDec() + pnPropHash.GetDec()

		sdk.Debug(
			fmt.Sprintf("(%d + %d + %d) -> %d",
				snPropHash.GetDec(),
				mnPropHash.GetDec(),
				pnPropHash.GetDec(),
				sum,
			))

		dec2Hex := fmt.Sprintf("%x", sum)
		dec2HexHash, _ := hash.NewHash(dec2Hex)

		sdk.Debug(fmt.Sprintf("Dec2Hex(%d) -> %s -> Hash(%s) -> %s", sum, dec2Hex, dec2Hex,dec2HexHash.GetHash()))

		sdk.Debug(fmt.Sprintf("Hash : %s", obitId.GetHash().GetHash()))
		sdk.Debug(fmt.Sprintf("Did : %s", obitId.GetDid()))
		sdk.Debug(fmt.Sprintf("Usn : Base58(%s) -> %s", obitId.GetHash().GetHash(), obitId.GetUsn()))
	}
}

func (sdk *Sdk) createStringProperty(str string) (properties.StringProperty, error) {
	var prop properties.StringProperty

	prop, err := properties.NewStringProperty(str)

	if err != nil {
		return prop, err
	}

	if sdk.debug {
		propHash := prop.GetHash()
		sdk.Debug(fmt.Sprintf("Hash(%q) -> Hash2Dec(%q) -> %d", prop.GetValue(), propHash.GetHash(), propHash.GetDec()))
	}

	return prop, nil
}
