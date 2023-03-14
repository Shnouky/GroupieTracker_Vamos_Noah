package main

import (
    "fmt"
    "net/http"
	"GroupieTracker_Vamos_Noah/controllers/"
)

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
	controllers.Carte1()
	r := gin.Default()
	r.GET("/carte", controllers.FindCarte)
	r.Run()
}


func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bienvenue sur mon site web en Golang !")
}

