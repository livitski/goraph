// Package dot converts JSON graph data
// into a DOT (graph description language) file.
package dot

import (
	"log"
	"os"
	"regexp"
	"strconv"

	"github.com/gyuho/gson/jgd"
)

// DeleteNonAlnum removes all alphanumeric characters.
func DeleteNonAlnum(str string) string {
	// alphanumeric (== [0-9A-Za-z])
	// \s is a white space character
	validID := regexp.MustCompile(`[^[:alnum:]\s]`)
	return validID.ReplaceAllString(str, "")
}

// convert converts input JSON to DOT format.
func convert(inputfile, graph string) string {
	nodes := jgd.GetNodes(inputfile, graph)
	gmap := jgd.MapGraph(inputfile, graph)
	// map[string]map[string][]float64

	graphname := DeleteNonAlnum(graph)
	result := "digraph " + graphname + " {" + "\n"

	// traverse all nodes in graph
	for _, node := range nodes {
		nm, ok := gmap[node]
		// if the node has neighbor(outgoing vertex)
		if ok {
			// traverse the map of outgoing vertices
			for key, value := range nm {
				// traverse the slice of weights
				for _, w := range value {
					wt := strconv.FormatFloat(w, 'f', -1, 64)
					result += "\t" + node + " -> " + key + " [label=" + wt + "]" + "\n"
				}
			}
		} else {
			// if the node has no outgoing vertex
			result += "\t" + node + "\n"
		}
	}
	result += "}"
	return result
}

// Convert converts input JSON graph data to DOT file.
func Convert(inputfile, graph, outputfile string) {
	file, err := os.Create(outputfile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	str := convert(inputfile, graph)
	file.WriteString(str)
}
