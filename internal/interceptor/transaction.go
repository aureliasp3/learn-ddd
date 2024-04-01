package interceptor

import (
	"context"

	"learn-ddd/lib/ctxctrl"

	"connectrpc.com/connect"
	"gorm.io/gorm"

	"learn-ddd/db"
)

func NewTransactionInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (res connect.AnyResponse, err error) {
			err = db.DB.Transaction(func(tx *gorm.DB) error {
				res, err = next(ctxctrl.SetTx(ctx, tx), req)
				if err != nil {
					return err
				}
				return nil
			})
			return res, err
		}
	}
}
