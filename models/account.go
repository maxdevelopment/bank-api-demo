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

type AccountList struct {
	list map[string]*Account
	sync.RWMutex
}

type Transfer struct {
	from *Account
	to   *Account
}

var al = AccountList{
	list: make(map[string]*Account),
}

func (al *AccountList) getAccount(accId string) (*Account, bool) {
	if acc, ok := al.list[accId]; !ok {
		return nil, false
	} else {
		return acc, true
	}
}

func (al *AccountList) setAccount(account *Account) { //LOCK
	al.RLock()
	defer al.RUnlock()
	al.list[account.ID] = account
}

func (al *AccountList) deleteAccount(accId string) {
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

	acc.mu.Lock()
	defer acc.mu.Unlock()

	if acc.Balance < sum {
		return nil, ErrNotEnoughMoney
	}

	acc.Balance -= sum

	return acc, nil
}

func DepositAccount(accId string, sum float64) (*Account, error) {

	acc, ok := al.getAccount(accId)
	if !ok {
		return nil, ErrAccountNotPreset
	}

	acc.mu.Lock()
	defer acc.mu.Unlock()
	acc.Balance += sum

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

	ta := Transfer{
		from: accFrom,
		to:   accTo,
	}

	return &ta, nil
}
