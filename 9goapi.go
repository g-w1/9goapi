package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"os"
)

func main() {
	authkey, isthere := os.LookupEnv("9apiauth")
	if !isthere {
		log.Fatal("need 9apiauth environment variable to authenticate requests")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// make sure the auth matches up
		auth := r.FormValue("auth")
		if auth == "" {
			fmt.Fprintf(w, "need auth key in form\n")
			return
		}
		
		if auth != authkey {
			fmt.Fprintf(w, "auth key in form is wrong\n")
			return
		}

		// parse the file from the form
		f, _, err := r.FormFile("toexec")
		if err != nil {
			fmt.Fprintf(w, "FormFile() err: %v\n", err)
			return
		}

		dat, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Fprintf(w, "ReadAll() err: %v\n", err)
			return
		}

		// write the data to a temporary file
		err = ioutil.WriteFile("/tmp/toexec", dat, 0777)
		if err != nil {
			fmt.Fprintf(w, "WriteFile() err: %v\n", err)
			return
		}

		// run it
		cmd := exec.Command("/tmp/toexec")
		out, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(w, "cmd.Output() err: %v\n", err)
			return
		}

		// print the output
		fmt.Fprintf(w, "%s", out)
	})

	log.Print(http.ListenAndServe(":8080", nil))
}