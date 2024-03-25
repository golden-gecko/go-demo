resource "aws_db_instance" "tutorial_database" {
  allocated_storage      = 10
  engine                 = "mysql"
  engine_version         = "8.0.36"
  instance_class         = "db.t3.micro"
  db_name                = "tutorial"
  username               = var.db_username
  password               = var.db_password
  db_subnet_group_name   = aws_db_subnet_group.tutorial_db_subnet_group.id
  vpc_security_group_ids = [aws_security_group.tutorial_db_sg.id]
  skip_final_snapshot    = true
}

resource "aws_instance" "tutorial_web" {
  count                  = var.settings.web_app.count
  ami                    = "ami-023adaba598e661ac"
  instance_type          = "t2.micro"
  subnet_id              = aws_subnet.tutorial_public_subnet[count.index].id
  key_name               = aws_key_pair.tutorial_kp.key_name
  vpc_security_group_ids = [aws_security_group.tutorial_web_sg.id]
}
