package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  
  //"time"
  "github.com/gofiber/template/html/v2"
  "github.com/gofiber/fiber/v2"
)

type Response struct {
  Height int `json:"height"`
  Hash string `json:"hash"`
  Timestamp string `json:"timestamp"`
}

type Binfo struct {
	Block string `json:"block"`	
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

  app.Post("/block", func (c *fiber.Ctx) error {

    newBlock := new(Binfo)	
    newBlock.Block = c.FormValue("block") 
    
		return c.Redirect("/block/"+newBlock.Block)

	  /*return c.Render("todo/create", fiber.Map{
		  "Page":          "Create Todo",
		  "FromProtected": fromProtected,
		  "UserID":        c.Locals("userId").(uint64),
		  "Username":      c.Locals("username").(string),
	  })*/
  })

  app.Get("/block/:block", func(c *fiber.Ctx) error {
   
    newBlock := new(Binfo)
    newBlock.Block = c.Params("block")
    resp, err := http.Get("https://mempool.space/api/v1/mining/blocks/timestamp/"+newBlock.Block)

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
