package models

import (
	"time"
	"github.com/maxdevelopment/bank-api-demo/utils"
	"sync"
	"errors"
)

//Balance returns value in cents
type Account struct {
	ID        string    `json:"id"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	mu        sync.Mutex
}

func (a *Account) increase(sum float64) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Balance += int64(sum * 100)
}

func (a *Account) decrease(sum float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	s := int64(sum * 100)
	if a.Balance < s {
		return ErrNotEnoughMoney
	}
	a.Balance -= s
	return nil
}

type AccountList struct {
	list map[string]*Account
	sync.Mutex
}

type Transfer struct {
	From *Account `json:"from"`
	To   *Account `json:"to"`
}

var al = AccountList{
	list: make(map[string]*Account),
}

func (al *AccountList) getAccount(accId string) (*Account, bool) {
	al.Lock()
	defer al.Unlock()
	if acc, ok := al.list[accId]; !ok {
		return nil, false
	} else {
		return acc, true
	}
}

func (al *AccountList) setAccount(account *Account) {
	al.Lock()
	defer al.Unlock()
	al.list[account.ID] = account
}

func (al *AccountList) deleteAccount(accId string) {
	al.Lock()
	defer al.Unlock()
	delete(al.list, accId)
}

var (
	ErrAccountNotPreset = errors.New("account not present")
	ErrNotEnoughMoney   = errors.New("not enough money")
	ErrSameAccounts     = errors.New("you cannot transfer money between the same accounts")
)

func CreateAccount() (*Account, error) {
	accId := utils.RandStringBytes(8)
	timestamp := time.Now()

	acc := Account{
		ID:        accId,
		Balance:   0,
		CreatedAt: timestamp,
	}
	al.setAccount(&acc)

	return &acc, nil
}

func DeleteAccount(accId string) (*Account, error) {

	acc, ok := al.getAccount(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}
	al.deleteAccount(accId)

	return acc, nil
}

func WithdrawAccount(accId string, sum float64) (*Account, error) {

	acc, ok := al.getAccount(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	err := acc.decrease(sum)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func DepositAccount(accId string, sum float64) (*Account, error) {

	acc, ok := al.getAccount(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	acc.increase(sum)

	return acc, nil
}

func TransferAccount(accIdFrom string, accIdTo string, sum float64) (*Transfer, error) {

	if accIdFrom == accIdTo {
		return nil, ErrSameAccounts
	}

	accFrom, ok := al.getAccount(accIdFrom)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	accTo, ok := al.getAccount(accIdTo)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	err := accFrom.decrease(sum)
	if err != nil {
		return nil, err
	}

	accTo.increase(sum)

	ta := Transfer{
		From: accFrom,
		To:   accTo,
	}

	return &ta, nil
}
