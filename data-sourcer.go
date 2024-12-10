package cloud_computer

import (
	"context"
	"sync"
)

type DataSourcer interface {
	GetInputChannel(ctx context.Context, name string) <-chan bool
	GetOutputChannel(ctx context.Context, name string) chan<- bool
	GetCurrentValue(ctx context.Context, name string) (bool, error)
	SetValue(ctx context.Context, name string, value any) error
	AddList(ctx context.Context, name string, value any) error
}

type MockupDataSource struct {
	inputChannels  map[string][]chan bool
	values         map[string]bool
	outputChannels map[string]chan bool
	writeMutex     *sync.Mutex
}

func (m *MockupDataSource) GetInputChannel(ctx context.Context, name string) <-chan bool {
	if m.inputChannels == nil {
		m.inputChannels = make(map[string][]chan bool)
	}

	if _, ok := m.inputChannels[name]; !ok {
		m.inputChannels[name] = make([]chan bool, 0)
	}

	m.inputChannels[name] = append(m.inputChannels[name], make(chan bool, 1000))

	go func() {
		for {
			select {
			case <-ctx.Done():
				for _, ch := range m.inputChannels[name] {
					close(ch)
				}
				return
			}
		}
	}()

	return m.inputChannels[name][len(m.inputChannels[name])-1]
}

func (m *MockupDataSource) GetOutputChannel(ctx context.Context, name string) chan<- bool {
	if m.values == nil {
		m.values = make(map[string]bool)
	}
	if m.outputChannels == nil {
		m.outputChannels = make(map[string]chan bool)
	}
	if ch, ok := m.outputChannels[name]; ok {
		return ch
	}

	m.outputChannels[name] = make(chan bool, 1000)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case v := <-m.outputChannels[name]:
				if m.writeMutex == nil {
					m.writeMutex = &sync.Mutex{}
				}

				m.writeMutex.Lock()
				m.values[name] = v
				m.writeMutex.Unlock()

				if _, ok := m.inputChannels[name]; ok {
					for _, ch := range m.inputChannels[name] {
						ch <- v
					}
				}
			}
		}
	}()
	return m.outputChannels[name]
}

func (m *MockupDataSource) GetCurrentValue(ctx context.Context, name string) (bool, error) {
	if v, ok := m.values[name]; ok {
		return v, nil
	}

	return false, nil
}

func (m *MockupDataSource) SetValue(ctx context.Context, name string, value any) error {
	if m.values == nil {
		m.values = make(map[string]bool)
	}
	m.values[name] = value.(bool)
	return nil
}

func (m *MockupDataSource) AddList(ctx context.Context, name string, value any) error {
	return nil
}
