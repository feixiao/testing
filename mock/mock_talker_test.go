package mock

import (
	"testing"
	"github.com/golang/mock/gomock"
	 mk "./mk"
)

func TestCompany_Meeting(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mk.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("这是自定义的返回值，可以是任意类型。")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
}
