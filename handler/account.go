package handler

import (
	"net/http"
	"account/model"
	"account/store"
	"errors"
	"fmt"
	"strings"
	"time"
	"account/domain"
)

const (
	hnd_List		= "list"
	hnd_Item		= "item"
	hnd_Open		= "open"
	hnd_Close		= "close"
	hnd_Delete		= "delete"
	hnd_Deposit		= "deposit"
	hnd_Withdraw	= "withdraw"
	hnd_Transfer	= "transfer"
)

// Interface to Account API
func getAccountApiHandler() http.Handler {
	hnd := &hndAccountAPI{}
	// Get interface to account storage
	hnd.keeper	= store.GetAccountKeeper()
	return hnd
}

// Account API handler
type hndAccountAPI struct {
	baseHandler						// Common functions inheritance
	keeper	model.AccountKeeper		// Interface to account storage
	command	string					// Internal command name
	err		error					// Internal error
}

// Serve HTTP request
func (this *hndAccountAPI)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Save begin time
	start := time.Now()
	// Check request command
	list := strings.Split(r.URL.Path, "/")
	if len(list) > 3 {
		this.command	= list[3]
		// Call command handler
		switch strings.ToLower(this.command) {
		case hnd_List :
			this.err = this.cmdAccountGetList(w, r)
		case hnd_Item :
			this.err = this.cmdAccountGetItem(w, r)
		case hnd_Open :
			this.err = this.cmdAccountOpen(w, r)
		case hnd_Close :
			this.err = this.cmdAccountClose(w, r)
		case hnd_Delete :
			this.err = this.cmdAccountDelete(w, r)
		case hnd_Deposit :
			this.err = this.cmdAccountDeposit(w, r)
		case hnd_Withdraw :
			this.err = this.cmdAccountWithdraw(w, r)
		case hnd_Transfer :
			this.err = this.cmdAccountTransfer(w, r)
		default:
			this.err = errors.New("Undefined account command")
			http.Error(w, this.err.Error(), http.StatusNotFound )
		}
	} else {
		this.err = errors.New("Request format is incorrect")
		http.Error(w, this.err.Error(), http.StatusBadRequest )
	}
	// Calculate execution time
	delta := time.Now().Sub(start)
	spend := int(delta/time.Microsecond)
	// Print command result
	if this.err == nil	{
		fmt.Printf("Account command %s works %d mcs;\n\n", this.command, spend )
	} else {
		fmt.Printf("Account command %s works %d mcs; Return error: %v\n\n", this.command, spend, this.err)
	}
}


///////////////////////////////////////////////////////////////////////

// Return list of all accounts in storage
func (this *hndAccountAPI)cmdAccountGetList(w http.ResponseWriter, r *http.Request) (err error) {
	// Get Account list from storage
	mngr := domain.GetAccountManager()
	list, err := mngr.GetAccountList()
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusNotFound )
	}
	// Write command response
	err = this.WriteJsonReply(w, list)
	return err
}

// Return account object by given account name
func (this *hndAccountAPI)cmdAccountGetItem(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}
	// Get Account from storage
	mngr := domain.GetAccountManager()
	item, err := mngr.GetAccount(*input.Account)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusNotFound )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Create new account
func (this *hndAccountAPI)cmdAccountOpen(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}
	// Check Account currency
	if input.Currency == nil || len(*input.Currency) != 3 {
		return this.WriteError(w, "Account currency is incorrect", http.StatusNotAcceptable)
	}
	// Check Account owner
	if input.Owner == nil || *input.Owner == "" {
		return this.WriteError(w, "Account owner is not set", http.StatusNotAcceptable)
	}
	// Check Account currency
	if input.Currency == nil || len(*input.Currency) != 3 {
		return this.WriteError(w, "Account currency is incorrect", http.StatusNotAcceptable)
	}

	// Open new Account into store
	mngr := domain.GetAccountManager()
	item, err := mngr.OpenAccount(*input.Account, *input.Owner, *input.Currency)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Change Account state to CLOSED
func (this *hndAccountAPI)cmdAccountClose(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}

	// Close account in the storage
	mngr := domain.GetAccountManager()
	item, err := mngr.CloseAccount(*input.Account)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Delete closed account from store
func (this *hndAccountAPI)cmdAccountDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}

	// Delete Account in the storage
	mngr := domain.GetAccountManager()
	item, err := mngr.DeleteAccount(*input.Account)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Deposit funds to account amount
func (this *hndAccountAPI)cmdAccountDeposit(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction currency
	if input.Currency == nil || len(*input.Currency) != 3 {
		return this.WriteError(w, "Transaction currency is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction amount
	if input.Amount == nil || *input.Amount <= 0.0 {
		return this.WriteError(w, "Transaction amount is incorrect", http.StatusNotAcceptable)
	}

	// Deposit fund from store
	mngr := domain.GetAccountManager()
	item, err := mngr.DepositFund(*input.Account, *input.Currency, *input.Amount)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Withdraw funds from account amount
func (this *hndAccountAPI)cmdAccountWithdraw(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Account name is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction currency
	if input.Currency == nil || len(*input.Currency) != 3 {
		return this.WriteError(w, "Transaction currency is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction amount
	if input.Amount == nil || *input.Amount <= 0.0 {
		return this.WriteError(w, "Transaction amount is incorrect", http.StatusNotAcceptable)
	}

	// Withdraw fund from store
	mngr := domain.GetAccountManager()
	item, err := mngr.WithdrawFund(*input.Account, *input.Currency, *input.Amount)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}

// Transfer amount from one account to other
func (this *hndAccountAPI)cmdAccountTransfer(w http.ResponseWriter, r *http.Request) (err error) {
	// Read query params
	input := &model.QueryParams{}
	err = this.AcceptInputQuery(w, r, input)
	if err != nil {	return err	}
	// Check source Account name
	if input.Account == nil || len(*input.Account) < 4 {
		return this.WriteError(w, "Source account name is incorrect", http.StatusNotAcceptable)
	}
	// Check target Account name
	if input.Target == nil || len(*input.Target) < 4 {
		return this.WriteError(w, "Target account name is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction currency
	if input.Currency == nil || len(*input.Currency) != 3 {
		return this.WriteError(w, "Transaction currency is incorrect", http.StatusNotAcceptable)
	}
	// Check Transaction amount
	if input.Amount == nil || *input.Amount <= 0.0 {
		return this.WriteError(w, "Transaction amount is incorrect", http.StatusNotAcceptable)
	}

	// Withdraw fund from store
	mngr := domain.GetAccountManager()
	item, err := mngr.TransferFund(*input.Account, *input.Target, *input.Currency, *input.Amount)
	if err != nil {
		return this.WriteError(w, err.Error(), http.StatusConflict )
	}
	// Write command response
	err = this.WriteJsonReply(w, item)
	return err
}


