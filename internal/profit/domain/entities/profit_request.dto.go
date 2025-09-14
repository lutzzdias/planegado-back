package entity

type ProfitRequest struct {
	Herd Herd `json:"herd"`
	Feed Feed `json:"feed"`
	Sale Sale `json:"sale"`
	Days int  `json:"days"`
}
