# echo-crud
- Go 1.20.7
  - Echo 4.11.1
- MySQL 8.1.0

## 起動
```bash
$ docker compose up -d
```

## DB マイグレーション
- データベースの更新を行う
- 以下の場合に実行
  - 初回立ち上げ時
  - テーブル作成、削除があった場合
  - カラム追加、変更、削除があった場合

```bash
$ sh ./migrate.sh
```