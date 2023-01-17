package middleware

import (
	errors2 "github.com/WesleyWu/ri-restful-api/util/errors"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gvalid"
	"net/http"
)

func ResponseJsonWrapper(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("服务器开小差了，请稍后再试吧！")
	} else {
		if err := r.GetError(); err != nil {
			r.Response.ClearBuffer()
			_, ok := err.(gvalid.Error)
			if ok {
				validationError := errors2.NewBadRequestErrorf(r.GetBodyString(), err.Error())
				r.Response.Status = validationError.Code
				r.Response.WriteJsonExit(validationError)
			}
			validationError, ok := err.(errors2.RequestError)
			if ok {
				r.Response.Status = validationError.Code
				r.Response.WriteJsonExit(validationError)
			}
			serviceError, ok := err.(errors2.ServiceError)
			if ok {
				r.Response.Status = serviceError.Code
				r.Response.WriteJsonExit(serviceError)
			}
			r.Response.Status = 500
			r.Response.WriteJsonExit(err)
		} else {
			handlerRes := r.GetHandlerResponse()
			r.Response.WriteJsonExit(handlerRes)
		}
	}
}
