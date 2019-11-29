package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func freqQuery(queries [][]int32) []int32 {
	data := make(map[int32]int32)
	freqSummary := make(map[int32]int32)
	freq := make([]int32, 0)
	var count int32

	for _, value := range queries {
		switch value[0] {
		case 1:
			// see if number exists
			if _, ok := data[value[1]]; ok {
				// see if frequency is in freqSummary
				if _, ok := freqSummary[data[value[1]]]; ok {
					// reduce frequency if not zero
					if freqSummary[data[value[1]]] > 0 {
						freqSummary[data[value[1]]]--
					}
				}
			}
			data[value[1]]++
			freqSummary[data[value[1]]]++
		case 2:
			if _, ok := data[value[1]]; ok {
				if data[value[1]] > 0 {
					if _, ok := freqSummary[data[value[1]]]; ok {
						if freqSummary[data[value[1]]] > 0 {
							freqSummary[data[value[1]]]--
						}
					}
					data[value[1]]--
					freqSummary[data[value[1]]]++
				}
			}
		case 3:
			count++
			if _, ok := freqSummary[value[1]]; ok {
				if freqSummary[value[1]] > 0 {
					freq = append(freq, 1)
				} else {
					freq = append(freq, 0)
				}
			} else {
				freq = append(freq, 0)
			}
		default:
			fmt.Println("encountered unexpected data!!! something went wrong.")
		}
	}

	return freq
}

func main() {
	file, err := os.Open("test-case-4-debug")
	checkError(err)

	reader := bufio.NewReaderSize(file, 16*1024*1024)

	stdout, err := os.Create("OUTPUT-test-case-4")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	var queries [][]int32
	for i := 0; i < int(q); i++ {
		queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != 2 {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	ans := freqQuery(queries)

	for i, ansItem := range ans {
		fmt.Fprintf(writer, "%d", ansItem)

		if i != len(ans)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
