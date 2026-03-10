output "public_ip" {
  value = aws_instance.ssh_portfolio.public_ip
}

output "instance_id" {
  value = aws_instance.ssh_portfolio.id
}
