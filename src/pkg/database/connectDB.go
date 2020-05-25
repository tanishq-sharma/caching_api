package database

import(
_ "cache-routine/src/github.com/lib/pq"
  "database/sql"
  "fmt"
  "os"
)

type manageDb struct {
	Db  *sql.DB
}

var state = manageDb{}

const (
host     = "localhost"
port     = 5432
user     = "tanishq"
password = "pass123"
dbname   = "tanishq"
)

func getDbUrl() string {
  p := os.Getenv("DATABASE_URL")
  if p != "" {
    return p
  }
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
      "password=%s dbname=%s sslmode=disable",
      host, port, user, password, dbname)
  return psqlInfo
}

func ConnectToDB(){
fmt.Printf(getDbUrl()+"\n")

db, err := sql.Open("postgres", getDbUrl())
  if err != nil {
    panic(err)
  }
  err = db.Ping()
  if err != nil {
    panic(err)
  }
    fmt.Println("Successfully connected to database")
    state.Db = db

    CreateTable();
}

func  exportDb() *sql.DB {
    return state.Db
}
