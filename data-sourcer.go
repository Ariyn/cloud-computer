package cloud_computer

import "context"

type DataSourcer interface {
	GetInputChannel(ctx context.Context, name string) <-chan bool
	GetCurrentValue(ctx context.Context, name string) (bool, error)
	GetOutputChannel(ctx context.Context, name string) chan<- bool
}
