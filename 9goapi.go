package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// parse the file from the form?
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v\n", err)
			return
		}

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
		err = ioutil.WriteFile("/tmp/9goapi", dat, 0777)
		if err != nil {
			fmt.Fprintf(w, "WriteFile() err: %v\n", err)
			return
		}

		// run it
		cmd := exec.Command("/tmp/9goapi")
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