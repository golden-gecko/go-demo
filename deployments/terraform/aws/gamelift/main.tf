resource "aws_s3_bucket" "iwamori111_test_bucket" {
  bucket = "iwamori111-test-bucket"
}

resource "aws_s3_object" "test_build" {
  bucket = aws_s3_bucket.iwamori111_test_bucket.id
  key    = "build-1-0-0"
  source = "data/server"
  etag   = filemd5("data/server")
}

resource "aws_iam_role" "some_role" {
  name = "my_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

resource "aws_gamelift_build" "test" {
  name             = "example-build"
  operating_system = "AMAZON_LINUX_2"

  storage_location {
    bucket   = aws_s3_bucket.iwamori111_test_bucket.id
    key      = aws_s3_object.test_build.key
    role_arn = aws_iam_role.some_role.arn
  }
}

resource "aws_gamelift_fleet" "example" {
  build_id          = aws_gamelift_build.test.id
  ec2_instance_type = "t2.micro"
  fleet_type        = "ON_DEMAND"
  name              = "example-fleet-name"

  runtime_configuration {
    server_process {
      concurrent_executions = 1
      launch_path           = "/local/game/code/server"
    }
  }

  ec2_inbound_permission {
    from_port = 8080
    to_port   = 8080
    ip_range  = "0.0.0.0/0"
    protocol  = "TCP"
  }
}
