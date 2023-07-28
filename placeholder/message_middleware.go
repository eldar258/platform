package placeholder

import (
	"platform/config"
	"platform/pipeline"
	"platform/templates"
)

type SimpleMessageComponent struct {
	Message string
	config.Configuration
}

func (lc *SimpleMessageComponent) ImplementsProcessRequestWithServices() {}

func (lc *SimpleMessageComponent) Init() {
	lc.Message = lc.Configuration.GetStringDefault("main:message", "Default Message")
}

func (lc *SimpleMessageComponent) ProcessRequestWithServices(
	ctx *pipeline.ComponentContext,
	next func(*pipeline.ComponentContext),
	executor templates.TemplateExecutor) {
	err := executor.ExecTemplate(ctx.ResponseWriter,
		"simple_message.html", lc.Message)
	if err != nil {
		ctx.Error(err)
	} else {
		next(ctx)
	}
}
