package gin_resp

import (
	"github.com/gin-gonic/gin"
)

type GinResp struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

var ginDefault = GinResp{
	StatusCode: 200,
	Message:    "success",
	Data:       nil,
}

type RespFn func(resp *GinResp)

/*
*@Author Administrator
*@Date 29/4/2021 15:12
*@desc
 */
func ginBase(ctx *gin.Context, fn ...RespFn) {
	if len(fn) > 0 {
		for _, respFn := range fn {
			respFn(&ginDefault)
		}
	}
	ctx.JSON(200, ginDefault)
	return
}

/*
*@Author Administrator
*@Date 29/4/2021 15:16
*@desc
 */
func Code(code int) RespFn {
	return func(resp *GinResp) {
		resp.StatusCode = code
	}
}

/*
*@Author Administrator
*@Date 29/4/2021 15:16
*@desc
 */
func Message(msg string) RespFn {
	return func(resp *GinResp) {
		resp.Message = msg
	}
}

/*
*@Author Administrator
*@Date 29/4/2021 15:16
*@desc
 */
func Data(data interface{}) RespFn {
	return func(resp *GinResp) {
		resp.Data = data
	}
}

/*
*@Author Administrator
*@Date 29/4/2021 15:16
*@desc 参数错误
 */
func ParameterError(ctx *gin.Context, fn ...RespFn) {
	var list []RespFn
	list = append(list, Code(422))
	list = append(list, fn...)
	ginBase(ctx, list...)
}

/*
*@Author Administrator
*@Date 29/4/2021 15:19
*@desc ServerError  服务器错误
 */
func ServerError(ctx *gin.Context, fn ...RespFn) {
	var list []RespFn
	list = append(list, Code(400))
	list = append(list, Message("服务器内部错误"))
	list = append(list, fn...)
	ginBase(ctx, list...)
}

/*
*@Author Administrator
*@Date 29/4/2021 15:34
*params Code 默认200 Message 默认success Data 默认nil
*@desc Success 操作成功
 */
func Success(ctx *gin.Context, fn ...RespFn) {
	ginBase(ctx, fn...)
}

/*
*@Author Administrator
*@Date 29/4/2021 15:39
*@desc
 */
func Error(ctx *gin.Context, fn ...RespFn) {

	var list []RespFn
	list = append(list, Code(400))
	list = append(list, fn...)

}
