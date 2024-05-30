package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type person struct {
	Firstname *string `json:"firstname"`
	Lastname  *string `json:"lastname"`
	ID        int     `json:"id"`
	Birthday  *string `json:"birthday"`
}

func main() {
	godotenv.Load()

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	var err error
	db, err = sql.Open("postgres", psqlconn)
	CheckErr(err)

	router := gin.Default()
	router.GET("/people", getPeople)
	router.POST("/people", createPeople)

	router.Run("localhost:8080")
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getPeople(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	rows, err := db.Query("SELECT id, firstname, lastname, birthday FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var people []person
	for rows.Next() {
		var p person
		err := rows.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Birthday)
		if err != nil {
			log.Fatal(err)
		}
		people = append(people, p)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.IndentedJSON(http.StatusOK, people)
}

func createPeople(c *gin.Context) {

	var newPerson person
	if err := c.BindJSON(&newPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		log.Println(err)
		return
	}

	stmt, err := db.Prepare("INSERT INTO people (firstname, lastname, id, birthday) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(newPerson.Firstname, newPerson.Lastname, newPerson.ID, newPerson.Birthday); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, newPerson)
}
