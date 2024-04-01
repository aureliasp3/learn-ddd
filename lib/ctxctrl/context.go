package ctxctrl

import (
	"context"

	"gorm.io/gorm"
)

type _CtxKeyTx struct{}

var ctxKeyTx = _CtxKeyTx{}

func SetTx(ctx context.Context, tx *gorm.DB) context.Context {
	return context.WithValue(ctx, ctxKeyTx, tx)
}

func GetTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(ctxKeyTx).(*gorm.DB); ok {
		return tx
	}
	return nil
}
