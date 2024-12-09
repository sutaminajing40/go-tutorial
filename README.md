# TODO アプリケーション

このプロジェクトは、クリーンアーキテクチャを採用した Golang ベースの TODO アプリケーションです。

## プロジェクト構造

```
.
├── .vscode/              # VSCode設定ファイル
├── backend/              # バックエンドのメインパッケージ
│   ├── cmd/              # アプリケーションのエントリーポイント
│   │   └── api/         # APIサーバーのメインパッケージ
│   │       └── main.go
│   ├── config/           # 設定関連のパッケージ
│   │   ├── domain/      # ドメインモデルとビジネスルール
│   │   ├── handler/     # HTTPハンドラー（プレゼンテーション層）
│   │   ├── repository/  # データアクセス層の実装
│   │   └── usecase/     # ユースケース（アプリケーション層）
│   ├── migrations/      # データベースマイグレーションファイル
│   ├── pkg/             # 公開可能な再利用可能なパッケージ
└── prompts/             # AIプロンプト用ファイル
```

## internal ディレクトリの詳細説明

### domain/

ビジネスロジックの中心となるエンティティとビジネスルールを定義します。

### handler/

HTTP リクエストの受け付けとレスポンスの整形を担当します。

- エンドポイントの定義
- リクエストのバリデーション
- レスポンスの生成

### repository/

データの永続化層を担当します。

### usecase/

アプリケーション固有のビジネスロジックを実装します。

- ドメインオブジェクトの操作
- トランザクション管理
- ビジネスルールの適用

## 技術スタック

- 言語: Go
- アーキテクチャ: クリーンアーキテクチャ
- データベース: GORM（ORM フレームワーク）

## 起動方法

### バックエンド

```
cd backend
go run cmd/api/main.go
```

### フロントエンド

```
cd frontend
npm install
npm run dev
```

## API ドキュメント

### API エンドポイント一覧

ベース URL: /api

### TODO 操作

| メソッド | エンドポイント | 説明                   | リクエスト                                  | レスポンス                            |
| -------- | -------------- | ---------------------- | ------------------------------------------- | ------------------------------------- |
| POST     | /todos         | TODO を作成            | `{ "title": string }`                       | TODO object                           |
| GET      | /todos         | 全 TODO を取得         | -                                           | TODO objects array                    |
| GET      | /todos/:id     | 指定 ID の TODO を取得 | -                                           | TODO object                           |
| PUT      | /todos/:id     | TODO を更新            | `{ "title": string, "completed": boolean }` | TODO object                           |
| DELETE   | /todos/:id     | TODO を削除            | -                                           | `{ "message": string, "id": number }` |

### レスポンス形式（TODO object）

{
"id": number,
"title": string,
"completed": boolean,
"created_at": string,
"updated_at": string
}

### エラーレスポンス

{
"error": string
}
