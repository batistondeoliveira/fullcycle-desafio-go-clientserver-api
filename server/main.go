package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	//"time"
	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/http_request"
	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/internal/exchange/entity"
	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/internal/infra/database"
	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/cotacao", DollarExchangeHandler)
	http.ListenAndServe(":8080", nil)
}

type ExchangeResponse struct {
	Bid string `json:"bid"`
}

func DollarExchangeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	exchange, err := doRequest()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	err = doSave(exchange)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(ExchangeResponse{
		Bid: exchange.Usdbrl.Bid,
	})
}

func doRequest() (*http_request.Exchange, error) {
	exchange, err := http_request.DollarExchangeRequest()
	if err != nil {
		return nil, err
	}
	return exchange, nil
}

func doSave(exchange *http_request.Exchange) error {
	db, err := sql.Open("sqlite3", "./exchange.db")
	defer db.Close()
	if err != nil {
		return err
	}

	newExchange := entity.NewExchange(
		uuid.New().String(),
		exchange.Usdbrl.Code,
		exchange.Usdbrl.Codein,
		exchange.Usdbrl.Name,
		exchange.Usdbrl.High,
		exchange.Usdbrl.Low,
		exchange.Usdbrl.VarBid,
		exchange.Usdbrl.PctChange,
		exchange.Usdbrl.Bid,
		exchange.Usdbrl.Ask,
		exchange.Usdbrl.Timestamp,
		exchange.Usdbrl.CreateDate,
	)

	repo := database.NewExchangeRepository(db)
	err = repo.Save(newExchange)
	if err != nil {
		return err
	}

	return nil
}
