package requests

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/neglarken/date_name_parser/model"
)

func (r *Requests) GetDataFromDB() (*model.Event, error) {
	r.Logger.Printf("starting GET request to %s\n", r.Cfg.GetArangoDbURL)

	jsonBody := []byte(`{
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
		"limit": 3
	}`)
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
