package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/lambovg/go-request-compose/examples/dataprovider"
)

func ListUsers(provider dataprovider.Users) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		users, err := provider.ListUsers()
		if err != nil {
			http.Error(rw, "Cannot fetch users", http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(rw).Encode(users); err != nil {
			http.Error(rw, "Cannot encode user response", http.StatusInternalServerError)
			return
		}
	})
}
