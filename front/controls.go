package front

// Get button html string
func GetButton(text, click, class string) string{
	str := 	`<button class="btn ` + class + `" onclick="` + click + `">` + text + `</button>`
	return str
}

// Get static html string
func GetStatic(name, text string) string{
	str := 	`<div class="form-group">
  <label class="control-label col-sm-3">` + text + `</label>
  <div class="col-sm-9"><p class="form-control-static">{{.` + name + `}}</p></div>
</div>
`
	return str
}

// Get input html string
func GetInput(name, text string) string{
	str := 	`<div class="form-group">
  <label for="` + name + `">` + text + `</label>
  <input type="text" class="form-control" id="` + name + `">
</div>
`
	return str
}

// Get currency select html string
func GetCurrency(name, text string) string{
	str := 	`<div class="form-group">
  <label for="` + name + `">` + text + `</label>
  <select  class="form-control" id="` + name + `">
    <option>USD</option>
    <option>EUR</option>
    <option>GBP</option>
    <option>UAH</option>
  </select>
</div>
`
	return str
}

