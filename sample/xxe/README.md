# XXE のテンプレート
対応言語
* java 
* php
* python

## example
Specify xml file in `data` directory.

```xml
go run main.go --service java --template sample/xxe/ data/xxe/test.xml
```

Python 3.7.1以降、外部の一般的なエンティティはデフォルトで処理されません
https://docs.python.org/3/library/xml.html#defused-packages
