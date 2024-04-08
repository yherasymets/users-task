package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// MockDB is a mock database implementation for testing
type MockDB struct{}

func (m *MockDB) getUsers() []User {
	return []User{
		{ID: 1, Name: "Alice", Role: "admin"},
		{ID: 2, Name: "Bob", Role: "developer"},
		{ID: 3, Name: "Sam", Role: "manager"},
	}
}

func TestGetUserHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "No query parameter",
			queryParam:     "",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Alice","role":"admin"},{"id":2,"name":"Bob","role":"developer"},{"id":3,"name":"Sam","role":"manager"}]`,
		},
		{
			name:           "Query parameter provided",
			queryParam:     "Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   `[{"id":1,"name":"Alice","role":"admin"}]`,
		},
		{
			name:           "User does not exist ",
			queryParam:     "Alex",
			expectedStatus: http.StatusOK,
			expectedBody:   `[]`,
		},
	}

	// Create a new instance of App with MockDB
	app := &App{db: &MockDB{}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/users?name="+tt.queryParam, nil)
			// Create response recorder to capture response
			w := httptest.NewRecorder()
			app.getUserHandler(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status code %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			body := w.Body.String()
			if body != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, body)
			}
		})
	}
}

func TestSearchByName(t *testing.T) {
	tests := []struct {
		name         string
		queryName    string
		users        []User
		expectedSize int
	}{
		{
			name:      "Name exists in the list",
			queryName: "Alice",
			users: []User{
				{ID: 1, Name: "Alice", Role: "admin"},
				{ID: 2, Name: "Bob", Role: "developer"},
				{ID: 3, Name: "Sam", Role: "manager"},
			},
			expectedSize: 1,
		},
		{
			name:      "Name does not exist in the list",
			queryName: "David",
			users: []User{
				{ID: 1, Name: "Alice", Role: "admin"},
				{ID: 2, Name: "Bob", Role: "developer"},
				{ID: 3, Name: "Sam", Role: "manager"},
			},
			expectedSize: 0,
		},
		{
			name:      "Incorrect name",
			queryName: "Jlkfn.sp",
			users: []User{
				{ID: 1, Name: "Alice", Role: "admin"},
				{ID: 2, Name: "Bob", Role: "developer"},
				{ID: 3, Name: "Sam", Role: "manager"},
			},
			expectedSize: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := searchByName(tt.queryName, tt.users)
			if len(result) != tt.expectedSize {
				t.Errorf("Expected search result size %d, got %d", tt.expectedSize, len(result))
			}
		})
	}
}

func TestSendJSON(t *testing.T) {
	tests := []struct {
		name           string
		statusCode     int
		data           interface{}
		expectedBody   string
		expectedHeader map[string]string
	}{
		{
			name:       "Test with valid data",
			statusCode: http.StatusOK,
			data: []User{
				{ID: 1, Name: "Alice", Role: "admin"},
				{ID: 2, Name: "Bob", Role: "developer"},
				{ID: 3, Name: "Sam", Role: "manager"},
			},
			expectedBody: `[{"id":1,"name":"Alice","role":"admin"},{"id":2,"name":"Bob","role":"developer"},{"id":3,"name":"Sam","role":"manager"}]`,
			expectedHeader: map[string]string{
				"Content-Type": "application/json",
			},
		},
		{
			name:         "Test with nil data",
			statusCode:   http.StatusInternalServerError,
			data:         []User{},
			expectedBody: `[]`,
			expectedHeader: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			err := sendJSON(w, tt.statusCode, tt.data)
			if err != nil {
				t.Errorf("sendJSON returned an error: %v", err)
			}
			contentType := w.Header().Get("Content-Type")
			if contentType != tt.expectedHeader["Content-Type"] {
				t.Errorf("Expected Content-Type header to be %q, got %q", tt.expectedHeader["Content-Type"], contentType)
			}
			resp := w.Result()
			if resp.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, resp.StatusCode)
			}
			if body := w.Body.String(); body != tt.expectedBody {
				t.Errorf("Expected body %q, got %q", tt.expectedBody, body)
			}
		})
	}
}
