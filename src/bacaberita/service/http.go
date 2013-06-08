package service

import (
	"fmt"
	"net/http"

	"appengine"

	"bacaberita/data"
)

func init() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/update", update)
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Use POST to register a new feed", http.StatusMethodNotAllowed)
		return
	}

	// FIXME figure out how to use PostFormValue instead of FormValue
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, "Feed URL is not specified", http.StatusBadRequest)
		return
	}

	c := appengine.NewContext(r)
	key, feed, err := data.RegisterFeed(c, url)

	fmt.Fprintf(w, "Register:\n")
	fmt.Fprintf(w, "- Url: %s\n", url)

	if err != nil {
		fmt.Fprintf(w, "- Error: %w\n", err)
	} else {
		fmt.Fprintf(w, "- Key: %w\n", key)
		fmt.Fprintf(w, "- Created: %w\n", feed.Created)
		fmt.Fprintf(w, "- Updated: %w\n", feed.Updated)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Use POST to update a feed", http.StatusMethodNotAllowed)
		return
	}

	// FIXME figure out how to use PostFormValue instead of FormValue
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, "Feed URL is not specified", http.StatusBadRequest)
		return
	}

	c := appengine.NewContext(r)
	feed, err := data.GetFeed(c, url)
	if err != nil {
		http.Error(w, "Feed URL is is not registered", http.StatusNotFound)
		return
	}

	key, err := data.UpdateFeed(c, feed)

	fmt.Fprintf(w, "Update:\n")
	fmt.Fprintf(w, "- Url: %s\n", url)

	if err != nil {
		fmt.Fprintf(w, "- Error: %w\n", err)
	} else {
		fmt.Fprintf(w, "- Key: %w\n", key)
		fmt.Fprintf(w, "- Created: %w\n", feed.Created)
		fmt.Fprintf(w, "- Updated: %w\n", feed.Updated)
	}
}
