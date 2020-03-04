package main
import (
	"net/http"
	"fmt"
	"time"
	"html/template"
	"path"
 	"runtime"
)

//Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

//Go application entrypoint
func main() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("No caller information")
	  }
	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	
	templates := template.Must(template.ParseFiles(path.Dir(filename)+"/templates/welcome.html"))

	

	http.Handle("/static/", //final url can be anything
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static")))) //Go looks in the relative static directory first, then matches it to a
			
	http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

		//Takes the name from the URL query e.g ?name=Martin, will set welcome.Name = Martin.
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name;
		}
		 
		if err := templates.ExecuteTemplate(w, "welcome.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	
	fmt.Println(http.ListenAndServe(":8080", nil));
}