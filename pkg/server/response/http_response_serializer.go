package response

import (
	"github.com/gofiber/fiber/v2"

	"github.com/mlflow/mlflow-go/pkg/contract"
	"github.com/mlflow/mlflow-go/pkg/protos"
	"github.com/mlflow/mlflow-go/pkg/utils/json"
)

// SerializeResponse serialises response into JSON using custom JSON serializer.
//
//nolint:wrapcheck
func SerializeResponse(ctx *fiber.Ctx, data interface{}) error {
	response, err := json.Marshal(data)
	if err != nil {
		return contract.NewError(protos.ErrorCode_BAD_REQUEST, err.Error())
	}

	ctx.Set("Content-Type", ctx.Get("Content-Type", "application/json; charset=utf-8"))

	return ctx.Send(response)
}
