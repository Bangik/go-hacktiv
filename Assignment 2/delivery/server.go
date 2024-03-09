package delivery

import (
	"fmt"
	"hacktiv-assignment-2/config"
	"hacktiv-assignment-2/delivery/controller"
	"hacktiv-assignment-2/manager"
	"log"

	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
}

func (a *appServer) initController() {
	controller.NewOrderController(a.engine, a.usecaseManager.OrderUsecase())
}

func (a *appServer) Run() {
	a.initController()

	err := a.engine.Run(a.host)
	if err != nil {
		panic(err.Error())
	}
}

func Server() *appServer {
	// Configurasi <=> Membuat Koneksi <=> Repository <=> Service <=> Controller
	engine := gin.Default()

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln("Error Config : ()", err.Error())
	}

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatalln("Error Conection : ", err.Error())
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	host := fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort)

	return &appServer{
		engine:         engine,
		host:           host,
		usecaseManager: useCaseManager,
	}
}
