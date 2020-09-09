# Check Interface

![Test](https://github.com/riita10069/check_interface/workflows/Build/badge.svg)
![Release](https://github.com/riita10069/check_interface/workflows/Release/badge.svg)

どの関数にも実装されていないインターフェイスに対して、警告を出す静的解析ツールです。
インターフェイスに実装されているメソッドと実装のメソッドのシグネチャが微妙に違うなどで、
インターフェイスの先のメソッドが呼び出せないような時にぬるぽが出るのが嫌だったので、作成しました。

## Getting Started

### How to install

```
go get github.com/riita10069/check_interface/cmd/check_interface
```

### Usage Example

```
 go vet -vettool=`which check_interface` [package path]
```

複数のパッケージのチェックをしたい場合は、

```
 go vet -vettool=`which check_interface` [package path1] [package path2] [package path3]
```

のように書くことで実行可能です。

