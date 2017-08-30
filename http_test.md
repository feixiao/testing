## HTTP Test

+ Golang为我们提供了httptest包进行http请求和应答的测试，不需要发起真正的请求。
+ http测试分成两个部分:
  + 客户端测试：针对客户端测试的时候，我们需要自己构建MockServer进行应答。
  + 服务端测试：针对服务端测试的时候，我们需要自己构建MockClient发起请求。

#### MockServer

```go
func TestMockServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()

	t.Logf("ts.Url : %s", ts.URL)
	res, err := http.Get(ts.URL) // 注意：这边使用的是ts.URL
	if err != nil {
		t.Fatal(err)
	}
	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%s", greeting)
}
```

#### MockClient

```go
func TestMockClient(t *testing.T) {
	//  需要被测试的handler
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}

	req := httptest.NewRequest("GET", "http://example.com", nil)
	w := httptest.NewRecorder()

	// 调用处理逻辑
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
```

