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

  app.Get("/blockinfo/:block", func(ctx *fiber.Ctx) error {
    block := ctx.Params("block")

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
    
    return ctx.Render("index", fiber.Map{
      "Height": result.Height,
      "Hash": result.Hash,
      "Timestamp": result.Timestamp,
    })
  })

  app.Post("/block", func(ctx *fiber.Ctx) error {
 
    block := ctx.FormValue("block")

    fmt.Println(block)
    
    return nil
  })

 app.Get("/get-block", func(ctx *fiber.Ctx) error {
    
    time.Sleep(1 *time.Second)
    block := ctx.FormValue("block")

    fmt.Println(block) 
    
    return ctx.Redirect("/blockinfo/"+block)
 }) 

  log.Fatal(app.Listen(":9000"))
}
