package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/beego/beego/v2/server/web" // Updated import for Beego v2
)

type BreedsController struct {
	web.Controller // Use Beego v2's Controller
}

type Breed struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	Temperament string `json:"temperament"`
}

type BreedImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (c *BreedsController) GetAllBreeds() {
	url := "https://api.thecatapi.com/v1/breeds"
	resp, err := http.Get(url)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch breeds"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to read response"}
		c.ServeJSON()
		return
	}

	var breeds []Breed
	err = json.Unmarshal(body, &breeds)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to parse breeds data"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = breeds
	c.ServeJSON()
}

func (c *BreedsController) GetBreedImages() {
	breedID := c.GetString("breed_id")
	if breedID == "" {
		c.Data["json"] = map[string]string{"error": "Breed ID is required"}
		c.ServeJSON()
		return
	}

	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_id=%s&limit=8", breedID)
	resp, err := http.Get(url)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch breed images"}
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to read response"}
		c.ServeJSON()
		return
	}

	var breedImages []BreedImage
	err = json.Unmarshal(body, &breedImages)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to parse breed images data"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = breedImages
	c.ServeJSON()
}
