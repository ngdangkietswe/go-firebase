/**
 * Author : ngdangkietswe
 * Since  : 10/25/2025
 */

package cache

import (
	"github.com/ngdangkietswe/swe-go-common-shared/cache"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	cache.NewRedisCache,
)
