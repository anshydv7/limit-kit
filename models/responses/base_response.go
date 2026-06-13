package responses

import (
	"auth-service/utils"
	"context"
)

type Response struct {
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`

	Message   string `json:"message"`
	RequestId string `json:"request_id"`
}

func (res *Response) NewErrorResponse(ctx context.Context, Message string) *Response {
	return &Response{
		Data:      nil,
		Message:   Message,
		Success:   false,
		RequestId: utils.GetRequestId(ctx),
	}
}

func (res *Response) NewSuccessResponse(ctx context.Context, Message string, data interface{}) *Response {
	return &Response{
		Data:      data,
		Message:   Message,
		Success:   true,
		RequestId: utils.GetRequestId(ctx),
	}
}
