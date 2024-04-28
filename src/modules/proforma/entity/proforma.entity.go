package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	Path    string `bson:"path"`
	Type    string `bson:"type"`
	Size    string `bson:"size"`
	Caption string `bson:"size"`
}

const (
	Draft string = "DRAFT"
)

// Proforma represents the schema for the "proformas" collection in MongoDB.
type Proforma struct {
	ID                          primitive.ObjectID   `bson:"_id,omitempty"` // MongoDB document ID
	ProformaInvoiceNumber       string               `bson:"proformaInvoiceNumber" validate:"required"`
	ProformaInvoiceDate         time.Time            `bson:"proformaInvoiceDate" validate:"required"`
	ProformaInvoiceValidityDate time.Time            `bson:"proformaInvoiceValidityDate"`
	TypeOfProductionUnits       string               `bson:"typeOfProductionUnits" validate:"required"`
	Buyer                       string               `bson:"buyer" validate:"required"`
	Seller                      string               `bson:"seller" validate:"required"`
	Currency                    string               `bson:"currency" validate:"required"`
	OrderRegistrationMode       string               `bson:"orderRegistrationMode"`
	CountryBeneficiary          string               `bson:"countryBeneficiary"`
	CountryOfOrigin             string               `bson:"countryOfOrigin"`
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
	Status                      string               `bson:"status" default:"Draft" validate:"oneof=Draft"`
	IsRemove                    bool                 `bson:"isRemove" default:"false"`
	CreatedAt                   time.Time            `bson:"createdAt,omitempty,default"`
	UpdatedAt                   time.Time            `bson:"updatedAt,omitempty,default"`
}
