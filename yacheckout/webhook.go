package yacheckout

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

//Webhook struct is Yandex.Checkout webhook object
type Webhook struct {
	ID    string `json:"id,omitempty"`
	Event string `json:"event"`
	URL   string `json:"url"`
}

//Webhooks struct is Yandex.Checkout webhooks object
type Webhooks struct {
	Type  string    `json:"type"`
	Items []Webhook `json:"items"`
}

//CreateWebhook func create webhook Yandex.Checkout
func (checkout *Checkout) CreateWebhook(client *http.Client, V4UUID *uuid.UUID, webhk *Webhook) (webhook *Webhook, apierr *Error, err error) {

	b, err := json.Marshal(webhk)
	if err != nil {
		return
	}

	b, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodPost, V4UUID, "webhooks", b)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &webhook)
	return
}

//GetWebhooks func receives webhooks Yandex.Checkout
func (checkout *Checkout) GetWebhooks(client *http.Client) (webhook *Webhooks, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "webhooks", nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &webhook)
	return
}

//DeleteWebhook func delete webhook Yandex.Checkout
func (checkout *Checkout) DeleteWebhook(client *http.Client, id string) (apierr *Error, err error) {

	_, apierr, err = checkout.Exec(APIEndpoint, client, http.MethodDelete, nil, "webhooks/"+id, nil)
	return
}
