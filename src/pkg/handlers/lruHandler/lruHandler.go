package lruHandler

import(
	"cache-routine/src/pkg/datastructures"
	"cache-routine/src/pkg/database"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"
)


func HandleSize(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newSize datastructures.Size_json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(reqBody, &newSize)
	size := int(newSize.Size.(float64))
	SetNewSize(size)
	w.WriteHeader(http.StatusCreated)
}



func HandleSet(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var newElement datastructures.Node_json
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

w.WriteHeader( http.StatusCreated )

json.Unmarshal(reqBody, &newElement)

id := newElement.ID
value := newElement.Value

Cache.Set(id,value)
database.SetToDB(id , value)
fmt.Println(newElement.ID)
fmt.Println(newElement.Value)

	w.WriteHeader( http.StatusCreated )
}

func HandleGet(w http.ResponseWriter,r *http.Request){
	key, ok := r.URL.Query()["a"]
	id := key[0]

	val, ok := Cache.Get(id,w,r)

	var newElement datastructures.Node_json
	if ok==true {
		newElement.ID=id
		newElement.Value=val;
	}	else {
	err,val :=	database.GetFromDB(id)
		if err==nil{
			newElement.ID=id
			newElement.Value=val;
			Cache.Set(id,val)
		}
	}

	json.NewEncoder(w).Encode(newElement)
		w.WriteHeader( http.StatusCreated )
}

func HandleState(w http.ResponseWriter,r *http.Request){
state := Cache.Dump()
	json.NewEncoder(w).Encode(state)
}
