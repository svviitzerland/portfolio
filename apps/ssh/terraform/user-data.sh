#!/bin/bash
set -e

# ── 1. Move SSHD to port 2222 (free up port 22 for container) ──
sed -i 's/^#Port 22/Port 2222/' /etc/ssh/sshd_config
sed -i 's/^Port 22$/Port 2222/' /etc/ssh/sshd_config

# Allow SELinux to use port 2222 for SSH
if command -v semanage &> /dev/null; then
  semanage port -a -t ssh_port_t -p tcp 2222 || true
fi

systemctl restart sshd

# ── 2. Install and start Docker ──
yum update -y
yum install -y docker
systemctl enable docker
systemctl start docker

# Add ec2-user to docker group
usermod -aG docker ec2-user

# ── 3. Run SSH portfolio container on port 22 ──
docker volume create ssh_keys

docker run -d \
  --name ssh-portfolio \
  --restart always \
  -p 22:2222 \
  -v ssh_keys:/app/.ssh \
  ${dockerhub_image}
