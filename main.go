package main

import (
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Prompt string
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	res := Chat("gemma3:270m", string(bytes))
	fmt.Fprintln(w, res)
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.ListenAndServe("0.0.0.0:6969", nil)
}
