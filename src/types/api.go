package types

type ApiResponse struct {
	Oficial     DolarData `json:"oficial"`
	Blue        DolarData `json:"blue"`
	OficialEuro DolarData `json:"oficial_euro"`
	BlueEuro    DolarData `json:"blue_euro"`
}

type DolarData struct {
	ValueAvg float64 `json:"value_avg"`
}
