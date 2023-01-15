# ddd-api-golang

API server configuration for Go language employing DDD

## DDDアーキテクチャの実装順序について

1. ドメイン層を設計する
ドメインモデルやインターフェイスなどの設計が含まれます。

ドメイン層のディレクトリ構成:

```
- domain
    - models
    - repositories
```

- domain: ドメイン層を示すフォルダ
  - models: アプリケーションのドメインモデルが格納されるフォルダ
  - repositories: リポジトリインターフェイスが格納されるフォルダ

インフラ層のディレクトリ構成:
```
- infrastructure
    - database
        - gorm
```

データベースや外部APIなどと対話するコードを格納する場所（外部システムとのアダプタ）

アプリケーション層
```
- application
    - handlers
    - usecases
```

- handlers: ルーティング、リクエスト・レスポンスを処理する
- usecases: ビジネスロジックを処理する

