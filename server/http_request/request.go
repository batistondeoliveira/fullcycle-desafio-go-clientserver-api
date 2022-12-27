package http_request

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/internal/_shared/constants"
)

type Exchange struct {
	Usdbrl Usdbrl `json:"USDBRL"`
}

type Usdbrl struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

func DollarExchangeRequest() (*Exchange, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		constants.REQUEST_MAX_TIMEOUT,
	)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, constants.URL, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Request Exception: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var e Exchange
	err = json.Unmarshal(body, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}
