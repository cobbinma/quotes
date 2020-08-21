package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/cobbinma/kanye-west-api/models"
	"io/ioutil"
	"net/http"
)

func Kanye(client Client) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := client.Get("https://api.kanye.rest/")
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

		fmt.Println(string(body))

		var quote models.Quote
		err = json.Unmarshal(body, &quote)
		if err != nil {
			fmt.Println(err)
			_, _ = w.Write([]byte("could not marshall kanye"))
			return
		}

		_, _ = w.Write([]byte(quote.Quote))
	}
}
