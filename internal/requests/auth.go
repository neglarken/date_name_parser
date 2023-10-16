package requests

import (
	"fmt"
	"net/http"
)

func (r *Requests) Auth(login, password string) (*http.Response, error) {
	r.Logger.Println("starting authentication")

	req, err := http.NewRequest("POST", fmt.Sprintf(r.Cfg.GetSessionURL, login, password), nil)
	if err != nil {
		return nil, err
	}

	res, err := r.Cl.Do(req)
	if err != nil {
		return nil, err
	}

	r.Logger.Printf("%s request to %s returned %s\nCookies: %v", req.Method, req.URL, res.Status, r.Cl.Jar)

	return res, nil
}
