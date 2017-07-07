## 表格驱动测试
​	TDT也叫表格驱动方法，有时也被归为关键字驱动测试(keyword-driven testing，是针对自动化测试的软件测试方法，它将创建测试程序的步骤分为规划及实现二个阶段。Go官方库中有些测试就使用了这种测试方法。

​	TDT中每个表格项就是一个完整的test case,包含输入和期望的输出，有时候还会加一些额外的信息比如测试名称。如果你发现你的测试中经常copy/paste操作，你就可以考虑把它们改造成TDT。测试代码就一块，但是可以测试表格中的每一项。

#### [gotests](https://github.com/cweill/gotests)

​	自动生成Table-Drived-Tests测试代码的工具。

+ 我们创建tdt目录，内容如下：

  ```go
  func Add(a, b int) int {
  	return a + b
  }

  func Min(nums []int) int {
  	var rst int
  	for i := range nums {
  		if rst > nums[i] {
  			rst = nums[i]
  		}
  	}
  	return rst
  }

  func Max(nums []int) int {
  	var rst int
  	for i := range nums {
  		if rst > nums[i] {
  			rst = nums[i]
  		}
  	}
  	return rst
  }
  ```

  ```
  // 生成上述代码的表格驱动测试代码，文件名为tdt_test.go
  gotests -all -w  tdt.go    
  ```

  以Add为例，分析生成的代码

  ```go
  func TestAdd(t *testing.T) {
    	// 需要传入的参数
  	type args struct {
  		a int
  		b int
  	}
  	tests := []struct {
  		name string		// 测试的名字
  		args args		// 需要传入的参数
  		want int		// 希望获取到的结果
  	}{
  	// TODO: Add test cases.
      // 添加自己的测试例子  
  	}
  	for _, tt := range tests {
  		t.Run(tt.name, func(t *testing.T) {
  			if got := Add(tt.args.a, tt.args.b); got != tt.want {
                	// 如果测试结果不是自己的预期，那么报错。
  				t.Errorf("Add() = %v, want %v", got, tt.want)
  			}
  		})
  	}
  }
  ```

  填充完上面的测试表格，完成测试即可。

#### 参考资料

+ [TableDrivenTests](https://github.com/golang/go/wiki/TableDrivenTests)
+ [关键字驱动测试](http://www.voidcn.com/blog/jaazure/article/p-4333140.html)