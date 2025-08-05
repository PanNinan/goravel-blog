package controllers

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/contracts/queue"
	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/collect"
	"github.com/goravel/framework/support/color"
	"github.com/goravel/framework/support/str"
	"goravel-blog/app/jobs"
	"time"
)

type TestController struct {
	// Dependent services
}

func NewTestController() *TestController {
	return &TestController{
		// Inject services
	}
}

func (r *TestController) Index(ctx http.Context) http.Response {
	color.Black().Println("TestController Index")
	color.Green().Println("TestController Index")
	color.Red().Println("TestController Index")
	color.Yellow().Println("TestController Index")
	color.Blue().Println("TestController Index")
	color.Magenta().Println("TestController Index")
	color.Cyan().Println("TestController Index")
	color.White().Println("TestController Index")
	color.Gray().Println("TestController Index")
	color.Default().Println("TestController Index")
	str.Of("hello world").After(" ").Append("!").Upper().Tap(func(s str.String) {
		color.Infoln(s.String())
	})
	collect.Count([]int{1, 2, 3, 4, 5})
	err := facades.Queue().Job(&jobs.TestJob{}, []queue.Arg{
		{"name", "string"},
		{"18", "uint"},
		{"true", "bool"},
		{[]string{"goravel", "framework"}, "[]string"},
	}).Delay(time.Now().Add(time.Second * 30)).Dispatch()
	if err != nil {
		return ctx.Response().Status(http.StatusOK).Json(http.Json{
			"message": err.Error(),
		})
	}
	return ctx.Response().Success().Json(http.Json{
		"message": "success",
	})
}
