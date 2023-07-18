package map_char

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

type Content_Set struct {
	Characters []struct {
		Name string `json:"name"`
		Set  string `json:"set"`
	} `json:"content"`
}

func Fill(name string) *Content_Set {
	var c *Content_Set
	file, err := os.Open(name)
	if err != nil {
		return nil
	}
	bytes, _ := ioutil.ReadAll(file)
	if json.Unmarshal(bytes, &c) != nil {
		return nil
	}
	return c
}

func generateUniqueRandomNumbers(n, max int) []int {
	rand.Seed(time.Now().UnixNano())
	set := make(map[int]bool)
	var result []int
	for len(set) < n {
		value := rand.Intn(max)
		if !set[value] {
			set[value] = true
			result = append(result, value)
		}
	}
	return result
}

func (c *Content_Set) Take_rand(value int) string {
	var ret string
	index := generateUniqueRandomNumbers(value, len(c.Characters))
	for i := 0; i < value; i++ {
		ret += c.Characters[index[i]].Name + "\n"
	}
	return ret
}
