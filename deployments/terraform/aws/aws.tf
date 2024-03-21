resource "aws_instance" "app_server" {
  ami           = "ami-01be94ae58414ab2e"
  instance_type = "t2.micro"
}

resource "aws_db_instance" "app_db" {
  allocated_storage   = 10
  engine              = "mysql"
  instance_class      = "db.t3.micro"
  username            = "go_user"
  password            = "go_password"
  skip_final_snapshot = true
}
