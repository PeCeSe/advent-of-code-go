# Run solutions for specific days
run:
	go run 2024/main.go

# Run tests for Day 2
test2:
	go test -v ./2024/solutions -run ^TestDay02

test4:
	go test -v ./2024/solutions -run ^TestDay04

test5:
	go test -v ./2024/solutions -run ^TestDay5

test7:
	go test -v ./2024/solutions -run ^TestDay7

# Run all tests
test:
	go test -v ./...