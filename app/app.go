package app

import (
	"encoding/json"
	"net/http"
)

// App struct contains a MockDB instance
type App struct {
	db mockDB
}

// NewApp initializes a new instance of App
func NewApp() *App {
	return &App{db: &User{}}
}

// Router returns a new HTTP request multiplexer configured with routes for the application
func (a *App) Router() *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc("GET /users", a.getUserHandler)
	return r
}

// getUserHandler handles GET requests to the "/users" route
func (a *App) getUserHandler(w http.ResponseWriter, r *http.Request) {
	users := a.db.getUsers()
	name := r.URL.Query().Get("name")
	// If the "name" query parameter is provided, filter users by name
	if name != "" {
		result := searchByName(name, users)
		if err := sendJSON(w, http.StatusOK, result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	// If no "name" query parameter provided, return all users
	if err := sendJSON(w, http.StatusOK, users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// searchByName filters users by name
func searchByName(name string, users []User) []User {
	result := make([]User, 0)
	for _, user := range users {
		if user.Name == name {
			result = append(result, user)
		}
	}
	return result
}

// sendJSON encodes data as JSON and writes it to the response writer
func sendJSON(w http.ResponseWriter, statusCode int, data any) error {
	json, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if _, err := w.Write(json); err != nil {
		return err
	}
	return nil
}
