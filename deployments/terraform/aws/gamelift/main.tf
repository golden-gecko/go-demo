resource "aws_s3_bucket" "build" {
  bucket = "my-tf-test-bucket"
}

resource "aws_s3_object" "build" {
  bucket = "build"
  key    = "build-1-0-0"
  source = "data/server"
  etag   = filemd5("data/server")
}

resource "aws_iam_role" "test_role" {
  name = "test_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "s3:GetObject"
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
  operating_system = "AMAZON_LINUX"

  storage_location {
    bucket   = aws_s3_bucket.build.id
    key      = aws_s3_object.build.key
    role_arn = aws_iam_role.test_role.arn
  }
}

resource "aws_gamelift_fleet" "example" {
  build_id          = "build-yourgame-abc1def"
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
