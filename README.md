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


API TESTING

Use curl commands from test_curl.txt file
Or use Web interface on http://localhost:8080/


