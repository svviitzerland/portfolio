data "aws_ami" "al2023" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["al2023-ami-*-x86_64"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }
}

resource "aws_key_pair" "admin" {
  key_name   = "ssh-portfolio-admin"
  public_key = var.ssh_public_key
}

resource "aws_instance" "ssh_portfolio" {
  ami                    = data.aws_ami.al2023.id
  instance_type          = var.instance_type
  key_name               = aws_key_pair.admin.key_name
  vpc_security_group_ids = [aws_security_group.ssh_portfolio.id]

  user_data = templatefile("${path.module}/user-data.sh", {
    dockerhub_image = var.dockerhub_image
  })

  tags = {
    Name = "ssh-portfolio"
  }
}
