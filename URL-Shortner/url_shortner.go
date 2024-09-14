package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	Id           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func main() {

	// Register the handle function to handle all requessts to the root URL ("/")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shorten", shortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	// Start the HTTP server on port 3000
	fmt.Println("Starting server on port 3000....")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error on starting server:", err)
	}
}

func shortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	generateShortURL := createURL(data.URL)

	response := struct {
		ShortURL string `json:"short_url"`
	}{ShortURL: generateShortURL}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func createURL(originalURL string) string {

	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		Id:           id,
		OriginalURL:  originalURL,
		ShortURL:     shortURL,
		CreationDate: time.Now(),
	}

	return shortURL
}

func generateShortURL(originalURL string) string {

	hasher := md5.New()
	hasher.Write([]byte(originalURL)) // It converts the originalURL to a byte slice
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)

	return hash[:8]
}

func getURL(id string) (URL, error) {

	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found")
	}

	return url, nil
}
