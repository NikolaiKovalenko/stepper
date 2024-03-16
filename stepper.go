package stepper

import (
	"context"
	"time"
)

type Stepper interface {
	TaskHandler(name string, handler Handler, middlewares ...MiddlewareHandler) HandlerStruct
	Listen(ctx context.Context) error
	Publish(ctx context.Context, name string, data []byte, options ...PublishOption) error
	UpdateTask(ctx context.Context, name string, updatedLaunchAt time.Time) error
	RegisterJob(ctx context.Context, config *JobConfig, h JobHandler) HandlerStruct
	UseMiddleware(h MiddlewareHandler)
}
