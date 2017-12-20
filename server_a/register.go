package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

type RegModel struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func registerImp(router *gin.Engine) {
	router.POST("/register", func(c *gin.Context) {
		var model RegModel
		if err := c.ShouldBindJSON(&model); err == nil {
			match, _ := regexp.MatchString("^1[3|4|5|7|8][0-9]{9}$", model.Phone)
			if !match {
				c.JSON(http.StatusBadRequest, gin.H{"error": "手机号不合法"})
				return
			}

			db := fetchDB(c)
			exists, err := checkExists(model, db)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}
			if exists {
				c.JSON(http.StatusBadRequest, gin.H{"error": "手机号已存在"})
				return
			}

			lastID, err := insert(model, db)
			if err != nil {
				fmt.Fprintln(gin.DefaultWriter, err.Error())
				c.JSON(http.StatusInternalServerError, gin.H{})
				return
			}

			c.JSON(http.StatusOK, gin.H{"id": lastID})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}

func checkExists(model RegModel, db *sql.DB) (bool, error) {
	stmt, err := db.Prepare("SELECT COUNT(*) FROM a_users WHERE phone = ?")
	if err != nil {
		return false, err
	}
	var count int
	err = stmt.QueryRow(model.Phone).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func insert(model RegModel, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO a_users(phone, name, password) VALUES(?, ?, ?)")
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(model.Phone, "默认名字", model.Password)
	if err != nil {
		return 0, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastID, nil
}
