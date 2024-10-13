package seeding

import (
  "fmt"
  "math/rand"
  "time"
  "github.com/google/uuid"
)

type Product struct { 
  Name string
  Price float64
  Quantity int
  Manufacturer string
  ID uint32
  MainID int
}

func Seeding() []Product {
  fmt.Println("Seeding...")
  products := []Product{}
  product_names := []string{"Apple", "Banana", "Orange", "Pineapple", "Mango", "Grapes", "Strawberry", "Blueberry", "Raspberry", "Blackberry"}
  product_prices := []float64{0.5, 0.3, 0.4, 1.0, 1.5, 0.8, 0.6, 0.7, 0.9, 1.2}
  product_quantities := []int{100, 200, 150, 50, 75, 125, 175, 225, 250, 300}
  product_manufacturers := []string{"Apple Inc.", "Banana Co.", "Orange Ltd.", "Pineapple Corp.", "Mango LLC", "Grapes Co.", "Strawberry Inc.", "Blueberry Ltd.", "Raspberry Corp.", "Blackberry LLC"}

  for i := 0; i < 15; i++ {
    rand.Seed(time.Now().UnixNano())
    product := Product{
      Name: product_names[rand.Intn(len(product_names))],
      Price: product_prices[rand.Intn(len(product_prices))],
      Quantity: product_quantities[rand.Intn(len(product_quantities))],
      Manufacturer: product_manufacturers[rand.Intn(len(product_manufacturers))],
      ID: uuid.New().ID(),
    }
    products = append(products, product)
  }

  return products
}
