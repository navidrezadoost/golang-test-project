package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	Fob string = "FOB"
	Fca string = "FCA"
)

const (
	Cfr string = "CFR"
	Cpt string = "CPT"
)

type ShippingTimes struct {
	TransportMethod string `bson:"transportMethod" validate:"required"`
	LoadingPort     string `bson:"loadingPort" validate:"required,uuid"`
	DischargePort   string `bson:"dischargePort" validate:"required,uuid"`
	EntranceBorder  string `bson:"entranceBorder" validate:"required,uuid"`
	TransportFleet  string `bson:"transportFleet" validate:"required"`
}

type Packages struct {
	PackageKind    string `bson:"packageKind" validate:"required"`
	Kind           string `bson:"kind" validate:"required,uuid"`
	PackageMaesure string `bson:"maesure" validate:"required"`
	Maesure        string `bson:"packageMaesure" validate:"required,uuid"`
}

type Product struct {
	ID primitive.ObjectID `bson:"_id,omitempty"` // MongoDB document ID
	//step 1
	RefrenceProductId string `bson:"refrenceProductId" validate:"required,uuid"`
	//step 2
	ContractType     string          `bson:"contractType" validate:"required"`
	PartialShipment  bool            `bson:"partialShipment"`
	OnePieceShipping bool            `bson:"onePieceShipping"`
	ShippingTimes    []ShippingTimes `bson:"shippingTimes"`
	//step3
	Qty                int64      `bson:"qty" validate:"required"`
	Unit               string     `bson:"unit" validate:"required,uuid"`
	NetWeight          float64    `bson:"netWeight" validate:"numeric"`
	MeasureNetWeight   string     `bson:"measureNetWeight" validate:"uuid"`
	GrossWeight        float64    `bson:"grossWeight" validate:"required,numeric"`
	MeasureGrossWeight float64    `bson:"measureGrossWeight" validate:"required,uuid"`
	FOB_FCA            string     `bson:"fob_fca" validate:"required,oneof=Fob,Fca"`
	FOB_FCA_UnitPrice  float64    `bson:"fOB_FCA_UnitPrice" validate:"required"`
	CFR_CPT            string     `bson:"cfr_cpt" validate:"required"`
	CFR_CPT_UnitPrice  float64    `bson:"cFR_CPT_UnitPrice" validate:"required,oneof=Cfr,Cpt"`
	OtherCharges       int64      `bson:"otherCharges" validate:"numeric"`
	Discount           int64      `bson:"discount" validate:"numeric"`
	Packages           []Packages `bson:"packages" validate:"required"`
	IsRemove           bool       `bson:"isRemove" default:"false"`
	CreatedAt          time.Time  `bson:"createdAt,omitempty,default"`
	UpdatedAt          time.Time  `bson:"updatedAt,omitempty,default"`
}
