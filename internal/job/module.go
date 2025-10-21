/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package job

import "go.uber.org/fx"

var Module = fx.Module("job",
	fx.Invoke(
		RegisterTopicTechJob,
	),
	// TODO: Add more jobs here
)
