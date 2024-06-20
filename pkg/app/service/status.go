package service

import "github.com/gin-gonic/gin"

type Status[T any] struct {
}
type IStatus interface {
	Status(ctx *gin.Context) (err error)
}
