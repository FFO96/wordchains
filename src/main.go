package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	word         string
	previousNode *Node
}

func main() {

	start := "dog"
	end := "cat"
	data := "data/dictionary"

	dictionary := make(map[string]struct{})

	file, err := os.Open(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = struct{}{}
	}

	if existInDictionary(start, dictionary) && existInDictionary(end, dictionary) {
		//TODO: findPath call
	} else {
		fmt.Println("The starting and ending words must be registered in the dictionary")
	}
}

// Checks if the word exists in the dictionary
func existInDictionary(word string, dictionary map[string]struct{}) bool {
	_, exist := dictionary[word]
	return exist
}

// TODO: Returns the neighbours of a word
func createNeighbours(word string, dictionary map[string]struct{}) map[string]struct{} {
	neighbours := make(map[string]struct{})
	for letter := 'a'; letter <= 'z'; letter++ {

		for i := range word {
			neighbour := word[:i] + string(letter) + word[i:]

			if _, exist := dictionary[neighbour]; exist && neighbour != word {
				neighbours[neighbour] = struct{}{}
			}
		}

		for i := range word {
			neighbour := word[:i] + word[i+1:]
			if _, exist := dictionary[neighbour]; exist && neighbour != word {
				neighbours[neighbour] = struct{}{}
			}
		}

		for i := range word {
			neighbour := word[:i] + string(letter) + word[i+1:]
			if _, exist := dictionary[neighbour]; exist && neighbour != word {
				neighbours[neighbour] = struct{}{}
			}
		}

	}
	return neighbours
}

//TODO: Returns the solution if exists
func getSolutionList() {}

//TODO: Find the path of the solution if exists
func findPath() {}
