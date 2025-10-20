/**
 * Author : ngdangkietswe
 * Since  : 10/18/2025
 */

package firebase

import "go.uber.org/fx"

var Module = fx.Provide(
	NewFirebaseClient,
	NewFCMClient,
	NewFAuthClient,
)
