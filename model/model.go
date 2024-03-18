package model

type CoinResp struct {
	Data map[string]Coin `json:"data"`
}

type Coin struct {
	Id     int64            `json:"id"`
	Symbol string           `json:"symbol"`
	Slug   string           `json:"slug"`
	Quote  map[string]Price `json:"quote"`
}

type Price struct {
	Price           float64 `json:"price"`
	VolumeChange24h float64 `json:"volume_change_24h"`
}
