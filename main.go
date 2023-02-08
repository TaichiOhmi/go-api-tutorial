package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 構造体のフィールド名は入れ子の内外に関わらず大文字で始めないと認識されない。
type catch struct {
	FishName string `json:"fish_name"`
	Quantity int64  `json:"quantity"`
}

type fishingResult struct {
	ID      string  `json:"id"`
	Angler  string  `json:"angler"`
	Results []catch `json:"results"`
}

var fishingResults = []fishingResult{
	{ID: "1", Angler: "Taichi", Results: []catch{{FishName: "アジ", Quantity: 30}}},
	{ID: "2", Angler: "Ichita", Results: []catch{{FishName: "アジ", Quantity: 10}, {FishName: "タチウオ", Quantity: 10}}},
	{ID: "3", Angler: "Chita", Results: []catch{{FishName: "イサキ", Quantity: 20}, {FishName: "カサゴ", Quantity: 10}, {FishName: "マダイ", Quantity: 1}}},
}

func getFishingResults(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, fishingResults)
}

func fishingResultById(c *gin.Context) {
	id := c.Query("id")
	fRes, err := getFishingResultById(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "result not found."})
		return
	}
	c.IndentedJSON(http.StatusOK, fRes)
}

func getFishingResultById(id string) (*fishingResult, error) {
	for i, res := range fishingResults {
		if res.ID == id {
			return &fishingResults[i], nil
		}
	}
	return nil, errors.New("result not found")
}

func createFishingResult(c *gin.Context) {
	var newResult fishingResult
	if err := c.BindJSON(&newResult); err != nil {
		return
	}
	fishingResults = append(fishingResults, newResult)
	c.IndentedJSON(http.StatusCreated, newResult)
}

func main() {
	router := gin.Default()
	router.GET("/fishing-results", getFishingResults)
	router.GET("/fishing-result", fishingResultById)
	router.POST("/fishing-results", createFishingResult)
	router.Run("localhost:8080")
}
