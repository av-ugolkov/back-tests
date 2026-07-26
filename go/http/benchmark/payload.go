package main

// User is the small JSON payload exchanged by every benchmark handler.
// ID is a string so that handlers can echo the path parameter back without
// an error branch on every framework.
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var sampleUser = User{
	ID:    "42",
	Name:  "Alexander",
	Email: "alexander@example.com",
}

var sampleUserJSON = []byte(`{"id":"42","name":"Alexander","email":"alexander@example.com"}`)
