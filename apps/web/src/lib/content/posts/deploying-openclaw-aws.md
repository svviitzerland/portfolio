---
title: 'Running Your Own AI Assistant on AWS with OpenClaw'
date: '2026-03-20'
description: 'I deployed OpenClaw on AWS to run my own AI assistant. Here is how I did it and why it was worth it.'
tags: ['aws', 'ai','infrastructure', 'openclaw']
---

<script>
    import GitHubCard from '$lib/components/GitHubCard.svelte';
</script>

# Why I Wanted My Own AI Assistant

So here's the thing. I use Claude, ChatGPT, Gemini, pretty much all of them on a daily basis. They're great tools, genuinely. But at some point I started thinking about where all my conversations actually live. Every prompt, every response, every piece of context sits on someone else's servers. I can't move it. I can't switch providers mid-conversation. I can't plug it into my own workflows the way I want.

For casual use that's totally fine. But I was building projects that needed AI capabilities integrated into my own infrastructure, and relying entirely on third-party web apps didn't feel right anymore. I wanted something I could run myself, on my own servers, with full control over the models and the data.

That's when I found OpenClaw.

<div style="display: flex; justify-content: center; margin: 2rem 0;">
  <svg viewBox="0 0 120 120" fill="none" xmlns="http://www.w3.org/2000/svg" width="80" height="80">
    <path d="M60 10 C30 10 15 35 15 55 C15 75 30 95 45 100 L45 110 L55 110 L55 100 C55 100 60 102 65 100 L65 110 L75 110 L75 100 C90 95 105 75 105 55 C105 35 90 10 60 10Z" fill="url(#lobster-grad)" />
    <path d="M20 45 C5 40 0 50 5 60 C10 70 20 65 25 55 C28 48 25 45 20 45Z" fill="url(#lobster-grad)" />
    <path d="M100 45 C115 40 120 50 115 60 C110 70 100 65 95 55 C92 48 95 45 100 45Z" fill="url(#lobster-grad)" />
    <path d="M45 15 Q35 5 30 8" stroke="#ff6b6b" stroke-width="2" stroke-linecap="round" />
    <path d="M75 15 Q85 5 90 8" stroke="#ff6b6b" stroke-width="2" stroke-linecap="round" />
    <circle cx="45" cy="35" r="6" fill="#050810" />
    <circle cx="75" cy="35" r="6" fill="#050810" />
    <circle cx="46" cy="34" r="2" fill="#00e5cc" />
    <circle cx="76" cy="34" r="2" fill="#00e5cc" />
    <defs>
      <linearGradient id="lobster-grad" x1="0%" y1="0%" x2="100%" y2="100%">
        <stop offset="0%" stop-color="#ff4444" />
        <stop offset="100%" stop-color="#cc2200" />
      </linearGradient>
    </defs>
  </svg>
</div>

## What OpenClaw Actually Is

OpenClaw is a self-hosted AI assistant platform. You install it on your own server, connect it to whatever AI model you want, and interact with it through messaging apps you already use. WhatsApp, Telegram, Slack, Discord, and a bunch of others. So instead of going to a web app, you just chat with your AI assistant directly from your phone or your team's Slack workspace.

The project has an interesting history too. It went through several name changes, from Warelay to Clawd to Molty, before finally landing on OpenClaw. The rename happened after Anthropic reached out about trademark concerns with "Clawd", and the team migrated everything in about three hours. Different name, same project. If you're curious about the full story, there's [the official lore](https://docs.openclaw.ai/start/lore) which is genuinely entertaining.

There are also lighter community forks like [NanoClaw](https://github.com/qwibitai/nanoclaw), [PicoClaw](https://github.com/sipeed/picoclaw), [ZeroClaw](https://github.com/zeroclaw-labs/zeroclaw), and many others. 

## Deploying It on AWS

For this deployment, I used an EC2 instance instead of Lightsail because I wanted the flexibility to manage my own security groups, networking, and instance types. I went with a `t4g.medium` which is ARM-based and gives you a good balance of performance and cost for running something like OpenClaw.

One thing I want to mention here is AWS Spot Instances. They're spare compute capacity that AWS sells at a steep discount, sometimes up to 90% cheaper than On-Demand pricing. The tradeoff is that AWS can reclaim the instance with a two-minute warning if they need the capacity back. For something like an AI assistant that you can quickly redeploy, Spot Instances are perfect. I use them in my setup and the savings are significant.

### Setting Up the Server

You start by creating an EC2 instance. You can just use the default VPC that AWS provides, no need to set up a custom one unless you have specific networking requirements. Make sure the security group allows SSH on port 22 and web traffic on port 8080. If you want a static IP, attach an Elastic IP to the instance so the address doesn't change when you reboot.

### Installing OpenClaw

Once I SSH'd in, OpenClaw has a one-liner installer that handles everything:

```bash
curl -fsSL https://openclaw.ai/install.sh | bash
```

If you prefer to configure things manually, you can skip the interactive onboarding:

```bash
curl -fsSL https://openclaw.ai/install.sh | bash -s -- --no-onboard
```

### Choosing an AI Provider

OpenClaw doesn't ship with its own model, so you connect it to whatever provider you prefer. In my case I went with OpenRouter because it lets me swap between Claude 4.6, GPT-5, Gemini, and others with just a config change. If you're already deep in the AWS ecosystem, Bedrock is another solid option. You can also plug in API keys directly from OpenAI, Anthropic, or Google.

```bash
openclaw config set env.OPENROUTER_API_KEY 'your-api-key'
openclaw config set agents.defaults.model.primary 'openrouter/anthropic/claude-sonnet-4'
```

### Getting It Online

By default OpenClaw only listens on localhost, so you need to open it up:

```bash
openclaw config set gateway.bind lan
```

I also set up Telegram integration since I wanted to chat with the assistant from my phone. You grab a bot token from [@BotFather](https://t.me/botfather) and add it to the config:

```bash
openclaw config set channels.telegram.enabled true
openclaw config set channels.telegram.botToken 'your-bot-token'
```

Then restart the gateway and you're done:

```bash
openclaw gateway restart
```

Your instance should be live at `http://your-instance-ip:8080`.

## Automating Everything

The steps above are straightforward, but I also packaged the entire deployment into a repository that automates all of it. Infrastructure provisioning, OpenClaw installation, configuration, Spot Instance setup, everything runs through a single deploy command.

## Why I Think This Is Worth Doing

Running your own AI assistant gives you something the hosted platforms can't: complete ownership. You pick the models, you decide where the data lives, you control who has access. When I want to experiment with a different LLM, I change one line in a config file instead of signing up for another service.

It's not for everyone, and the hosted platforms are excellent for what they are. But if you're the kind of person who likes to own your infrastructure and understand exactly how things work under the hood, this is a really satisfying project.

## The Repository

I've open-sourced my complete setup. Everything you need to deploy OpenClaw on AWS is in here.

<GitHubCard repo="svviitzerland/AWS-OpenClaw" />

Clone it, deploy it, tweak it for your own needs. The repo includes the full infrastructure code, automated setup scripts, example configs, and a Makefile for common operations.
