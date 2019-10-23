package yacheckout

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//Payment struct is Yandex.Checkout payment object
type Payment struct {
	ID                   string                `json:"id,omitempty"`
	Status               string                `json:"status,omitempty"`
	Amount               *Amount               `json:"amount,omitempty"`
	Description          string                `json:"description,omitempty"`
	Receipt              *Receipt              `json:"receipt,omitempty"`
	Recipient            *Recipient            `json:"recipient,omitempty"`
	Requestor            *Requestor            `json:"requestor,omitempty"`
	PaymentToken         string                `json:"payment_token,omitempty"`
	PaymentMethodID      string                `json:"payment_method_id,omitempty"`
	PaymentMethodData    *PaymentMethod        `json:"payment_method_data,omitempty"`
	PaymentMethod        *PaymentMethod        `json:"payment_method,omitempty"`
	CapturedAt           *time.Time            `json:"captured_at,string,omitempty"`
	CreatedAt            *time.Time            `json:"created_at,string,omitempty"`
	ExpiresAt            *time.Time            `json:"expires_at,string,omitempty"`
	Confirmation         *Confirmation         `json:"confirmation,omitempty"`
	Test                 bool                  `json:"test,omitempty"`
	RefundedAmount       *RefundedAmount       `json:"refunded_amount,omitempty"`
	Paid                 bool                  `json:"paid,omitempty"`
	Refundable           bool                  `json:"refundable,omitempty"`
	ReceiptRegistration  string                `json:"receipt_registration,omitempty"`
	SavePaymentMethod    bool                  `json:"save_payment_method,omitempty"`
	Capture              bool                  `json:"capture,omitempty"`
	ClientIP             string                `json:"client_ip,omitempty"`
	Metadata             interface{}           `json:"metadata,omitempty"`
	CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorization_details,omitempty"`
	Airline              *Airline              `json:"airline,omitempty"`
}

//Amount struct is payment.amount object
type Amount struct {
	Value    float64 `json:"value,string"`
	Currency string  `json:"currency"`
}

//Recipient struct is payment.recipient object
type Recipient struct {
	AccountID uint32 `json:"account_id,string,omitempty"`
	GatewayID uint32 `json:"gateway_id,string"`
}

//Requestor struct is payment.requestor object
type Requestor struct {
	Type       string `json:"type"`
	AccountID  uint32 `json:"account_id,string,omitempty"`
	ClientID   string `json:"client_id,omitempty"`
	ClientName string `json:"client_name,omitempty"`
}

//PaymentMethod struct is payment.payment_method object
type PaymentMethod struct {
	Type                string            `json:"type"`
	ID                  string            `json:"id,omitempty"`
	Saved               bool              `json:"saved,omitempty"`
	Title               string            `json:"title,omitempty"`
	Login               string            `json:"login,omitempty"`
	Phone               uint64            `json:"phone,string,omitempty"`
	Card                *Card             `json:"card,omitempty"`
	PayerBankDetails    *PayerBankDetails `json:"payer_bank_details,omitempty"`
	PaymentPurpose      string            `json:"payment_purpose,omitempty"`
	VatData             *VatData          `json:"vat_data,omitempty"`
	AccountNumber       uint64            `json:"account_number,string,omitempty"`
	PaymentData         string            `json:"payment_data,omitempty"`
	GoogleTransactionID string            `json:"google_transaction_id,omitempty"`
	PaymentMethodToken  string            `json:"payment_method_token,omitempty"`
}

//Card struct is payment.payment_method.card object
type Card struct {
	Number        uint64 `json:"number,string"`
	ExpiryYear    uint16 `json:"expiry_year,string"`
	ExpiryMonth   string `json:"expiry_month"`
	CSC           string `json:"csc,omitempty"`
	Cardholder    string `json:"cardholder,omitempty"`
	First6        uint32 `json:"first6,string,omitempty"`
	Last4         string `json:"last4,omitempty"`
	CardType      string `json:"card_type,omitempty"`
	IssuerCountry string `json:"issuer_country,omitempty"`
	IssuerName    string `json:"issuer_name,omitempty"`
	Source        string `json:"source,omitempty"`
}

//PayerBankDetails struct is payment.payment_method.payer_bank_details object
type PayerBankDetails struct {
	FullName   string `json:"full_name"`
	ShortName  string `json:"short_name"`
	Address    string `json:"address"`
	INN        string `json:"inn"`
	KPP        string `json:"kpp"`
	BankName   string `json:"bank_name"`
	BankBranch string `json:"bank_branch"`
	BankBik    string `json:"bank_bik"`
	Account    string `json:"account"`
}

//VatData struct is payment.payment_method.vat_data object
type VatData struct {
	Type   string  `json:"type"`
	Amount *Amount `json:"amount,omitempty"`
	Rate   uint8   `json:"rate,string,omitempty"`
}

//Confirmation struct is payment.confirmation object
type Confirmation struct {
	Type             string `json:"type"`
	ConfirmationData string `json:"confirmation_data,omitempty"`
	Locale           string `json:"locale,omitempty"`
	ConfirmationURL  string `json:"confirmation_url,omitempty"`
	Enforce          bool   `json:"enforce,omitempty"`
	ReturnURL        string `json:"return_url,omitempty"`
}

//RefundedAmount struct is payment.refunded_amount object
type RefundedAmount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

//CancellationDetails struct is payment.cancellation_details object
type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

//AuthorizationDetails struct is payment.authorization_details object
type AuthorizationDetails struct {
	RRN      string `json:"rrn,omitempty"`
	AuthCode string `json:"auth_code,omitempty"`
}

//Airline struct is payment.airline object
type Airline struct {
	TicketNumber     string      `json:"ticket_number,omitempty"`
	BookingReference string      `json:"booking_reference,omitempty"`
	Passengers       []Passenger `json:"passengers,omitempty"`
	Legs             []Leg       `json:"legs,omitempty"`
}

//Passenger struct is payment.airline.passenger object
type Passenger struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

//Leg struct is payment.airline.leg object
type Leg struct {
	DepartureAirport   string     `json:"departure_airport"`
	DestinationAirport string     `json:"destination_airport"`
	DepartureDate      *time.Time `json:"departure_date,string"`
	CarrierCode        string     `json:"carrier_code,omitempty"`
}

//CreatePayment func create payment Yandex.Checkout
func (checkout *Checkout) CreatePayment(client *http.Client, V4UUID *uuid.UUID, pay *Payment) (payment *Payment, apierr *Error, err error) {

	b, err := json.Marshal(pay)
	if err != nil {
		return
	}

	b, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "payments", b)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &payment)
	return
}

//GetPayment func receives payment information Yandex.Checkout
func (checkout *Checkout) GetPayment(client *http.Client, id string) (payment *Payment, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "payments/"+id, nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &payment)
	return
}

//CapturePayment func confirm payment Yandex.Checkout
func (checkout *Checkout) CapturePayment(client *http.Client, V4UUID *uuid.UUID, id string, pay *Payment) (payment *Payment, apierr *Error, err error) {

	b, err := json.Marshal(pay)
	if err != nil {
		return
	}

	b, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "payments/"+id+"/capture", b)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &payment)
	return
}

//CancelPayment func cancel payment Yandex.Checkout
func (checkout *Checkout) CancelPayment(client *http.Client, V4UUID *uuid.UUID, id string) (payment *Payment, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "payments/"+id+"/cancel", []byte("{ }"))
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &payment)
	return
}
