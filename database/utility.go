package database

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) {
	if _, err := ioutil.ReadAll(r.Body); err != nil {
		if err := json.Unmarshal([]byte(err.Error()), x); err != nil {
			return
		}
	}
}
