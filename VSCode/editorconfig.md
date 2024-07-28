# このドキュメントについて

EditorConfig と MarkDown で不可解なフォーマットされていたので、調べた結果をまとめる

# 詳細

## 発生した事象

* EditorConfig for VSCode プラグインを入れる
* Format on Save を有効化する
* フォーマットを実行
    * 半角文字と全角文字の間にスペースが挿入される
    * `*` が `-   ` 等に置換される

## 解決方法

* `.prettierignore` ファイルを導入し `**/*.md` を指定して MarkDown ファイルを無視させた


## 調査内容

* [tats-u/prettier-plugin-md-nocjsp: Prettier Markdown Enhancement plugin - Don't insert spaces between alphanumeric and Han, Hiragana, or Katakana / Prettier Markdown修正プラグイン (英数字・漢字仮名間に半角スペースが挿入されないようになります)](https://github.com/tats-u/prettier-plugin-md-nocjsp)
    * これを入れるまでは無かったが、このプラグインの存在で割とよくある問題だと把握できた
* [Markdownを書くときはPrettierを窓から投げ捨てろ](https://zenn.dev/nullheader/articles/fb24ead4938c9a)
* [Ignoring Code · Prettier](https://prettier.io/docs/en/ignore.html)
