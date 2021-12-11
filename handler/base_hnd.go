package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const (
	hnd_Method_GET    = "GET"
	hnd_Method_POST   = "POST"
	hnd_Method_PUT    = "PUT"
	hnd_Method_DELETE = "DELETE"
	hnd_ErrBadMethod  = "Wrong request method"
)

///////////////////////////////////////////////////////////////////////
// Empty structure to hold common handler functions
type baseHandler struct {
}

// Marshal to JSON response unit and send it as http reply
func (this *baseHandler) WriteJsonReply(w http.ResponseWriter, unit interface{}) error {
	body, err := json.Marshal(unit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	fmt.Printf("Http reply - %s\n", string(body))
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

// Send http error to client
func (this *baseHandler) WriteError(w http.ResponseWriter, text string, code int) error {
	http.Error(w, text, code)
	return errors.New(text)
}

// Parse request query or body to input container
func (this *baseHandler) AcceptInputQuery(w http.ResponseWriter, r *http.Request, unit interface{}) (err error) {
	if r.Method == hnd_Method_GET || r.Method == hnd_Method_DELETE {
		err = this.readUrlQuery(w, r, unit)
	} else if r.Method == hnd_Method_POST || r.Method == hnd_Method_PUT {
		err = this.readJsonBody(w, r, unit)
	} else {
		http.Error(w, hnd_ErrBadMethod, http.StatusBadRequest)
		return errors.New(hnd_ErrBadMethod)
	}
	return err
}

// Read and parse request body as JSON string
func (this *baseHandler) readJsonBody(w http.ResponseWriter, r *http.Request, unit interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	fmt.Printf("Http json query - %s\n", string(body))
	if len(body) > 0 {
		err = json.Unmarshal(body, unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return err
		}
	}
	return nil
}

// Get and parse request query
func (this *baseHandler) readUrlQuery(w http.ResponseWriter, r *http.Request, unit interface{}) error {
	fmt.Printf("Http URL query - %+v\n", r.URL.RawQuery)
	//	Fill input structure using reflect library
	iterateInputContainer(reflect.ValueOf(unit), r.URL.Query())
	return nil
}

//	Fill input structure using reflect library
func iterateInputContainer(value reflect.Value, qry url.Values) {
	if value.Kind() == reflect.Ptr {
		iterateInputContainer(value.Elem(), qry)
	}
	if value.Kind() == reflect.Struct {
		t := value.Type()
		// Iterate through struct fields
		for i := 0; i < t.NumField(); i++ {
			field := value.Field(i)
			fld_t := t.Field(i)
			// Process included anonymous structure
			if fld_t.Anonymous && fld_t.Type.Kind() == reflect.Struct {
				iterateInputContainer(reflect.Indirect(field), qry)
				continue
			}
			// Find field key name
			tag := fld_t.Tag.Get("json")
			if tag == "" {
				tag = strings.ToLower(fld_t.Name)
			}
			// Check out key in query set
			keys, ok := qry[tag]
			if !ok || len(keys) < 1 {
				continue
			}
			// Get query field as string
			str := keys[0]
			// Find or create value placeholder it input structure
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem()))
			}
			val := reflect.Indirect(field)
			// Convert query field string to structure field value
			switch val.Interface().(type) {
			case string:
				val.SetString(str)
			case int, int8, int16, int32, int64:
				i, err := strconv.ParseInt(str, 10, 64)
				if err == nil {
					val.SetInt(i)
				}
			case uint, uint8, uint16, uint32, uint64:
				u, err := strconv.ParseUint(str, 10, 64)
				if err == nil {
					val.SetUint(u)
				}
			case float32, float64:
				f, err := strconv.ParseFloat(str, 64)
				if err == nil {
					val.SetFloat(f)
				}
			case bool:
				b, err := strconv.ParseBool(str)
				if err == nil {
					val.SetBool(b)
				}
			}
		}
	}
}
