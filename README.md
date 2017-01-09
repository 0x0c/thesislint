# thesislint

![](./ss.png)

.texファイルに記述された文章の表現をチェックします。

## 使い方

### 1. textlintをインストールする

```
npm i -g textlint
npm i -g textlint-rule-max-ten textlint-rule-spellcheck-tech-word textlint-rule-no-mix-dearu-desumasu ja-no-redundant-expression
```

### 2. pandocをインストールする

[ここ](http://pandoc.org)

### 3. texファイルがあるディレクトリで`go run thesislint.go`を実行する

```
go run thesislint.go
```

### 4. 自戒の念を込めて文章を修正する
