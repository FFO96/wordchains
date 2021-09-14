package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestCreateNeighbours(t *testing.T) {

	expectedNeighbours := make(map[string]struct{})
	dictionary := make(map[string]struct{})

	// Setting the expected neighbours for the word "dog"
	data := "data/expected_neighbours"

	file, err := os.Open(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		expectedNeighbours[strings.ToLower(scanner.Text())] = struct{}{}
	}
	// Setting the dictionary
	data = "data/dictionary"

	file, err = os.Open(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = struct{}{}
	}

	// Comparing solution and expected solution
	if !reflect.DeepEqual(createNeighbours("dog", dictionary), expectedNeighbours) {
		t.Fail()
	}
}

func TestGetSolutionArray(t *testing.T) {

	// Setting the expected solution
	expectedSolution := []string{"dog", "dot", "cot", "cat"}

	// Setting the graph parameter
	testGraph := Node{word: "cat", previousNode: &Node{word: "cot", previousNode: &Node{word: "dot", previousNode: &Node{word: "dog", previousNode: nil}}}}
	var solution []string

	getSolutionArray(&solution, &testGraph)

	// Comparing solution and expected solution
	if !reflect.DeepEqual(solution, expectedSolution) {
		t.Fail()
	}
}

func TestExistInDictionary(t *testing.T) {

	// Setting the dictionary
	dictionary := make(map[string]struct{})
	data := "data/dictionary"

	file, err := os.Open(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = struct{}{}
	}

	// Testing existing word in dictionary. Expected: TRUE
	if existInDictionary("dog", dictionary) == false {
		t.Fail()
	}

	// Testing not existing word in dictionary. Expected: FALSE
	if existInDictionary("perro", dictionary) == true {
		t.Fail()
	}
}

func TestFindPath(t *testing.T) {

	var solution []string
	expectedSolution := []string{"dog", "dot", "cot", "cat"}

	// Setting the dictionary
	dictionary := make(map[string]struct{})
	data := "data/dictionary"

	file, err := os.Open(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		dictionary[strings.ToLower(scanner.Text())] = struct{}{}
	}

	// Testing expected TRUE and expected solution
	found := findPath("dog", "cat", dictionary, &solution)
	if found == false || !reflect.DeepEqual(solution, expectedSolution) {
		t.Fail()
	}

	// Testing expected FALSE
	found = findPath("paratrooper", "cat", dictionary, &solution)
	if found == true {
		t.Fail()
	}
}
