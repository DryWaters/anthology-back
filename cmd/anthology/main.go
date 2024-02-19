package anthology

import (
	"flag"
	"net/http"
)

type User struct{}
type Account struct{}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = user
}
func handleGetAccount(w http.ResponseWriter, r *http.Request) {
	var acc Account
	_ = acc
}

func main() {
	// get config

	// connect to DB

	// register handlers
	listenAddr := flag.String("listenaddr", ":49999", "todo")
	flag.Parse()

	http.HandleFunc("/user", handleGetUser)
	http.HandleFunc("/account", handleGetAccount)

	http.ListenAndServe(*listenAddr, nil)
}
