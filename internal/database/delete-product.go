package database

func DeleteProduct(id string) (int64, error) {
  row, err := Db.Exec("DELETE FROM products WHERE id = $1", id)
  if err != nil {
    panic(err)
  }
  return row.RowsAffected()
}
