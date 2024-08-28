package controllers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

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

func (c *VotingController) AddFavorite() {
	var req FavoriteRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	url := "https://api.thecatapi.com/v1/favourites"
	body, _ := json.Marshal(req)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
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