# brewManager
homebrewでインストールしたパッケージをSQLiteで管理するアプリ

## 使い方
1. frontendのビルド
```sh
cd frontend
npm install
npm run build
```

2. Goアプリのビルド
```sh
cd ..
go build -o bm
```

3. パスを通す
インストールしたディレクトリのパスを通し、ターミナルで`bm`コマンドが使えるようにします。
以下のコマンドを.zshrcなどのシェルの設定ファイルに追加してください。
```sh
echo "export PATH=\"$PATH:$(pwd)\"" >> ~/.zshrc
```

4. アプリの初期化
ターミナルで以下のコマンドを実行して、アプリを初期化します。
```sh
bm init
```

### コマンド一覧
- `bm init`: アプリの初期化
- `bm gui`: GUIの起動
- `bm add`: パッケージの追加
    - `bm add category <追加したいカテゴリ名>`: カテゴリの追加
    - `bm add package <追加したいパッケージ名> <カテゴリID>`: パッケージの追加
- `bm list <categories|packages>`: テーブルデータの一覧を表示
