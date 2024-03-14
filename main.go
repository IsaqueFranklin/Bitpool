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

app.Get("/", func(c *fiber.Ctx) error {
   

    return c.Render("index", fiber.Map{})
  })


  app.Post("/get-block/", func(c *fiber.Ctx) error {
    time.Sleep(1 *time.Second)
    block := c.FormValue("Block")

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

    if err := c.BodyParser(&result); err != nil {
      return err
    }
    
    return c.Render("index", fiber.Map{
      "Height": result.Height,
      "Hash": result.Hash,
      "Timestamp": result.Timestamp,
    })
  })

  log.Fatal(app.Listen(":9000"))
}
