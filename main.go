package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/neviim/golang-estudo/internal/entity/usecase"
	"github.com/neviim/golang-estudo/internal/infra/database"
)

func main() {
	// println("Jesus é o Senhor!")

	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close() // defer espera tudo fechar e depois roda o close
	OrderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(OrderRepository)

	input := usecase.OrderInput{
		ID:    "125",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)

}
