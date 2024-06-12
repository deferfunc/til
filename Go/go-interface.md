GO 言語のインターフェースについて

* 特定のメソッドシグネチャのセットを定義する型システムの一部
* インターフェースを明示的に実装する必要はない
* ある型がインターフェースに定義されたすべてのメソッドを持っていれば、自動的にそのインターフェースを実装しているとみなされる
    * Person 型が Talker インターフェースを実装する例


```go
type Talker interface {
    Talk() string
}

type Person struct {
    Name string
}

func (p Person) Talk() string {
    return "Hello, my name is " + p.Name
}
```

使い方

```go
func Greet(t Talker) {
    fmt.Println(t.Talk())
}
```