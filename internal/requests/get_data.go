package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/neglarken/date_name_parser/model"
)

var reqBody = `{
	"filter": {
		"field": {
			"key": "type",
			"sign": "LIKE",
			"values": [
				"MATRIX_REQUEST"
			]
		}
	},
	"sort": {
		"fields": [
			"time"
		],
		"direction": "DESC"
	},
	"limit": %v
}`

func (r *Requests) GetDataFromDB(limit int) (*model.Event, error) {
	r.Logger.Printf("starting GET request to %s\n", r.Cfg.GetArangoDbURL)

	jsonBody := []byte(fmt.Sprintf(reqBody, limit))
	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("GET", r.Cfg.GetArangoDbURL, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := r.Cl.Do(req)
	if err != nil {
		return nil, err
	}

	var resBody *model.Event
	err = json.NewDecoder(res.Body).Decode(&resBody)

	r.Logger.Printf("%s request to %s returned %v with code %s\n", req.Method, req.URL, resBody, res.Status)
	return resBody, nil
}
