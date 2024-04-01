package interceptor

import (
	"context"
	"log"

	"connectrpc.com/connect"
	"github.com/morikuni/failure"

	"learn-ddd/lib/errctrl"
)

func NewErrorResponseInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				log.Printf("%+v\n", err)
				return nil, getConnectError(err)
			}
			return res, nil
		}
	}
}

func getConnectError(err error) *connect.Error {
	code, ok := failure.CodeOf(err)
	if !ok {
		return connect.NewError(connect.CodeUnknown, err)
	}
	switch code {
	case errctrl.NotFound:
		return connect.NewError(connect.CodeNotFound, err)
	case errctrl.InvalidArgument:
		return connect.NewError(connect.CodeInvalidArgument, err)
	case errctrl.Internal:
		return connect.NewError(connect.CodeInternal, err)
	default:
		return connect.NewError(connect.CodeUnknown, err)
	}
}
