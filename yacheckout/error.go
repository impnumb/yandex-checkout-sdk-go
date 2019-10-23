package yacheckout

//Error struct is response error Yandex.Checkout
type Error struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter"`
	RetryAfter  uint32 `json:"retry_after"`
}
