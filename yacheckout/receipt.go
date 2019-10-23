package yacheckout

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//Receipt struct is Yandex.Checkout receipt object
type Receipt struct {
	ID                   string       `json:"id,omitempty"`
	Type                 string       `json:"type"`
	PaymentID            string       `json:"payment_id,omitempty"`
	RefundID             string       `json:"refund_id,omitempty"`
	Status               string       `json:"status,omitempty"`
	FiscalDocumentNumber string       `json:"fiscal_document_number,omitempty"`
	FiscalStorageNumber  string       `json:"fiscal_storage_number,omitempty"`
	FiscalAttribute      string       `json:"fiscal_attribute,omitempty"`
	RegisteredAt         *time.Time   `json:"registered_at,string,omitempty"`
	FiscalProviderID     string       `json:"fiscal_provider_id,omitempty"`
	Customer             *Customer    `json:"customer,omitempty"`
	Items                []Item       `json:"items"`
	TaxSystemCode        uint8        `json:"tax_system_code,omitempty"`
	Send                 bool         `json:"send"`
	Settlements          []Settlement `json:"settlements,omitempty"`
}

//Customer struct is receipt.customer object
type Customer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

//Item struct is receipt.items object
type Item struct {
	Description              string  `json:"description"`
	Quantity                 string  `json:"quantity"`
	Amount                   Amount  `json:"amount"`
	VatCode                  uint8   `json:"vat_code"`
	PaymentSubject           string  `json:"payment_subject,omitempty"`
	PaymentMode              string  `json:"payment_mode,omitempty"`
	ProductCode              string  `json:"product_code,omitempty"`
	CountryOfOriginCode      string  `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber string  `json:"customs_declaration_number,omitempty"`
	Excise                   float64 `json:"excise,string,omitempty"`
}

//Settlement struct is receipt.settlements object
type Settlement struct {
	Type   string `json:"type"`
	Amount Amount `json:"amount"`
}

//Receipts struct is Yandex.Checkout receipts object
type Receipts struct {
	Type  string    `json:"type"`
	Items []Receipt `json:"items"`
}

//CreateReceipt func create receipt Yandex.Checkout
func (checkout *Checkout) CreateReceipt(client *http.Client, V4UUID *uuid.UUID, rcpt *Receipt) (receipt *Receipt, apierr *Error, err error) {

	b, err := json.Marshal(rcpt)
	if err != nil {
		return
	}

	b, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "receipts", b)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &receipt)
	return
}

//GetReceipts func receives receipts Yandex.Checkout
func (checkout *Checkout) GetReceipts(client *http.Client, refund bool, id string) (receipts *Receipts, apierr *Error, err error) {

	method := "payment_id="

	if refund {
		method = "refund_id="
	}

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "receipts?"+method+id, nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &receipts)
	return
}

//GetReceipt func receives receipt Yandex.Checkout
func (checkout *Checkout) GetReceipt(client *http.Client, id string) (receipt *Receipt, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "receipts/"+id, nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &receipt)
	return
}
