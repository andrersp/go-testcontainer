package repository

import (
	"context"

	"github.com/andrersp/test-containers/customer"
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	conn *pgx.Conn
}

func NewRepository(ctx context.Context, connStr string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}

	return &Repository{
		conn: conn,
	}, nil
}

func (r Repository) CreateCustomer(ctx context.Context, customer customer.Customer) (customer.Customer, error) {
	query := "INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id"
	row := r.conn.QueryRow(ctx, query, customer.Name, customer.Email)
	err := row.Scan(&customer.Id)
	return customer, err
}

func (r Repository) GetCustomerByEmail(ctx context.Context, email string) (customer.Customer, error) {

	var customer customer.Customer
	query := "SELECT * from customers WHERE email = $1"
	row := r.conn.QueryRow(ctx, query, email)
	err := row.Scan(&customer.Id, &customer.Name, &customer.Email)

	if err != nil {
		return customer, err
	}
	return customer, nil
}
