/**
 * Author : ngdangkietswe
 * Since  : 10/21/2025
 */

package cron

import (
	"context"
	"time"

	"github.com/ngdangkietswe/swe-go-common-shared/logger"
	"github.com/robfig/cron/v3"
	"go.uber.org/fx"
)

type CronProps struct {
	fx.In
	Lifecycle fx.Lifecycle
	Logger    *logger.Logger
	Cron      *cron.Cron
}

func NewCron() *cron.Cron {
	return cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.Recover(cron.DefaultLogger),
			cron.DelayIfStillRunning(cron.DefaultLogger),
		),
		cron.WithLocation(time.FixedZone("Asia/Ho_Chi_Minh", 7*3600)),
	)
}

func StartCron(props CronProps) {
	props.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			props.Logger.Info("Cronjob started")
			props.Cron.Start()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			props.Logger.Info("Cronjob stopped")
			newCtx := props.Cron.Stop()
			<-newCtx.Done()
			return nil
		},
	})
}

var Module = fx.Module("cron",
	fx.Provide(
		NewCron,
	),
	fx.Invoke(
		StartCron,
	),
)
