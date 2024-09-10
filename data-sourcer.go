package cloud_computer

import "context"

type DataSourcer interface {
	GetInputChannel(ctx context.Context, name string) <-chan bool
	GetOutputChannel(ctx context.Context, name string) chan<- bool
	GetCurrentValue(ctx context.Context, name string) (bool, error)
	SetValue(ctx context.Context, name string, value any) error
	AddList(ctx context.Context, name string, value any) error
}
