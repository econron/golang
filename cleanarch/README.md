## DDDをgoで実践するためのコンセプトプロジェクト

### ドメインモデリング

#### 対象サービス

英会話事業を行う企業

- 生徒
- 講師
- レッスン

基本的にこの３ドメインで構成される。
ここにビジネスロジックが幾分か絡んでくるとする。

レッスンを提供しAIによるレッスン分析を通じ、発音・文法・語彙・表現力　を評価しフィードバックするサービスにする。

#### フロー

#### 初期フロー

生徒
- 会員登録する
- 認証メールが来る
- 認証メールのリンクを開いて認証を完了する
- レッスンを予約する場合、決済情報を登録する

講師
- 講師が応募してくる
- 講師の応募情報を閲覧できるようにする
- 講師を面接する
- 講師が合格する
- 講師のレベルを独自判定する
- 講師登録をする
  - 銀行口座の番号も登録する

講師がレッスンスロットを開く
- レッスンスロットを開けるのは現在時刻より1時間前とする
- カスタマーサポート経由でブラックリストになってない講師とする
生徒がレッスンを予約する
- レッスンの予約が可能なのはスロットの開始時刻の30分前とする
- 講師にとってブラックリスト入りした生徒以外にする
レッスンを開始する
- 生徒はレッスンルームに5分前に入れる
- 講師も同様
  - 講師はどんなレッスンをすれば良いかのアドバイスをAIから受ける
- レッスンルームでは生徒と講師で教材を画面共有できる
レッスンの映像を分析しCEFR、改善点を算出する
- レッスン完了後、動画データを深層学習エンジンを搭載したソフトウェアに渡す
- ソフトウェアが解析処理を実行し、CEFRと改善点を算出する
生徒はレッスン終了後、改善点を参照し次のレッスンに活かす
- 生徒はレッスン完了から2分以内に改善レポートを閲覧可能
生徒、講師ともにブラックリスト判定をする
生徒はレッスンを評価する


### イベントストーミング

イベントを付箋に書き出す
イベントを並び替える
コマンドや集約・外部システムを追加する
ポリシーを追加する
全体の確認とディスカッション

### ドメインモデリング

講師 Entity
- 講師応募情報 
- 講師の銀行口座情報
- 講師の独自レベル
- 名前
- 年齢
生徒 Entity
- 名前
- 年齢
スロット Entity
- 講師のID
- 開始時刻
- 終了時刻
レッスン Entity
- スロットの情報
- 生徒のID
レッスンレポート Entity
- 生徒のID
- レッスンのID
- CEFR
- 改善アドバイス


### コーディング方針

- ドメインオブジェクトの生成をするのがリポジトリや外部サービスのデータなどである
- ハンドラーでリクエストバリデーション、ユースケース呼び出し、レスポンス生成をする

という感じ。

### エンドポイント一覧

### student

POST /student/register
POST /student/send_registering_email
POST /student/check_registering_email

### tutor

POST /tutor/apply
GET /tutor/application_info
POST /tutor/our_unique_level
POST /tutor/register

### slot & tutor

POST /tutor/{tutor_id}/slot/open

### slot & student

POST /student/{student_id}/slot/book

### lesson

POST /lesson/{lesson_id}/start
POST /lesson/{lesson_id}/end

### lesson_report

POST /lesson/{lesson_id}/lesson_report

student & lesson_report

GET /student/{student_id}/lesson_report
GET /student/{student_id}/lesson_report/{lesson_report_id}

