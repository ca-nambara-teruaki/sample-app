## entとは
* entは、Facebook（現Meta）が開発した、Go（Golang）向けのオープンソースのORM（オブジェクトリレーショナルマッピング）ライブラリ
* Goコードでデータベースのスキーマ（構造）を定義し、その定義に基づいてデータベースとやり取りするコードを自動生成できる

https://entgo.io/ja/docs/getting-started

## entでUserエンティティを作成する
### 準備
1. Go mod init, Postgres起動
```
cd entdemo
go mod init entdemo
docker-compose up -d
```

### entの設定
2. Userエンティティの作成
  * `ent/schema/user.go` が作成される
```
go run -mod=mod entgo.io/ent/cmd/ent new User
```

3. Userエンティティのフィールドを定義
  * `ent/schema/user.go` の `func (User) Fields()` 内を編集

4. アセットの作成
  * `ent/user` にアセットが作成される
```
go generate ./ent
```

### マイグレーション
5. マイグレーション
  * `entdemo/user.go` にマイグレーション処理 `Schema.Create` を記述する
```
dsn := "host=" + cfg.Host + " port=" + cfg.Port + " user=" + cfg.User + " dbname=" + cfg.DBName + " password=" + cfg.Password + " sslmode=" + cfg.SSLMode
client, err := ent.Open("postgres", dsn)
client.Schema.Create(context.Background())
```
  
  * 以下の通り実行してマイグレーション
```
go mod tidy
go run user.go
```

6. テーブル確認 
* 以下のコマンドでテーブルの内容を確認する。
  * この時点ではテーブルは作成されているが、データはない。
```
docker exec -it postgres.local psql -U admin -d sampledb -c "\dt"
       List of relations
 Schema | Name  | Type  | Owner
--------+-------+-------+-------
 public | users | table | admin
(1 row)

docker exec -it postgres.local psql -U admin -d sampledb -c "select * from users;"
 id | age | name 
----+-----+------
(0 rows)
```

### User登録
7. User登録
* `entdemo/user.go` にUser作成処理 `client.User.Create()` を記述する
  * これは、4. アセットの作成で作成されたメソッド
```
client.User.Create().SetAge(30).SetName("Alice").Save(ctx)
```

* 以下の通り実行してUserを登録する。

```
go mod tidy
go run user.go
```

8. テーブル確認 
* 以下のコマンドでテーブルの内容を確認する。
  * Userテーブルにレコードが追加されている
```
docker exec -it postgres.local psql -U admin -d sampledb -c "select * from users;"
 id | age | name
----+-----+-------
  1 |  30 | Alice
(1 row)
```
