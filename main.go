package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama Nur Laily Lathifa")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9008")
http.ListenAndServe(":9008", nil)

}
