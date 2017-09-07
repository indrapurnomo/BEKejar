package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

type Person struct {
	ID         uint   `json:"id"`
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Phone      string `json:"phone"`
	Hobby      string `json:"hobby"`
	Website    string `json:"website"`
	AvatarUrl  string `json:"avatarurl"`
}

// START OMIT
func main() {
	db, err = gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=apa sslmode=disable password=091289")

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.POST("/user", CreatePerson)
	r.GET("/user", GetPeople) // HL
	r.GET("/user/:id", GetPerson)
	r.PUT("/user/:id", UpdatePerson)
	r.PATCH("/user/:id", UpdatePerson)
	r.DELETE("/user/:id", DeletePerson)
	r.Run(":8080")
}

// END OMIT

func CreatePerson(c *gin.Context) {
	var person Person
	c.BindJSON(&person)
	db.Create(&person)
	c.JSON(200, person)
}

func GetPeople(c *gin.Context) {
	var allpost []Person
	if err := db.Find(&allpost).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, allpost)
	}

}

func GetPerson(c *gin.Context) {
	id := c.Params.ByName("id")
	fmt.Println(id)
	var person Person
	if err := db.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, person)
	}
}

func UpdatePerson(c *gin.Context) {
	var person Person
	id := c.Params.ByName("id")
	if err := db.Where("id= ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&person)
	db.Save(&person)
	c.JSON(200, person)
}

func DeletePerson(c *gin.Context) {
	id := c.Params.ByName("id")
	var person Person
	d := db.Where("id = ?", id).Delete(&person)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
