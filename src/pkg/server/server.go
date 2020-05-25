package server

import (
	"log"
	"net/http"
	"cache-routine/src/github.com/gorilla/handlers"
	"cache-routine/src/github.com/gorilla/mux"
	"cache-routine/src/pkg/handlers/lruHandler"
	"cache-routine/src/pkg/database"
	"os"
	"fmt"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	 http.ServeFile(w, r, "./build/")
}

func Size(w http.ResponseWriter,r *http.Request){
	lruHandler.HandleSize(w,r)
}

func Set(w http.ResponseWriter,r *http.Request) {
	lruHandler.HandleSet(w,r)
}

func Get(w http.ResponseWriter, r *http.Request){
	lruHandler.HandleGet(w,r)
}

func State(w http.ResponseWriter,r *http.Request){
	lruHandler.HandleState(w,r)
}

func getPort() string {
  p := os.Getenv("PORT")
  if p != "" {
    return ":" + p
  }
  return ":8080"
}

func StartServer(){
	lruHandler.Initialize()
	database.ConnectToDB()


	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/",homeLink).Methods("GET")
	router.HandleFunc("/cache/size",Size).Methods("POST")
	router.HandleFunc("/cache/set",Set).Methods("POST")
	router.HandleFunc("/cache/get/",Get).Methods("GET")
	router.HandleFunc("/cache/state",State).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./build/"))))

	originsOk := handlers.AllowedOrigins([]string{"*"})
fmt.Printf("running on port "  + getPort())
	log.Fatal(http.ListenAndServe((getPort()), handlers.CORS(originsOk)(router)))

}
