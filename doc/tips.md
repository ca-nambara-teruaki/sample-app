# backend
## 使用技術
* 言語
　* Go
* Web
  * Echo
* ORM
  * ent
* DB
  * PostgreSQL

## ディレクトリ構成
ディレクトリ構成はクリーンアーキテクチャを参考にしている。
具体的には以下の通り

```
backend/
├── config/
├── db/
├── handlers/
├── services/
├── repository/
└── ent/
    └── schema/
```

参考：
* https://github.com/kuma-coffee/go-clean-archi
* https://github.com/zett-8/go-clean-echo/tree/master


## ER図
1. ER図を作成する
2. entでエンティティを作成する
3. repository層のインターフェースを設計する
4. repository層の処理を実装する
5. repository層のテストを実装する
6. service層の処理を実装する


## シーケンス図



## テスト
