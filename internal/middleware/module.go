/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package middleware

import "go.uber.org/fx"

var Module = fx.Provide(
	fx.Annotate(NewAuthMiddleware, fx.ResultTags(`group:"middlewares"`)),
	fx.Annotate(NewCORSMiddleware, fx.ResultTags(`group:"middlewares"`)),
	fx.Annotate(NewLoggerMiddleware, fx.ResultTags(`group:"middlewares"`)),
)
