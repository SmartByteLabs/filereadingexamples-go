#!/bin/bash

# Define the base directory relative to the script's location
BASE_DIR="../"  # Adjust this if your script location changes

# Find all test files and run the benchmark in their respective directories
find "$BASE_DIR" -type f -name "*_test.go" | while read -r test_file; do
  dir=$(dirname "$test_file")
  filename=$(basename "$test_file")
  
  # Print a separator before starting the benchmark
  echo "-------------------------------------------"
  echo "Running benchmark for $filename"
  echo "-------------------------------------------"
  
  # Navigate into the directory and run the benchmark
  (cd "$dir" && go test -bench=. -benchmem -run=^$ .)
  
  # Print a separator after completing the benchmark
  echo "-------------------------------------------"
  echo "Completed benchmark for $filename"
  echo "-------------------------------------------"
  echo
done

