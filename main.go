package main

import (
	"NotionHasPlayedMe/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("PID: %d\n", os.Getpid())
	gin.SetMode(gin.ReleaseMode)

	var cache []model.ResultPage

	fmt.Println("Populating Cache")
	cache = TransformResponseToResultPage(QueryNotionDB())
	fmt.Println("Cache Populated")

	r := gin.Default()

	r.GET("notionDB", func(c *gin.Context) {
		c.JSON(200, cache)
	})

	r.GET("/refreshCache", func(c *gin.Context) {
		fmt.Println("Refreshing Cache")
		cache = TransformResponseToResultPage(QueryNotionDB())
		fmt.Println("Cache Refreshed")

		c.JSON(200, "refreshed")
	})

	err := r.Run()
	if err != nil {
		log.Fatal("Failed to run server")
	}
}

func QueryNotionDB() model.Response {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var notionAPIKey = os.Getenv("API_KEY")
	var bear = fmt.Sprintf("Bearer %s", notionAPIKey)

	var url = "https://api.notion.com/v1/databases/10ff6a0b63fd80a4be89e442149409db/query"
	var method = "POST"

	var client = &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte("{}")))
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	req.Header.Add("Authorization", bear)
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("Close: ", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
	}

	var result model.Response
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatal("Unmarshal: ", err)
	}

	return result
}

func TransformResponseToResultPage(response model.Response) []model.ResultPage {
	var resultPages []model.ResultPage

	for _, page := range response.Results {
		var resultPage model.ResultPage

		resultPage.Name = page.Properties.Name.Title[0].Text.Content

		if len(page.Properties.Picture.Files) > 0 {
			resultPage.Picture = page.Properties.Picture.Files[0].File.Url
		}

		resultPage.Price = float64(page.Properties.Price.Number)

		resultPage.Size = page.Properties.Size.Select.Name

		resultPage.Ingredients = ParseStringToIngredients(page.Properties.Ingredients.RichText[0].Text.Content)

		resultPages = append(resultPages, resultPage)
	}

	return resultPages
}

func ParseStringToIngredients(input string) []model.Ingredient {
	var result []model.Ingredient
	var ingredients = strings.Replace(input, "[", "", 1)
	ingredients = strings.Replace(ingredients, "]", "", 1)

	var splittedIngredients = strings.Split(ingredients, ",")

	for _, ingredient := range splittedIngredients {
		ingredient = strings.Replace(ingredient, "{", "", 1)
		ingredient = strings.Replace(ingredient, "}", "", 1)

		var elems = strings.Split(ingredient, ":")
		var name = elems[0]
		quantity, _ := strconv.Atoi(strings.TrimSpace(elems[1]))
		result = append(result, model.Ingredient{Name: name, Quantity: quantity})
	}
	return result
}
