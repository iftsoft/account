package front

import (
	"net/http"
	"account/store"
)

// Interface to Preview handler
func GetPreviewHandler() http.Handler {
	hnd := &guiPreviewHandler{}
	return hnd
}

// Preview handler definition
type guiPreviewHandler struct {
	guiBaseHandler
}

// Preview handler executor
func (this *guiPreviewHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	// Find account name
	qry := r.URL.Query()
	keys, ok := qry["account"]
	if !ok || len(keys) < 1 {
		http.Error(w, "No account specified", http.StatusBadRequest)
		return
	}
	// Get account from store
	keeper := store.GetAccountKeeper()
	item, err := keeper.GetAccount(keys[0])
	if err != nil {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	}
	// Prepare template data
	this.data = make(map[string]interface{})
	this.data["Account"]	= item
	this.tpls = []string { baseFrameTpl, previewContentTpl, previewScriptTpl, previewFormTpl, previewWorkTpl }
	this.execTemplates(w)
}

var previewContentTpl = `{{define "content"}}
{{template "preview" .Account}}<br>
{{template "transact" .Account}}
{{end}}`

var previewFormTpl = `{{define "preview"}}
<div class="panel panel-primary">
  <div class="panel-heading">Account property</div>
  <div class="panel-body">
<form  class="form-horizontal">` +
	GetStatic("Name", "Account name") +
	GetStatic("Owner", "Account owner") +
	GetStatic("State", "Account state") +
	GetStatic("Currency", "Account currency") +
	GetStatic("Amount", "Account amount") +
	GetStatic("Created", "Account created") +
	GetStatic("Updated", "Account updated") +
`</form></div>
<div class="panel-footer">` +
	GetButton("Accounts list", "gotoAccountList();", "btn-success") +
	GetButton("Close", "closeAccount();", "btn-warning") +
	GetButton("Delete", "deleteAccount();", "btn-danger") +
`</div>
</div> 
{{end}}`

var previewWorkTpl = `{{define "transact"}}
<div class="panel panel-primary">
  <div class="panel-heading">Transaction property</div>
  <div class="panel-body">
<form class="form">` +
	GetCurrency("currency", "Transaction currency") +
	GetInput("amount", "Transaction amount") +
	GetInput("target", "Target account") +
`</form></div>
<div class="panel-footer">` +
	GetButton("Deposit", "transDeposit();", "btn-info") +
	GetButton("Withdraw", "transWithdraw();", "btn-info") +
	GetButton("Transfer", "transTransfer();", "btn-info") +
`</div>
</div> 
{{end}}`

var previewScriptTpl = `{{define "scripts"}}
<script type="text/javascript">
{{if .Account}}
var accName = "{{.Account.Name}}";
{{end}}
function gotoAccountList() {
	window.location.href = "/acc_list.html";
}
function closeAccount() {
    event.preventDefault();
    var obj = {};
    obj.account  = accName;
    accountApiQuery("close", obj);
}
function deleteAccount() {
    event.preventDefault();
    var obj = {};
    obj.account  = accName;
    accountApiQuery("delete", obj);
}
function transDeposit() {
    event.preventDefault();
    var obj = {};
    obj.account  = accName;
    obj.amount   = parseFloat($("#amount")[0].value);
    obj.currency = $("#currency")[0].value;
    accountApiQuery("deposit", obj);
}
function transWithdraw() {
    event.preventDefault();
    var obj = {};
    obj.account  = accName;
    obj.amount   = parseFloat($("#amount")[0].value);
    obj.currency = $("#currency")[0].value;
    accountApiQuery("withdraw", obj);
}
function transTransfer() {
    event.preventDefault();
    var obj = {};
    obj.account  = accName;
    obj.amount   = parseFloat($("#amount")[0].value);
    obj.currency = $("#currency")[0].value;
    obj.target   = $("#target")[0].value;
    accountApiQuery("transfer", obj);
}

function accountApiQuery(cmd, obj) {
    var url = "/api/account/" + cmd;
    var body = JSON.stringify(obj);
    $.post(url, body, function(data, status){
        if( status === "success" ){
	        window.location.href = "/acc_list.html";
        }
    }, "json");
};
</script>
{{end}}`


