package store

import (
	"account/model"
	"errors"
	"sync"
)

// Error strings
const (
	errAccountIsExist  = "Account is already exist"
	errAccountNotExist = "Account is not exist"
)

// Return interface to account storage
func GetAccountKeeper() model.AccountKeeper {
	return &storage
}

// Singleton instance of account storage
var storage = accountStorage{items: make(map[string]*model.Account)}

// Account storage definition
type accountStorage struct {
	sync.RWMutex                           // For access synchronization
	items        map[string]*model.Account // Account name is a map key
}

// Put new account to the storage
func (this *accountStorage) InsertAccount(item *model.Account) error {
	this.Lock() // Lock storage for write
	defer this.Unlock()

	// Check storage for account name
	_, found := this.items[item.Name]
	if found {
		// Account redefinition error
		return errors.New(errAccountIsExist)
	}
	// Put account into map by name
	this.items[item.Name] = item
	return nil
}

// Set new account data by account name
func (this *accountStorage) UpdateAccount(item *model.Account) error {
	this.Lock() // Lock storage for write
	defer this.Unlock()

	// Check storage for account name
	_, found := this.items[item.Name]
	if !found {
		// Account is not found error
		return errors.New(errAccountNotExist)
	}
	// Replace map entry with new account data
	this.items[item.Name] = item
	return nil
}

// Remove account from the storage by name
func (this *accountStorage) DeleteAccount(name string) error {
	this.Lock() // Lock storage for write
	defer this.Unlock()

	// Check storage for account name
	_, found := this.items[name]
	if !found {
		// Account is not found error
		return errors.New(errAccountNotExist)
	}
	// Delete map entry
	delete(this.items, name)
	return nil
}

// Get account from the storage by name
func (this *accountStorage) GetAccount(name string) (item *model.Account, err error) {
	this.RLock() // Lock storage for read
	defer this.RUnlock()

	// Check storage for account name
	item, found := this.items[name]
	if !found {
		// Account is not found error
		return nil, errors.New(errAccountNotExist)
	}
	// Return map entry
	return item, nil
}

// Get all accounts from the storage
func (this *accountStorage) GetAccountList() model.AccountList {
	this.RLock() // Lock storage for read
	defer this.RUnlock()

	// Init empty account list
	list := make(model.AccountList, 0)
	// Iterate through all map entries
	for _, item := range this.items {
		// Add account to list
		list = append(list, item)
	}
	return list
}
