#!/bin/bash
set -e

# Build Keplars CLI for all platforms

VERSION=$(cat ../VERSION)
OUTPUT_DIR="dist"

echo "Building Keplars CLI v${VERSION} for all platforms..."

# Create output directory
mkdir -p ${OUTPUT_DIR}

# Build for Linux (amd64)
echo "Building for Linux (amd64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-linux-amd64 .

# Build for Linux (arm64)
echo "Building for Linux (arm64)..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-linux-arm64 .

# Build for macOS (amd64 - Intel)
echo "Building for macOS (amd64 - Intel)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-darwin-amd64 .

# Build for macOS (arm64 - Apple Silicon)
echo "Building for macOS (arm64 - Apple Silicon)..."
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-darwin-arm64 .

# Build for Windows (amd64)
echo "Building for Windows (amd64)..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-windows-amd64.exe .

# Build for Windows (arm64)
echo "Building for Windows (arm64)..."
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -ldflags "-X github.com/Swing-Technologies/keplars-cli/cmd.Version=${VERSION}" -o ${OUTPUT_DIR}/keplars-windows-arm64.exe .

echo ""
echo "✓ Build complete! Binaries are in ${OUTPUT_DIR}/"
echo ""
ls -lh ${OUTPUT_DIR}/
