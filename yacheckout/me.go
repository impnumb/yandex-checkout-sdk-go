package yacheckout

import (
	"encoding/json"
	"net/http"
)

//Me struct is Yandex.Checkout me object
type Me struct {
	AccountID            int  `json:"account_id"`
	Test                 bool `json:"test"`
	FiscalizationEnabled bool `json:"fiscalization_enabled"`
}

//GetMe func receives me information Yandex.Checkout
func (checkout *Checkout) GetMe(client *http.Client) (me *Me, apierr *Error, err error) {

	b, apierr, err := checkout.Exec(APIEndpoint, client, http.MethodGet, nil, "me", nil)
	if err != nil || apierr != nil {
		return
	}

	err = json.Unmarshal(b, &me)
	return
}
