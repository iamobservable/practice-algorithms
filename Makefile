.PHONY: usage all

usage:
	@echo The following algorithm practice challenges are available. After
	@echo adding an implementation for the challenge, execute the associated test
	@echo command listed below it.
	@echo
	@tput setaf 2; echo Heapsort; tput sgr0
	@echo make test-heapsort
	@echo
	@tput setaf 2; echo Insertion Sort; tput sgr0
	@echo make test-insertionsort
	@echo
	@tput setaf 2; echo Merge Sort; tput sgr0
	@echo make test-mergesort
	@echo
	@tput setaf 2; echo Quick Sort; tput sgr0
	@echo make test-quicksort
	@echo
	@tput setaf 2; echo All Sorts; tput sgr0
	@echo make test-sorts
	@echo
	@tput setaf 2; echo Reversing A String; tput sgr0
	@echo make test-reverse-string
	@echo
	@tput setaf 2; echo Creating A Stack; tput sgr0
	@echo make test-stack
	@echo
	@tput setaf 2; echo Reverse Integer; tput sgr0
	@echo make test-leetcode-00007
	@echo
	@tput setaf 2; echo Rotate Image; tput sgr0
	@echo make test-leetcode-00048
	@echo
	@tput setaf 2; echo Gas Station Starting Index; tput sgr0
	@echo make test-leetcode-00134
	@echo
	@tput setaf 2; echo Candy Distribution; tput sgr0
	@echo make test-leetcode-00135
	@echo
	@tput setaf 2; echo Perfect Squares; tput sgr0
	@echo make test-leetcode-00279
	@echo
	@tput setaf 2; echo Wiggle Subsequence; tput sgr0
	@echo make test-leetcode-00376
	@echo
	@tput setaf 2; echo Remove K Digits; tput sgr0
	@echo make test-leetcode-00402
	@echo
	@tput setaf 2; echo Can Place Flowers; tput sgr0
	@echo make test-leetcode-00605

all: test-reverse-string test-sorts test-stack

# Leetcode
test-leetcode-00007:
	@echo executing leetcode 00007 reverse integer
	@go test -v lc/00007_reverse_integer_test.go lc/00007_reverse_integer.go

test-leetcode-00048:
	@echo executing leetcode 00048 rotate image
	@go test -v lc/00048_rotate_image_test.go lc/00048_rotate_image.go

test-leetcode-00134:
	@echo executing leetcode 00134 gas station tests
	@go test -v lc/00134_gas_station_test.go lc/00134_gas_station.go

test-leetcode-00135:
	@echo executing leetcode 00135 candy tests
	@go test -v lc/00135_candy_test.go lc/00135_candy.go

test-leetcode-00279:
	@echo executing leetcode 00279 perfect squares
	@go test -v lc/00279_perfect_squares_test.go lc/00279_perfect_squares.go

test-leetcode-00376:
	@echo executing leetcode 00376 wiggle subsequence tests
	@go test -v lc/00376_wiggle_subsequence_test.go lc/00376_wiggle_subsequence.go

test-leetcode-00402:
	@echo executing leetcode 00402 remove k digits tests
	@go test -v lc/00402_remove_k_digits_test.go lc/00402_remove_k_digits.go

test-leetcode-00605:
	@echo executing leetcode 00605 can place flowers tests
	@go test -v lc/00605_can_place_flowers_test.go lc/00605_can_place_flowers.go


# Sorts
test-sorts: test-countingsort test-heaptsort test-insertionsort test-mergesort test-quicksort
	@echo executing sort tests
	@go test -v sorts/countingsort_test.go sorts/countingsort.go \
		sorts/heapsort_test.go sorts/heapsort.go \
		sorts/insertionsort_test.go sorts/insertionsort.go \
		sorts/mergesort_test.go sorts/mergesort.go \
		sorts/quicksort_test.go sorts/quicksort.go

test-countingsort:
	@echo executing countingsort tests
	@go test -v sorts/countingsort_test.go sorts/countingsort.go

test-heaptsort:
	@echo executing heapsort tests
	@go test -v sorts/heapsort_test.go sorts/heapsort.go

test-insertionsort:
	@echo executing insertionsort tests
	@go test -v sorts/insertionsort_test.go sorts/insertionsort.go

test-mergesort:
	@echo executing mergesort tests
	@go test -v sorts/mergesort_test.go sorts/mergesort.go

test-quicksort:
	@echo executing quicksort tests
	@go test -v sorts/quicksort_test.go sorts/quicksort.go


# Extra
test-reverse-string:
	@echo executing reverse string tests
	@go test -v strings/reverse_string_test.go strings/reverse_string.go

test-stack:
	@echo executing stack tests
	@go test -v stacks/stack_test.go stacks/stack.go
