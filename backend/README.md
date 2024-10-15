## DB起動

```
docker-compose up -d
```

## マイグレーション

```
# migration
go run cmd/migrate.go 

# テーブル確認
docker exec -it postgres.local psql -U admin -d sampledb -c "\dt"
       List of relations
 Schema | Name  | Type  | Owner
--------+-------+-------+-------
 public | tasks | table | admin

# クエリ実行
docker exec -it postgres.local psql -U admin -d sampledb -c "<query>"

# ex.
docker exec -it postgres.local psql -U admin -d sampledb -c "select * from tasks;"
 id | title | description | created_by | is_deleted 
----+-------+-------------+------------+------------
(0 rows)

```

### テスト
```
# キャッシュを削除
go clean -testcache

# test
go test -v ./...
```
