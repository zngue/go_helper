package response

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var (
	SuccessMessage = "操作成功"
	FailMessage    = "操作失败"
	SuccessCode    = 200
	FailCode       = 400
	ParameterCodeError=422
	ParameterError ="参数错误"

)
func HttpSuccessWithError(ctx *gin.Context,err error, Data interface{}) {
	if err!=nil{
		ReturnCode(ctx, FailCode, err.Error(), Data)
	}else{
		ReturnCode(ctx, SuccessCode, SuccessMessage, Data)
	}
}
func HttpParameterError(ctx *gin.Context,err error) {
	ReturnCode(ctx, FailCode, err.Error(),nil)
	return
}

func ReturnCode(ctx *gin.Context, StatusCode int, Message string, Data interface{}) {
	ctx.JSON(200, &Response{StatusCode, Message, Data})
	return
}
func HttpOk(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, SuccessCode, SuccessMessage, data[0])
	} else {
		ReturnCode(ctx, SuccessCode, SuccessMessage, nil)
	}
}
func HttpOkWithMessage(ctx *gin.Context, message string, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, SuccessCode, message, data[0])
	} else {
		ReturnCode(ctx, SuccessCode, message, nil)
	}
}
func HttpFail(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, FailCode, FailMessage, data[0])
	} else {
		ReturnCode(ctx, FailCode, FailMessage, nil)
	}
}
func HttpFailWithMessage(ctx *gin.Context, message string, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, FailCode, message, data[0])
	} else {
		ReturnCode(ctx, FailCode, message, nil)
	}
}
func HttpFailWithCodeAndMessage(code int, message string, ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, code, message, data[0])
	} else {
		ReturnCode(ctx, code, message, nil)
	}
}
func HttpFailWithParameter(ctx *gin.Context, data ...interface{}) {
	if len(data) > 0 {
		ReturnCode(ctx, 422, ParameterError, data[0])
	} else {
		ReturnCode(ctx, 422, ParameterError, nil)
	}
}
func HttpFailWithErr(ctx *gin.Context, data ...error)  {
	if len(data) > 0 {
		if len(data)==1 {
			ReturnCode(ctx, FailCode, FailMessage, data[0].Error())
		}else{
			var userArr []string
			for _,err:=range  data {
				userArr=append(userArr, err.Error())
			}
			ReturnCode(ctx, FailCode, FailMessage, userArr)
		}

	} else {
		ReturnCode(ctx, FailCode, FailMessage, nil)
	}
}

