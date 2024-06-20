package service

import "github.com/gin-gonic/gin"

type Status[T any] struct {
}

func (s *Status[T]) Status(ctx *gin.Context) (err error) {
	//TODO implement me
	panic("implement me")
}

type IStatus interface {
	Status(ctx *gin.Context) (err error)
}

func NewStatus[T any]() IStatus {
	return new(Status[T])

}
