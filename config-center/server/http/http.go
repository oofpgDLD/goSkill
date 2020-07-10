package http

import (
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"github.com/oofpgDLD/goSkill/config-center/service"
)

var (
	log = log15.New("server", "http")
	Svc *service.Service
)

func Init(c *conf.Config, s *service.Service) {
	api.SetConfig(c.AuthServer)
	//
	initService(s)
	r := Default()
	SetupEngine(r)
	setupInnerEngine(r)
	go func(c *conf.Config, e *gin.Engine) {
		err := r.Run(c.Server.Addr)
		if err != nil {
			log.Error("main run", "err", err)
		}
	}(c, r)
}

func initService(s *service.Service) {
	Svc = s
}

// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *gin.Engine {
	router := gin.New()
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.LoggerWithFormatter(api.Chat33GinLogFormatter))
	router.Use(gin.Recovery())
	return router
}

func SetupEngine(e *gin.Engine) *gin.Engine {
	root := e.Group("", api.RespMiddleWare())
	/*root.Use(api.WorkAuthMiddleWare())
	root.Any("/user", UserInfo)*/

	team := root.Group("/team")
	team.POST("/info", EPInfo)
	team.POST("/search", EPSearch)
	team.Use(api.WorkAuthMiddleWare())
	{
		team.POST("/create", CreateEP)
		
	}

	role := team.Group("/role")
	{
		role.POST("/users", RoleUsers)
		role.POST("/edit", EditRoleUser)
	}

	dep := team.Group("/department")
	{
		dep.POST("/members", DepMembers)
		dep.POST("/delDepMember", DelDepMembers)
		dep.POST("/addMembers", AddDepMembers)
		dep.POST("/add", CreateDepartment)
		dep.POST("/info", DepartmentInfo)
		dep.POST("/edit", EditDepartment)
		dep.POST("/delete", DelDepartment)
	}
	return e
}

func setupInnerEngine(e *gin.Engine) *gin.Engine {
	root := e.Group("/inner", api.RespMiddleWare())
	root.Any("/checkToken", CheckToken)
	root.Any("/checkUser", CheckUser)
	root.Any("/users", UserList)
	return e
}
