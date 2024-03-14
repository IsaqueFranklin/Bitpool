 package main
import (
  "fmt"
  "html/template"
  "time"
  "encoding/json"
  "io/ioutil"
  //"io"
  "log"
  "net/http"
)

type Response struct {
  Height int `json:"height"`
  Hash string `json:"hash"`
  Timestamp string `json:"timestamp"`
}

func main(){
  fmt.Println("Hello.")
  
  h1 := func(w http.ResponseWriter, r *http.Request){
    tmpl := template.Must(template.ParseFiles("index.html"))

    tmpl.Execute(w, nil)
  }  

  h2 := func (w http.ResponseWriter, r *http.Request) {
    time.Sleep(1 * time.Second)
    
    //block := r.PostFormValue("block") 
    resp, err := http.Get("https://mempool.space/api/v1/mining/blocks/timestamp/1")
    
    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result Response

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Can not unmarshal JSON.")
    }

    //htmlStr := fmt.Sprintf("<div class='border border-gray-700 rounded-xl p-4 mt-4'><p>%s - %s</p></div>", title, director)
    //tmpl, _ := template.New("t").Parse(htmlStr)
    //tmpl.Execute(w, nil)

    tmpl := template.Must(template.ParseFiles("index.html"))

    tmpl.ExecuteTemplate(w, "get-block", Response{Height: result.Height, Hash: result.Hash, Timestamp: result.Timestamp})
  }


  http.HandleFunc("/", h1)
  http.HandleFunc("/get-block/", h2)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
