package server

import (
	"context"
	"crypto/rsa"
	"github.com/cyantarek/ghost-clone-project/cache"
	"github.com/cyantarek/ghost-clone-project/config"
	"github.com/cyantarek/ghost-clone-project/db"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
	"gopkg.in/macaron.v1"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	oauth2Models "gopkg.in/oauth2.v3/models"
	oauth2 "gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	SignKey     *rsa.PrivateKey
	VerifyKey   *rsa.PublicKey
	ContextData interface{}
	DB          db.Db
	Cache       cache.Cache
	Settings    Settings
	Router      *macaron.Macaron
	HttpServer  *http.Server
	Oauth2      *oauth2.Server
	ClientStore *store.ClientStore
}

type Settings struct {
	Title             string       `json:"title"`
	Description       string       `json:"description"`
	Logo              string       `json:"logo"`
	Icon              string       `json:"icon"`
	CoverImage        string       `json:"cover_image" bson:"cover_image"`
	Facebook          string       `json:"facebook"`
	Twitter           string       `json:"twitter"`
	Lang              string       `json:"lang"`
	Timezone          string       `json:"timezone"`
	CodeinjectionHead string       `json:"codeinjection_head" bson:"codeinjection_head"`
	CodeinjectionFoot string       `json:"codeinjection_foot" bson:"codeinjection_foot"`
	Navigation        []Navigation `json:"navigation"`
}

type Navigation struct {
	Label string
	URL   string
}

var server *Server

func New() (*Server, error) {
	// instantiate new server
	server = new(Server)

	// load config
	config.InitConfig()

	var err error
	var portStr string

	// check which port to run the application
	if viper.GetString("server.port") != "" {
		portStr = viper.GetString("server.port")
	} else {
		portStr = "3036"
	}

	// instantiate the router/app
	server.Router = macaron.New()

	// check mode and run logger according to mode
	if len(os.Args) > 1 && os.Args[1] == "development" || len(os.Args) == 1 {
		server.Router.Use(macaron.Logger())
		log.Println(chalk.Yellow, "[development mode]")
	} else if len(os.Args) > 1 && os.Args[1] == "production" {
		log.Println(chalk.Yellow, "[production mode]")
	}

	// instantiate a new httpServer for the Server struct instance with proper timeout and port
	server.HttpServer = &http.Server{
		Addr:         ":" + portStr,
		Handler:      server.Router,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
	}

	// start db
	server.DB, err = db.NewMongo()
	if err != nil {
		return nil, err
	}

	// start routes
	server.InitRoutes()

	// setup cache server based on development needs
	if viper.GetBool("cache.enabled") {
		log.Println(chalk.Green, "cache enabled")
		redisAddr := viper.GetString("cache.host") + ":" + viper.GetString("cache.port")
		server.Cache, err = cache.NewMemoryStorage(redisAddr, "", 1)
		if err != nil {
			return nil, err
		}
	} else {
		log.Println(chalk.Red, "cache disabled")

	}

	// use json rendering middleware
	server.Router.Use(macaron.Renderer())

	// setup settings
	server.Settings = Settings{"Mits Articles",
		"the professional publishing platform",
		"https://static.ghost.org/v1.0.0/images/ghost-logo.svg",
		"https://static.ghost.org/favicon.ico",
		"https://static.ghost.org/v1.0.0/images/blog-cover.jpg",
		"Mits",
		"Mits",
		"en",
		"ETC/UTC",
		"",
		"",
		[]Navigation{{"Home", "/"}, {"Tag", "/tag"}},}

	// setup oauth2 server based on development needs
	if viper.GetBool("server.oauth2") {
		if err := server.SetupOAuth2(); err != nil {
			return nil, err
		}
	} else {
		log.Println(chalk.Red, "Oauth2 disabled")
	}

	// setup jwt server based on development needs
	if viper.GetBool("server.jwt") {
		if err := server.SetupJwt(); err != nil {
			return nil, err
		}
	} else {
		log.Println(chalk.Red, "JWT disabled")
	}

	return server, nil
}

// GetSettings returns setting as JSON
func (s *Server) GetSettings(c *macaron.Context) {
	c.JSON(200, s.Settings)
}

// setup OAuth2
func (s *Server) SetupOAuth2() error {
	log.Println(chalk.Green, "Oauth2 enabled")
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	s.ClientStore = store.NewClientStore()
	manager.MapClientStorage(s.ClientStore)

	s.Oauth2 = oauth2.NewDefaultServer(manager)
	s.Oauth2.SetAllowGetAccessRequest(true)
	s.Oauth2.SetClientInfoHandler(oauth2.ClientFormHandler)
	s.Oauth2.SetUserAuthorizationHandler(s.userAuthorizationHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	s.Oauth2.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	s.Oauth2.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	if err := s.ClientStore.Set("default", &oauth2Models.Client{
		ID:     "default",
		Secret: "default",
		Domain: "http://localhost",
	}); err != nil {
		return err
	}
	return nil
}

func GetServer() (*Server, error) {
	if server == nil {
		return New()
	} else {
		return server, nil
	}
}

func GetRouter() (*macaron.Macaron, error) {
	srv, err := GetServer()
	if err != nil {
		return nil, err
	}
	return srv.Router, nil
}

func (s *Server) Run() error {
	log.Println("starting server: API " + config.VERSION)
	return s.HttpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	s.DB.Close()
	return s.HttpServer.Shutdown(context.Background())
}
