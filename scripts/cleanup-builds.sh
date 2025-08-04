#!/bin/bash

# Cleanup script for cube-orchestrator builds
# Keeps only the most recent N builds to save disk space

set -e

# Configuration
BUILD_DIR="builds"
BINARY_NAME="cube-orchestrator"
KEEP_BUILDS=1

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧹 Cleaning up old builds...${NC}"

if [ ! -d "${BUILD_DIR}" ]; then
    echo -e "${YELLOW}⚠️  No builds directory found${NC}"
    exit 0
fi

cd "${BUILD_DIR}"

# Count timestamped builds (exclude latest symlink and debug builds)
BUILD_COUNT=$(ls -1 ${BINARY_NAME}_[0-9]* 2>/dev/null | wc -l)

echo -e "${BLUE}📊 Found ${BUILD_COUNT} timestamped builds${NC}"

if [ "$BUILD_COUNT" -le "$KEEP_BUILDS" ]; then
    echo -e "${GREEN}✅ No cleanup needed (keeping up to ${KEEP_BUILDS} builds)${NC}"
    exit 0
fi

# Calculate how many to delete
DELETE_COUNT=$((BUILD_COUNT - KEEP_BUILDS))
echo -e "${YELLOW}🗑️  Will delete ${DELETE_COUNT} old builds${NC}"

# Get oldest builds to delete
TO_DELETE=$(ls -1t ${BINARY_NAME}_[0-9]* | tail -n ${DELETE_COUNT})

# Delete old builds
for build in $TO_DELETE; do
    echo -e "${RED}🗑️  Deleting: ${build}${NC}"
    rm -f "$build"
done

# Update latest symlink to point to most recent build
LATEST_BUILD=$(ls -1t ${BINARY_NAME}_[0-9]* | head -n 1)
if [ -n "$LATEST_BUILD" ]; then
    rm -f "${BINARY_NAME}_latest"
    ln -s "$LATEST_BUILD" "${BINARY_NAME}_latest"
    echo -e "${GREEN}🔗 Updated latest symlink to: ${LATEST_BUILD}${NC}"
fi

echo -e "${GREEN}✅ Cleanup complete!${NC}"
echo -e "${BLUE}📋 Remaining builds:${NC}"
ls -la ${BINARY_NAME}_*
