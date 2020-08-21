package handlers

import (
	"encoding/json"
	"github.com/cobbinma/kanye-west-api/models"
	"io/ioutil"
	"net/http"
)

func Kanye(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.kanye.rest/")
	if err != nil {
		_, _ = w.Write([]byte("error getting kanye quote"))
		return
	}

	if resp == nil || resp.StatusCode != http.StatusOK {
		_, _ = w.Write([]byte("bad kanye response"))
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		_, _ = w.Write([]byte("bad kanye body"))
		return
	}

	var quote models.Quote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		_, _ = w.Write([]byte("could not marshall kanye"))
		return
	}

	_, _ = w.Write([]byte(quote.Quote))
}
