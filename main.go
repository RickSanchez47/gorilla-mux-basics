package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	go func() {
		fmt.Println("Server listening on port 80...")
		if err := http.ListenAndServe(":80", r); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	}()

	// Prompt message after localhost started
	fmt.Println("Localhost server started successfully!")

	// Prevent the main goroutine from exiting
	select {}
}
