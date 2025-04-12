package main

//import "fmt"
import "net/http"

func handler(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html");
}

func main() {
		/*http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, "Web serives are easy in Go!");
		})*/


		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, "home.html");
		})

		http.ListenAndServe(":3000", nil)
}
