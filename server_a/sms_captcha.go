package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func smsCaptchaImp(router *gin.Engine) {
	router.GET("/smsCaptcha/:phone", func(c *gin.Context) {
		ip := c.ClientIP()
		phone := c.Param("phone")
		code := "110112"

		db := fetchDB(c)

		var updatedAt time.Time
		err := db.QueryRow("select updated_at from a_sms_captcha where phone = ?", phone).Scan(&updatedAt)
		if err != nil {
			log.Println(err)
			return
		}
		//updatedAt.Unix() //Unix 时间戳
		local, err := time.LoadLocation("Local")
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(updatedAt)
		fmt.Println(updatedAt.In(local).String())
		fmt.Println(reflect.TypeOf(updatedAt))
		c.JSON(http.StatusOK, gin.H{"ip": ip, "code": code, "updatedAt": updatedAt.Unix()})

		/*
			stmt, err := db.Prepare("INSERT INTO a_sms_captcha(ip, phone, code) VALUES(?, ?, ?)")
			if err != nil {
				log.Println(err)
				return
			}

			res, err := stmt.Exec(ip, phone, code)
			if err != nil {
				log.Println(err)
				return
			}

			lastID, err := res.LastInsertId()
			if err != nil {
				log.Println(err)
				return
			}
			c.JSON(http.StatusOK, gin.H{"id": lastID})
		*/
	})
}
