# Keplars Email CLI

A command-line tool for sending transactional emails and managing your Keplars account.

## Cross-Platform Support

Works on all major operating systems:
- **Linux** (x64, ARM64)
- **macOS** (Intel, Apple Silicon)
- **Windows** (x64, ARM64)

## Installation

### Install Script (Recommended)

**Linux, macOS, and Windows (Git Bash/WSL):**

```bash
curl -fsSL https://keplars.com/install.sh | bash
```

Install a specific version:
```bash
curl -fsSL https://keplars.com/install.sh | bash -s -- 0.1.0
```

### Download Binaries

Download pre-built binaries from the [GitHub Releases](https://github.com/KeplarsHQ/cli/releases) page.

### Build from Source

```bash
cd cli
go build -o keplars
```

## Configuration

### Option 1: Config file (recommended)

```bash
keplars config set api-key kms_xxx.live_yyy
```

This saves your key to `~/.keplars/config.json` and is used automatically for all commands.

### Option 2: Environment variable

```bash
export KEPLARS_API_KEY="kms_xxx.live_yyy"
```

### Option 3: Flag per command

```bash
keplars --api-key kms_xxx.live_yyy send ...
```

**Resolution order (highest wins):** `--api-key` flag → `KEPLARS_API_KEY` env var → `~/.keplars/config.json`

## Commands

### `config` — Manage configuration

```bash
keplars config set api-key <key>       # Save API key to config file
keplars config set base-url <url>      # Save custom base URL
keplars config get                     # Print resolved config with source
keplars config delete                  # Remove config file
```

### `send` — Send an email

```bash
keplars send \
  --to user@example.com \
  --from hello@yourdomain.com \
  --subject "Hello" \
  --html "<h1>Hi!</h1>"
```

**Required:** `--to`, `--from`, `--subject`, and `--html` or `--text`

**Optional:** `--cc`, `--bcc`, `--reply-to`, `--timeout`, `--json`

### `status` — Get email status

```bash
keplars status <email-id>
keplars status <email-id> --json
```

### `contacts` — Manage contacts

```bash
keplars contacts add --email user@example.com --name "Alice" --audience-id <id>
keplars contacts get user@example.com
keplars contacts list --audience-id <id> --page 1 --limit 20
keplars contacts update user@example.com --name "Alice Smith"
keplars contacts delete user@example.com
```

### `audiences` — Manage audiences

```bash
keplars audiences create --name "Newsletter" --description "Weekly newsletter"
keplars audiences list --page 1 --limit 20
keplars audiences get <id>
keplars audiences delete <id>
```

### `automations` — Manage automations

```bash
keplars automations list
keplars automations get <id>
keplars automations enroll <id> --email user@example.com
keplars automations unenroll <id> --email user@example.com
```

### `domains` — Manage sending domains

```bash
keplars domains add yourdomain.com
keplars domains list
keplars domains status <domain-id>
keplars domains verify <domain-id>
keplars domains delete <domain-id>
keplars domains create-api-key <domain-id> --name "Production Key"
```

## Global Flags

- `--api-key`: Keplars API key
- `--base-url`: Custom API base URL
- `--help`: Help for any command
- `--version`: Version information

## Environment Variables

- `KEPLARS_API_KEY`: Your Keplars API key
- `KEPLARS_BASE_URL`: Custom API base URL (optional)

## Support

- GitHub: https://github.com/Swing-Technologies/keplars-sdk
- Email: support@keplars.com
