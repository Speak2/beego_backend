package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/server/web"
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

type APIResponse struct {
	Body  []byte
	Error error
}

func makeAPICall(method, url string, body []byte, apiKey string) chan APIResponse {
	responseChan := make(chan APIResponse)

	go func() {
		client := &http.Client{}
		req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("failed to create request: %v", err)}
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("x-api-key", apiKey)

		resp, err := client.Do(req)
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("failed to make API call: %v", err)}
			return
		}
		defer resp.Body.Close()

		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			responseChan <- APIResponse{nil, fmt.Errorf("failed to read response body: %v", err)}
			return
		}

		responseChan <- APIResponse{responseBody, nil}
	}()

	return responseChan
}

func (c *VotingController) AddFavorite() {

	body, err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		logs.Error("Failed to read request body:", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Failed to read request body"}
		c.ServeJSON()
		return
	}

	var req FavoriteRequest
	if err := json.Unmarshal(body, &req); err != nil {
		logs.Error("Failed to unmarshal request body:", err)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	requestBody, _ := json.Marshal(req)
	responseChan := makeAPICall("POST", url, requestBody, apiKey)

	response := <-responseChan
	if response.Error != nil {
		logs.Error("Failed to add favorite:", response.Error)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to add favorite"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = json.RawMessage(response.Body)
	c.ServeJSON()
}

func (c *VotingController) DeleteFavorite() {
	favoriteID := c.Ctx.Input.Param(":favoriteId")

	url := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s", favoriteID)
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	responseChan := makeAPICall("DELETE", url, nil, apiKey)

	response := <-responseChan
	if response.Error != nil {
		logs.Error("Failed to delete favorite:", response.Error)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to delete favorite"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"message": "Favorite deleted successfully"}
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
	apiKey, err := web.AppConfig.String("cat_api_key")
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to get API key"}
		c.ServeJSON()
		return
	}

	body, _ := json.Marshal(req)
	responseChan := makeAPICall("POST", url, body, apiKey)

	response := <-responseChan
	if response.Error != nil {
		logs.Error("Failed to submit vote:", response.Error)
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = map[string]string{"error": "Failed to submit vote"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = json.RawMessage(response.Body)
	c.ServeJSON()
}
