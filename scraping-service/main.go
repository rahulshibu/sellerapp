package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	mongo "github.com/rahulshibu/sellerapp/database"
)

type (
	scrapData struct {
		Product string `json:"product"  bson:"product"`
		Reviews string `json:"reviews" bson:"reviews"`
		Rating  string `json:"rating" bson:"rating"`
		Price   string `json:"price" bson:"price"`
		Image   string `json:"image" bson:"image"`
	}
	postScrap struct {
		URL string `json:"url" valid:"required,url"`
	}
	saveScrap struct {
		URL       string      `bson:"url"`
		Data      []scrapData `bson:"data"`
		CreatedAt time.Time   `bson:"created_at"`
	}
)

var (
	err              error
	postRequestScrap postScrap
	scrap            scrapData
	scrapArr         []scrapData
)

func main() {
	//gin router with default middlewares
	router := gin.Default()
	fmt.Println("scrapping service")
	//Establishing mongo connection. If the connection is not correct, will throw error
	db := mongo.GetSharedConnection()
	//POST api to scrap data
	router.POST("/scrap", func(c *gin.Context) {

		c.Bind(&postRequestScrap)

		if postRequestScrap.URL == "" {
			c.JSON(400, gin.H{
				"error": "URL is mandatory",
			})
			return
		}

		co := colly.NewCollector()
		co.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL)
		})

		co.OnHTML("div.s-result-list.s-search-results.sg-row", func(e *colly.HTMLElement) {
			e.ForEach("div.a-section.a-spacing-medium", func(_ int, e *colly.HTMLElement) {

				scrap.Product = e.ChildText("span.a-size-medium.a-color-base.a-text-normal")
				if scrap.Product == "" {
					// If we can't get any name, we return and go directly to the next element
					return
				}
				scrap.Image = e.ChildAttr("img", "src")

				scrap.Reviews = e.ChildText("span.a-size-base")

				scrap.Rating = e.ChildText("span.a-icon-alt")

				scrap.Price = e.ChildText("span.a-price  span.a-price-symbol") + e.ChildText("span.a-price  span.a-price-whole")

				scrapArr = append(scrapArr, scrap)

			})
		})
		co.Visit(postRequestScrap.URL)
		var save saveScrap
		save.Data = scrapArr
		save.URL = postRequestScrap.URL
		save.CreatedAt = time.Now()
		//Inserting the data in to database
		insertResult, err := db.Collection("scraps").InsertOne(context.TODO(), save)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		fmt.Println("Inserted post with ID:", insertResult.InsertedID)

		c.JSON(200, gin.H{
			"success": true,
			"data":    scrapArr,
		})
	})

	router.Run(":8081")

}
