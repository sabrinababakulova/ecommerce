package database

import (
  _ "github.com/lib/pq"
  "github.com/sabrinababakulova/golang-crud/internal/types"
)

func CreateProduct(product types.ProductJson) (int64, error) {
  row, err := Db.Exec("INSERT INTO products (name, price, quantity, manufacturer, mainid) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Price, product.Quantity, product.Manufacturer, product.MainID)
  if err != nil {
    panic(err)
  }
  return row.RowsAffected()
}
