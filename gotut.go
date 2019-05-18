package main

import ("fmt"
        "net/http"
        "time")


func index_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Whoa, Go is neat!\n")

  dt := time.Now()
  fmt.Fprintln(w, "Current date and time is: ", dt.String(), "\n")
}


func main() {
  http.HandleFunc("/", index_handler)
  http.ListenAndServe(":8000", nil)


}
