# Run solutions for specific days
run:
	go run 2024/main.go

# Run tests for Day 2
test2:
	go test -v ./2024/solutions -run ^TestDay02

# Run all tests
test:
	go test -v ./...