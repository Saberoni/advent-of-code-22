package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	e := make([]int, 0)
	var s int = 0

	for scanner.Scan() {
		ln := scanner.Text()

		if ln != "" {
			i, err := strconv.Atoi(ln)

			if err != nil {
				log.Fatal(err)
			}

			s += i
		} else {
			e = append(e, s)
			s = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(e)

	fmt.Println("Part 1:")
	fmt.Println(e[len(e)-1])

	var topThree int
	for _, tope := range e[len(e)-3:] {
		topThree += tope
	}
	fmt.Println("Part 2:")
	fmt.Println(topThree)
}
