# thesislint

![](./ss.png)

.texファイルに記述された文章の表現をチェックします。

## 使い方

### 1. textlintをインストールする

```
npm i textlint
npm i textlint-rule-max-ten textlint-rule-spellcheck-tech-word textlint-rule-no-mix-dearu-desumasu textlint-rule-ja-no-redundant-expression textlint-rule-no-doubled-joshi textlint-rule-no-double-negative-ja textlint-rule-ja-no-abusage textlint-rule-no-dropping-the-ra textlint-rule-no-doubled-conjunctive-particle-ga textlint-rule-no-doubled-conjunction textlint-rule-ja-hiragana-keishikimeishi textlint-rule-ja-hiragana-fukushi textlint-rule-ja-hiragana-hojodoushi textlint-rule-general-novel-style-ja textlint-rule-ja-no-mixed-period
```

### 2. texファイルをcloneしたディレクトリにコピーし、`go run thesislint.go`を実行する
texファイルが`thesislint.go`と同じディレクトリにないとチェックしてくれない仕様です。

```
go run thesislint.go
```

### 3. 自戒の念を込めて文章を修正する

`.tex`を`.md`に無理やり置き換えてtextlintするとチェックしてくれる。