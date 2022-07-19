package router

import (
	// Uncomment below imports as needed

	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codility/rest_api_go/database"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", UsersHandler).Methods("GET")
	return r
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users := database.GetUsers()
	name := r.URL.Query().Get("name")

	if name != "" {
		users = filterUsersByName(users, name)
	}

	usersJson, _ := json.Marshal(users)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", usersJson)
}

func filterUsersByName(users []database.User, name string) (filteredUsers []database.User) {
	for _, user := range users {
		if user.Name == name {
			filteredUsers = append(filteredUsers, user)
		}
	}
	return
}
