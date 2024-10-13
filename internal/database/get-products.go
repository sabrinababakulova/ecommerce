package database 

import (
  "github.com/sabrinababakulova/golang-crud/pkg/seeding"
  _ "github.com/lib/pq"
)

func GetProducts() []seeding.Product {
    products := []seeding.Product{}
    rows, err := Db.Query(`Select * FROM products`)

    if err != nil {
      panic(err)
    }
    defer rows.Close()

    for rows.Next() {
      var product seeding.Product
      errRows := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Quantity, &product.Manufacturer, &product.MainID)
      if errRows != nil {
        panic(errRows)
      }
      products = append(products, product)
    }
    return products
}
