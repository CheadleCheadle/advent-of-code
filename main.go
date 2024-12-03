package main

import (
	"sort"
	"fmt"
	"io"
	"net/http"
	"os"
	"log"
	"strings"
	"strconv"

	"github.com/joho/godotenv"
)

func getInput(day int) string {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

sessionCookie := os.Getenv("SESSION")

 url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
 req, _ := http.NewRequest("GET", url, nil)

 req.AddCookie(&http.Cookie{
	Name: "session",
	Value: sessionCookie,
 })


 client := &http.Client{}
 resp, _ := client.Do(req)

 defer resp.Body.Close()

 body, _ := io.ReadAll(resp.Body)

 return string(body)



}

func main() {
	day := 1
	input := getInput(day)

	newInput := strings.Split(input, " ")

	var numbers[]int

	for _, str := range newInput {
		cleaned := strings.ReplaceAll(str, " ", "")
		cleaned = strings.ReplaceAll(str, "\n", "")

		for i := 0; i < len(cleaned); i += 5 {
			numberStr := cleaned[i : i+5]

			num, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Error converting to number", err)
				continue
			}

			numbers = append(numbers, num)
		}

	}

	var left[]int
	var right[]int

	for i := 0; i < len(numbers); i++ {
		if i % 2 == 0 {
			left = append(left, numbers[i])
		} else {
			right = append(right, numbers[i])
		}
	}

	sort.Ints(left)
	sort.Ints(right)

	var total int

	//Day 1 Part 1 O(n)
	// for x := 0; x < len(right); x++ {
	// 	diff := right[x] - left[x]
	// 	if diff < 0 {
	// 		diff = -diff
	// 	}
	// 	total += diff
	// }


	// Day 1 Part 2 O(n^2)
	for x := 0; x < len(left); x++ {
		curLeft := left[x]
		freq := 0
		for y := 0; y < len(right); y++ {
		curRight:= right[y]
		if curLeft == curRight {
			freq++
		}
		}
		total += curLeft * freq
	}

	fmt.Println(total)
}
