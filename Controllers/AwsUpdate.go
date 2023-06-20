package Controllers

import (
	"223987-235861-184019-providers/Service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AwsCreds struct {
	AccessKeyId     string `json:"accessKeyId" validate:"required"`
	SecretAccessKey string `json:"secretAccessKey" validate:"required"`
	SessionToken    string `json:"sessionToken" validate:"required"`
}

func UpdateAwsCreds(c *gin.Context) {
	var awsCreds AwsCreds
	c.BindJSON(&awsCreds)

	Service.UpdateSession(awsCreds.AccessKeyId, awsCreds.SecretAccessKey, awsCreds.SessionToken)

	c.Status(http.StatusOK)
}
