package entity

type Exchange struct {
	ID         string
	Code       string
	Codein     string
	Name       string
	High       string
	Low        string
	VarBid     string
	PctChange  string
	Bid        string
	Ask        string
	Timestamp  string
	CreateDate string
}

func NewExchange(
	id string,
	code string,
	codein string,
	name string,
	high string,
	low string,
	varBid string,
	pctChange string,
	bid string,
	ask string,
	timestamp string,
	createDate string,
) *Exchange {
	return &Exchange{
		ID:         id,
		Code:       code,
		Codein:     codein,
		Name:       name,
		High:       high,
		Low:        low,
		VarBid:     varBid,
		PctChange:  pctChange,
		Bid:        bid,
		Ask:        ask,
		Timestamp:  timestamp,
		CreateDate: createDate,
	}
}
