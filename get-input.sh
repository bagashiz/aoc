#!/usr/bin/env bash

# Load environment variables from .env
if [ -f .env ]; then
  export "$(grep -v '^#' .env | xargs)"
else
  echo "Error: .env file not found!"
  exit 1
fi

# Check if AOC_TOKEN is set
if [ -z "$AOC_TOKEN" ]; then
  echo "Error: AOC_TOKEN is not set in the .env file."
  exit 1
fi

# Check if year and day are provided as arguments
if [ -z "$1" ] || [ -z "$2" ]; then
  echo "Usage: $0 <year> <day>"
  exit 1
fi

YEAR=$1
DAY=$2
DAY_PADDED=$(printf "%02d" "$DAY") # Pad single-digit days for the folder structure
URL="https://adventofcode.com/$YEAR/day/$DAY/input"
OUTPUT_DIR="${YEAR}/day${DAY_PADDED}"
OUTPUT_FILE="${OUTPUT_DIR}/input.txt"

# Create the directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Fetch input from the API
echo "Fetching input for year $YEAR, day $DAY..."
HTTP_STATUS=$(curl -s -w "%{http_code}" --cookie "session=$AOC_TOKEN" "$URL" -o "$OUTPUT_FILE.tmp")

if [ "$HTTP_STATUS" -eq 200 ]; then
  mv "$OUTPUT_FILE.tmp" "$OUTPUT_FILE"
  echo "Input for year $YEAR, day $DAY saved to $OUTPUT_FILE."
else
  rm "$OUTPUT_FILE.tmp"
  rmdir "$OUTPUT_DIR"
  echo "Error: Failed to fetch input for year $YEAR, day $DAY. HTTP status code: $HTTP_STATUS."
  exit 1
fi
