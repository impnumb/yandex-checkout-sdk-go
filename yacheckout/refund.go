package yacheckout

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//Refund struct is Yandex.Checkout refund object
type Refund struct {
	ID          string     `json:"id,omitempty"`
	PaymentID   string     `json:"payment_id"`
	Requestor   *Requestor `json:"requestor,omitempty"`
	Status      string     `json:"status,omitempty"`
	CreatedAt   *time.Time `json:"created_at,string,omitempty"`
	Amount      *Amount    `json:"amount,omitempty"`
	Description string     `json:"description,omitempty"`
	Receipt     *Receipt   `json:"receipt,omitempty"`
}

//CreateRefund func create refund Yandex.Checkout
func (checkout *Checkout) CreateRefund(client *http.Client, V4UUID *uuid.UUID, rfd *Refund) (refund *Refund, apierr *Error, err error) {

	b, err := json.Marshal(rfd)
	if err != nil {
		return
	}

	b, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "refunds", b)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &refund)
	return
}

//GetRefund func receives refund information Yandex.Checkout
func (checkout *Checkout) GetRefund(client *http.Client, id string) (refund *Refund, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "refunds/"+id, nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &refund)
	return
}
