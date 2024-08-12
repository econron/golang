## 資料

https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/cloud-design-patterns/cloud-design-patterns.pdf

## 資料メモ

ワークフロー手順

- api-gatwayに対して乗車予約リクエスト
- 乗車予約リクエストを変換してkinesis data streamsへ
- kinesis data streamsに来たイベントをs3に保存する
- kinesis data streamsに来たイベントを変換および処理し、rdsに保存する
- 決済が完了するとkinesis data streamsに送信される
- 顧客が乗車するとルートと乗車履歴がrdsに入る

推奨ブログ記事

https://aws.amazon.com/jp/blogs/compute/new-for-aws-lambda-sqs-fifo-as-an-event-source/

## 推奨ブログ記事メモ

sqs fifoキューは標準キューと違う

- MessageGroupIDを複数指定すると、同じキューの中で複数のメッセージグループが有効になる
- MessageDeduplicationIDにより、5分間隔で重複メッセージを排除できる