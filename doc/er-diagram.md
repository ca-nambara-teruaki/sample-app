erDiagram

  TASK {
    int id PK "ID"
    string title "タイトル"
    string description "説明"
    enum status "ステータス"
    enum priority "優先度"
    datetime start_date "開始日"
    datetime due_date "期限"
    datetime updated_at "更新日時"
    int assigned_user_id FK "担当者"
    boolean is_deleted  "削除フラグ"
  }
