provider "aws" {
  region = "eu-north-1"
}

resource "aws_vpc" "roberts_vpc" {
  cidr_block = "10.0.0.0/16"
}

resource "aws_internet_gateway" "roberts_gateway" {
  vpc_id = aws_vpc.roberts_vpc.id
}

resource "aws_subnet" "subnet_logger2_api" {
  vpc_id            = aws_vpc.roberts_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "eu-north-1a"
}

resource "aws_subnet" "subnet_logger2_logger" {
  vpc_id            = aws_vpc.roberts_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "eu-north-1b"
}

resource "aws_route_table" "public_table" {
  vpc_id = aws_vpc.roberts_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.roberts_gateway.id
  }
}

resource "aws_route_table_association" "api" {
  subnet_id      = aws_subnet.subnet_logger2_api
  route_table_id = aws_route_table.public_table
}

resource "aws_route_table_association" "logger" {
  subnet_id      = aws_subnet.subnet_logger2_logger
  route_table_id = aws_route_table.public_table
}
