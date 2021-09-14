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
		var solution []string
		found := findPath(start, end, dictionary, &solution)

		if found {
			fmt.Println(solution)
		} else {
			fmt.Println("There is no path between the words")
		}
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

// Returns the solution if exists
func getSolutionArray(solution *[]string, expandedNodes *Node) {

	if expandedNodes.previousNode != nil {
		getSolutionArray(solution, expandedNodes.previousNode)
	}
	*solution = append(*solution, expandedNodes.word)
}

// Find the path of the solution if exists
func findPath(start, end string, dictionary map[string]struct{}, solution *[]string) bool {

	graph := []Node{{word: start, previousNode: nil}}
	wordsRecord := make(map[string]struct{})
	neighbours := make(map[string]struct{})

	for i := 0; i < len(graph); i++ {
		neighbours = nil

		if graph[i].word == end {
			getSolutionArray(solution, &graph[i])
			return true
		} else {
			neighbours = createNeighbours(graph[i].word, dictionary)

			for key := range neighbours {
				if _, exist := wordsRecord[key]; !exist {
					wordsRecord[key] = struct{}{}
					graph = append(graph, Node{word: key, previousNode: &graph[i]})
				}
			}
		}

	}
	// If there is no path between the words you can print the words record to see the words found
	//fmt.Println(wordsRecord)
	return false
}
