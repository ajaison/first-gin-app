package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprint(w, `
				<html>
					<body>
						<form method="post">
							<input type="text" id="textbox" name="textbox"/>
							<button type="submit" id="button">Submit</button>
						</form>
					</body>
				</html>
			`)
		} else if r.Method == http.MethodPost {
			r.ParseForm()
			text := r.FormValue("textbox")
			fmt.Fprintf(w, "You entered: %s", text)
		}
	}).Methods(http.MethodGet, http.MethodPost)

	http.ListenAndServe(":8080", r)
}
