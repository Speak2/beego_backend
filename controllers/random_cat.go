// controllers/random_cat_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/beego/beego/v2/server/web"
    "io/ioutil"
)

type RandomCatController struct {
    web.Controller
}

type CatImage struct {
    ID     string `json:"id"`
    URL    string `json:"url"`
    Width  int    `json:"width"`
    Height int    `json:"height"`
}

// Get handles the GET request for fetching a random cat image
func (c *RandomCatController) Get() {
    apiURL := "https://api.thecatapi.com/v1/images/search"

    resp, err := http.Get(apiURL)
    if err != nil {
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
        c.Ctx.WriteString("Failed to fetch cat image")
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
        c.Ctx.WriteString("Failed to read response body")
        return
    }

    var catImages []CatImage
    err = json.Unmarshal(body, &catImages)
    if err != nil {
        c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
        c.Ctx.WriteString("Failed to parse JSON")
        return
    }

    if len(catImages) == 0 {
        c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
        c.Ctx.WriteString("No cat images found")
        return
    }

    c.Ctx.Output.SetStatus(http.StatusOK)
    c.Data["json"] = catImages[0]
    c.ServeJSON()
}
