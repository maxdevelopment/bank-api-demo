package account

import (
	"net/http"
	"bank-api-demo/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func HandlerCreate(w http.ResponseWriter, r *http.Request) {
	acc, _ := models.CreateAccount()
	json.NewEncoder(w).Encode(&acc)
}

func HandlerClose(w http.ResponseWriter, r *http.Request) {
	accId := mux.Vars(r)["id"]
	acc, _ := models.DeleteAccount(accId)
	json.NewEncoder(w).Encode(&acc)
}

func HandlerWithdraw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		return
	}

	acc, _ := models.WithdrawAccount(id, sum)
	json.NewEncoder(w).Encode(&acc)

}

func HandlerDeposit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		return
	}

	acc, _ := models.DepositAccount(id, sum)
	json.NewEncoder(w).Encode(&acc)
}

func HandlerTransfer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idFrom := r.FormValue("idFrom")
	idTo := r.FormValue("idTo")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		return
	}

	acc, _ := models.TransferAccount(idFrom, idTo, sum)
	json.NewEncoder(w).Encode(&acc)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
