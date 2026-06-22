package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Group struct {
	students int
	index    int
}

type Room struct {
	computers int
	index     int
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err == nil {
		defer inputFile.Close()
		os.Stdin = inputFile
	}

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output.txt:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nextInt := func() int {
		scanner.Scan()
		var x int
		fmt.Sscanf(scanner.Text(), "%d", &x)
		return x
	}

	N, M := nextInt(), nextInt()

	groups := make([]Group, N)
	rooms := make([]Room, M)

	for i := 0; i < N; i++ {
		groups[i] = Group{students: nextInt() + 1, index: i}
	}

	for j := 0; j < M; j++ {
		rooms[j] = Room{computers: nextInt(), index: j + 1}
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].students < groups[j].students
	})

	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i].computers < rooms[j].computers
	})

	assignment := make([]int, N)
	p := 0
	j := 0

	for _, group := range groups {
		for j < M && rooms[j].computers < group.students {
			j++
		}
		if j < M {
			assignment[group.index] = rooms[j].index
			p++
			j++
		}
	}

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()
	fmt.Fprintln(writer, p)
	for _, val := range assignment {
		fmt.Fprint(writer, val, " ")
	}
	fmt.Fprintln(writer)
}
