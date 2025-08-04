#!/bin/bash

# Cleanup script for cube-orchestrator builds
# Keeps only the latest N builds (default: 5)

set -e

# Configuration
BUILD_DIR="builds"
BINARY_PREFIX="cube-orchestrator_"
KEEP_COUNT=${1:-5}  # Default to keeping 5 builds, can be overridden

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${BLUE}🧹 Cleaning up old builds...${NC}"
echo -e "${YELLOW}Keeping latest ${KEEP_COUNT} builds${NC}"

# Change to project root
cd "$(dirname "$0")/.."

# Check if builds directory exists
if [ ! -d "${BUILD_DIR}" ]; then
    echo -e "${YELLOW}⚠️  No builds directory found. Nothing to clean.${NC}"
    exit 0
fi

# List all timestamped builds (excluding symlinks)
cd "${BUILD_DIR}"
BUILDS=($(ls -t ${BINARY_PREFIX}[0-9]* 2>/dev/null | head -20))

if [ ${#BUILDS[@]} -eq 0 ]; then
    echo -e "${YELLOW}⚠️  No timestamped builds found. Nothing to clean.${NC}"
    exit 0
fi

echo -e "${BLUE}📋 Found ${#BUILDS[@]} builds:${NC}"
for build in "${BUILDS[@]}"; do
    echo "  - $build"
done

# If we have more builds than we want to keep, delete the old ones
if [ ${#BUILDS[@]} -gt ${KEEP_COUNT} ]; then
    BUILDS_TO_DELETE=(${BUILDS[@]:${KEEP_COUNT}})
    
    echo -e "\n${RED}🗑️  Deleting ${#BUILDS_TO_DELETE[@]} old builds:${NC}"
    for build in "${BUILDS_TO_DELETE[@]}"; do
        echo -e "  ${RED}Removing: $build${NC}"
        rm -f "$build"
    done
    
    echo -e "\n${GREEN}✅ Cleanup completed!${NC}"
    echo -e "${GREEN}📁 Kept latest ${KEEP_COUNT} builds${NC}"
else
    echo -e "\n${GREEN}✅ No cleanup needed.${NC}"
    echo -e "${GREEN}📁 Current build count (${#BUILDS[@]}) is within limit (${KEEP_COUNT})${NC}"
fi

# Show remaining builds
echo -e "\n${BLUE}📋 Remaining builds:${NC}"
ls -la ${BINARY_PREFIX}* 2>/dev/null || echo "  No builds found"
