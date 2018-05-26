package model

import "time"

// Enumeration of account state
const (
	State_Active    int = iota
	State_Closed
)
// Account object declaration
type Account struct {
	Name 		string		`json:"name"`		// Unique account ID number
	State 		int			`json:"state"`		// State of account
	Owner 		string		`json:"owner"`		// Name of account owner
	Currency	string		`json:"currency"`	// Currency ISO code
	Amount		float32		`json:"amount"`		// Amount of currency
	Created		time.Time	`json:"created"`	// Account creation time
	Updated		time.Time	`json:"updated"`	// Account last update time
}
// Account list declaration
type AccountList []*Account

// Interface to account storage
type AccountKeeper interface {
	InsertAccount(item *Account) error			// Put new account to the storage
	UpdateAccount(item *Account) error			// Set new account data by account name
	DeleteAccount(name string) error			// Remove account from the storage by name
	GetAccount(name string) (*Account, error)	// Get account from the storage by name
	GetAccountList() AccountList				// Get all accounts from the storage
}

// Placeholder for query params
type QueryParams struct {
	Account		*string		`json:"account"`	// Active account name
	Owner		*string		`json:"owner"`		// Name of account owner
	Currency	*string		`json:"currency"`	// Currency ISO code
	Amount		*float32	`json:"amount"`		// Transaction amount
	Target		*string		`json:"target"`		// Transfer to account
}
