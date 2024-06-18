package main

import (
	"example/hello/controllers"
	"example/hello/initializers"
	"example/hello/middleware"

	// "example/hello/models"
	// "log"
	// "net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run("localhost:8080")
	// database.ConnectToDB()

	// router := gin.Default()
	// router.GET("/users", getUsers)
	// router.GET("/users/:id", getUser)
	// router.POST("/users", createUser)
	// router.PUT("/users/:id", updateUser)
	// router.DELETE("/users/:id", deleteUser)

	// router.Run("localhost:8080")
}

// func getUsers(c *gin.Context) {
// 	c.Header("Content-Type", "application/json")

// 	rows, err := initializers.DB.Query("SELECT id, username, password FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var user []models.User
// 	for rows.Next() {
// 		var u models.User
// 		err := rows.Scan(&u.ID, &u.Username, &u.Password)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		user = append(user, u)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.IndentedJSON(http.StatusOK, user)
// }

// func getUser(c *gin.Context) {
// 	id := c.Param("id")
// 	sqlStatement := `SELECT * FROM users WHERE id=$1;`
// 	rows, err := db.Query(sqlStatement, id)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	c.IndentedJSON(http.StatusOK, rows)
// }

// func createUser(c *gin.Context) {

// 	var newUser models.User
// 	if err := c.BindJSON(&newUser); err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
// 		log.Println(err)
// 		return
// 	}

// 	stmt, err := db.Prepare("INSERT INTO people (username, id, password) VALUES ($1, $2, $3)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	if _, err := stmt.Exec(newUser.Username, newUser.Password, newUser.ID); err != nil {
// 		log.Fatal(err)
// 	}

// 	c.JSON(http.StatusCreated, newUser)
// }

// func updateUser(c *gin.Context) {
// 	id := c.Param("id")

// 	var updatedUser models.User

// 	if err := c.ShouldBindJSON(&updatedUser); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
// }

// func deleteUser(c *gin.Context) {

// }
