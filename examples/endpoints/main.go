package endpoints

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/lambovg/go-request-compose/examples/dataprovider"
)

func main() {
	db, err := sql.Open("", "")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	provider := dataprovider.NewUsersDatabase(db)
	mux := http.NewServeMux()
	mux.Handle("/users", ListUsers(provider))

	if err := http.ListenAndServe("", mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
