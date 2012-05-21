package main

import (
        "github.com/hoisie/web"
        "github.com/abrookins/go_iching"
        "fmt"
        "encoding/json"
        "strconv"
)

func reading(ctx *web.Context) string {
        ctx.Request.ParseForm()
        question := ctx.Request.Form.Get("question")
        ctx.ContentType("json")

        if question == "" {
                question = "No question provided"
        }

        reading := iching.GetReading(question)
        js, _ := json.Marshal(reading)

        return string(js)
}

func hexagram(ctx *web.Context, num string) string {
        n, _ := strconv.Atoi(num)
        hexagram, found := iching.GetHexagramByNum(n)
        ctx.ContentType("json")

        if found {
                js, _ := json.Marshal(hexagram) 
                return string(js)
        }

        return "Hexagram not found"
}

func allHexagrams(ctx *web.Context) string {
        hexagrams := iching.GetAllHexagrams()
        ctx.ContentType("json")
        js, _ := json.Marshal(hexagrams) 
        return string(js)
}

func main() {
        fmt.Printf("Starting server ... ")
        web.Get("/iching/hexagrams/(.*)", hexagram)
        web.Get("/iching/hexagrams", allHexagrams)
        web.Get("/iching/reading", reading)
        web.Run("0.0.0.0:9999")
}