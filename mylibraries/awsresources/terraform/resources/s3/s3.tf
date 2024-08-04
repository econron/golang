resource "aws_s3_bucket" "testing-bucket" {
    bucket = "testing-bucket-testtest1234"

    tags = {
        Name = "for testing"
        Environment = "Dev"
    }
}