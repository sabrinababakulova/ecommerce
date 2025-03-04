package database

import (
   "database/sql"
   "fmt"
   "github.com/joho/godotenv"
   _ "github.com/lib/pq"
   "os"
   "strconv"
)

var Db *sql.DB

func ConnectDatabase() {
   err := godotenv.Load()
   if err != nil {
      fmt.Println("Error is occurred  on .env file please check")
   }

   host := os.Getenv("DB_HOST")
   port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
   user := os.Getenv("DB_USERNAME")
   dbname := os.Getenv("DB_NAME")
   pass := os.Getenv("DB_PASSWORD")

   // set up postgres sql to open it.
   psqlSetup := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
       host, port, user, dbname, pass)
   db, errSql := sql.Open("postgres", psqlSetup)
   if errSql != nil {
      fmt.Println("There is an error while connecting to the database ", err)
      panic(err)
   } else {
      Db = db
      fmt.Println("Successfully connected to database!")
   }
}
