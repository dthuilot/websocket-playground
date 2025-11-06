#!/bin/bash

# WebSocket Connection Test Script
# This script tests the WebSocket connection using websocat or wscat

set -e

HOST="${1:-localhost:8080}"
PATH_URL="${2:-/ws}"
FULL_URL="ws://${HOST}${PATH_URL}"

echo "==============================================="
echo "WebSocket Connection Test"
echo "==============================================="
echo "Target: ${FULL_URL}"
echo ""

# Check if websocat is installed
if command -v websocat &> /dev/null; then
    echo "Using websocat for testing..."
    echo "Type messages to send (Ctrl+C to exit)"
    echo "-----------------------------------------------"
    websocat "${FULL_URL}"
# Check if wscat is installed
elif command -v wscat &> /dev/null; then
    echo "Using wscat for testing..."
    echo "Type messages to send (Ctrl+C to exit)"
    echo "-----------------------------------------------"
    wscat -c "${FULL_URL}"
else
    echo "ERROR: Neither websocat nor wscat is installed"
    echo ""
    echo "Please install one of the following:"
    echo "  - websocat: https://github.com/vi/websocat"
    echo "    Installation: cargo install websocat"
    echo ""
    echo "  - wscat: https://github.com/websockets/wscat"
    echo "    Installation: npm install -g wscat"
    echo ""
    echo "Alternatively, you can:"
    echo "  1. Use the Go client: go run cmd/client/main.go"
    echo "  2. Open http://localhost:8080 in a browser"
    echo "  3. Use Postman's WebSocket feature"
    exit 1
fi
