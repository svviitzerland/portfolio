variable "aws_region" {
  default = "ap-southeast-1"
}

variable "instance_type" {
  default = "t3.micro"
}

variable "dockerhub_image" {
  description = "Docker image for SSH portfolio"
  type        = string
}

variable "ssh_public_key" {
  description = "Public key for admin SSH access to EC2"
  type        = string
}
