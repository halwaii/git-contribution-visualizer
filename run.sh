#!/bin/bash
set -e
VERSION="v1.0.0"
REPO="halwaii/git-contribution-visualizer"
BASE_URL="https://github.com/${REPO}/releases/download/${VERSION}"
TARGET_DIR="${1:-.}"

OS="$(uname -s)"
case "$OS" in
  Linux*)  BIN="git-contribution-visualizer-linux" ;;
  Darwin*) BIN="git-contribution-visualizer-mac" ;;
  MINGW*|MSYS*|CYGWIN*) BIN="git-contribution-visualizer.exe" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

URL="${BASE_URL}/${BIN}"
TMP_BIN="/tmp/${BIN}"
echo "→ Detected OS: $OS"
echo "→ Downloading $BIN ..."
curl -L -o "$TMP_BIN" "$URL" || wget -O "$TMP_BIN" "$URL"
chmod +x "$TMP_BIN"
echo "→ Running for: $TARGET_DIR"
"$TMP_BIN" "$TARGET_DIR"