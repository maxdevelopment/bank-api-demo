package routes

import (
	"net/http"
	"io/ioutil"
	"github.com/gorilla/mux"
	"log"
	"bank-api-demo/api/account"
)

func NewRoutes() *mux.Router {
	m := mux.NewRouter().StrictSlash(true)
	m.HandleFunc("/", GetIndex).Methods("GET")
	m.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("web/dist/"))))

	api := m.PathPrefix("/api/").Subrouter()

	api.HandleFunc("/open-account", account.HandlerCreate).Methods("POST")
	api.HandleFunc("/close-account/{id}", account.HandlerClose).Methods("DELETE")

	api.HandleFunc("/withdraw", account.HandlerWithdraw).Methods("PUT")
	api.HandleFunc("/deposit", account.HandlerDeposit).Methods("PUT")
	api.HandleFunc("/transfer", account.HandlerTransfer).Methods("POST")

	return m
}

func GetIndex(writer http.ResponseWriter, request *http.Request) {
	indexFile, err := ioutil.ReadFile("web/index.html")
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Write(indexFile)
}