## Mock

​	Mock测试就是在测试过程中，对于某些不容易构造或者不容易获取的对象，用一个虚拟的对象来创建以便测试的测试方法。这个虚拟的对象就是mock对象。mock对象就是真实对象在调试期间的代替品。

+ mock对象使用范畴

  真实对象具有不可确定的行为，产生不可预测的效果，真实对象很难被创建的、真实对象的某些行为很难被触发、真实对象实际上还不存在的（和其他开发小组或者和新的硬件打交道）等等。

+ 使用mock对象测试的关键步骤

  使用一个接口来描述这个对象。在产品代码中实现这个接口，在测试代码中实现这个接口，在被测试代码中只是通过接口来引用对象，所以它不知道这个引用的对象是真实对象，还是mock对象。

#### gomock

+ [gomock](https://github.com/golang/mock) golang/mock下面的mock包，目前还没有发布release版本。

+ 创建接口

  ```go
  type Talker interface {
      SayHello(word string)(response string)
  }
  ```

+ 生成mock程序

  ```shell
  mkdir mk
  mockgen -source=talker.go > mk/mock_Talker.go
  ```

  ```go
  // MockTalker is a mock of Talker interface
  type MockTalker struct {
  	ctrl     *gomock.Controller
  	recorder *MockTalkerMockRecorder
  }
  // MockTalkerMockRecorder is the mock recorder for MockTalker
  type MockTalkerMockRecorder struct {
  	mock *MockTalker
  }
  // NewMockTalker creates a new mock instance
  func NewMockTalker(ctrl *gomock.Controller) *MockTalker {
  	mock := &MockTalker{ctrl: ctrl}
  	mock.recorder = &MockTalkerMockRecorder{mock}
  	return mock
  }
  // EXPECT returns an object that allows the caller to indicate expected use
  func (_m *MockTalker) EXPECT() *MockTalkerMockRecorder {
  	return _m.recorder
  }
  // SayHello mocks base method
  func (_m *MockTalker) SayHello(word string) string {
  	ret := _m.ctrl.Call(_m, "SayHello", word)
  	ret0, _ := ret[0].(string)
  	return ret0
  }
  // SayHello indicates an expected call of SayHello
  func (_mr *MockTalkerMockRecorder) SayHello(arg0 interface{}) *gomock.Call {
  	return _mr.mock.ctrl.RecordCall(_mr.mock, "SayHello", arg0)
  }
  ```

+ 测试程序

  ```go
  func TestCompany_Meeting(t *testing.T) {
  	ctl := gomock.NewController(t)
    	// 创建一个mock对象
  	mock_talker := mk.NewMockTalker(ctl)
    	// 设置方法的返回值
  	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("这是自定义的返回值，可以是任意类型。")

    	// 传入mock对象
  	company := NewCompany(mock_talker)
  	t.Log(company.Meeting("王尼玛"))
  }
  ```

#### 参考资料

+ [用gomock进行mock测试](https://segmentfault.com/a/1190000009894570)
+ [Go单元测试进阶篇](http://blog.csdn.net/qian_xiaoqian/article/details/54344856)
+ [GOLANG最容易做测试MOCK](http://blog.csdn.net/win_lin/article/details/72967636)