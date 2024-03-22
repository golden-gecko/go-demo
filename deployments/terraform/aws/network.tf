resource "aws_vpc" "tutorial_vpc" {
  cidr_block           = var.vpc_cidr_block
  enable_dns_hostnames = true
}

resource "aws_internet_gateway" "tutorial_igw" {
  vpc_id = aws_vpc.tutorial_vpc.id
}

resource "aws_subnet" "tutorial_public_subnet" {
  // count is the number of resources we want to create
  // here we are referencing the subnet_count.public variable which
  // current assigned to 1 so only 1 public subnet will be created
  count = var.subnet_count.public

  // Put the subnet into the "tutorial_vpc" VPC
  vpc_id = aws_vpc.tutorial_vpc.id

  // We are grabbing a CIDR block from the "public_subnet_cidr_blocks" variable
  // since it is a list, we need to grab the element based on count,
  // since count is 1, we will be grabbing the first cidr block 
  // which is going to be 10.0.1.0/24
  cidr_block = var.public_subnet_cidr_blocks[count.index]

  // We are grabbing the availability zone from the data object we created earlier
  // Since this is a list, we are grabbing the name of the element based on count,
  // so since count is 1, and our region is us-east-2, this should grab us-east-2a
  availability_zone = data.aws_availability_zones.available.names[count.index]
}

resource "aws_subnet" "tutorial_private_subnet" {
  // count is the number of resources we want to create
  // here we are referencing the subnet_count.private variable which
  // current assigned to 2, so 2 private subnets will be created
  count = var.subnet_count.private

  // Put the subnet into the "tutorial_vpc" VPC
  vpc_id = aws_vpc.tutorial_vpc.id

  // We are grabbing a CIDR block from the "private_subnet_cidr_blocks" variable
  // since it is a list, we need to grab the element based on count,
  // since count is 2, the first subnet will grab the CIDR block 10.0.101.0/24
  // and the second subnet will grab the CIDR block 10.0.102.0/24
  cidr_block = var.private_subnet_cidr_blocks[count.index]

  // We are grabbing the availability zone from the data object we created earlier
  // Since this is a list, we are grabbing the name of the element based on count,
  // since count is 2, and our region is us-east-2, the first subnet should
  // grab us-east-2a and the second will grab us-east-2b
  availability_zone = data.aws_availability_zones.available.names[count.index]
}

resource "aws_route_table" "tutorial_public_rt" {
  // Put the route table in the "tutorial_vpc" VPC
  vpc_id = aws_vpc.tutorial_vpc.id

  // Since this is the public route table, it will need
  // access to the internet. So we are adding a route with
  // a destination of 0.0.0.0/0 and targeting the Internet 	 
  // Gateway "tutorial_igw"
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.tutorial_igw.id
  }
}

resource "aws_route_table_association" "public" {
  // count is the number of subnets we want to associate with
  // this route table. We are using the subnet_count.public variable
  // which is currently 1, so we will be adding the 1 public subnet
  count = var.subnet_count.public

  // Here we are making sure that the route table is
  // "tutorial_public_rt" from above
  route_table_id = aws_route_table.tutorial_public_rt.id

  // This is the subnet ID. Since the "tutorial_public_subnet" is a 
  // list of the public subnets, we need to use count to grab the
  // subnet element and then grab the id of that subnet
  subnet_id = aws_subnet.tutorial_public_subnet[count.index].id
}

resource "aws_route_table" "tutorial_private_rt" {
  // Put the route table in the "tutorial_VPC" VPC
  vpc_id = aws_vpc.tutorial_vpc.id

  // Since this is going to be a private route table, 
  // we will not be adding a route
}

resource "aws_route_table_association" "private" {
  // count is the number of subnets we want to associate with
  // the route table. We are using the subnet_count.private variable
  // which is currently 2, so we will be adding the 2 private subnets
  count = var.subnet_count.private

  // Here we are making sure that the route table is
  // "tutorial_private_rt" from above
  route_table_id = aws_route_table.tutorial_private_rt.id

  // This is the subnet ID. Since the "tutorial_private_subnet" is a
  // list of private subnets, we need to use count to grab the
  // subnet element and then grab the ID of that subnet
  subnet_id = aws_subnet.tutorial_private_subnet[count.index].id
}

// Create a db subnet group named "tutorial_db_subnet_group"
resource "aws_db_subnet_group" "tutorial_db_subnet_group" {
  // The name and description of the db subnet group
  name        = "tutorial_db_subnet_group"
  description = "DB subnet group for tutorial"

  // Since the db subnet group requires 2 or more subnets, we are going to
  // loop through our private subnets in "tutorial_private_subnet" and
  // add them to this db subnet group
  subnet_ids = [for subnet in aws_subnet.tutorial_private_subnet : subnet.id]
}

resource "aws_eip" "tutorial_web_eip" {
  count    = var.settings.web_app.count
  instance = aws_instance.tutorial_web[count.index].id
}
