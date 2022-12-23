package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const REQUEST_MAX_TIMEOUT = 300 * time.Millisecond

type Exchange struct {
	Bid string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), REQUEST_MAX_TIMEOUT)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)
	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var exchange Exchange
	err = json.Unmarshal(body, &exchange)
	if err != nil {
		panic(err)
	}

	doGravar(exchange)
}

func doGravar(exchange Exchange) {
	f, err := os.Create("cotacao.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	tamanho, err := f.WriteString(fmt.Sprintf("Dólar: {%s}", exchange.Bid))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! Tamanho: %d bytes", tamanho)
}
