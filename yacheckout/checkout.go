package yacheckout

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

//APIEndpoints
const (
	APIEndpoint      = "https://payment.yandex.net/api/v3/"
)

//Payment methods.See https://kassa.yandex.ru/developers/payment-methods/overview
const (
	Alfabank      = "alfabank"
	MobileBalance = "mobile_balance"
	BankCard      = "bank_card"
	Installments  = "installments"
	Cash          = "cash"
	B2BSberbank   = "b2b_sberbank"
	Sberbank      = "sberbank"
	TinkoffBank   = "tinkoff_bank"
	YandexMoney   = "yandex_money"
	ApplePay      = "apple_pay"
	GooglePay     = "google_pay"
	Qiwi          = "qiwi"
	WeChat        = "wechat"
	Webmoney      = "webmoney"
)

//Bank card types.See https://kassa.yandex.ru/developers/api#payment_object_payment_method_bank_card_card_card_type
const (
	MasterCard      = "MasterCard"
	Visa            = "Visa"
	Mir             = "Mir"
	UnionPay        = "UnionPay"
	JCB             = "JCB"
	AmericanExpress = "AmericanExpress"
	DinersClub      = "DinersClub"
	Unknown         = "Unknown"
)

//Cancellation details party.See https://kassa.yandex.ru/developers/payments/declined-payments#cancellation-details-party
const (
	YandexCheckout = "yandex_checkout"
	PaymentNetwork = "payment_network"
	Merchant       = "merchant"
)

//Cancellation details reason.See https://kassa.yandex.ru/developers/payments/declined-payments#cancellation-details-reason
const (
	DSecureFailed              = "3d_secure_failed"
	CallIssuer                 = "call_issuer"
	CardExpired                = "card_expired"
	CountryForbidden           = "country_forbidden"
	FraudSuspected             = "fraud_suspected"
	GeneralDecline             = "general_decline"
	IdentificationRequired     = "identification_required"
	InsufficientFunds          = "insufficient_funds"
	InvalidCardNumber          = "invalid_card_number"
	InvalidCSC                 = "invalid_csc"
	IssuerUnavailable          = "issuer_unavailable"
	PaymentMethodLimitExceeded = "payment_method_limit_exceeded"
	PaymentMethodRestricted    = "payment_method_restricted"
	PermissionRevoked          = "permission_revoked"
)

//Webhook events.See https://kassa.yandex.ru/developers/using-api/webhooks#events
const (
	PaymentWaitingForCapture = "payment.waiting_for_capture"
	PaymentSucceeded         = "payment.succeeded"
	PaymentCanceled          = "payment.canceled"
	RefundSucceeded          = "refund.succeeded"
)

//Payment Statuses.See https://kassa.yandex.ru/developers/payments/basics/payment-process#payment-statuses
const (
	Pending           = "pending"
	WaitingForCapture = "waiting_for_capture"
	Succeeded         = "succeeded"
	Canceled          = "canceled"
)

//Tax systems.See https://kassa.yandex.ru/developers/payments/54fz/parameters-values#tax-systems
const (
	GeneralTaxationSystem = iota + 1
	USNIncome
	USNProfit
	ENVD
	ESN
	PatentTaxSystem
)

//VAT codes.See https://kassa.yandex.ru/developers/payments/54fz/parameters-values#vat-codes
const (
	WithoutVAT = iota + 1
	VAT0
	VAT10
	VAT20
	VAT10110
	VAT20120
)

//Payment subject.See https://kassa.yandex.ru/developers/54fz/parameters-values#payment-subject
const (
	Commodity            = "commodity"
	Excise               = "excise"
	Job                  = "job"
	Service              = "service"
	GamblingBet          = "gambling_bet"
	GamblingPrize        = "gambling_prize"
	Lottery              = "lottery"
	LotteryPrize         = "lottery_prize"
	IntellectualActivity = "intellectual_activity"
	Paymentc             = "payment"
	AgentCommission      = "agent_commission"
	PropertyRight        = "property_right"
	NonOperatingGain     = "non_operating_gain"
	InsurancePremium     = "insurance_premium"
	SalesTax             = "sales_tax"
	ResortFee            = "resort_fee"
	Composite            = "composite"
	Another              = "another"
)

//Payment mode.See https://kassa.yandex.ru/developers/54fz/parameters-values#payment-mode
const (
	FullPrepayment    = "full_prepayment"
	PartialPrepayment = "partial_prepayment"
	Advance           = "advance"
	FullPayment       = "full_payment"
	PartialPayment    = "partial_payment"
	Credit            = "credit"
	CreditPayment     = "credit_payment"
)

//Settlement type.See https://kassa.yandex.ru/developers/54fz/parameters-values#settlement-type
const (
	Cashless      = "cashless"
	Prepayment    = "prepayment"
	Postpayment   = "postpayment"
	Consideration = "consideration"
)

//Checkout struct is Yandex.Checkout
type Checkout struct {
	ShopID        int
	SecurityToken string
	OAuthToken    string
}

//NewCheckout func return Checkout struct
func NewCheckout(id int, stoken, oatoken string) *Checkout {
	return &Checkout{ShopID: id, SecurityToken: stoken, OAuthToken: oatoken}
}

//Exec func is custom execution
//V4UUID - https://checkout.yandex.com/docs/checkout-api/#idempotence
//https://kassa.yandex.ru/developers/using-api/basics#idempotence
//Use https://godoc.org/github.com/google/uuid#NewRandom
func (checkout *Checkout) Exec(endpoint string, client *http.Client, httpMethod string, V4UUID *uuid.UUID, method string, data []byte) (b []byte, apierr *Error, err error) {

	var req *http.Request

	switch httpMethod {
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, endpoint+method, nil)
	case http.MethodPost:
		req, err = http.NewRequest(http.MethodPost, endpoint+method, bytes.NewBuffer(data))
		req.Header.Set("Idempotence-Key", V4UUID.String())
		req.Header.Set("Content-Type", "application/json")
	case http.MethodDelete:
		req, err = http.NewRequest(http.MethodDelete, endpoint+method, nil)
		req.Header.Set("Content-Type", "application/json")
	default:
		err = errors.New("Unknown HTTP method")
		return
	}

	if err != nil {
		return
	}

	if checkout.OAuthToken == "" {
		req.SetBasicAuth(strconv.Itoa(checkout.ShopID), checkout.SecurityToken)
	} else {
		req.Header.Set("Authorization", "Bearer "+checkout.OAuthToken)
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	b, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = json.Unmarshal(b, &apierr)
	}

	return
}
