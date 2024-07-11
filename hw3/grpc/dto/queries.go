package dto

import "fmt"

func NewAS() *AccountStorage {
	return &AccountStorage{Storage: make(map[string]*Account)}
}

func (as *AccountStorage) CreateAccount(account *Account) {
	as.Mtx.Lock()
	defer as.Mtx.Unlock()

	as.Storage[account.Name] = account
}

func (as *AccountStorage) GetAccount(name string) (*Account, error) {
	as.Mtx.RLock()
	defer as.Mtx.RUnlock()

	account, ok := as.Storage[name]
	if !ok {
		return nil, fmt.Errorf(name)
	}

	return account, nil
}

func (as *AccountStorage) RenameAccount(newName, oldName string) error {
	as.Mtx.Lock()
	defer as.Mtx.Unlock()

	account, ok := as.Storage[oldName]
	if !ok {
		return fmt.Errorf(oldName)
	}

	if _, exists := as.Storage[newName]; exists {
		return fmt.Errorf(newName)
	}

	delete(as.Storage, oldName)

	account.Name = newName
	as.Storage[newName] = account

	return nil
}

func (as *AccountStorage) UpdateBalance(account *Account) error {
	as.Mtx.Lock()
	defer as.Mtx.Unlock()

	storedAccount, ok := as.Storage[account.Name]
	if !ok {
		return fmt.Errorf(account.Name)
	}

	storedAccount.Balance = account.Balance

	return nil
}

func (as *AccountStorage) DeleteAccount(name string) error {
	as.Mtx.Lock()
	defer as.Mtx.Unlock()

	if _, ok := as.Storage[name]; !ok {
		return fmt.Errorf(name)
	}
	delete(as.Storage, name)

	return nil
}
