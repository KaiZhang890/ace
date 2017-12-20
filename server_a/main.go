package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DataSourceName string
}

func main() {
	// for release
	/*
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f)
		gin.DisableConsoleColor()
		gin.SetMode(gin.ReleaseMode)
	*/
	var conf Config
	if _, err := toml.DecodeFile("./config.toml", &conf); err != nil {
		fmt.Println(err)
	}
	db, err := sql.Open("mysql", conf.DataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := gin.Default()
	router.Use(bindDB(db))

	v1 := router.Group("/v1/todos")
	{
		v1.POST("/", createTodo)        // http://127.0.0.1:8080/v1/todos?content=读报
		v1.GET("/", fetchAllTodo)       // http://127.0.0.1:8080/v1/todos
		v1.GET("/:id", fetchSingleTodo) // http://127.0.0.1:8080/v1/todos/14
		v1.PUT("/:id", updateTodo)      // http://127.0.0.1:8080/v1/todos/14?content=readNewspaper&status=2
		v1.DELETE("/:id", deleteTodo)   // http://127.0.0.1:8080/v1/todos/14
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "API not found"})
	})

	router.Run(":8080")
}

func bindDB(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("databaseConn", db)
		c.Next()
	}
}

func fetchDB(c *gin.Context) *sql.DB {
	db, ok := c.MustGet("databaseConn").(*sql.DB)
	if !ok {
		log.Fatal("There is no databaseConn")
	}
	return db
}

func createTodo(c *gin.Context) {
	content := c.Query("content")

	db := fetchDB(c)
	stmt, err := db.Prepare("INSERT INTO todos(content) VALUES(?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(content)
	if err != nil {
		log.Fatal(err)
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	var status int
	err = db.QueryRow("SELECT content, status FROM todos WHERE id = ?", lastID).Scan(&content, &status)
	c.JSON(http.StatusOK, gin.H{"id": lastID, "content": content, "status": 0})
}

func fetchAllTodo(c *gin.Context) {
	db := fetchDB(c)
	rows, err := db.Query("SELECT id, content, status FROM todos")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var (
		id      int
		content string
		status  int
	)
	var todos []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(&id, &content, &status)
		if err != nil {
			log.Fatal(err)
		}
		m := make(map[string]interface{})
		m["id"] = id
		m["content"] = content
		m["status"] = status
		todos = append(todos, m)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

func fetchSingleTodo(c *gin.Context) {
	db := fetchDB(c)
	todoID := c.Param("id")
	var (
		content string
		status  int
	)
	stmt, err := db.Prepare("SELECT content, status FROM todos WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(todoID).Scan(&content, &status)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"content": content, "status": status})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	content := c.Query("content")
	status := c.Query("status")

	db := fetchDB(c)
	stmt, err := db.Prepare("UPDATE todos SET content=?, status=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(content, status, id)
	if err != nil {
		log.Fatal(err)
	}

	err = db.QueryRow("SELECT content, status FROM todos WHERE id = ?", id).Scan(&content, &status)
	c.JSON(http.StatusOK, gin.H{"id": id, "content": content, "status": status})
}

func deleteTodo(c *gin.Context) {
	db := fetchDB(c)
	todoID := c.Param("id")

	_, err := db.Exec("DELETE FROM todos WHERE id = ?", todoID)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "deleted"})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
}
