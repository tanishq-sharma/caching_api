package database

import(
_ "cache-routine/src/github.com/lib/pq"
  "cache-routine/src/pkg/datastructures"
  "fmt"
 "encoding/json"
 "log"
)

type T interface{}

const (

CreateTableStatement = `CREATE TABLE IF NOT EXISTS cache (id text PRIMARY KEY,

                    value JSONb);`

InsertStatement = `INSERT INTO cache (id,value)
                  SELECT $1, $2
                  WHERE NOT EXISTS (SELECT 1 FROM cache WHERE id=$1);`

UpdateStatement = `UPDATE cache SET value=$2 WHERE id=$1 ;`


QueryStatement = `SELECT * FROM cache WHERE id = ($1)`
)

func SetToDB(key T,value T){
    db:=exportDb()
    record := datastructures.Object{
        Value : value,
      }
    var jsonData []byte
    jsonData, err1 := json.Marshal(record)
    if err1 != nil {
        log.Println(err1)
    }

    str := fmt.Sprint(key)

    _, err1 = db.Exec(UpdateStatement,str,jsonData)
    if err1 != nil {
      panic(err1)
    }

    _, err1 = db.Exec(InsertStatement,str,jsonData)
    if err1 != nil {
      panic(err1)
    }
}


func GetFromDB(key T) (error, T) {
  db:=exportDb()
  str := fmt.Sprint(key)

  sqlStatement := `SELECT value FROM cache WHERE id = $1;`
  var user []byte
  row := db.QueryRow(sqlStatement,str)
  err2 := row.Scan(&user)

  var jsonValue datastructures.Object
    json.Unmarshal(user,&jsonValue)

    return err2,jsonValue.Value


}

func CreateTable(){
  db:=exportDb()

if _, err := db.Exec(CreateTableStatement); err != nil {
                fmt.Sprintf("Error creating database table: %q", err)
            return
        }
}
