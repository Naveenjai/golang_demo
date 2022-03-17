package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var path = "C:/Users/naveen-pt4930/files/Naveenkumar_Registration_form.pdf"

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}
	return (err != nil)
}

func main() {
	log.SetFlags(log.Lshortfile)
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		modtime := time.Now()
		var file, err = os.OpenFile(path, os.O_RDWR, 0644)
		if isError(err) {
			return
		}
		defer file.Close()
		// ServeContent uses the name for mime detection
		const name = "random.pdf"

		// tell the browser the returned content should be downloaded
		w.Header().Add("Content-Disposition", "Attachment")

		http.ServeContent(w, req, name, modtime, file)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
