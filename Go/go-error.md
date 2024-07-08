go の error はインターフェース

```go
type error interface {
    Error() string
}
```

カスタムエラー型を作成する場合はこんな感じ

```go
package main

import (
    "fmt"
)

// MyError はカスタムエラー型です。
type MyError struct {
    Code int
    Msg  string
}

// Error メソッドは error インターフェースを実装します。
func (e *MyError) Error() string {
    return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Msg)
}

// 例: カスタムエラーの使用
func someFunction() error {
    // エラーを生成して返します。
    return &MyError{Code: 404, Msg: "Not Found"}
}

func main() {
    err := someFunction()
    if err != nil {
        fmt.Println(err)
    }
}
```