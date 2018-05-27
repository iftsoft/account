package domain

import (
	"account/store"
	"account/model"
	"time"
	"errors"
)

const (
	err_AccountNotEmpty		= "This account is not empty"
	err_AccountNotClosed	= "This account is not closed"
	err_AccountIsClosed 	= "This account is closed"
	err_SourceIsClosed		= "Source account is closed"
	err_TargetIsClosed		= "Target account is closed"
	err_CurrencyMismatch	= "Account and Transaction currency mismatch"
	err_AmountToBig			= "Transaction amount is too big"
)

// Interface to Account API
func GetAccountManager() *AccountManager {
	mng := &AccountManager{}
	// Get interface to account storage
	mng.keeper	= store.GetAccountKeeper()
	return mng
}

// Account manager
type AccountManager struct {
	keeper	model.AccountKeeper		// Interface to account storage
}

// Get Account list from storage
func (this *AccountManager)GetAccountList() (model.AccountList, error) {
	list := this.keeper.GetAccountList()
	return list, nil
}

// Get Account from storage by name
func (this *AccountManager)GetAccount(name string) (*model.Account, error) {
	return this.keeper.GetAccount(name)
}

// Open new account in the storage
func (this *AccountManager)OpenAccount(name, owner, curr string) (*model.Account, error) {
	// Create new account
	item := &model.Account{}
	item.Name		= name
	item.State		= model.State_Active
	item.Owner		= owner
	item.Currency	= curr
	item.Amount		= 0.0
	item.Updated	= time.Now()
	item.Created	= item.Updated

	// Insert Account into store
	err := this.keeper.InsertAccount(item)
	return item, err
}

// Close account in the storage
func (this *AccountManager)CloseAccount(name string) (*model.Account, error) {
	// Get Account from store
	item, err := this.keeper.GetAccount(name)
	if err != nil {
		return item, err
	}
	// Check for account emptiness
	if item.Amount > 0.0 {
		return item, errors.New(err_AccountNotEmpty)
	}
	// Change account state
	item.State = model.State_Closed
	item.Updated	= time.Now()
	// Save changed Account into store
	err = this.keeper.UpdateAccount(item)
	return item, err
}

// Delete account in the storage
func (this *AccountManager)DeleteAccount(name string) (*model.Account, error) {
	// Get Account from store
	item, err := this.keeper.GetAccount(name)
	if err != nil {
		return item, err
	}
	// Check for account is closed
	if item.State != model.State_Closed {
		return item, errors.New(err_AccountNotClosed)
	}
	// Delete Account in store
	err = this.keeper.DeleteAccount(name)
	return item, err
}

// Deposit funds to account amount
func (this *AccountManager)DepositFund(name, curr string, amount float32) (*model.Account, error) {
	// Get Account from store
	item, err := this.keeper.GetAccount(name)
	if err != nil {
		return item, err
	}
	// Check for account is closed
	if item.State == model.State_Closed {
		return item, errors.New(err_AccountIsClosed)
	}
	// Check for account currency
	if item.Currency != curr {
		return item, errors.New(err_CurrencyMismatch)
	}
	// Change account amount
	item.Amount 	+= amount
	item.Updated	= time.Now()
	// Save changed Account into store
	err = this.keeper.UpdateAccount(item)
	return item, err
}

// Withdraw funds to account amount
func (this *AccountManager)WithdrawFund(name, curr string, amount float32) (*model.Account, error) {
	// Get Account from store
	item, err := this.keeper.GetAccount(name)
	if err != nil {
		return item, err
	}
	// Check for account is closed
	if item.State == model.State_Closed {
		return item, errors.New(err_AccountIsClosed)
	}
	// Check for account currency
	if item.Currency != curr {
		return item, errors.New(err_CurrencyMismatch)
	}
	// Check for account balance
	if item.Amount < amount {
		return item, errors.New(err_AmountToBig)
	}
	// Change account amount
	item.Amount 	-= amount
	item.Updated	= time.Now()
	// Save changed Account into store
	err = this.keeper.UpdateAccount(item)
	return item, err
}

// Transfer amount from one account to other
func (this *AccountManager)TransferFund(account, target, currency string, amount float32) (*model.Account, error) {
	// Get source Account from store
	item1, err := this.keeper.GetAccount(account)
	if err != nil {
		return nil, err
	}
	// Check for account is closed
	if item1.State == model.State_Closed {
		return nil, errors.New(err_SourceIsClosed)
	}
	// Check for account currency
	if item1.Currency != currency {
		return nil, errors.New(err_CurrencyMismatch)
	}
	// Check for account balance
	if item1.Amount < amount {
		return nil, errors.New(err_AmountToBig)
	}

	// Get target Account from store
	item2, err := this.keeper.GetAccount(target)
	if err != nil {
		return nil, err
	}
	// Check for account is closed
	if item2.State == model.State_Closed {
		return nil, errors.New(err_TargetIsClosed)
	}
	// Check for account currency
	if item2.Currency != currency {
		return nil, errors.New(err_CurrencyMismatch)
	}

	// Change account amount
	item1.Amount -= amount
	item1.Updated	= time.Now()
	item2.Amount += amount
	item2.Updated	= time.Now()
	// Save changed Accounts into store
	err = this.keeper.UpdateAccount(item1)
	if err != nil {
		return nil, err
	}
	err = this.keeper.UpdateAccount(item2)
	if err != nil {
		return nil, err
	}
	return item1, err
}

