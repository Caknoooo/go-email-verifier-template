package main

import (
	"go-email-verifier-tool/dto"
	"go-email-verifier-tool/utils"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/send-mail", sendMail)

	router.Run(":8080")
}

func sendMail(ctx *gin.Context) {
	var req dto.SendMailRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	html, err := os.ReadFile("index.html")
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	draftEmail, err := utils.MakeVerificationEmail(req.Email, html)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := utils.SendMail(req.Email, draftEmail["subject"], draftEmail["body"]); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "success send email",
	})
}