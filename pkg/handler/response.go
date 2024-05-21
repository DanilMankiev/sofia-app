package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"github.com/oklog/ulid"
	"strings"
	"fmt"
	"time"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}
type tokenResponse struct{
	Token string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type imageURLResponse struct{
	ImageURL []string `json:"imageurl"`
}
type idResponse struct{
	ID int `json:"id"`
}


func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}


func randomFilename() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	return strings.ToLower(fmt.Sprintf("%v", ulid.MustNew(ulid.Timestamp(t), entropy)))
}