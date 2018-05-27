package front

import (
	"net/http"
)

// Interface to Index handler
func GetIndexHandler() http.Handler {
	hnd := &guiIndexHandler{}
	return hnd
}

// Index handler definition
type guiIndexHandler struct {
	guiBaseHandler
}

// Index handler executor
func (this *guiIndexHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	this.data = make(map[string]interface{})
	this.tpls = []string { baseFrameTpl, indexContentTpl, indexScriptTpl }
	this.execTemplates(w)
}

var indexContentTpl = `{{define "content"}}
<p>You may show list of all accounts:</p>` +
	GetButton("Accounts list", "gotoAccountList();", "btn-success") +
`<p>Or create new account:</p>` +
	GetButton("Create Account", "gotoCreateAccount();", "btn-warning") +
`{{end}}`

var indexScriptTpl = `{{define "scripts"}}
<script type="text/javascript">
function gotoAccountList() {
	window.location.href = "/acc_list.html";
}
function gotoCreateAccount() {
	window.location.href = "/create.html";
}
</script>
{{end}}`

