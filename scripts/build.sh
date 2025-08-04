#!/bin/bash

# Build script for cube-orchestrator
# Generates timestamped executables in the builds/ directory

set -e

# Get current timestamp
TIMESTAMP=$(date +"%Y%m%d_%H%M%S")
BUILD_DIR="builds"
BINARY_NAME="cube-orchestrator"
OUTPUT_FILE="${BUILD_DIR}/${BINARY_NAME}_${TIMESTAMP}"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${BLUE}🔨 Building cube-orchestrator...${NC}"
echo -e "${YELLOW}Timestamp: ${TIMESTAMP}${NC}"

# Create builds directory if it doesn't exist
mkdir -p "${BUILD_DIR}"

# Change to src directory for build
cd src/cmd/orchestrator

# Build the application
echo -e "${BLUE}📦 Compiling...${NC}"
go build -o "../../../${OUTPUT_FILE}" .

# Make executable (for Unix systems)
chmod +x "../../../${OUTPUT_FILE}"

echo -e "${GREEN}✅ Build successful!${NC}"
echo -e "${GREEN}📁 Executable: ${OUTPUT_FILE}${NC}"

# Show file info
ls -lh "../../../${OUTPUT_FILE}"

# Create/update latest symlink for convenience
cd ../../..
rm -f "${BUILD_DIR}/${BINARY_NAME}_latest"
ln -s "${BINARY_NAME}_${TIMESTAMP}" "${BUILD_DIR}/${BINARY_NAME}_latest"

echo -e "${GREEN}🔗 Latest symlink updated: ${BUILD_DIR}/${BINARY_NAME}_latest${NC}"

# Show builds directory contents
echo -e "${BLUE}📋 Builds directory:${NC}"
ls -la "${BUILD_DIR}/"

echo -e "${GREEN}🎉 Build complete!${NC}"
