# Keplars CLI

Homebrew tap for the official Keplars CLI — send emails, manage contacts, audiences, automations, and domains from your terminal.

## Install

```bash
brew tap KeplarsHQ/cli
brew install keplars
```

Or in one line:

```bash
brew install KeplarsHQ/cli/keplars
```

## Upgrade

```bash
brew upgrade keplars
```

## Quick start

```bash
keplars config set api-key kms_xxx.live_yyy

keplars send \
  --to user@example.com \
  --from hello@yourdomain.com \
  --subject "Hello" \
  --text "Hello from Keplars"
```

Full docs at [keplars.com](https://keplars.com) · [CLI reference](https://github.com/KeplarsHQ/cli)
