# Check Interface
どの関数にも実装されていないインターフェイスに対して、警告を出す静的解析ツールです。
インターフェイスに実装されているメソッドと実装のメソッドのシグネチャが微妙に違うなどで、
インターフェイスの先のメソッドが呼び出せないような時にぬるぽが出るのが嫌だったので、作成しました。

## Getting Started

### How to install

```
go get github.com/riita10069/check_interface
```

### Usage Example

```
 go vet -vettool=`which check_interface` [package name]
```

複数のパッケージのチェックをしたい場合は、

```
 go vet -vettool=`which check_interface` [package name1] [package_name2] [package_name3]
```

のように書くことで実行可能です。

