package requests

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/neglarken/date_name_parser/model"
)

var (
	superTags string = `[
		{
			"tag":{
				"id":2,
				"name":"Клиент",
				"key":"client",
				"values_source":0
			},
			"value":"%s"
		}
	]`
	comment string = `{
		"indicator_to_mo_id":%s,
		"platform":"%s"}`
)

func (r *Requests) PostDataToMySQL(data *model.Event) error {
	form := url.Values{}
	form.Add("period_start", "2023-09-01")
	form.Add("period_end", "2023-09-30")
	form.Add("period_key", "month")
	form.Add("indicator_to_mo_id", "315914")
	form.Add("indicator_to_mo_fact_id", "0")
	form.Add("value", "1")
	form.Add("fact_time", "2023-08-11")
	form.Add("is_plan", "0")
	form.Add("supertags", superTags)
	form.Add("auth_user_id", "40")
	form.Add("comment", comment)

	for i := 0; i < data.RowsCount; i++ {
		r.Logger.Printf("starting sending request to %s\n", r.Cfg.PostMySqlURL)

		form.Set("supertags", fmt.Sprintf(superTags, data.Rows[i].Author.UserName))
		form.Set("fact_time", strings.Replace(data.Rows[i].Time[0:19], "T", " ", 1))
		form.Set("comment", fmt.Sprintf(
			comment,
			data.Rows[i].Params.IndicatorToMyId,
			data.Rows[i].Params.Platform,
		))
		form.Set("period_start", data.Rows[i].Period.Start)
		form.Set("period_end", data.Rows[i].Period.End)

		req, err := http.NewRequest("POST", r.Cfg.PostMySqlURL, strings.NewReader(form.Encode()))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		res, err := r.Cl.Do(req)
		if err != nil {
			return err
		}

		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		r.Logger.Printf("%s request to %s returned %s with code %s\n", req.Method, req.URL, string(resBody), res.Status)
	}

	return nil
}
