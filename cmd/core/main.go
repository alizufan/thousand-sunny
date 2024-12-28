package main

import (
	"go.uber.org/fx"

	"github.com/Mind2Screen-Dev-Team/thousand-sunny/app/dependency"
	"github.com/Mind2Screen-Dev-Team/thousand-sunny/app/registry"
)

func main() {
	fx.New(
		registry.Fx,
		registry.Dependency,
		registry.DependencyStartUp,
		registry.Http,
		registry.HttpStartUp,
		registry.HttpGlobalMiddleware,
	).Run()

	defer dependency.RotateLog()
}
