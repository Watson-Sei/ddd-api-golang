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
    - usecases
```

- domain: ドメイン層を示すフォルダ
  - models: アプリケーションのドメインモデルが格納されるフォルダ
  - repositories: リポジトリインターフェイスが格納されるフォルダ
  - usecase: usecaseが格納されるフォルダ

