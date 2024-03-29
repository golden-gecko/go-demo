resource "aws_key_pair" "tutorial_kp" {
  key_name   = "tutorial_kp"
  public_key = file("tutorial_kp.pub")
}

resource "aws_security_group" "tutorial_web_sg" {
  name        = "tutorial_web_sg"
  description = "Security group for tutorial web servers"
  vpc_id      = aws_vpc.tutorial_vpc.id

  ingress {
    description = "Allow all traffic through HTTP"
    from_port   = "80"
    to_port     = "80"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow SSH from my computer"
    from_port   = "22"
    to_port     = "22"
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all outbound traffic"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "tutorial_db_sg" {
  name        = "tutorial_db_sg"
  description = "Security group for tutorial databases"
  vpc_id      = aws_vpc.tutorial_vpc.id

  ingress {
    description     = "Allow MySQL traffic from only the web sg"
    from_port       = "3306"
    to_port         = "3306"
    protocol        = "tcp"
    security_groups = [aws_security_group.tutorial_web_sg.id]
  }
}
