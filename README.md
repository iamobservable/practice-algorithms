# Algorithms

This repository has been created using the programming language Go. Its goal
is to help in practicing algorithm skills. It includes algorithm tests that 
can be run directly via Makefile actions. How the algorithms are implemented 
is up to you.

Enjoy!

# Setup / Installation

You can install both golang and make individually or [install and use mise](https://mise.jdx.dev/installing-mise.html)
to handle installation for you.

- [ ] Go 1.26.2
- [ ] make 4.4.1

If mise is installed, just run the following the the project directory:

```sh
mise install
```



## Leetcode

Run leetcode tests for for sorts algorithm.

| Practice                                                                                                       | Test Command             |
| :---                                                                                                           |                     ---: |
| [Reverse Integer](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00007_reverse_integer.go)       | make test-leetcode-00007 |
| [Rotate Image](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00048_rotate_image.go)             | make test-leetcode-00048 |
| [Gas Station](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00134_gas_station.go)               | make test-leetcode-00134 |
| [Candy](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00135_candy.go)                           | make test-leetcode-00135 |
| [Perfect Squares](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00279_perfect_squares.go)       | make test-leetcode-00279 |
| [Wiggle Subsequence](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00376_wiggle_subsequence.go) | make test-leetcode-00376 |
| [Remove K Digits](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00402_remove_k_digits.go)       | make test-leetcode-00402 |
| [Can Place Flowers](https://github.com/iamobservable/practices-algorithms/blob/main/lc/00605_can_place_flowers.go)   | make test-leetcode-00605 |


## Reverse

Run tests for reverse algorithm.

| Practice                                                                                                       | Test Command             |
| :---                                                                                                           |                     ---: |
| [Reverse String](https://github.com/iamobservable/practices-algorithms/blob/main/strings/reverse_string.go)          | make test-reverse-string |


## Sorts

Run all or individual tests for sort algorithms.

```sh
make test-sorts
```

| Practice                                                                                                       | Test Command             |
| :---                                                                                                           |                     ---: |
| [Countingsort](https://github.com/iamobservable/practices-algorithms/blob/main/sorts/countingsort.go)                | make test-countsorting   |
| [Heapsort](https://github.com/iamobservable/practices-algorithms/blob/main/sorts/heapsort.go)                        | make test-heapsort       |
| [Insertionsort](https://github.com/iamobservable/practices-algorithms/blob/main/sorts/insertionsort.go)              | make test-insertionsort  |
| [Mergesort](https://github.com/iamobservable/practices-algorithms/blob/main/sorts/mergesort.go)                      | make test-mergesort      |
| [Quicksort](https://github.com/iamobservable/practices-algorithms/blob/main/sorts/quicksort.go)                      | make test-quicksort      |


## Stack

Run tests for stack algorithm.

| Practice                                                                                                       | Test Command             |
| :---                                                                                                           |                     ---: |
| [Stack](https://github.com/iamobservable/practices-algorithms/blob/main/stacks/stack.go)                             | make test-stack          |
