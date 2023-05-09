package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type foodItem struct {
	id    int
	count int
}

func main() {
	// Open the log file
	// file-names of sample log files are "log.txt", "log1.txt" and "log2.txt"
	file, err := os.Open("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a map to store the counts of each foodItem
	counts := make(map[int]int)

	// Read the log file line by line
	var dda [][]int
	var sa []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		eater_id, foodmenu_id, err := parseLine(line)
		sa = []int{eater_id, foodmenu_id}
		dda = append(dda, sa)
		if err != nil {
			fmt.Println(err)
			continue
		}
		key := foodmenu_id
		counts[key]++
	}

	// show error when eater_id and foodmenu_id appears more than once
	err = checkDuplicateRows(dda)
	if err != nil {
		fmt.Println("eater_id and foodmenu_id cannot be ")
		return
	}

	// Sort the food items by count in descending order
	var foodItems []foodItem
	for key, count := range counts {
		foodItems = append(foodItems, foodItem{key, count})
	}
	sort.Slice(foodItems, func(i, j int) bool {
		return foodItems[i].count > foodItems[j].count
	})

	// Print the top 3 food items consumed
	fmt.Println("Top 3 menu items consumed:")
	for i := 0; i < 3 && i < len(foodItems); i++ {
		fmt.Printf("%d. Foodmenu_id %d was consumed %d times\n", i+1, foodItems[i].id, foodItems[i].count)
	}
}

// Parses a line of the log file and returns the eater_id and foodmenu_id.
func parseLine(line string) (int, int, error) {
	fields := bufio.NewScanner(strings.NewReader(line))
	fields.Split(bufio.ScanWords)
	var eater_id, foodmenu_id int
	for i := 0; fields.Scan(); i++ {
		if i == 0 {
			id, err := strconv.Atoi(fields.Text())
			if err != nil {
				return 0, 0, err
			}
			eater_id = id
		} else if i == 1 {
			id, err := strconv.Atoi(fields.Text())
			if err != nil {
				return 0, 0, err
			}
			foodmenu_id = id
		} else {
			return 0, 0, fmt.Errorf("unexpected field in line: %s", line)
		}

	}

	return eater_id, foodmenu_id, nil
}

// Check if there are any duplicate rows in the double dimensional array (log-file)
func checkDuplicateRows(arr [][]int) error {
	seen := make(map[string]bool)
	for _, row := range arr {
		rowStr := fmt.Sprintf("%v", row)
		if seen[rowStr] {
			return fmt.Errorf("duplicate row found: %v", row)
		}
		seen[rowStr] = true
	}
	return nil
}
