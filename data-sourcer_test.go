package cloud_computer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMockupDataSource_SetValue(t *testing.T) {
	t.Run("새로운 이름으로 저정할 경우, 변수가 저장된다", func(t *testing.T) {
		m := MockupDataSource{}
		err := m.SetValue(nil, "test", true)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("기존의 이름으로 저장할 경우, 변수가 업데이트된다", func(t *testing.T) {
		m := MockupDataSource{}
		err := m.SetValue(nil, "test", true)
		if err != nil {
			t.Error(err)
		}

		err = m.SetValue(nil, "test", false)
		if err != nil {
			t.Error(err)
		}

		assert.False(t, m.values["test"])
	})
}

func TestMockupDataSource_GetCurrentValue(t *testing.T) {
	t.Run("존재하지 않는 이름을 조회할 경우, False가 반환된다.", func(t *testing.T) {
		m := MockupDataSource{}
		v, err := m.GetCurrentValue(nil, "test")
		assert.NoError(t, err)
		assert.False(t, v)
	})

	t.Run("존재하는 이름을 조회할 경우, 변수가 반환된다", func(t *testing.T) {
		m := MockupDataSource{}
		m.values = map[string]bool{
			"test": true,
		}

		v, err := m.GetCurrentValue(nil, "test")

		assert.NoError(t, err)
		assert.True(t, v)
	})

	t.Run("setValue로 저장된 값이 getCurrentValue로 조회된다", func(t *testing.T) {
		m := MockupDataSource{}
		err := m.SetValue(context.TODO(), "test", true)
		if err != nil {
			t.Error(err)
		}

		v, err := m.GetCurrentValue(nil, "test")
		assert.NoError(t, err)
		assert.True(t, v)
	})
}

func TestMockupDataSource_GetInputChannel(t *testing.T) {
	t.Run("존재하지 않는 이름을 조회할 경우, 채널이 생성된다", func(t *testing.T) {
		m := MockupDataSource{}
		ch := m.GetInputChannel(context.TODO(), "test")
		assert.NotNil(t, ch)
	})

	t.Run("특정 채널의 output 채널에 값을 쓸 경우, 해당 이름을 가진 input channel로 데이터가 나온다.", func(t *testing.T) {
		m := MockupDataSource{}
		outCh := m.GetOutputChannel(context.TODO(), "test")
		inCh := m.GetInputChannel(context.TODO(), "test")

		outCh <- true
		v := <-inCh
		assert.True(t, v)
	})

	// TODO: input channel은 각각 1000개의 버퍼를 가지는데, 이 이상의 listener가 있을때 경우에 따라서는 데드락이 발생할 수 있음.
	t.Run("하나의 인풋에 대해 두 개 이상이 listen하는 경우, 모든 listener에게 값이 전달된다.", func(t *testing.T) {
		m := MockupDataSource{}
		outCh := m.GetOutputChannel(context.TODO(), "test")
		inCh := m.GetInputChannel(context.TODO(), "test")
		inCh2 := m.GetInputChannel(context.TODO(), "test")

		outCh <- true
		v2 := <-inCh2
		v := <-inCh
		assert.True(t, v)
		assert.True(t, v2)
	})
}
