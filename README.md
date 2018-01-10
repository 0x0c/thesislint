# thesislint

![](./ss.png)

.texファイルに記述された文章の表現をチェックします。

## 使い方

### 1. textlintをインストールする

```
npm i
```

### 2. texファイルをcloneしたディレクトリにコピーし、`go run thesislint.go`を実行する
texファイルが`thesislint.go`と同じディレクトリにないとチェックしてくれない仕様です。

```
go run thesislint.go
```

### 3. 自戒の念を込めて文章を修正する

`.tex`を`.md`に無理やり置き換えてtextlintするとチェックしてくれました。

## チェック内容について
`.textlintrc`を見てください。`rules`以下の各フィールド名でググるとチェックする内容を解説したページが見つかると思います。

適宜ルールは有効無効を切り替えて使ってください。
