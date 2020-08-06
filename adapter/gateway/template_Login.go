package gateway

import (
	"net/http"
)

func (t *templateClient) Login(w http.ResponseWriter, r *http.Request) error {
	err := t.tpl.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		return err
	}
	return nil
}
