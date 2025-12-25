package main

//import "fmt"
import "net/http"
//import "encoding/json"
import "log"

func main() {
	//fmt.Println("Hello, World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(w, r);
	})

	log.Println("Server started on :8085")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

