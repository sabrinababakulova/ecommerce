package database 

import (
  "github.com/sabrinababakulova/golang-crud/pkg/seeding"
  _ "github.com/lib/pq"
)

func Migrate() {
  products := seeding.Seeding()

    _, err := Db.Exec(`DROP TABLE IF EXISTS products`)
    if err != nil {
      panic(err)
    }
    createTableSQL := `
    CREATE TABLE products (
      id SERIAL PRIMARY KEY,
      name TEXT,
      price FLOAT,
      quantity INT,
      manufacturer TEXT,
      mainid INT
    );`

    _, createErr := Db.Exec(createTableSQL)

    if createErr != nil {
      panic(createErr)
    }


  for _, product := range products {
     Db.Exec("INSERT INTO products (name, price, quantity, manufacturer, mainid) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Price, product.Quantity, product.Manufacturer, product.ID)
  }

}
