package initialize

import(
	"pkg/datastructures"
	"lru"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)
var C *lru.LRU

func HandleSet(w http.ResponseWriter,r *http.Request){
	var newElement datastructures.Node_json

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader( http.StatusCreated )

	json.Unmarshal(reqBody, &newElement)

	C.Set(newElement.ID,newElement.Value)

	w.WriteHeader( http.StatusCreated ) 	// Not WORKING ********(fixed)
}

func HandleGet(w http.ResponseWriter,r *http.Request){
/*	r.ParseForm()											// add this in front end
	id := fmt.Sprint(r.Form["Key"])							//same
	id = strings.Trim(id,"[")								//add formatting of input response in front end 
	id = strings.Trim(id,"]")
*/
	id:= mux.Vars(r)["id"]	
	val, ok := C.Get(id,w,r)
	var newElement datastructures.Node_json
	if ok==true{											
		newElement.ID=id
		newElement.Value=val;
	}

	json.NewEncoder(w).Encode(newElement)
		w.WriteHeader( http.StatusCreated )
}

func HandleState(w http.ResponseWriter,r *http.Request){
	var currentState datastructures.State_json
	currentState.State = C.Dump();
	json.NewEncoder(w).Encode(currentState)
		w.WriteHeader( http.StatusCreated )
}


func Initialize(){
	fmt.Printf("Enter the size of LRU cache : ")
	var count int
	fmt.Scan(&count)
	C = lru.New(count);
	fmt.Printf("LRU cache of size %v has been created  \n ",count)

}

