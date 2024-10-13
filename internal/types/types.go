package types

type ProductJson struct {
  Name        string  `json:"name" binding:"required"`
  Price float64  `json:"price" binding:"required"`
  Quantity       int `json:"quantity" binding:"required"`
  Manufacturer string  `json:"manufacturer" binding:"required"`
  MainID int  `json:"mainid" binding:"required"`
}

