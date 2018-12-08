package day07part1

import (
	"regexp"
	"sort"
	"strings"
	"time"

	"../../utils"
)

type work struct {
	task     string
	timeLeft int
}

func instructionOrder(steps []string) string {
	defer utils.TimeTaken(time.Now())

	var graph = make(map[string][]string)
	var freeNodes = make([]string, 0)
	var lastFreeNode = ""
	var instructions strings.Builder

	var re = regexp.MustCompile("([A-Z]\\b)")

	for _, step := range steps {
		var res = re.FindAllString(step, -1)
		var from = res[0]
		var to = res[1]
		graph[to] = append(graph[to], from)
		if _, exist := graph[from]; !exist {
			graph[from] = make([]string, 0)
		}
	}

	for len(graph) > 0 {
		for node, childrens := range graph {
			for index, children := range childrens {
				if children == lastFreeNode {
					graph[node] = append(graph[node][:index], graph[node][index+1:]...)
				}
			}

			if len(graph[node]) == 0 {
				freeNodes = append(freeNodes, node)
			}
		}

		if len(freeNodes) > 0 {
			sort.Strings(freeNodes)
			lastFreeNode = freeNodes[0]
			delete(graph, lastFreeNode)
			instructions.WriteString(lastFreeNode)
			freeNodes = make([]string, 0)
		}

	}

	return instructions.String()
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getLengthOfTask(task string) int {
	return int([]rune(task)[0] - 4)
}
