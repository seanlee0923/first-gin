#!/bin/bash

GO_VERSION="1.22"
GO_OS_ARCH="linux-amd64"
GO_DOWNLOAD_URL="https://golang.org/dl/go$GO_VERSION.$GO_OS_ARCH.tar.gz"
GO_FILENAME="go$GO_VERSION.$GO_OS_ARCH.tar.gz"

INSTALL_DIR="/usr/local"
GO_INSTALL_DIR="$INSTALL_DIR/go"

if command -v go &> /dev/null
then
    echo "Go already installed."
    exit 0
fi

echo "Downloading Go..."
curl -fsSL "$GO_DOWNLOAD_URL" -o "$GO_FILENAME"
sudo tar -C "$INSTALL_DIR" -xzf "$GO_FILENAME"
rm "$GO_FILENAME"

echo "export PATH=\$PATH:$GO_INSTALL_DIR/bin" >> ~/.bashrc

source ~/.bashrc

echo "Go installation complete."
