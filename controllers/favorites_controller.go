package controllers

import (
	"encoding/json"
	"net/http"
	"io"

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

func (c *FavoritesController) GetFavorites() {
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error creating request"}
		c.ServeJSON()
		return
	}

	req.Header.Add("x-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error fetching favorites"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error reading response body"}
		c.ServeJSON()
		return
	}

	var favorites []Favorite
	err = json.Unmarshal(body, &favorites)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error decoding response: " + err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = favorites
	c.ServeJSON()
}