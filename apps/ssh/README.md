# Farhan Aulianda - Terminal Portfolio (SSH)

This is the backend SSH server for my interactive portfolio, providing a unique command-line experience to explore my work.

## Live Access

- **Command:** `ssh ssh.farhanaulianda.my.id`

## Overview

Built with Go, utilizing the [Wish](https://github.com/charmbracelet/wish) SSH server framework and [Bubble Tea](https://github.com/charmbracelet/bubbletea) for a sophisticated Terminal User Interface (TUI). This application runs seamlessly inside a Docker container.

## Local Development

To run the SSH server locally:

```sh
go run .
```

Then connect via SSH on your local machine:

```sh
ssh -p 2222 localhost
```

### Docker

You can also build and run the server using Docker:

```sh
docker build -t ssh-portfolio .
docker run -p 2222:2222 ssh-portfolio
```

## Deployment & Infrastructure

Deployments are automated via GitHub Actions to an EC2 instance. The application is containerized and managed securely. There is also a `terraform/` folder provided for provisioning the EC2 infrastructure if desired.

### Security Note

For production environments, it is recommended to keep management ports private. For example, use Tailscale or Cloudflare Tunnels for secure admin/EC2 access, completely avoiding exposing administrative ports. Only the primary application SSH port needs to be public.
