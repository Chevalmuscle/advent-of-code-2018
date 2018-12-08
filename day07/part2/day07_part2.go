package day07

import (
	"regexp"
	"sort"
	"strconv"
	"time"

	"../../utils"
)

type work struct {
	task     string
	timeLeft int
}

func instructionOrder(steps []string, workersAmount int) int {
	defer utils.TimeTaken(time.Now())

	var workers = make(map[int]work)
	var tasksDone = make(map[string]bool)

	var queue = make([]string, 0)
	var current = make(map[string]bool)

	for i := 0; i < workersAmount; i++ {
		workers[i] = work{task: "", timeLeft: 0}
	}
	var time = -1

	var historique = make(map[int][]string)

	var graph = make(map[string][]string)
	var freeNodes = make([]string, 0)
	var lastFreeNode = ""

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
		time++
		for node, childrens := range graph {
			for index, children := range childrens {
				if _, isDone := tasksDone[children]; isDone {
					graph[node] = append(graph[node][:index], graph[node][index+1:]...)
				}
			}

			if len(graph[node]) == 0 {
				freeNodes = append(freeNodes, node)
				if !contains(queue, node) {
					if _, isDone := tasksDone[node]; !isDone {
						if _, isDoing := current[node]; !isDoing {
							queue = append([]string{node}, queue...)
						}
					}
				}
			}
		}

		if len(freeNodes) > 0 {
			sort.Strings(freeNodes)
			lastFreeNode = freeNodes[0]
			delete(graph, lastFreeNode)
			freeNodes = make([]string, 0)
		}

		if len(queue) > 0 {
			sort.Strings(queue)

			for worker, work := range workers {
				if work.timeLeft <= 0 { //free worker
					for i, task := range queue {
						current[task] = true
						work.task = task
						work.timeLeft = getLengthOfTask(task)
						workers[worker] = work
						historique[worker] = append(historique[worker], task+strconv.Itoa(work.timeLeft)+" "+strconv.Itoa(time))
						queue = append(queue[:i], queue[i+1:]...)
						break
					}
				}
			}
		}

		for worker, work := range workers {
			work.timeLeft--
			workers[worker] = work

			if work.timeLeft == 0 {
				tasksDone[work.task] = true
				delete(current, work.task)
			}
		}

	}

	var workIsDone = false
	for !workIsDone {

		time++
		workIsDone = true
		for worker, work := range workers {
			if work.timeLeft > 0 {
				work.timeLeft--
				workers[worker] = work
				workIsDone = false
			}
		}
	}
	return time
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
