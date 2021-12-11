package front

import (
	"account/store"
	"net/http"
)

// Interface to Account list handler
func GetAccListHandler() http.Handler {
	hnd := &guiAccListHandler{}
	return hnd
}

// Account list handler definition
type guiAccListHandler struct {
	guiBaseHandler
}

// Account list handler executor
func (this *guiAccListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get account list from store
	keeper := store.GetAccountKeeper()
	list := keeper.GetAccountList()
	// Prepare template data
	this.data = make(map[string]interface{})
	this.data["lines"] = list
	this.tpls = []string{baseFrameTpl, listContentTpl, listScriptTpl}
	this.execTemplates(w)
}

var listContentTpl = `{{define "content"}}
<table class="table table-striped">
<thead><tr><th>Account</th><th>Owner</th><th>Amount</th><th>Curr</th><th>State</th></tr></thead>
{{if .lines}}<tbody>{{range .lines}}<tr onclick='gotoAccountForm("{{.Name}}");'>
  <td>{{.Name}}</td><td>{{.Owner}}</td><td>{{.Amount}}</td><td>{{.Currency}}</td><td>{{.State}}</td>
</tr>{{end}}</tbody>{{end}}
</table>
<p>Click on account row or create new account</p>` +
	GetButton("Create New Account", "gotoCreateAccount();", "btn-warning") +
	`{{end}}`

var listScriptTpl = `{{define "scripts"}}
<script type="text/javascript">
function gotoAccountForm(name) {
    console.log(name);
	window.location.href = "/preview.html/?account=" + name;
}
function gotoCreateAccount() {
	window.location.href = "/create.html";
}
</script>
{{end}}`
