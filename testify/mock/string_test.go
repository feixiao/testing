package mock

import (
	"testing"
	"./mocks"
)

func TestString(t *testing.T) {
	testObj := new(mocks.Stringer)
	// 设置String方法的返回值
	testObj.On("String").Return("Hello Mock Testing")

	// assert that the expectations were met

	t.Log(testObj.String())
	testObj.AssertExpectations(t)
}