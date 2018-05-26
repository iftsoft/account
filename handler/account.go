package handler

import (
	"net/http"
	"account/model"
	"account/store"
	"errors"
	"fmt"
	"strings"
	"time"
)

const (
	hnd_Item		= "item"
	hnd_List		= "list"
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
	if len(list) > 2 {
		this.command	= list[2]
		// Call command handler
		switch strings.ToLower(this.command) {
		case hnd_Item :
			this.err = this.cmdAccountGetItem(w, r)
		case hnd_List :
			this.err = this.cmdAccountGetList(w, r)
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
		fmt.Printf("Account command %s works %d mcs;\n", this.command, spend )
	} else {
		fmt.Printf("Account command %s works %d mcs; Return error: %v\n", this.command, spend, this.err)
	}
}


///////////////////////////////////////////////////////////////////////

func (this *hndAccountAPI)cmdAccountGetItem(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountGetList(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountOpen(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountClose(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountDelete(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountDeposit(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountWithdraw(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

func (this *hndAccountAPI)cmdAccountTransfer(w http.ResponseWriter, r *http.Request) (err error) {
	return err
}

//func (this *hndAccountAPI)basicSelect(w http.ResponseWriter, r *http.Request) (err error) {
//	input := &PrimaryKeys{}
//	err = this.readJsonBody(w, r, input)
//	if err != nil {	return err	}
//
//	if this.check != nil {
//		err = this.check.CheckSelectInput(input)
//		if err != nil {
//			this.WriteBadReply(w, http.StatusExpectationFailed, err.Error())
//			return err
//		}
//	}
//	dao := this.store.GetObjectKeeper(this.name)
//	if dao == nil {
//		this.WriteBadReply(w, http.StatusInternalServerError, hnd_ErrResolveDAO)
//		return errors.New(hnd_ErrResolveDAO)
//	}
//	keys := makeParamList(input)
//	err = dao.Select(this.unit, keys)
//	if err != nil {
//		this.WriteBadReply(w, http.StatusInternalServerError, err.Error())
//		return err
//	}
//	this.WriteJsonReply(w, this.unit)
//	return err
//}
//
//func (this *hndAccountAPI)basicUpdate(w http.ResponseWriter, r *http.Request) (err error) {
//	err = this.readJsonBody(w, r, this.unit)
//	if err != nil {	return err	}
//
//	if this.check != nil {
//		err = this.check.CheckModifyInput(this.unit, false)
//		if err != nil {
//			this.WriteBadReply(w, http.StatusExpectationFailed, err.Error())
//			return err
//		}
//	}
//	dao := this.store.GetObjectKeeper(this.name)
//	if dao == nil {
//		this.WriteBadReply(w, http.StatusInternalServerError, hnd_ErrResolveDAO)
//		return errors.New(hnd_ErrResolveDAO)
//	}
//	err = dao.Update(this.unit)
//	if err != nil {
//		this.WriteBadReply(w, http.StatusInternalServerError, err.Error())
//		return err
//	}
//	err = dao.Select(this.unit, lla.ParamList{})
//	if err != nil {
//		this.WriteBadReply(w, http.StatusInternalServerError, err.Error())
//		return err
//	}
//	this.WriteJsonReply(w, this.unit)
//	return err
//}


