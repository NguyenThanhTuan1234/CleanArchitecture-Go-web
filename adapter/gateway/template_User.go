package gateway

import (
	"net/http"
)

func (t *templateClient) User(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "user.html", nil)
	if err != nil {
		return err
	}
	return nil
}
