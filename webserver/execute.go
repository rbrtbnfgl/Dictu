

package main

import (
	"fmt"
	"log"
	"net/http"
        "bytes"
        "os/exec"
	"io"
	"os"
)
func getRequest(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/styles.css" && r.URL.Path != "/index.html"{
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "/index.html")
		} else {
			http.ServeFile(w, r, r.URL.Path)
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		b, _ := io.ReadAll(r.Body)
		
		file, err := os.Create("/prog.du")
    		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
    		}
		file.Write(b)
		file.Close()
		cmd := exec.Command("dictu", "/prog.du")
		var out bytes.Buffer
                cmd.Stdout = &out
		err = cmd.Run()
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintf(w, out.String())
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", getRequest)

	fmt.Printf("Starting dictu server\n")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
