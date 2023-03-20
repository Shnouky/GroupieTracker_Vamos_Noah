package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Character struct {
	Name      string `json:"name"`
	ImageURL  string `json:"imageUrl"`
	Thumbnail string `json:"thumbnailUrl"`
}

type DisneyAPIResponse struct {
	Data []Character `json:"data"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://api.disneyapi.dev/characters")
		if err != nil {
			http.Error(w, "Error retrieving data", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Error reading response body", http.StatusInternalServerError)
			return
		}

		var apiResponse DisneyAPIResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			http.Error(w, "Error unmarshalling response body", http.StatusInternalServerError)
			return
		}

		// Create a slice of character names
		characterNames := make([]string, len(apiResponse.Data))
		for i, character := range apiResponse.Data {
			characterNames[i] = character.Name
		}

		// Define a simple template that displays the character names
		tmpl, err := template.New("index").Parse("<html><body><h1>Disney Characters</h1><ul>{{range .}}<li>{{.}}</li>{{end}}</ul></body></html>")
		if err != nil {
			http.Error(w, "Error creating template", http.StatusInternalServerError)
			return
		}

		// Execute the template with the character names and write the result to the response
		err = tmpl.Execute(w, characterNames)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	// Define a handler for serving a separate web page
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("about").Parse("<html><body><h1>About this page</h1><p>This is a simple web page served by a Go HTTP server.</p></body></html>")
		if err != nil {
			http.Error(w, "Error creating template", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	})

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
