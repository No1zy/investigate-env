# investigate-env
各検証環境を構築してテストするためのフレームワーク的な何か

デフォルトで対応している言語 (バージョンは全て最新版のもの)
* ruby
* Go
* Python
* PHP (including apache)
* Java
* Perl
* Nodejs

フレームワーク
* Ruby on Rails
* Spring Boot

# How to Use

## 基本的な使い方

以下のコマンドを実行すると、全てのサービスでペイロードがテストされます  
```sh
go run main.go [Payload]
```

サービスを指定する場合は`--service` オプションを指定してください  

```sh
go run main.go --service python [Payload]
```


## Use original template
テンプレートは自分で作成することができます  

例: pythonの場合

```sh
mkdir mytemplate
cd mytemplate
vim template.py
```

コマンドラインから渡すペイロードは`{{ .VAR }}` の形式でテンプレートに記述してください  

```
# template.py

url = '{{ .URL }}'

print(url)
```

指定した変数はconfig/var.goファイルのVariable構造体として定義してください  
**構造体のメンバ名とテンプレートの変数名は同じにしてください。**

```go
// config/var.go

type Variable struct {
    URL string
}
```

オリジナルのテンプレートを作るときはディレクトリ構成は以下のように**検証したいサービスのディレクトリを作成するように**してください

```
mytemplate
├── go
│   └── main.go
├── java
│   └── Main.java
├── node
│   └── main.js
├── perl
│   └── main.pl
├── php
│   └── main.php
├── python
│   └── main.py
├── rails
│   ├── Gemfile
│   ├── README.md
│   ├── Rakefile
│   ├── app
│   ...
│
└── ruby
    └── main.rb
```

オリジナルのテンプレートを指定する場合は`--template`を使用してください  
```sh
$ go run main.go --service python --template mytemplate http://example.com/
Attaching to investigate-env_python_1
python_1  | http://example.com/
Removing container...
```

ファイルを読み込む場合は`data`ディレクトリ以下にファイルを置いて`[Payload]`引数として指定してください  
```sh
$ go run main.go --service java --template sample/xxe data/xxe/test.xml
```
