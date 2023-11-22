package main 

import (
	"database/sql"
	"fmt"
	"github.com/GabrielMessiasdaRosa/golang-intensivo/internal/infra/database"
	"github.com/GabrielMessiasdaRosa/golang-intensivo/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	useCase := usecase.CalculateFinalPriceFactory(orderRepository)

	input := usecase.OrderInput{
		ID: "123",
		Price: 14,
		Tax: 1,
	}
	output, err := useCase.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}
