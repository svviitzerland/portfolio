# SSH Portfolio

An interactive portfolio you access through your terminal. Built with Go, using Wish and Bubble Tea.

```
ssh ssh.farhanaulianda.my.id
```

That's it. Just SSH.

## How it works

The app runs inside a Docker container on an EC2 instance. When you SSH in, you get a TUI (terminal user interface) that lets you browse through my profile, experience, and projects interactively.

Built with:
- [Wish](https://github.com/charmbracelet/wish) - SSH server framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- Docker for deployment
- Terraform for infrastructure
- GitHub Actions for CI/CD

## Running locally

```
go run .
```

Then connect:

```
ssh -p 2222 localhost
```

Or with Docker:

```
docker build -t ssh-portfolio .
docker run -p 2222:2222 ssh-portfolio
```

## Deployment

Push to `main` triggers GitHub Actions which builds the Docker image, pushes to Docker Hub, and deploys to EC2 via SSH.

There's also a `terraform/` folder if you want to provision the EC2 instance with Terraform. Totally optional though - you can set up the server however you like. As long as Docker is installed and the container is running, it'll work.

## A note on security

If you're deploying something like this on AWS, here's what I'd recommend:

The EC2 admin SSH runs on a non-standard port (2222), separate from the portfolio on port 22. This works fine, but ideally you don't want that port open to the internet at all.

A better setup is using Tailscale or Cloudflare Tunnel for admin access. Install Tailscale on the EC2 instance, and you can SSH into it over your private Tailscale network without exposing port 2222 publicly. Same idea with Cloudflare Tunnel using `cloudflared` - it creates an outbound tunnel so you don't need any inbound ports for admin.

For CI/CD (GitHub Actions deploying via SSH), you'd either run a self-hosted runner inside your Tailscale network, or use Tailscale's GitHub Action to join the network temporarily during deployment. Cloudflare Tunnel works similarly with their `cloudflared` action.

The point is: only port 22 (the portfolio) needs to be public. Everything else can stay private.
