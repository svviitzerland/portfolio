---
title: 'When Terraform Forked: The OpenTofu Story'
date: '2026-01-30'
description: 'HashiCorp changed a license. The community built OpenTofu. Here is what actually happened and why it keeps happening in open source.'
tags: ['terraform', 'opentofu', 'opensource', 'infrastructure']
---

# The Open Source Licensing Drama

You know what's funny? We keep seeing the same story play out over and over again in the open source world.

A company builds something amazing. The community loves it. Everyone uses it. Then one day... the license changes.

And suddenly, we have a fork.

## It's Happened Before

Remember Redis? Yeah, that super fast in-memory database everyone uses. In 2024, they switched from BSD to a proprietary license. The community wasn't happy. So what happened? [Valkey](https://valkey.io/) was born.

MySQL did it too. Oracle acquired them, and the community got nervous. Result? [MariaDB](https://mariadb.org/).

Elasticsearch? Same story. They changed their license in 2021. Amazon and the community responded with [OpenSearch](https://opensearch.org/).

See the pattern?

## Then Came Terraform

August 2023. HashiCorp announced they're changing Terraform's license from Mozilla Public License (MPL) to Business Source License (BSL).

For context, MPL is a true open source license. BSL? Not so much. It restricts commercial use. Basically, "you can use it, but don't compete with us."

The infrastructure-as-code community wasn't thrilled.

## Why Did HashiCorp Do This?

Look, I get it. Building and maintaining software costs money. HashiCorp saw cloud providers (especially AWS, GCP, Azure) offering managed Terraform services without contributing back.

They were essentially saying: "We built this. You're making money from it. We're not. That needs to change."

Fair point, honestly. But the execution? That's where things got messy.

## Enter OpenTofu

Within weeks of the announcement, the community rallied. The Linux Foundation stepped in. And [OpenTofu](https://opentofu.org/) was born.

A true open source fork of Terraform. MPL licensed. Community-driven.

Big names joined immediately: Spacelift, env0, Scalr, Gruntwork. Even cloud providers showed support.

## So... What Actually Changed?

Here's the thing that surprised me: **not much, actually**.

I've been using both. Let me show you.

### The Basics Still Work

Your existing Terraform code? It just works with OpenTofu.

```hcl
# This works in both Terraform and OpenTofu
terraform {
  required_version = ">= 1.0"
}

provider "aws" {
  region = "ap-southeast-1"
}

resource "aws_instance" "web" {
  ami           = "ami-0f74c08b8b5effa56"
  instance_type = "t3.micro"

  tags = {
    Name = "my-server"
  }
}
```

Same syntax. Same providers. Same state files (mostly).

### Installation

**Terraform:**
```bash
# Using tfenv
tfenv install 1.7.0
tfenv use 1.7.0
```

**OpenTofu:**
```bash
# Using tofuenv (yes, they made their own version manager)
tofuenv install 1.6.0
tofuenv use 1.6.0

# Or just use the binary
brew install opentofu
```

### Running Commands

This is where it gets interesting. OpenTofu kept the `terraform` command for compatibility:

```bash
# Both work with OpenTofu
tofu plan
terraform plan  # yep, this works too

tofu apply
terraform apply
```

They made migration dead simple. Smart move.

## What OpenTofu Actually Added

While keeping compatibility, OpenTofu started adding their own features:

### 1. State Encryption (Native)

```hcl
terraform {
  encryption {
    key_provider "pbkdf2" "mykey" {
      passphrase = var.state_passphrase
    }

    method "aes_gcm" "mymethod" {
      keys = key_provider.pbkdf2.mykey
    }

    state {
      method = method.aes_gcm.mymethod
    }
  }
}
```

Terraform requires Terraform Cloud for this. OpenTofu? Built-in.

### 2. Better Testing Framework

```hcl
# test.tftest.hcl
run "test_instance_creation" {
  command = plan

  assert {
    condition     = aws_instance.web.instance_type == "t3.micro"
    error_message = "Instance type should be t3.micro"
  }
}
```

More intuitive than Terraform's testing approach.

### 3. For-Each on Modules

```hcl
module "vpc" {
  for_each = toset(["dev", "staging", "prod"])
  
  source = "./modules/vpc"
  env    = each.key
}
```

This was a community request for years. OpenTofu shipped it.

## The Real Difference: Philosophy

Here's what actually matters.

**Terraform** is now a product. HashiCorp needs to make money. Features get prioritized based on business needs. Want the good stuff? Pay for Terraform Cloud.

**OpenTofu** is a project. Community-driven. Features get prioritized based on what users actually need. Everything stays open source.

Neither approach is wrong. They're just different.

## Should You Switch?

Honestly? Depends.

**Stick with Terraform if:**
- You're already using Terraform Cloud
- You want official HashiCorp support
- You don't care about the license
- Your company has a HashiCorp contract

**Consider OpenTofu if:**
- You want true open source
- You're building commercial tools around IaC
- You like community-driven development
- You want features like native state encryption

**The good news?** You don't have to decide right now. They're compatible. You can switch later if you want.

## My Take

I'm not switching. Yet.

Not because OpenTofu isn't good. It is. But because Terraform still works for what I'm building. The license change doesn't affect me personally.

But I'm watching OpenTofu closely. The community momentum is real. The features are compelling. And the philosophy aligns with what I believe open source should be.

Plus, having competition is healthy. It keeps both projects honest.

## The Pattern Continues

This won't be the last time we see this happen. Some company will build something amazing, change the license, and the community will fork it.

It's the circle of open source life.

Redis → Valkey  
MySQL → MariaDB  
Elasticsearch → OpenSearch  
Terraform → OpenTofu  

Who's next? MongoDB? Docker? We'll see.

## What This Means for You

If you're learning infrastructure-as-code right now, learn the concepts, not the tool. The syntax might change. The tool might fork. But the principles stay the same.

Understand state management. Learn about providers. Master modules. Get good at organizing infrastructure code.

Those skills transfer. Whether you're using Terraform, OpenTofu, Pulumi, or whatever comes next.

## Final Thoughts

The Terraform/OpenTofu split isn't a tragedy. It's not even that dramatic, really.

It's just... open source being open source.

A company made a business decision. The community made a different choice. Now we have options.

And honestly? Options are good.
