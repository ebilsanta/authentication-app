terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  profile = "cs301"
  region  = "ap-southeast-1"
}

# VPC
resource "aws_vpc" "cs301" {
  cidr_block = "10.0.0.0/16"
}

# Internet Gateway
resource "aws_internet_gateway" "igw" {
  vpc_id = aws_vpc.cs301.id
}

# Public Subnet
resource "aws_subnet" "cs301_public_1a" {
  vpc_id                  = aws_vpc.cs301.id
  cidr_block              = "10.0.0.0/24"
  availability_zone       = "ap-southeast-1a"
  map_public_ip_on_launch = true
}

# Elastic IP
resource "aws_eip" "nat_eip" {
  domain = "vpc"
}

# NAT Gateway
resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.nat_eip.id
  subnet_id     = aws_subnet.cs301_public_1a.id
}

# Private Subnets
resource "aws_subnet" "cs301_private_1b" {
  vpc_id            = aws_vpc.cs301.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "ap-southeast-1b"
}

resource "aws_subnet" "cs301_private_1c" {
  vpc_id            = aws_vpc.cs301.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "ap-southeast-1c"
}

# Route Table, Routes and Corresponding Route Associations for Private Subnets for Internet Bound Traffic
resource "aws_route_table" "private_route_table" {
  vpc_id = aws_vpc.my_vpc.id
}

resource "aws_route" "private_route" {
  route_table_id         = aws_route_table.private_route_table.id
  destination_cidr_block = "0.0.0.0/0"
  nat_gateway_id         = aws_nat_gateway.nat_gateway.id
}

resource "aws_route_table_association" "private_subnet_association_1b" {
  subnet_id      = aws_subnet.cs301_private_1b.id
  route_table_id = aws_route_table.private_route_table.id
}

resource "aws_route_table_association" "private_subnet_association_1c" {
  subnet_id      = aws_subnet.cs301_private_1c.id
  route_table_id = aws_route_table.private_route_table.id
}

# Security Group for Authentication and Authorisation Service
resource "aws_security_group" "internal_services" {
  vpc_id = aws_vpc.cs301.id

  egress {
    from_port   = 0
    to_port     = 65535
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = -1
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group_rule" "allow_all_ssh" {
  security_group_id = aws_security_group.internal_services.id

  type        = "ingress"
  from_port   = 22
  to_port     = 22
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "allow_all_https_egress" {
  security_group_id = aws_security_group.internal_services.id

  type        = "egress"
  from_port   = 443
  to_port     = 443
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "allow_all_smtp_egress" {
  security_group_id = aws_security_group.internal_services.id

  type        = "egress"
  from_port   = 465
  to_port     = 465
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}


# Authorisation Instances
resource "aws_launch_configuration" "authorisation" {
  image_id        = "ami-0c94855ba95c71c99"
  instance_type   = var.authorisation_ec2_instance_type
  security_groups = [aws_security_group.internal_services.id]
  key_name        = var.key_pair

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "authorisation" {
  launch_configuration = aws_launch_configuration.authorisation.name
  min_size             = 4
  max_size             = 12
  desired_capacity     = 4
  vpc_zone_identifier  = [aws_subnet.cs301_private_1b, aws_subnet.cs301_private_1c]
  health_check_type    = "EC2"
  availability_zones   = ["ap-southeast-1b", "ap-southeast-1c"]
}


# Authentication Instances
resource "aws_kms_key" "ecs_log_key" {
  description             = "ecs_log_key"
  deletion_window_in_days = 7
}

resource "aws_cloudwatch_log_group" "authentication_cluster" {
  name = "authentication_cluster"
}

resource "aws_ecs_cluster" "authentication_cluster" {
  name = "authentication_cluster"

  configuration {
    execute_command_configuration {
      kms_key_id = aws_kms_key.ecs_log_key.arn
      logging    = "OVERRIDE"

      log_configuration {
        cloud_watch_encryption_enabled = true
        cloud_watch_log_group_name     = aws_cloudwatch_log_group.example.name
      }
    }
  }
}

resource "aws_launch_configuration" "authentication" {
  image_id        = "ami-0c94855ba95c71c99"
  instance_type   = var.authentication_ec2_instance_type
  security_groups = [aws_security_group.internal_services.id]
  key_name        = var.key_pair

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_autoscaling_group" "authentication" {
  launch_configuration = aws_launch_configuration.authentication.name
  min_size             = 4
  max_size             = 8
  desired_capacity     = 4
  vpc_zone_identifier  = [aws_subnet.cs301_private_1b, aws_subnet.cs301_private_1c]
  health_check_type    = "EC2"
  availability_zones   = ["ap-southeast-1b", "ap-southeast-1c"]
}