package domain

import (
	"testing"
	"account/store"
	"account/model"
	"time"
)

func TestGetAccount(t *testing.T) {
	account := &model.Account{ Name:"Key_1", Owner:"Owner 1", Currency:"USD", Amount:25.0, State:0, Created:time.Now(), Updated:time.Now()}
	keeper := store.GetAccountKeeper()
	mngr := GetAccountManager()

	_, err := mngr.GetAccount("Key_1")
	if err == nil {
		t.Error("Expected Error for GetAccount: ", "Key_1")
	}
	keeper.InsertAccount(account)
	item, err := mngr.GetAccount("Key_1")
	if err != nil {
		t.Error("Unexpected Error for GetAccount: ", err.Error())
	}
	if item.Name != "Key_1" {
		t.Error("Unexpected mismatch for item.Name: ", item.Name)
	}
	if item.Owner != "Owner 1" {
		t.Error("Unexpected mismatch for item.Owner: ", item.Owner)
	}
	if item.Currency != "USD" {
		t.Error("Unexpected mismatch for item.Currency: ", item.Currency)
	}
	if item.Amount != 25.0 {
		t.Error("Unexpected mismatch for item.Amount: ", item.Amount)
	}
}