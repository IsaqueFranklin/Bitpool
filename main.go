package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
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
    resp, err := http.Get("https://mempool.space/api/v1/mining/blocks/timestamp/1672531200")

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

    return c.Render("index", fiber.Map{
      "Height": result.Height,
      "Hash": result.Hash,
      "Timestamp": result.Timestamp,
    })
  })

  log.Fatal(app.Listen(":9000"))
}
