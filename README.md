# account
Test Golang example

ORIGINAL REQUIREMENTS:

* code should be provided in a zip or in a public github repository
* code should be written in golang
* Makefile should be provided to build the solution

1. Create a REST API to be used as the backend for a bank website
2. API should support the following functions
    a. OPEN and CLOSE account
    b. WITHDRAW funds
    c. DEPOSIT funds
    d. TRANSFER funds


ADDITIONAL REQUIREMENTS:

* account name is unique string and not less then 4 symbol
* closed account must be stored as well as open one
* only empty account can be closed
* only closed account can be deleted
* accounts must have a currency
* transaction and account currency must be same


SOLUTION

Application contains several packages:
* model   - domain objects definition
* store   - in-memory storage for accounts
* domain  - business logic manager
* handler - dispatcher of API http requests
* front   - dispatcher of GUI http requests
* server  - http port listener


BUILDING

Use IDE or Makefile to build application


ACCOUNT OBJECT

Account object declaration
{
	"name":"A00123",		// Unique account ID number
	"state":0,				// State of account (0-active, 1-closed)
	"owner":"Ivan Rybakov",	// Name of account owner
	"currency":"UAH",		// Currency ISO code
	"amount":125.00,		// Amount of currency
	"created":"04/05/2018T18:00:00",	// Account creation time
	"updated":"06/05/2018T19:30:00"	// Account last update time
}


API DESCRIPTION

1. Get account list		- /api/account/list"
2. Get account item 	- /api/account/item"
3. Open new account		- /api/account/open"
2. Close account	 	- /api/account/close"
4. Delete account		- /api/account/delete"
5. Deposit transaction 	- /api/account/deposit"
6. Withdraw transaction	- /api/account/withdraw"
7. Transfer transaction	- /api/account/transfer"


API TESTING

Use curl commands from test_curl.txt file
Or use Web interface on http://localhost:8080/


Author: Ivan Rybakov

