package front

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"text/template"
	"time"
)

///////////////////////////////////////////////////////////////////////
// Base Front Handler
type guiBaseHandler struct {
	tpls []string               // list of template name
	data map[string]interface{} // data for templates
}

// Default executor
func (this *guiBaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request path not found %s\n", r.URL.Path)
	http.Error(w, "Not found", http.StatusNotFound)
}

// Prepare and execute all templates on the page
func (this *guiBaseHandler) execTemplates(rw http.ResponseWriter) {
	tmpl := template.New("HtmlPage")
	for _, tpl := range this.tpls {
		tmpl = template.Must(tmpl.Parse(tpl))
	}
	var s string
	buf := bytes.NewBufferString(s)
	err := tmpl.Execute(buf, this.data)
	if err != nil {
		fmt.Printf("Front Error %s\n", err.Error())
	}
	rw.Write(buf.Bytes())
	//	fmt.Printf("Front Body %s", buf.String())
}

///////////////////////////////////////////////////////////////////////

func rootPanicRecover(w http.ResponseWriter) {
	if r := recover(); r != nil {
		text := fmt.Sprintf("Panic is recovered: %+v", r)
		fmt.Println(text)
		http.Error(w, text, http.StatusInternalServerError)
	}
}

// Root Web GUI handler
func GuiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get root template for Front GUI
		hnd := &rootGuiHandler{}
		// Serve HTTP request
		hnd.ServeHTTP(w, r)
	})
}

// Root handler description
type rootGuiHandler struct {
	guiBaseHandler
}

// Root handler executor
func (this *rootGuiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer rootPanicRecover(w)
	// Ignore favorite ico
	if r.URL.Path == "/favicon.ico" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	// Save begin time
	start := time.Now()
	fmt.Printf("Admin GUI %s\n", r.URL.Path)
	// Dispatch GUI pages
	this.DispatchRequest(w, r)
	// Calculate execution time
	delta := time.Now().Sub(start)
	spend := int(delta / time.Microsecond)
	// Print execution result
	fmt.Printf("Front GUI %s works %d mcs;\n\n", r.URL.Path, spend)
}

func (this *rootGuiHandler) DispatchRequest(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	if url == "/" {
		url = "/index.html"
	}
	list := strings.Split(url, "/")
	if len(list) >= 2 {
		switch list[1] {
		case "index.html":
			hnd := GetIndexHandler()
			hnd.ServeHTTP(w, r)
		case "acc_list.html":
			hnd := GetAccListHandler()
			hnd.ServeHTTP(w, r)
		case "preview.html":
			hnd := GetPreviewHandler()
			hnd.ServeHTTP(w, r)
		case "create.html":
			hnd := GetCreateHandler()
			hnd.ServeHTTP(w, r)
		default:
			hnd := GetIndexHandler()
			hnd.ServeHTTP(w, r)
		}
	} else {
		hnd := GetIndexHandler()
		hnd.ServeHTTP(w, r)
	}
}

// Default script template
var defaultScriptTpl = `{{define "scripts"}}
<script type="text/javascript">
</script>
{{end}}`

// The base FrameTemplate
var baseFrameTpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
<link href="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/css/bootstrap.min.css" rel="stylesheet">
<title>Bank Accounting System</title>
</head>
<body>
<div class="container">
<h1>Bank Accounting System</h1>
{{template "content" .}}
</div>
<script src="//code.jquery.com/jquery-1.11.1.min.js"></script>
<script src="//maxcdn.bootstrapcdn.com/bootstrap/3.2.0/js/bootstrap.min.js"></script>
{{template "scripts" .}}
</body>
</html>`
