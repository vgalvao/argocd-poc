package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Println("# Hello ArgoCD # Env_HELLO_WORLD = ", os.Getenv("HELLO_WORLD"))))		
	})
	http.ListenAndServe(":8181", nil)
}
