# biletojy
issue tracking system with tag

## コンセプト
* チケットの属性はタグとして表現される
* チケットにmarkdown, mermaid記法を採用する
* タグはグループを持つことができる
* タグは階層構造を持つことができる
* 認証/認可の機能を持たない
* データベースにはSQLite3を使用する
* 操作にショートカットキーが使える

## ショートカット
* `ctrl+n` チケット作成
* `ctrl+e` 表示中のチケット編集
* `ctrl+l` チケット一覧へ移動
* `ctrl+t` タグ一覧へ移動
* `ctrl+shift+n` タグ作成

## タグ
* タグの中間に:を含めた場合、そのタグの左辺はタググループとなる
* 同じタググループの複数のタグは、タグ付けする際にプルダウンメニューのように振る舞う
* タググループを使用することで、ステータス、カテゴリー、マイルストーンなどの属性を表現できる
* タグには色を指定できる
* タグの中間に/を含めた場合、そのタグは階層構造を持つと見なされる
* 階層構造を持つタグは、検索時にプルダウンメニューのように振る舞う
* タググループも末尾を@にした場合、日時を表すタグとして扱う  
  例えば、タグ登録/タグ編集時に `due-date@:` と入力すると、日付ピッカーが表示され日時を選択できる