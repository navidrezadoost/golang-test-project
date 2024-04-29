package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	DirectoryName string   `bson:"directoryName" json:"directoryName" binding:"required"`
	Caption       string   `bson:"caption" json:"caption"`
	FileIDs       []string `bson:"fileIDs" json:"fileIDs" binding:"required"`
}

// Proforma represents the schema for the "proformas" collection in MongoDB.
type Proforma struct {
	ID                          primitive.ObjectID   `bson:"_id,omitempty"` // MongoDB document ID
	ProformaInvoiceNumber       string               `bson:"proformaInvoiceNumber" binding:"required" validate:"required"`
	ProformaInvoiceDate         time.Time            `bson:"proformaInvoiceDate" binding:"required" validate:"required"`
	ProformaInvoiceValidityDate time.Time            `bson:"proformaInvoiceValidityDate" binding:"required"`
	TypeOfProductionUnits       string               `bson:"typeOfProductionUnits" validate:"required"`
	Buyer                       string               `bson:"buyer" binding:"required" validate:"required"`
	Seller                      string               `bson:"seller" binding:"required" validate:"required"`
	Currency                    string               `bson:"currency" validate:"required"`
	OrderRegistrationMode       string               `bson:"orderRegistrationMode" binding:"required"`
	CountryBeneficiary          string               `bson:"countryBeneficiary" binding:"required"`
	CountryOfOrigin             string               `bson:"countryOfOrigin" binding:"required"`
	ContractType                string               `bson:"contractType" validate:"required"`
	LoadingPort                 string               `bson:"loadingPort" validate:"required"`
	DischargePort               string               `bson:"dischargePort" validate:"required"`
	PaymentTerm                 string               `bson:"paymentTerm" validate:"required"`
	GeneralDiscount             int64                `bson:"generalDiscount"`
	GeneralOtherCharge          int64                `bson:"generalOtherCharge"`
	TotalAmount                 int64                `bson:"totalAmount"`
	TotalDiscount               int64                `bson:"totalDiscount"`
	FileNumber                  int64                `bson:"fileNumber"`
	Documents                   []File               `bson:"documents"`
	Products                    []primitive.ObjectID `bson:"products"`
	Status                      string               `bson:"status" default:"DRAFT" validate:"oneof=DRAFT READY_FOR_PROCESS"`
	IsRemove                    bool                 `bson:"isRemove" default:"false"`
	CreatedAt                   time.Time            `bson:"createdAt,omitempty"`
	UpdatedAt                   time.Time            `bson:"updatedAt,omitempty"`
}
