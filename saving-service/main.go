package main

import (
	"context"
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	mongo "github.com/rahulshibu/sellerapp/database"
)

type (
	postData struct {
		URL     string  `json:"url"  bson:"url" valid:"required,url"`
		Product product `json:"product"  bson:"product" valid:"required"`
	}
	product struct {
		Name         string `json:"name"  bson:"name" valid:"required"`
		ImageURL     string `json:"imageURL"  bson:"imageURL" valid:"required,url"`
		Description  string `json:"description"  bson:"description" valid:"required"`
		Price        string `json:"price"  bson:"price" valid:"required"`
		TotalReviews int    `json:"totalReviews"  bson:"totalReviews" valid:"required"`
	}
)

var (
	err  error
	post postData
)

func main() {
	//gin router with default middlewares
	router := gin.Default()
	fmt.Println("saving service")

	//Establishing mongo connection. If the connection is not correct, will throw error
	db := mongo.GetSharedConnection()

	//POST api to save data
	router.POST("/save", func(c *gin.Context) {
		c.Bind(&post)
		_, err := govalidator.ValidateStruct(post)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		_, err = govalidator.ValidateStruct(post.Product)
		if err != nil {
			fmt.Print(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		insertResult, err := db.Collection("scraps").InsertOne(context.TODO(), post)
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
		})
		return

	})

	router.Run(":8080")

}
