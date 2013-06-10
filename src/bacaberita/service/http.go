package service

import (
	"fmt"
	"net/http"

	"appengine"
	"appengine/user"

	"bacaberita/data"
)

func init() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/update", update)
	http.HandleFunc("/subscribe", subscribe)
	http.HandleFunc("/", main)
}

func main(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	if u == nil {
		url, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(c, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)<br/>`, u, url)

	fmt.Fprintf(w, `<form method="post" action="/register">Register <input type="text" name="url"/><input type="submit"/></form>`)
	fmt.Fprintf(w, `<form method="post" action="/update">Update <input type="text" name="url"/><input type="submit"/></form>`)
	fmt.Fprintf(w, `<form method="post" action="/subscribe">Subscribe <input type="text" name="url"/><input type="submit"/></form>`)
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
	key, feed, err := data.GetFeed(c, url)
	if err != nil {
		http.Error(w, "Feed URL is is not registered", http.StatusNotFound)
		return
	}

	key, err = data.UpdateFeed(c, feed)

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

func subscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Use POST to subscribe to a feed", http.StatusMethodNotAllowed)
		return
	}

	// FIXME figure out how to use PostFormValue instead of FormValue
	url := r.FormValue("url")
	if len(url) == 0 {
		http.Error(w, "Feed URL is not specified", http.StatusBadRequest)
		return
	}

	c := appengine.NewContext(r)
	key, _, err := data.RegisterFeed(c, url)
	if err != nil {
		c.Errorf("Error when registering feed: %w", err)
		http.Error(w, "Error when registering feed", http.StatusInternalServerError)
		return
	}

	u := user.Current(c)
	key, s, err := data.SubscribeFeed(c, u, key)

	fmt.Fprintf(w, "Subscribe:\n")
	fmt.Fprintf(w, "- Url: %s\n", url)
	fmt.Fprintf(w, "- UserID: %s\n", u.ID)
	fmt.Fprintf(w, "- Email: %s\n", u.Email)

	if err != nil {
		fmt.Fprintf(w, "- Error: %w\n", err)
	} else {
		fmt.Fprintf(w, "- Key: %w\n", key)
		fmt.Fprintf(w, "- Created: %w\n", s.Created)
		fmt.Fprintf(w, "- Updated: %w\n", s.Updated)
		fmt.Fprintf(w, "- Feeds: %d\n", len(s.FeedList))
	}
}
