package main

import(
	"errors"
	"database/sql"
)

type Product struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
}

func (p*Product)GetProduct(db *sql.DB) error {
	return errors.New("to be implemented")
}

func (p*Product)UpdateProduct(db *sql.DB) error {
	return errors.New("to be implemented")
}

func (p*Product)DeleteProduct(db *sql.DB) error {
	return errors.New("to be implemented")
}

func (p*Product)CreateProduct(db *sql.DB)error  {
	return errors.New("to be implemented")
}

func (p*Product)FetchAllProduct(db *sql.DB, count, start int)([]Product, error)  {
	return nil, errors.New("to be implemented")
}