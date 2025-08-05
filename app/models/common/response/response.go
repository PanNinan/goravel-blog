package response

import "github.com/goravel/framework/contracts/http"

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(ctx http.Context, code int, data interface{}, msg string) http.Response {
	return ctx.Response().Json(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx http.Context) http.Response {
	return Result(ctx, SUCCESS, map[string]interface{}{}, "success")
}

func OkWithMessage(ctx http.Context, message string) http.Response {
	return Result(ctx, SUCCESS, map[string]interface{}{}, message)
}

func OkWithData(ctx http.Context, data interface{}) http.Response {
	return Result(ctx, SUCCESS, data, "success")
}

func OkWithDetailed(ctx http.Context, data interface{}, message string) http.Response {
	return Result(ctx, SUCCESS, data, message)
}

func Fail(ctx http.Context) http.Response {
	return Result(ctx, ERROR, map[string]interface{}{}, "fail")
}

func FailWithMessage(ctx http.Context, message string) http.Response {
	return Result(ctx, ERROR, map[string]interface{}{}, message)
}

func FailWithDetailed(ctx http.Context, data interface{}, message string) http.Response {
	return Result(ctx, ERROR, data, message)
}

func NoAuth(ctx http.Context, message string) http.Response {
	return ctx.Response().Json(http.StatusUnauthorized, Response{
		http.StatusUnauthorized,
		nil,
		message,
	})
}

func InternalServerError(ctx http.Context, message string) http.Response {
	return ctx.Response().Json(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		nil,
		message,
	})
}
