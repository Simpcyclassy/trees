package main

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
	"regexp"
)

const (
	port = "8000"
)

func favoriteTreeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html") // Set the header

	t, err := template.ParseFiles("./templates/favoriteTree.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n, err := template.ParseFiles("./templates/noTree.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	v, err := template.ParseFiles("./templates/validTree.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Bad request, please go to /", http.StatusBadRequest)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "404 Method not found.", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	favoriteTree := query.Get("favoriteTree")

	if favoriteTree == "" {
		err = n.Execute(w, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	if validName(favoriteTree) {
		err = t.Execute(w, favoriteTree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return	
	} else {
		err = v.Execute(w, favoriteTree)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return	
	}
}

func main() {
	http.HandleFunc("/", favoriteTreeHandler)

	fmt.Printf("Service started at %v", port)
	err := http.ListenAndServe("0.0.0.0:"+port, http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}

}

func validName (name string) bool {
	// Escape the special characters and numbers
	var re = regexp.MustCompile(`^[A-Za-z]+$`)
	fmt.Printf("This is the name: %v, this is the value: %v", name, re.MatchString(name))
	return re.MatchString(name)
}