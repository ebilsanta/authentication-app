variable "authentication_ec2_instance_type" {
  description = "AWS EC2 Instance Type for Authentication Service"
  type        = string
  default     = "t3.micro"
}

variable "authorisation_ec2_instance_type" {
  description = "AWS EC2 Instance Type for Authentication Service"
  type        = string
  default     = "t3.small"
}

variable "wazuh_ec2_instance_type" {
  description = "AWS EC2 Instance Type for Authentication Service"
  type        = string
  default     = "t3.medium"
}

variable "key_pair" {
  description = "Name of Key Pair for SSH"
  type        = string
  default     = "cs301-weibin"
}