@echo off
REM Build Keplars CLI for all platforms

set /p VERSION=<..\VERSION
set OUTPUT_DIR=dist

echo Building Keplars CLI v%VERSION% for all platforms...

REM Create output directory
if not exist %OUTPUT_DIR% mkdir %OUTPUT_DIR%

REM Build for Linux (amd64)
echo Building for Linux (amd64)...
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=amd64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-linux-amd64 .

REM Build for Linux (arm64)
echo Building for Linux (arm64)...
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-linux-arm64 .

REM Build for macOS (amd64 - Intel)
echo Building for macOS (amd64 - Intel)...
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=amd64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-darwin-amd64 .

REM Build for macOS (arm64 - Apple Silicon)
echo Building for macOS (arm64 - Apple Silicon)...
set CGO_ENABLED=0
set GOOS=darwin
set GOARCH=arm64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-darwin-arm64 .

REM Build for Windows (amd64)
echo Building for Windows (amd64)...
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-windows-amd64.exe .

REM Build for Windows (arm64)
echo Building for Windows (arm64)...
set CGO_ENABLED=0
set GOOS=windows
set GOARCH=arm64
go build -ldflags "-X github.com/KeplarsHQ/cli/cmd.Version=%VERSION%" -o %OUTPUT_DIR%\keplars-windows-arm64.exe .

echo.
echo Build complete! Binaries are in %OUTPUT_DIR%\
echo.
dir %OUTPUT_DIR%
