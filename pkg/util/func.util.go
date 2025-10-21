/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package util

import (
	"context"
	"fmt"
	"log"
)

func SafeFunc[Req any, Res any](ctx context.Context, req Req, fn func(ctx context.Context, req Req) (Res, error)) (res Res, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v", r)
			err = fmt.Errorf("internal server error")
		}
	}()
	return fn(ctx, req)
}

func SafeFuncNoReq[Res any](ctx context.Context, fn func(ctx context.Context) (Res, error)) (res Res, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic recovered: %v", r)
			err = fmt.Errorf("internal server error")
		}
	}()
	return fn(ctx)
}
