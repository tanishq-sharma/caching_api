package server

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"pkg/initialize"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = "/home/tanishq/Desktop/go/static/web.html" 			//implement this in Front end
	http.ServeFile(w, r, r.URL.Path)
}

func Set(w http.ResponseWriter,r *http.Request) {
	initialize.HandleSet(w,r)
}


func Get(w http.ResponseWriter, r *http.Request){
	initialize.HandleGet(w,r)
}


func StartServer(){
	initialize.Initialize()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/",homeLink)
	router.HandleFunc("/cache/set",Set).Methods("POST")
	router.HandleFunc("/cache/get/{id}",Get)
	log.Fatal(http.ListenAndServe(":8080",router))
	
}
