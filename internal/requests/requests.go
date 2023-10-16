package requests

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"github.com/neglarken/date_name_parser/config"
)

type Requests struct {
	Cfg    *config.Config
	Cl     *http.Client
	Logger *log.Logger
}

func NewRequests(config *config.Config) (*Requests, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	return &Requests{
		Cfg: config,
		Cl: &http.Client{
			Jar: jar,
		},
		Logger: log.New(os.Stdout, "INFO\t", log.LUTC),
	}, nil
}
