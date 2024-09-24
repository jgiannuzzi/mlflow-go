package response

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/mlflow/mlflow-go/pkg/contract"
	"github.com/mlflow/mlflow-go/pkg/protos"
)

// SerializeResponse serialises response into JSON using custom JSON serializer.
//
//nolint:wrapcheck
func SerializeResponse(ctx *fiber.Ctx, data proto.Message) error {
	response, err := protojson.Marshal(data)
	if err != nil {
		return contract.NewError(protos.ErrorCode_BAD_REQUEST, err.Error())
	}

	ctx.Set("Content-Type", ctx.Get("Content-Type", "application/json; charset=utf-8"))

	return ctx.Send(response)
}
