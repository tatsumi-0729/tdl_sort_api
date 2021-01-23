package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func main() {
	file, err := os.Open("./list.csv")
	if err != nil {
		log.Fatalln("csvの読み込みに失敗しました。")
	}
	defer file.Close()

	var line []string
	reader := csv.NewReader(file)
	result := map[string]int{}

	for {
		line, err = reader.Read()
		if err != nil {
			break
		}
		tmpKey := ""
		tmpValue := ""
		valFlag := false
		for _, l := range line[0] {
			if valFlag {
				if l == '.' {
					break
				}
				tmpValue += string(l)
				continue
			}
			if l == '-' || l == ' ' {
				result[tmpKey] = 1
				continue
			}
			if l == '$' {
				valFlag = true
				continue
			}
			tmpKey += string(l)
		}
		result[tmpKey], _ = strconv.Atoi(tmpValue)
	}

	p := make(PairList, len(result))

	i := 0
	for k, v := range result {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p)

	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, "$"+strconv.Itoa(k.Value)+".00")
	}
}
