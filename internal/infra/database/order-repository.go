package database

import (
	"database/sql"
    "github.com/GabrielMessiasdaRosa/golang-intensivo/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		Db: db,
	}
}

func createTable(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS orders (id TEXT PRIMARY KEY NOT NULL UNIQUE, price REAL, tax REAL, final_price REAL)")
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) Save(order *entity.Order) error {
	errOnCreateTable := createTable(orderRepository.Db)
	if errOnCreateTable != nil {
		return errOnCreateTable
	}
	_, err := orderRepository.Db.Exec("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (orderRepository *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := orderRepository.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

