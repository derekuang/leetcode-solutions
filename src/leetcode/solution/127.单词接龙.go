/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

package solution

import "math"

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !isContain(wordList, endWord) {
		return 0
	}
	wordList = append(append([]string{}, beginWord), wordList...)
	distance := make([][]int, len(wordList))
	for i := 0; i < len(distance); i++ {
		distance[i] = make([]int, len(wordList))
	}
	for i := 1; i < len(wordList); i++ {
		for j := 0; j < i; j++ {
			if canBeConvert(wordList[i], wordList[j]) {
				distance[j][i] = 1
				distance[i][j] = 1
			}
		}
	}
	distMap := map[string]int{}
	for i := 1; i < len(wordList); i++ {
		distMap[wordList[i]] = math.MaxInt64
	}
	distMap[beginWord] = 1
	queue := make(chan int, len(distance)-1)
	queue <- 0
	// BFS
	for len(queue) > 0 {
		p := <-queue
		d0 := distMap[wordList[p]]
		for q := 0; q < len(distance); q++ {
			if distance[p][q] != 0 && distance[p][q]+d0 < distMap[wordList[q]] {
				distMap[wordList[q]] = distance[p][q] + d0
				queue <- q
			}
		}
	}
	if distMap[endWord] == math.MaxInt64 {
		return 0
	}
	return distMap[endWord]
}

func isContain(wordList []string, word string) bool {
	flag := false
	for _, s := range wordList {
		if s == word {
			flag = true
		}
	}
	return flag
}

func canBeConvert(word1, word2 string) bool {
	cnt := 0
	for i := 0; i < len(word1); i++ {
		if word1[i] != word2[i] {
			cnt++
		}
	}
	return cnt == 1
}

// @lc code=end
