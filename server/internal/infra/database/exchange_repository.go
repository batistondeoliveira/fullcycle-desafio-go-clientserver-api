package database

import (
	"context"
	"database/sql"

	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/internal/_shared/constants"
	"github.com/batistondeoliveira/fullcycle_desafio_go_client_server_api/server/internal/exchange/entity"
)

type ExchangeRepository struct {
	Db *sql.DB
}

func NewExchangeRepository(db *sql.DB) *ExchangeRepository {
	return &ExchangeRepository{
		Db: db,
	}
}

func (r *ExchangeRepository) Save(exchange *entity.Exchange) error {
	ctx, cancel := context.WithTimeout(context.Background(), constants.DB_MAX_TIMEOUT)
	defer cancel()

	stmt, err := r.Db.PrepareContext(ctx,
		"INSERT INTO exchange ("+
			" id, "+
			"	code, "+
			" codein, "+
			"	name, "+
			"	high, "+
			"	low, "+
			"	varbid, "+
			"	pctchange, "+
			" bid, "+
			" ask, "+
			" timestamp, "+
			" create_date "+
			")VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx,
		exchange.ID,
		exchange.Code,
		exchange.Codein,
		exchange.Name,
		exchange.High,
		exchange.Low,
		exchange.VarBid,
		exchange.PctChange,
		exchange.Bid,
		exchange.Ask,
		exchange.Timestamp,
		exchange.CreateDate,
	)
	if err != nil {
		return err
	}

	return nil
}
