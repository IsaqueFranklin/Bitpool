package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "time"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
)

type Response struct {
  Height int `json:"height"`
  Hash string `json:"hash"`
  Timestamp string `json:"timestamp"`
}

type Price struct {
  Time int `json:"time"`
  USD int `json:USD`
  EUR int `json:EUR`
}

func main() {

  app := fiber.New(fiber.Config{
    Views: html.New("./views", ".html"),
  })

  app.Static("/", "./public", fiber.Static{
    Compress: true,
  }) 

  app.Get("/", func(ctx *fiber.Ctx) error {
    return ctx.Render("index", fiber.Map{})
  })
 
  app.Post("/block", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)
    block := ctx.FormValue("block")

    fmt.Println(block)
    
     resp, err := http.Get("https://mempool.space/api/v1/mining/blocks/timestamp/"+block)

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

    fmt.Println(result) 
    
    return ctx.Render("comps/block", fiber.Map{
      "Height": result.Height,
      "Hash": result.Hash,
      "Timestamp": result.Timestamp,
    })
  })

  app.Post("/price", func(ctx *fiber.Ctx) error {
    time.Sleep(1 *time.Second)
    
    resp, err := http.Get("https://mempool.space/api/v1/prices")

    if err != nil {
      log.Fatalln(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatalln(err)
    }

    var result Price 

    if err := json.Unmarshal(body, &result); err != nil {
      fmt.Println("Cannot unmarshal JSON.")
    }

    fmt.Println(result)

    return ctx.Render("comps/price", fiber.Map{
      "Time": result.Time,
      "USD": result.USD,
      "EUR": result.EUR,
    })
  }) 

  log.Fatal(app.Listen(":9000"))
}
