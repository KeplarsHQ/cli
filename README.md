# Keplars CLI

Official command-line interface for [Keplars](https://keplars.com) — send transactional emails, manage contacts, audiences, automations, and domains from your terminal.

## Installation

### macOS / Linux — Homebrew
```bash
brew tap KeplarsHQ/cli
brew install keplars
```

### macOS / Linux — Install script
```bash
curl -fsSL https://cli.keplars.com/install.sh | bash
```

### Windows — Scoop
```powershell
scoop bucket add keplars https://github.com/KeplarsHQ/cli
scoop install keplars
```

### Windows — PowerShell
```powershell
irm https://cli.keplars.com/install.ps1 | iex
```

## Setup

```bash
keplars config set api-key kms_xxx.live_yyy
```

Get your API key from [keplars.com/dashboard](https://keplars.com/dashboard).

## Usage

### Send an email
```bash
keplars send \
  --to user@example.com \
  --from hello@yourdomain.com \
  --subject "Welcome!" \
  --html "<h1>Hello!</h1>"
```

### Contacts
```bash
keplars contacts add --email user@example.com --name "Jane Doe"
keplars contacts list
keplars contacts get user@example.com
keplars contacts update user@example.com --name "Jane Smith"
keplars contacts delete user@example.com
```

### Audiences
```bash
keplars audiences create --name "Newsletter"
keplars audiences list
keplars audiences get <id>
keplars audiences delete <id>
```

### Automations
```bash
keplars automations list
keplars automations get <id>
keplars automations enroll <id> --email user@example.com
keplars automations unenroll <id> --email user@example.com
```

### Domains
```bash
keplars domains add yourdomain.com
keplars domains list
keplars domains status <domain-id>
keplars domains verify <domain-id>
keplars domains delete <domain-id>
keplars domains create-api-key <domain-id> --name "Production"
```

### Config
```bash
keplars config set api-key <key>
keplars config set base-url <url>
keplars config get
keplars config delete
```

## Global flags

| Flag | Description |
|---|---|
| `--api-key` | Override API key for this command |
| `--base-url` | Override API base URL |
| `--json` | Output raw JSON response |

## Key types

| Key | Format | Used for |
|---|---|---|
| Regular | `kms_xxx.live_yyy` | Sending emails |
| Admin | `kms_xxx.adm_yyy` | Contacts, audiences, automations, domains |

## License

MIT
