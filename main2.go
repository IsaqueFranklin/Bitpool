 package main
import (
  "fmt"
  "html/template"
  "time"
  //"io"
  "log"
  "net/http"
)

type Film struct {
  Title       string
  Director    string
}

func main(){
  fmt.Println("Hello.")
  
  h1 := func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("index.html"))

    films := map[string][]Film{
     "Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},    
    }
    tmpl.Execute(w, films)
  }  

  h2 := func (w http.ResponseWriter, r *http.Request) {
    time.Sleep(1 * time.Second)
    title := r.PostFormValue("title") 
    director := r.PostFormValue("director")
    //htmlStr := fmt.Sprintf("<div class='border border-gray-700 rounded-xl p-4 mt-4'><p>%s - %s</p></div>", title, director)
    //tmpl, _ := template.New("t").Parse(htmlStr)
    //tmpl.Execute(w, nil)

    tmpl := template.Must(template.ParseFiles("index.html"))

    tmpl.ExecuteTemplate(w, "film-list-element", Film{Title: title, Director: director})
  }


  http.HandleFunc("/", h1)
  http.HandleFunc("/add-film/", h2)
  log.Fatal(http.ListenAndServe(":8000", nil))
}
