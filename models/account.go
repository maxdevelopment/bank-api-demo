package models

import (
	"time"
	"bank-api-demo/utils"
	"fmt"
)

type Account struct {
	ID        string    `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

var accountList = make(map[string]Account)

func isPresent(accId string) bool {
	if _, ok := accountList[accId]; !ok {
		return false
	}
	return true
}

func CreateAccount() (*Account, error) {
	accId := utils.RandStringBytes(8)
	timestamp := time.Now()

	acc := Account{
		ID:        accId,
		Balance:   0,
		CreatedAt: timestamp,
	}

	accountList[accId] = acc

	return &acc, nil
}

func DeleteAccount(accId string) (*Account, error) {

	if !isPresent(accId) {
		return nil,nil
	}

	delete(accountList, accId)
	return nil,nil
}

func WithdrawAccount(accId string, sum float64) (*Account, error) {

	fmt.Println(accId)
	fmt.Println(sum)

	return nil,nil
}

func DepositAccount(accId string, sum float64) (*Account, error) {
	fmt.Println(accId)
	fmt.Println(sum)

	return nil,nil
}

func TransferAccount(accIdFrom string, accIdTo string, sum float64) (*Account, error) {
	fmt.Println(accIdFrom)
	fmt.Println(accIdTo)
	fmt.Println(sum)

	return nil,nil
}
