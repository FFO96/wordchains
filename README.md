# Wordchains
The word-chain puzzle is a type of puzzle where the challenge is to build a chain of words, starting with one particular word and ending with another. Successive entries in the chain must all be real words, and each can differ from the previous word by just one letter. For example, you can get from “dog” to “cat” using the following chain: 

**dog** :arrow_right: dot :arrow_right: cot :arrow_right: **cat**

This program needs to be set with a start word, an end word and a dictionary to determine that the words we go through actually exist. It is important to highlight that the start word and the end word must appear in the dictionary. Also, some words in the dictionary do not have [neighbours](#neighbour), so in some cases there may not be a solution.

This code has been developed to work even if the start and end words are not the same length. A single change is understood to be the modification, deletion or insertion of a letter.

## Problem solving
To solve the problem we have to focus on two key concepts: neighbours search and the way to go through the [graph](#graph).

#### Neighbor search
To find the neighbors, the word is modified by adding, deleting or replacing a letter in each of the positions of the word string, with all the letters of the alphabet.  Once the string is formed, it is checked that it appears in the dictionary and if it does not, it is excluded. If the word is valid, it is added to the graph to be expanded if necessary.

#### Path of the network
The decision to go through the graph with a [breadth-first search](#bfs) is to find the shortest path. In a [depth-first search](#dfs) we may find a path but we do not know for sure if it is the shortest.

## Keywords

- <a id="neighbour">**Neighbour**</a>: A word that differ from the previous word by just one letter.
- <a id="graph">**Graph**</a>: Network of words that express the connection between a word and its neighbours.
- <a id="bfs">**Breadth-first search**</a>: Exploration of all nodes at the present depth prior to moving on to the nodes at the next depth level. 
- <a id="dfs">**Depth-first search**</a>: Exploration of a full branch prior to visit the same level nodes.

## Run the progam

- **Normal program**</a>: go run src/main.go
- **Testing**</a>: go test src/main_test.go src/main.go -v

## Reference
Word-chain problem: [http://codekata.com/kata/kata19-word-chains/](http://codekata.com/kata/kata19-word-chains/)
