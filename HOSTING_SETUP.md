# CLI Installation Script Hosting Setup

This guide explains how to set up the custom domain redirect for the CLI installation script.

## Overview

Instead of using the long GitHub raw URL:
```bash
curl -fsSL https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh | bash
```

We use a clean custom domain:
```bash
curl -fsSL https://keplars.com/install.sh | bash
```

## Implementation Options

### Option 1: Simple HTTP Redirect (Recommended)

Configure your web server to redirect `https://keplars.com/install.sh` to the GitHub raw URL.

#### Nginx Configuration

Add this to your `keplars.com` server block:

```nginx
server {
    listen 443 ssl http2;
    server_name keplars.com;

    # SSL configuration
    ssl_certificate /path/to/cert.pem;
    ssl_certificate_key /path/to/key.pem;

    # Redirect install.sh to GitHub
    location = /install.sh {
        return 302 https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh;
    }

    # Your other location blocks...
}
```

#### Apache Configuration

Add to your `.htaccess` or Apache config:

```apache
RewriteEngine On
RewriteRule ^install\.sh$ https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh [R=302,L]
```

#### Caddy Configuration

```caddy
keplars.com {
    redir /install.sh https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh 302
}
```

### Option 2: Cloudflare Workers

Create a Cloudflare Worker to handle the redirect:

```javascript
export default {
  async fetch(request) {
    const url = new URL(request.url);

    if (url.pathname === '/install.sh') {
      return Response.redirect(
        'https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh',
        302
      );
    }

    // Handle other routes
    return fetch(request);
  }
}
```

Deploy this worker and route `keplars.com/install.sh` to it.

### Option 3: Vercel/Netlify Redirects

#### Vercel (`vercel.json`)

```json
{
  "redirects": [
    {
      "source": "/install.sh",
      "destination": "https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh",
      "permanent": false
    }
  ]
}
```

#### Netlify (`_redirects` file)

```
/install.sh https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh 302
```

### Option 4: Host the File Directly

Copy `cli/install.sh` to your web server and serve it directly from `https://keplars.com/install.sh`.

**Pros:**
- Fastest response time
- No dependency on GitHub
- Full control

**Cons:**
- Need to manually update when script changes
- Requires CI/CD to keep in sync

**Example with GitHub Actions:**

```yaml
name: Update Install Script

on:
  push:
    branches: [main]
    paths:
      - 'cli/install.sh'

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Deploy to web server
        run: |
          scp cli/install.sh ${{ secrets.WEB_SERVER_USER }}@${{ secrets.WEB_SERVER_HOST }}:/var/www/keplars.com/install.sh
```

## Testing the Setup

After configuration, test with:

```bash
# Test the redirect
curl -I https://keplars.com/install.sh

# Expected response should include:
# HTTP/2 302
# Location: https://raw.githubusercontent.com/...

# Test the actual installation
curl -fsSL https://keplars.com/install.sh | bash -s -- --help
```

## Recommended Setup

For `keplars.com`, we recommend **Option 1 (Simple HTTP Redirect)** because:

1. ✅ Simple to implement
2. ✅ Always serves the latest version from GitHub
3. ✅ No manual updates needed
4. ✅ Minimal maintenance
5. ✅ Works with existing infrastructure

## Security Considerations

1. **Always use HTTPS** - The installation script should only be served over HTTPS
2. **Use 302 (Temporary) redirects** - Allows changing the target URL if needed
3. **Consider Content Security** - If hosting directly, ensure proper file permissions
4. **Monitor access logs** - Track installation attempts and potential abuse

## Additional URLs to Consider

You might also want to set up redirects for:

```nginx
# Short URL for documentation
location = /docs {
    return 302 https://github.com/Swing-Technologies/keplars-sdk;
}

# CLI-specific docs
location = /cli {
    return 302 https://github.com/Swing-Technologies/keplars-sdk/blob/main/cli/README.md;
}

# Latest CLI release
location = /cli/latest {
    return 302 https://github.com/Swing-Technologies/keplars-sdk/releases/latest;
}
```

## Fallback Strategy

Always document the GitHub raw URL as a fallback in case the custom domain experiences issues:

```bash
# Primary (custom domain)
curl -fsSL https://keplars.com/install.sh | bash

# Fallback (direct GitHub)
curl -fsSL https://raw.githubusercontent.com/Swing-Technologies/keplars-sdk/main/cli/install.sh | bash
```

This ensures users can always install the CLI even if there are temporary domain or server issues.
