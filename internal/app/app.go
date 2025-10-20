/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package app

import (
	"go-firebase/internal/controller"
	"go-firebase/internal/data/repository"
	"go-firebase/internal/firebase"
	"go-firebase/internal/handler"
	"go-firebase/internal/helper"
	"go-firebase/internal/mapper"
	"go-firebase/internal/middleware"
	"go-firebase/internal/route"
	"go-firebase/internal/server"
	"go-firebase/internal/service"
	"go-firebase/pkg/logger"

	"github.com/ngdangkietswe/swe-go-common-shared/config"

	"go.uber.org/fx"
)

func Start() {
	config.Init()
	fxApp := fx.New(
		logger.Module,
		firebase.Module,
		repository.Module,
		mapper.Module,
		helper.Module,
		service.Module,
		handler.Module,
		controller.Module,
		route.Module,
		middleware.Module,
		server.Module,
	)
	fxApp.Run()
}
