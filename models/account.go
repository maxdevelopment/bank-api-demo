package models

import (
	"time"
	"bank-api-demo/utils"
	"sync"
	"errors"
)

type Account struct {
	ID        string    `json:"id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	mu        sync.Mutex
}

var accountList = make(map[string]*Account)


var (
	ErrAccountNotPreset = errors.New("account not present")
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrSameAccounts = errors.New("you cannot transfer money between the same accounts")
)

func isPresent(accId string) (*Account, bool) {
	if acc, ok := accountList[accId]; !ok {
		return nil, false
	} else {
		return acc, true
	}
}

func CreateAccount() (*Account, error) {
	accId := utils.RandStringBytes(8)
	timestamp := time.Now()

	acc := Account{
		ID:        accId,
		Balance:   0,
		CreatedAt: timestamp,
	}

	accountList[accId] = &acc

	return &acc, nil
}

func DeleteAccount(accId string) (*Account, error) {

	acc, ok := isPresent(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}
	delete(accountList, accId)

	return acc, nil
}

func WithdrawAccount(accId string, sum float64) (*Account, error) {

	acc, ok := isPresent(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.Balance < sum {
		return nil, ErrNotEnoughMoney
	}

	acc.Balance -= sum

	return acc, nil
}

func DepositAccount(accId string, sum float64) (*Account, error) {

	acc, ok := isPresent(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.Balance += sum

	return acc, nil
}

func TransferAccount(accIdFrom string, accIdTo string, sum float64) (interface{}, error) {

	if accIdFrom == accIdTo {
		return nil, ErrSameAccounts
	}

	accFrom, ok := isPresent(accIdFrom)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	accTo, ok := isPresent(accIdTo)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	accFrom.mu.Lock()
	accTo.mu.Lock()
	defer func() {
		accFrom.mu.Unlock()
		accTo.mu.Unlock()
	}()

	if accFrom.Balance < sum {
		return nil, ErrNotEnoughMoney
	}

	accFrom.Balance -= sum
	accTo.Balance += sum

	return &accountList, nil
}

func GetAccounts() interface{} {
	return &accountList
}