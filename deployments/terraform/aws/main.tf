resource "aws_ecr_repository" "repository" {
  name = "repository"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_vpc" "app_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
}

resource "aws_internet_gateway" "app_gateway" {
  vpc_id = aws_vpc.app_vpc.id
}

resource "aws_subnet" "app_public_subnet" {
  vpc_id     = aws_vpc.app_vpc.id
  cidr_block = "10.0.1.0/24"
}

resource "aws_subnet" "app_private_subnet" {
  vpc_id     = aws_vpc.app_vpc.id
  cidr_block = "10.0.2.0/24"
}

resource "aws_route_table" "app_public_route_table" {
  vpc_id = aws_vpc.app_vpc.id
}

resource "aws_route_table_association" "app_public_route_table_association" {
  route_table_id = aws_route_table.app_public_route_table.id
  subnet_id      = aws_subnet.app_public_subnet.id
}

resource "aws_route_table" "app_private_route_table" {
  vpc_id = aws_vpc.app_vpc.id
}

resource "aws_route_table_association" "app_private_route_table_association" {
  route_table_id = aws_route_table.app_private_route_table.id
  subnet_id      = aws_subnet.app_private_subnet.id
}

resource "aws_key_pair" "app_key_pair" {
  key_name   = "app_key_pair"
  public_key = file("go-demo.pub")
}

resource "aws_eip" "app_ip" {
  instance = aws_instance.app_server.id
}

resource "aws_security_group" "app_security_group" {
  name = "app_security_group"

  ingress {
    from_port   = "22"
    to_port     = "22"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = "443"
    to_port     = "443"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = "0"
    to_port     = "0"
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "db_security_group" {
  name = "app_security_group"

  ingress {
    from_port   = "5432"
    to_port     = "5432"
    protocol    = "tcp"
    cidr_blocks = [aws_security_group.app_security_group.id]
  }
}

resource "aws_db_instance" "db_server" {
  allocated_storage   = 10
  engine              = "postgres"
  instance_class      = "db.t3.micro"
  username            = var.db_username
  password            = var.db_password
  skip_final_snapshot = true
}

resource "aws_instance" "app_server" {
  ami           = "ami-01be94ae58414ab2e"
  instance_type = "t2.micro"
  key_name      = aws_key_pair.app_key_pair.key_name
}
