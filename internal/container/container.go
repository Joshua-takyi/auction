package container

import "log/slog"

type Container struct {
	Logger *slog.Logger
}

func NewContainer(logger *slog.Logger) (*Container, error) {
	return &Container{
		Logger: logger,
	}, nil
}
