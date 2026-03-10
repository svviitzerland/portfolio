---
title: 'I Built LLMAO: A Faster LLM Orchestrator in Rust'
date: '2026-01-02'
description: 'Started as a fix for rate limits. Ended up building something way faster than expected.'
tags: ['rust', 'python', 'llm', 'opensource']
---

<script>
    import GitHubCard from '$lib/components/GitHubCard.svelte';
</script>

# I Built LLMAO

Okay, so let me tell you how this happened.

## It Started with LiteLLM

A few weeks ago, I was digging through LiteLLM's codebase. Found a config issue, [submitted a fix](https://github.com/BerriAI/litellm/pull/18284#issuecomment-3677708531). Nothing crazy, just wanted to help out.

But that got me thinking.

## The Real Problem

Here's the thing. As a student, I use free tier API keys a lot. And so do my friends. We're always hitting rate limits. Like, constantly.

You know that feeling when you're in the middle of working on something, and suddenly... *429 Too Many Requests*. Annoying, right?

So we'd swap API keys manually. Or wait. Or just give up for the day.

There had to be a better way.

## Building LLMAO

I decided to build my own orchestrator. Something simple. Something that would:

1. **Rotate API keys automatically** when one hits the limit
2. **Work with multiple providers** (OpenAI, Anthropic, Groq, Cerebras, all of them)
3. **Be fast**. Actually fast. I wrote it in Rust because... well, I wanted to learn Rust. And Python bindings are easy with PyO3.

The result? **LLMAO** - Lightweight LLM API Orchestrator.

## The Surprise

Here's where it gets interesting.

I ran some benchmarks against LiteLLM. Expected maybe a small improvement since it's compiled code.

I was wrong.

![LLMAO vs LiteLLM Performance Benchmark](/images/posts/llmao/benchmark.png)

Look at those numbers:
- **Cold start**: ~100x faster
- **Memory usage**: ~7x less
- **Latency (p99)**: nearly half

I had to run the benchmarks multiple times because I didn't believe it at first.

## How It Works

Super simple. Install it:

```bash
pip install llmao-py
```

Use it:

```python
from llmao import LLMClient

client = LLMClient()
response = client.completion(
    model="groq/llama-3.1-70b-versatile",
    messages=[{"role": "user", "content": "Hello!"}]
)
```

Want multiple API keys? Just set them up:

```bash
export GROQ_API_KEY="key1"
export GROQ_API_KEY_2="key2"
export GROQ_API_KEY_3="key3"
```

LLMAO handles the rotation automatically. When one key hits the limit, it moves to the next one. No manual switching. No waiting.

## For My Fellow Students

If you're a student using free tier APIs... this is for you. Stack up those free API keys from different accounts, let LLMAO rotate through them, and focus on what you're actually building.

No more rate limit interruptions.

## Try It Out

The whole thing is open source. Check it out, break it, improve it, whatever you want.

<GitHubCard repo="svviitzerland/LLMAO" />

Would love to hear what you think. Found a bug? Open an issue. Have an idea? Pull requests welcome.

Let's make LLM development less painful.
