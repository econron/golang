## どう実現するか？

api-gateway + lambda -> ride request with curl

lambda -> kinesis data streams

kinesis data streams -> s3

kinesis data streams -> save event to rds

## 個人的な疑問点

支払いが完了しない限り予約したことにしたくない

同じイベントの中に支払いがある・・・？

## 作るもの

api-gateway + lambda

lambda -> kinesis data streams

kinesis data streams -> s3

kinesis data streams -> lambda

lambda -> rds

の流れを作れるようにする

## リソースタイプについて

api-gateway

rest vs http

-> restで良いのではないか。

https://docs.aws.amazon.com/apigateway/latest/developerguide/http-api-vs-rest.html


記事を読む感じだと

https://zenn.dev/ncdc/articles/24c411930bc047

- http api -> ecs : OK!
- rest api -> ecs : NO!
- rest api -> nlb -> ecs : OK!

rest apiであればapi keyが利用可能