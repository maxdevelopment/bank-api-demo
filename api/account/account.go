package account

import (
	"net/http"
	"github.com/maxdevelopment/bank-api-demo/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func HandlerCreate(w http.ResponseWriter, r *http.Request) {
	acc, _ := models.CreateAccount()
	respondWithJSON(w, http.StatusOK, acc)
}

func HandlerClose(w http.ResponseWriter, r *http.Request) {
	accId := mux.Vars(r)["id"]
	acc, err := models.DeleteAccount(accId)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, acc)
}

func HandlerWithdraw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	acc, err := models.WithdrawAccount(id, sum)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, acc)

}

func HandlerDeposit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	acc, err := models.DepositAccount(id, sum)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, acc)
}

func HandlerTransfer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	idFrom := r.FormValue("idFrom")
	idTo := r.FormValue("idTo")
	sum, err := strconv.ParseFloat(r.FormValue("sum"), 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}

	accounts, err := models.TransferAccount(idFrom, idTo, sum)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, &accounts)
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
