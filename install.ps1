#Requires -Version 5.1
<#
.SYNOPSIS
    Keplars CLI Windows Installer
.DESCRIPTION
    Downloads and installs the Keplars CLI for Windows.
.PARAMETER Version
    Specific version to install (e.g. "1.11.2"). Defaults to latest.
.PARAMETER InstallDir
    Directory to install the binary. Defaults to $env:USERPROFILE\.keplars\bin
.EXAMPLE
    irm https://keplars.com/install.ps1 | iex
.EXAMPLE
    & ([scriptblock]::Create((irm https://keplars.com/install.ps1))) -Version 1.11.2
#>
param(
    [string]$Version = "",
    [string]$InstallDir = ""
)

$Repo = "KeplarsHQ/cli"
$BinaryName = "keplars.exe"

if ($InstallDir -eq "") {
    $InstallDir = Join-Path $env:USERPROFILE ".keplars\bin"
}

function Write-Info  { param($msg) Write-Host "  [i] $msg" -ForegroundColor Cyan }
function Write-Ok    { param($msg) Write-Host "  [+] $msg" -ForegroundColor Green }
function Write-Warn  { param($msg) Write-Host "  [!] $msg" -ForegroundColor Yellow }
function Write-Fail  { param($msg) Write-Host "  [x] $msg" -ForegroundColor Red }

function Get-Arch {
    $arch = $env:PROCESSOR_ARCHITECTURE
    if ($env:PROCESSOR_ARCHITEW6432) { $arch = $env:PROCESSOR_ARCHITEW6432 }
    switch ($arch) {
        "AMD64" { return "amd64" }
        "ARM64" { return "arm64" }
        default {
            Write-Fail "Unsupported architecture: $arch"
            exit 1
        }
    }
}

function Get-LatestVersion {
    Write-Info "Fetching latest version..."
    try {
        $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$Repo/releases/latest" -UseBasicParsing
        return $release.tag_name -replace '^v', ''
    } catch {
        Write-Fail "Could not fetch latest version: $_"
        exit 1
    }
}

function Add-ToUserPath {
    param([string]$Dir)
    $currentPath = [Environment]::GetEnvironmentVariable("PATH", "User")
    if ($currentPath -notlike "*$Dir*") {
        [Environment]::SetEnvironmentVariable("PATH", "$currentPath;$Dir", "User")
        $env:PATH = "$env:PATH;$Dir"
        Write-Ok "Added $Dir to your user PATH"
        Write-Warn "Restart your terminal for PATH changes to take effect"
    }
}

Write-Host ""
Write-Host "  ======================================="  -ForegroundColor Blue
Write-Host "        Keplars CLI Installer (Win)       " -ForegroundColor Blue
Write-Host "  ======================================="  -ForegroundColor Blue
Write-Host ""

$arch = Get-Arch

if ($Version -eq "") {
    $Version = Get-LatestVersion
}

$assetName = "keplars-windows-$arch.exe"
$downloadUrl = "https://github.com/$Repo/releases/download/v$Version/$assetName"

Write-Info "Version  : v$Version"
Write-Info "Arch     : $arch"
Write-Info "Install  : $InstallDir"
Write-Host ""

if (-not (Test-Path $InstallDir)) {
    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
}

$targetPath = Join-Path $InstallDir $BinaryName
$tmpPath = Join-Path $env:TEMP "keplars-install-$Version.exe"

Write-Info "Downloading $downloadUrl..."

try {
    Invoke-WebRequest -Uri $downloadUrl -OutFile $tmpPath -UseBasicParsing
} catch {
    Write-Fail "Download failed: $_"
    Write-Fail "Version v$Version may not exist or is unavailable for windows-$arch"
    exit 1
}

Move-Item -Path $tmpPath -Destination $targetPath -Force
Write-Ok "Binary installed to $targetPath"

Add-ToUserPath $InstallDir

Write-Host ""
try {
    $installedVersion = & $targetPath --version 2>&1
    Write-Ok "Keplars CLI $installedVersion installed successfully!"
} catch {
    Write-Ok "Keplars CLI v$Version installed successfully!"
}

Write-Host ""
Write-Info "Get started:"
Write-Host "    keplars config set api-key kms_xxx.live_yyy" -ForegroundColor White
Write-Host "    keplars send --to user@example.com --from hello@yourdomain.com --subject `"Test`" --text `"Hello!`"" -ForegroundColor White
Write-Host ""
Write-Info "Docs: https://github.com/$Repo"
Write-Host ""
