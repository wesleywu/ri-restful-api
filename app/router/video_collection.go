package router

import (
	"github.com/WesleyWu/ri-restful-api/app/service"
	"github.com/WesleyWu/ri-restful-api/middleware"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// 加载路由
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.Group("/app", func(group *ghttp.RouterGroup) {
			group.Group("/video-collection", func(group *ghttp.RouterGroup) {
				group.Middleware(middleware.ResponseJsonWrapper)
				group.Bind(
					service.VideoCollection,
					service.VideoCollectionRepo,
				)
			})
		})
	})
}
