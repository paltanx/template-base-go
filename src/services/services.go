package services

import (
	"get-otp-go/src/utils"
)

type Container struct {
	Logger          utils.ILogger
	Environment     utils.IEnvironment
}

func NewContainer(
	logger utils.ILogger,
	env utils.IEnvironment,
) *Container {
	return &Container{
		logger,
		env,
	}
}
