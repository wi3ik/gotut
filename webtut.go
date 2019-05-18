package main

import ("fmt"
		"net/http"
		)
		//"encoding/json")
// log ~=  fmt.   !!!!


func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there</h1>\n")
}


type test_struct struct {
	TEST string
}


func post_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hi there -> post_handler</h1>\n")

	//fmt.Println(r)

	//r.ParseForm()
	fmt.Println(r.Form)

 	for key, value := range r.Form {
		fmt.Println("index: ", key, " value: ", value)
	}


	/*
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t.TEST)
	*/

}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/v2/pushes/", post_handler)
	http.ListenAndServe(":8000", nil)
}

/*
 curl -X POST -d "{test=value}" http://localhost:8000/v2/pushes/
 curl -X POST --data-binary "{test=value}" http://localhost:8000/v2/pushes/



 */

