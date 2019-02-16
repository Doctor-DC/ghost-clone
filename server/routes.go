package server

import (
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/go-macaron/toolbox"
	"gopkg.in/macaron.v1"
	"net/http/pprof"
)

func (s *Server) InitRoutes() {
	s.Router.Use(toolbox.Toolboxer(s.Router))
	s.Router.Get("/debug/pprof/", pprof.Index)
	s.Router.Get("/debug/pprof/cmdline", pprof.Cmdline)
	s.Router.Get("/debug/pprof/profile", pprof.Profile)
	s.Router.Get("/debug/pprof/symbol", pprof.Symbol)
	s.Router.Get("/debug/pprof/trace", pprof.Trace)
	//s.Router.Handle("GET", "/debug/pprof/heap", []macaron.Handler{pprof.Handler("heap")})
	//s.Router.Get("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	//s.Router.Get("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	//s.Router.Get("/debug/pprof/block", pprof.Handler("block"))
	s.Router.Get("/health", s.HealthEndpoint)

	s.Router.Get("/settings", s.GetSettings)

	s.Router.Group("/api/"+config.VERSION, func() {
		s.Router.Get("", func(c *macaron.Context) {
			c.JSON(200, map[string]string{"version":config.VERSION})
		})
		s.Router.Group("/posts", func() {
			s.Router.Get("", s.GetAllPosts)
			s.Router.Get("/slug/:postSlug", s.GetPostBySlug)
			s.Router.Get("/:postId", s.GetPostById)
		//})
		}, s.ValidateOauthToken)
		s.Router.Group("/tags", func() {
			s.Router.Get("", s.GetAllTags)
			s.Router.Get("/:tagId", s.GetTagById)
			s.Router.Get("/slug/:tagSlug", s.GetTagBySlug)
		//})
		}, s.ValidateOauthToken)
		s.Router.Group("/authors", func() {
			s.Router.Get("", s.GetAllAuthors)
			s.Router.Get("/:authorId", s.GetAuthorById)
			s.Router.Get("/slug/:authorSlug", s.GetAuthorBySlug)
		//})
		}, s.ValidateOauthToken)
		s.Router.Group("/oauth", func() {
			s.Router.Get("/credentials", s.GetCredentials)
			s.Router.Get("/token", s.GetToken)
		})

	})

	//authGroup := server.Router.Group("/auth")
	//authGroup.POST("/login", server.Login)
	//authGroup.POST("/signup", server.Signup)
}
