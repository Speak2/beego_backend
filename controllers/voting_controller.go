package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http" 
	"github.com/beego/beego/v2/server/web"
	"github.com/astaxie/beego/logs"
)

type VotingController struct {
	web.Controller
}

type VoteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
	Value   int    `json:"value"`
}

type FavoriteRequest struct {
	ImageID string `json:"image_id"`
	SubID   string `json:"sub_id"`
}

func (c *VotingController) AddFavorite() {
	logs.Info("AddFavorite function called")

	// Read the request body
	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		logs.Error("Failed to read request body:", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Failed to read request body"}
		c.ServeJSON()
		return
	}

	// Log the received body for debugging
	logs.Info("Received body:", string(body))

	var req FavoriteRequest
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error("Failed to unmarshal request body:", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	logs.Info("Parsed request:", req)

	url := "https://api.thecatapi.com/v1/favourites"
	apiKey,err := web.AppConfig.String("cat_api_key") // Make sure to add this to your app.conf
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	requestBody, _ := json.Marshal(req)
	client := &http.Client{}
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		logs.Error("Failed to create request:", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to create request"}
		c.ServeJSON()
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", apiKey)

	resp, err := client.Do(httpReq)
	if err != nil {
		logs.Error("Failed to add favorite:", err)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to add favorite"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	c.Data["json"] = json.RawMessage(responseBody)
	c.ServeJSON()
}

func (c *VotingController) Vote() {
	var req VoteRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/votes"
	body, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to submit vote"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	responseBody, _ := ioutil.ReadAll(resp.Body)
	c.Data["json"] = json.RawMessage(responseBody)
	c.ServeJSON()
}