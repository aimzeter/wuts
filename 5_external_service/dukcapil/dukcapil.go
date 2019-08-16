package dukcapil

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/aimzeter/wuts/5_external_service"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// interface validator
var cs zone.CitizenStore = &Dukcapil{}

type Dukcapil struct {
	client HTTPClient
	host   string
}

func NewDukcapil(c HTTPClient, host string) *Dukcapil {
	return &Dukcapil{c, host}
}

func (d *Dukcapil) get(path string, dest interface{}) error {
	path = d.host + path

	req, err := http.NewRequest(http.MethodGet, path, strings.NewReader(""))
	if err != nil {
		m := fmt.Sprintf("build new request got '%s'", err.Error())
		return errors.New(m)
	}

	resp, err := d.client.Do(req)
	if err != nil {
		m := fmt.Sprintf("call %s got '%s'", req.URL.String(), err.Error())
		return errors.New(m)
	}

	if resp.StatusCode != 200 {
		return errors.New("got non ok status code")
	}

	err = json.NewDecoder(resp.Body).Decode(dest)
	if err != nil {
		m := fmt.Sprintf("decode response body got '%s'", err.Error())
		return errors.New(m)
	}

	return nil
}
