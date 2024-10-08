// Code generated by mlflow/go/cmd/generate/main.go. DO NOT EDIT.

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mlflow/mlflow-go/pkg/server/parser"
	"github.com/mlflow/mlflow-go/pkg/contract/service"
	"github.com/mlflow/mlflow-go/pkg/utils"
	"github.com/mlflow/mlflow-go/pkg/protos"
)

func RegisterModelRegistryServiceRoutes(service service.ModelRegistryService, parser *parser.HTTPRequestParser, app *fiber.App) {
	app.Post("/mlflow/registered-models/get-latest-versions", func(ctx *fiber.Ctx) error {
		input := &protos.GetLatestVersions{}
		if err := parser.ParseBody(ctx, input); err != nil {
			return err
		}
		output, err := service.GetLatestVersions(utils.NewContextWithLoggerFromFiberContext(ctx), input)
		if err != nil {
			return err
		}
		return ctx.JSON(output)
	})
	app.Get("/mlflow/registered-models/get-latest-versions", func(ctx *fiber.Ctx) error {
		input := &protos.GetLatestVersions{}
		if err := parser.ParseQuery(ctx, input); err != nil {
			return err
		}
		output, err := service.GetLatestVersions(utils.NewContextWithLoggerFromFiberContext(ctx), input)
		if err != nil {
			return err
		}
		return ctx.JSON(output)
	})
}
