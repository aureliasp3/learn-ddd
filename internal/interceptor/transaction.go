package interceptor

import (
	"context"

	"connectrpc.com/connect"
	"gorm.io/gorm"

	"learn-ddd/db"
	"learn-ddd/internal/common"
)

func NewTransactionInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (res connect.AnyResponse, err error) {
			err = db.DB.Transaction(func(tx *gorm.DB) error {
				res, err = next(common.SetTx(ctx, tx), req)
				if err != nil {
					return err
				}
				return nil
			})
			return res, err
		}
	}
}
