#!/bin/bash
set -e

# Keplars CLI Installation Script
# Works on Linux, macOS, and Windows (via Git Bash/WSL)
#
# Usage:
#   curl -fsSL https://cli.keplars.com/install.sh | bash
#   curl -fsSL https://cli.keplars.com/install.sh | bash -s -- 1.11.2

REPO="KeplarsHQ/cli"
BINARY_NAME="keplars"

# Detect if running on Windows (Git Bash or WSL)
if [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" || "$OSTYPE" == "cygwin" ]]; then
    # Windows (Git Bash)
    INSTALL_DIR="${INSTALL_DIR:-$HOME/bin}"
    IS_WINDOWS=true
    CURL_OPTS="--ssl-no-revoke"
elif [[ -n "$WSL_DISTRO_NAME" ]] || grep -qi microsoft /proc/version 2>/dev/null; then
    # WSL
    INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
    IS_WINDOWS=false
    CURL_OPTS=""
else
    # Linux/macOS
    INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
    IS_WINDOWS=false
    CURL_OPTS=""
fi

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Functions
print_info() {
    echo -e "${BLUE}ℹ${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠${NC} $1"
}

# Detect OS and Architecture
detect_platform() {
    OS="$(uname -s)"
    ARCH="$(uname -m)"

    case "$OS" in
        Linux*)
            PLATFORM="linux"
            ;;
        Darwin*)
            PLATFORM="darwin"
            ;;
        MINGW*|MSYS*|CYGWIN*)
            PLATFORM="windows"
            ;;
        *)
            print_error "Unsupported operating system: $OS"
            exit 1
            ;;
    esac

    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac

    if [ "$PLATFORM" = "windows" ]; then
        BINARY_NAME_WITH_PLATFORM="${BINARY_NAME}-${PLATFORM}-${ARCH}.exe"
        BINARY_NAME="${BINARY_NAME}.exe"
    else
        BINARY_NAME_WITH_PLATFORM="${BINARY_NAME}-${PLATFORM}-${ARCH}"
    fi
}

# Get latest version from GitHub
get_latest_version() {
    print_info "Fetching latest version..." >&2

    LATEST_VERSION=$(curl -fsSL $CURL_OPTS -H "User-Agent: keplars-cli-installer" -H "Accept: application/vnd.github.v3+json" "https://api.github.com/repos/${REPO}/releases/latest" | \
        grep '"tag_name":' | \
        sed -E 's/.*"v([^"]+)".*/\1/')

    if [ -z "$LATEST_VERSION" ]; then
        print_error "Could not determine latest version" >&2
        exit 1
    fi

    echo "$LATEST_VERSION"
}

# Setup PATH for Windows (Git Bash)
setup_windows_path() {
    # Create $HOME/bin if it doesn't exist
    if [ ! -d "$HOME/bin" ]; then
        mkdir -p "$HOME/bin"
    fi

    # Check if $HOME/bin is in PATH
    if [[ ":$PATH:" != *":$HOME/bin:"* ]]; then
        print_warning "Adding $HOME/bin to PATH..."

        # Add to .bashrc if it exists
        if [ -f "$HOME/.bashrc" ]; then
            echo '' >> "$HOME/.bashrc"
            echo '# Added by Keplars CLI installer' >> "$HOME/.bashrc"
            echo 'export PATH="$HOME/bin:$PATH"' >> "$HOME/.bashrc"
            print_success "Added to ~/.bashrc"
        fi

        # Add to .bash_profile if it exists
        if [ -f "$HOME/.bash_profile" ]; then
            echo '' >> "$HOME/.bash_profile"
            echo '# Added by Keplars CLI installer' >> "$HOME/.bash_profile"
            echo 'export PATH="$HOME/bin:$PATH"' >> "$HOME/.bash_profile"
            print_success "Added to ~/.bash_profile"
        fi

        # Update current session
        export PATH="$HOME/bin:$PATH"

        print_warning "Please restart your terminal or run: source ~/.bashrc"
    fi
}

# Download and install
install_cli() {
    VERSION="${1:-$(get_latest_version)}"

    print_info "Installing Keplars CLI v${VERSION} for ${PLATFORM}-${ARCH}..."

    # Download URL
    DOWNLOAD_URL="https://github.com/${REPO}/releases/download/v${VERSION}/${BINARY_NAME_WITH_PLATFORM}"

    # Create temporary directory
    TMP_DIR=$(mktemp -d)
    TMP_FILE="${TMP_DIR}/${BINARY_NAME}"

    # Download binary
    print_info "Downloading from ${DOWNLOAD_URL}..."
    if ! curl -fsSL $CURL_OPTS "$DOWNLOAD_URL" -o "$TMP_FILE"; then
        print_error "Failed to download binary"
        print_error "Version v${VERSION} may not exist or binary not available for ${PLATFORM}-${ARCH}"
        rm -rf "$TMP_DIR"
        exit 1
    fi

    # Make executable
    chmod +x "$TMP_FILE"

    # Ensure install directory exists
    mkdir -p "$INSTALL_DIR"

    # Check if we need sudo (never on Windows)
    if [ "$IS_WINDOWS" = true ] || [ -w "$INSTALL_DIR" ]; then
        SUDO=""
    else
        SUDO="sudo"
        print_warning "Installing to ${INSTALL_DIR} requires sudo access"
    fi

    # Install binary
    print_info "Installing to ${INSTALL_DIR}/${BINARY_NAME}..."
    $SUDO mv "$TMP_FILE" "${INSTALL_DIR}/${BINARY_NAME}"

    # Cleanup
    rm -rf "$TMP_DIR"

    # Setup PATH for Windows (Git Bash)
    if [ "$IS_WINDOWS" = true ]; then
        setup_windows_path
    fi

    # Verify installation
    VERIFY_CMD="${INSTALL_DIR}/${BINARY_NAME}"
    if command -v "${BINARY_NAME%.exe}" >/dev/null 2>&1 || [ -f "$VERIFY_CMD" ]; then
        if [ -f "$VERIFY_CMD" ]; then
            INSTALLED_VERSION=$("$VERIFY_CMD" --version 2>&1 | grep -oE '[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?' | head -n1)
        else
            INSTALLED_VERSION=$("${BINARY_NAME%.exe}" --version 2>&1 | grep -oE '[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.]+)?' | head -n1)
        fi

        print_success "Keplars CLI v${INSTALLED_VERSION} installed successfully!"
        echo ""
        print_info "Get started:"

        if [ "$PLATFORM" = "windows" ]; then
            echo "  export KEPLARS_API_KEY='your-api-key'"
            echo "  keplars.exe send --to user@example.com --from hello@yourdomain.com --subject 'Test' --text 'Hello!'"
        else
            echo "  export KEPLARS_API_KEY='your-api-key'"
            echo "  keplars send --to user@example.com --from hello@yourdomain.com --subject 'Test' --text 'Hello!'"
        fi

        echo ""
        print_info "Documentation: https://github.com/${REPO}#readme"
    else
        print_error "Installation completed but ${BINARY_NAME} not found in PATH"
        print_warning "You may need to add ${INSTALL_DIR} to your PATH"
        echo "  export PATH=\"\$PATH:${INSTALL_DIR}\""
    fi
}

# Main
main() {
    echo ""
    echo "╔═══════════════════════════════════════╗"
    echo "║       Keplars CLI Installer          ║"
    echo "╚═══════════════════════════════════════╝"
    echo ""

    detect_platform
    install_cli "$1"
}

main "$@"
