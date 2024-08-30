// controllers/random_cat_controller.go

package controllers

import (
    "encoding/json"
    "net/http"
    "github.com/beego/beego/v2/server/web"
    "io/ioutil"
    "time"
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

type APIResult struct {
    CatImages []CatImage
    Error     error
}

// Get handles the GET request for fetching a random cat image
func (c *RandomCatController) Get() {
    resultChan := make(chan APIResult)
    go fetchCatImage(resultChan)

    select {
    case result := <-resultChan:
        if result.Error != nil {
            c.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
            c.Ctx.WriteString("Failed to fetch cat image: " + result.Error.Error())
            return
        }

        if len(result.CatImages) == 0 {
            c.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
            c.Ctx.WriteString("No cat images found")
            return
        }

        c.Ctx.Output.SetStatus(http.StatusOK)
        c.Data["json"] = result.CatImages[0]
        c.ServeJSON()

    case <-time.After(10 * time.Second):
        c.Ctx.ResponseWriter.WriteHeader(http.StatusGatewayTimeout)
        c.Ctx.WriteString("Request timed out")
    }
}

func fetchCatImage(resultChan chan<- APIResult) {
    apiURL := "https://api.thecatapi.com/v1/images/search"

    resp, err := http.Get(apiURL)
    if err != nil {
        resultChan <- APIResult{Error: err}
        return
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        resultChan <- APIResult{Error: err}
        return
    }

    var catImages []CatImage
    err = json.Unmarshal(body, &catImages)
    if err != nil {
        resultChan <- APIResult{Error: err}
        return
    }

    resultChan <- APIResult{CatImages: catImages}
}