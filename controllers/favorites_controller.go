package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/beego/beego/v2/server/web"
)

type FavoritesController struct {
	web.Controller
}

type Favorite struct {
	ID        int    `json:"id"`
	UserID    string `json:"user_id"`
	ImageID   string `json:"image_id"`
	SubID     string `json:"sub_id"`
	CreatedAt string `json:"created_at"`
	Image     struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"image"`
}


func makeAPICalls(method, url, apiKey string) chan APIResponse {
	responseChan := make(chan APIResponse)

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			responseChan <- APIResponse{nil, err}
			return
		}

		req.Header.Add("x-api-key", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			responseChan <- APIResponse{nil, err}
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		responseChan <- APIResponse{body, err}
	}()

	return responseChan
}

func (c *FavoritesController) GetFavorites() {
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	
	responseChan := makeAPICalls("GET", url, apiKey)
	
	response := <-responseChan
	if response.Error != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error fetching favorites"}
		c.ServeJSON()
		return
	}

	var favorites []Favorite
	err = json.Unmarshal(response.Body, &favorites)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error decoding response: " + err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = favorites
	c.ServeJSON()
}