package gateway

import (
	"net/http"
)

func (t *templateClient) User(w http.ResponseWriter, r *http.Request, xs string) error {
	err := t.tpl.ExecuteTemplate(w, "user.html", xs)
	if err != nil {
		return err
	}
	return nil
}
