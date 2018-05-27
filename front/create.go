package front

import (
	"net/http"
)

// Interface to Create handler
func GetCreateHandler() http.Handler {
	hnd := &guiCreateHandler{}
	return hnd
}

// Create handler definition
type guiCreateHandler struct {
	guiBaseHandler
}

// Create handler executor
func (this *guiCreateHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	this.data = make(map[string]interface{})
	this.tpls = []string { baseFrameTpl, createContentTpl, createScriptTpl }
	this.execTemplates(w)
}

var createContentTpl = `{{define "content"}}
<form role="form">` +
	GetInput("account", "Account name") +
	GetInput("owner", "Account owner") +
	GetCurrency("currency", "Account currency") +
	GetButton("Create", "createAccount();", "btn-warning") +
`</form> 
{{end}}`

var createScriptTpl = `{{define "scripts"}}
<script type="text/javascript">
function createAccount() {
    event.preventDefault();
    var obj = {};
    obj.account  = $("#account")[0].value;
    obj.owner    = $("#owner")[0].value;
    obj.currency = $("#currency")[0].value;
    accountApiQuery("open", obj);
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
