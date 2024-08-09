package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	shopowner "ginpractice1/internal/usecase/shopowner"
)

type album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title:"Blue train", Artist:"aaaa", Price:54},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func updateProfile (c *gin.Context) {
	var request shopowner.UpdateProfileRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}
	usecase := shopowner.ShopOwnerUsecase{}
	result := usecase.UpdateMyProfile(request)

	if result == false {
		c.IndentedJSON(http.StatusBadRequest, "karioki error")
	}
	c.IndentedJSON(http.StatusOK, "suceeded")
}

// type 

func main(){
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.POST("/owner/profile", updateProfile)
	router.Run("localhost:8080")
}