#!/bin/bash

# Check if a version parameter has been provided
if [ -z "$1" ]; then
  echo "Usage: $0 <version>"
  exit 1
fi

VERSION=$1
BUILD_DIR="build"
OUTPUT_DIR="output"

# Create build and output directories if they don't exist
mkdir -p $BUILD_DIR
mkdir -p $OUTPUT_DIR

# Build for macOS
echo "Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o $BUILD_DIR/toolbox-cli-darwin-amd64-$VERSION
tar -cvzf $OUTPUT_DIR/toolbox-cli-darwin-amd64-$VERSION.tar.gz -C $BUILD_DIR toolbox-cli-darwin-amd64-$VERSION

# Build for Linux
echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o $BUILD_DIR/toolbox-cli-linux-amd64-$VERSION
tar -cvzf $OUTPUT_DIR/toolbox-cli-linux-amd64-$VERSION.tar.gz -C $BUILD_DIR toolbox-cli-linux-amd64-$VERSION

# Build for Windows
# echo "Building for Windows..."
# GOOS=windows GOARCH=amd64 go build -o $BUILD_DIR/toolbox-cli-windows-amd64-$VERSION.exe
# zip -r $OUTPUT_DIR/toolbox-cli-windows-amd64-$VERSION.zip $BUILD_DIR/toolbox-cli-windows-amd64-$VERSION.exe

# Clean up temporary build files
echo "Cleaning up..."
rm -rf $BUILD_DIR

echo "Build and packaging complete. Output files are in the $OUTPUT_DIR directory."
