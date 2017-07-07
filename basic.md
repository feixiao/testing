## Go单元测试基础
​	Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试。在包目录内，以_test.go为后缀名的源文件都是go test的一部分，而不是go build的构建部分。

```shell
# 查看帮助
go test -h
# 命令形式
go test [build/test flags] [packages] [build/test flags & test binary flags]
```

####  单元测试文件

+ 文件名必须是`_test.go`结尾的，这样在执行`go test`的时候才会执行到相应的代码。

- 你必须`import testing`这个包。
- 所有的测试用例函数必须是`Test`开头。
- 测试用例会按照源代码中写的顺序依次执行。
- 测试函数`TestX()`的参数是`testing.T`，我们可以使用该类型来记录错误或者是测试状态。
- 测试格式：`func TestXxx (t *testing.T)`,`Xxx`部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如`Testintdiv`是错误的函数名。
- 函数中通过调用`testing.T`的`Error`, `Errorf`, `FailNow`, `Fatal`, `FatalIf`方法，说明测试不通过，调用`Log`方法用来记录测试的信息。

#### 测试函数

+ 每个测试函数必须导入testing包，测试函数有如下签名：

  ```go
  func TestXXX(t *testing.T) {
      // ...
  }
  ```


+ 我们使用[《The Go Programming Language》](http://gopl.io/) 中的例子说明，测试代码我们放置在word目录下面。

  ```go
  // 被测试的函数是IsPalindrome，功能是判断传入的字符窜是否对称
  func TestPalindrome(t *testing.T) {
  	if !IsPalindrome("detartrated") {
  		t.Error(`IsPalindrome("detartrated") = false`)
  	}
  	if !IsPalindrome("kayak") {
  		t.Error(`IsPalindrome("kayak") = false`)
  	}
  }

  func TestNonPalindrome(t *testing.T) {
  	if IsPalindrome("palindrome") {
  		t.Error(`IsPalindrome("palindrome") = true`)
  	}
  }
  ```

+ 运行测试

  ```go
  go test ./word  	// 指定目录
  go test word		// 指定包名（包需要在GOPTAH/src或者GOROOT/src目录下面）

  cd word				// 进入包测试
  go test  

  // 输出结果
  PASS
  ok      _/home/frank/workspace/testing/word     0.004s
  ```

  注：参数`-v`可用于打印每个测试函数的名字和运行时间：

+ 添加异常测试

  ```go
  func TestCanalPalindrome(t *testing.T) {
      input := "A man, a plan, a canal: Panama"
      if !IsPalindrome(input) {
          t.Errorf(`IsPalindrome(%q) = false`, input)
      }
  }
  ```

  参数`-run`对应一个正则表达式，只有测试函数名被它正确匹配的测试函数才会被`go test`测试命令运行：

  ```shell
  go test -timeout 30s -tags -v -run ^TestCanalPalindrome$

  // 异常输出
  --- FAIL: TestCanalPalindrome (0.00s)
          word_test.go:23: IsPalindrome("A man, a plan, a canal: Panama") = false
  FAIL
  exit status 1
  FAIL    _/home/frank/workspace/testing/word     0.003s
  ```

#### 覆盖率测试

​	**语句的覆盖率**是指在测试中至少被运行一次的代码占总代码数的比例。我们使用`go test`命令中集成的测试覆盖率工具，来度量下面代码的测试覆盖率，帮助我们识别测试和我们期望间的差距。

+ 我们使用gopl中的eval包进行说明。我们放置在eval目录中，以表格驱动的测试进行。

  ```shell
  // 我们先保证测试可以通过
  go test -v
  // 生成覆盖率测试报告
  go test -coverprofile=c.out
  // 转换成html(需要在GOPATH/src目录下面)
  go tool cover -html=c.out -o coverage.html
  ```

#### 基准测试

​	基准测试是测量一个程序在固定工作负载下的性能。在Go语言中，基准测试函数和普通测试函数写法类似，但是以Benchmark为前缀名，并且带有一个`*testing.B`类型的参数；`*testing.B`参数除了提供和`*testing.T`类似的方法，还有额外一些和性能测量相关的方法。它还提供了一个整数N，用于指定操作执行的循环次数。

```
func BenchmarkIsPalindrome(b *testing.B) {
    for i := 0; i < b.N; i++ {
        IsPalindrome("A man, a plan, a canal: Panama")
    }
}
```

```go
go test -bench=. -run BenchmarkIsPalindrome
go test -bench=. -run BenchmarkIsPalindrome -benchmem
......
```

#### 参考资料

+ [《Go怎么写测试用例》](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md)
+ [《编写可测试的Go代码》](http://tabalt.net/blog/golang-testing/)
+ [《golang test说明解读》](http://www.cnblogs.com/yjf512/archive/2013/01/22/2870927.html)
+ [《Go单元测试》](https://rootsongjc.github.io/go-practice/docs/go_unit_test.html)
+ [《Go语言圣经》](http://books.studygolang.com/gopl-zh/)