package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	logger := log.New(os.Stdout, "[GOEXRCS]\t", log.Ldate|log.Ltime)

	var from int
	var to int
	var list string

	flag.IntVar(&from, "from", 1, "From (Inclusive)")
	flag.IntVar(&to, "to", 100, "To (Exclusive)")

	flag.StringVar(&list, "divisible-by", "[]", "Divisible By")
	flag.Parse()

	numbers, err := parse(list)
	if err != nil {
		logger.Printf("failed to parse disible-by list: %s\n", err)
		os.Exit(1)
	}

	items := make(map[int]bool)
	for i := from; i < to; i++ {
		items[i] = true
	}

	for _, number := range numbers {
		for key, value := range items {
			if !divisible(key, number) {
				if value {
					items[key] = false
				}
			}
		}
	}

	for key, value := range items {
		if value {
			fmt.Printf("%d. ", key)
		}
	}

	fmt.Println()
}

func parse(list string) ([]int, error) {
	var numbers []int

	list = strings.ReplaceAll(list, " ", "")
	if !(strings.HasPrefix(list, "[") && strings.HasSuffix(list, "]")) {
		return nil, errors.New("invalid divisible-by list format")
	}

	list = strings.ReplaceAll(list, "[", "")
	list = strings.ReplaceAll(list, "]", "")

	if list == "" {
		return numbers, nil
	}

	parts := strings.Split(list, ",")
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, errors.New("invalid divisible-by list element: can't be converted to int")
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func divisible(a int, b int) bool {
	return a%b == 0
}
