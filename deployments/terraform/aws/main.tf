resource "aws_ecr_repository" "repository" {
  name = "repository"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_key_pair" "app_key_pair" {
    key_name   = "go-demo"
    public_key = file("go-demo.pub")
}

resource "aws_eip" "app_ip" {
    count = 1
    instance = aws_instance.app_server.id
}

resource "aws_instance" "app_server" {
  ami           = "ami-01be94ae58414ab2e"
  instance_type = "t2.micro"
  key_name      = aws_key_pair.app_key_pair.key_name
}

resource "aws_db_instance" "app_db" {
  allocated_storage   = 10
  engine              = "postgres"
  instance_class      = "db.t3.micro"
  username            = var.db_username
  password            = var.db_password
  skip_final_snapshot = true
}
