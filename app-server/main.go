package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("hello world")
	httpPort := os.Getenv("PORT")
	http.HandleFunc("/", RootPath)

	http.ListenAndServe(":"+httpPort, nil)
}

type Post struct {
	User           string
	Day            time.Time
	Interpretation string
}

func RootPath(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	_, _ = fmt.Fprint(w, "Hello from my ECS instance!")
}
