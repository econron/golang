## レシピ

### AWS

https://docs.aws.amazon.com/ja_jp/prescriptive-guidance/latest/cloud-design-patterns/cloud-design-patterns.pdf

- api gateway + lambda + dynamoDB
- route53 + alb on public subnet + ecs on fargate on private subnet
- api gateway http proxy mode
- api gateway -> lambda -> kinesis data stream -> s3, lambda -> rds
- api gateway -> lambda -> sns topic -> lambda -> aurora
- api gateway -> lambda -> event bridge -> (lambda, sns) -> (aurora, email)
- ecs cluster has multiple services 

### パターン

- []腐敗防止
- []apiルーティング
- []サーキットブレーカー
- []イベントソーシング
- []六角形アーキテクチャ
- []pub-sub
- []バックオフ
- []Saga
- []散布図
- []ストラングラーフィグ
- []トランザクションアウトボックス
- []
- []
