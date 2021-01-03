package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/triple-sum/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=search

func removeDuplicates(sl []int32) []int32 {
	addedMap := make(map[int32]bool)
	var resSlice []int32

	for _, el := range sl {
		if b, ok := addedMap[el]; !ok && !b {
			resSlice = append(resSlice, el)
			addedMap[el] = true
		}
	}

	return resSlice
}

// Complete the triplets function below.
func triplets(a []int32, b []int32, c []int32) int64 {
	a = removeDuplicates(a)
	b = removeDuplicates(b)
	c = removeDuplicates(c)

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(c, func(i, j int) bool {
		return c[i] < c[j]
	})

	a_len := len(a)
	c_len := len(c)

	res := 0

	for _, be := range b {
		next_a := a_len/2
		s_a := 0
		e_a := a_len-1
		for e_a >= s_a {
			if a[next_a] <= be {
				s_a = next_a+1
			} else {
				e_a = next_a-1
			}
			next_a = s_a + (e_a-s_a)/2
		}

		next_c := c_len/2
		s_c := 0
		e_c := c_len - 1
		for e_c >= s_c {
			if c[next_c] <= be {
				s_c = next_c+1
			} else {
				e_c = next_c-1
			}
			next_c = s_c + (e_c-s_c)/2
		}

		res += (next_a) * (next_c)

	}
	return int64(res)
}

func TripleSumApp() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	lenaLenbLenc := strings.Split(readLine(reader), " ")

	lenaTemp, err := strconv.ParseInt(lenaLenbLenc[0], 10, 64)
	checkError(err)
	lena := int32(lenaTemp)

	lenbTemp, err := strconv.ParseInt(lenaLenbLenc[1], 10, 64)
	checkError(err)
	lenb := int32(lenbTemp)

	lencTemp, err := strconv.ParseInt(lenaLenbLenc[2], 10, 64)
	checkError(err)
	lenc := int32(lencTemp)

	arraTemp := strings.Split(readLine(reader), " ")

	var arra []int32

	for i := 0; i < int(lena); i++ {
		arraItemTemp, err := strconv.ParseInt(arraTemp[i], 10, 64)
		checkError(err)
		arraItem := int32(arraItemTemp)
		arra = append(arra, arraItem)
	}

	arrbTemp := strings.Split(readLine(reader), " ")

	var arrb []int32

	for i := 0; i < int(lenb); i++ {
		arrbItemTemp, err := strconv.ParseInt(arrbTemp[i], 10, 64)
		checkError(err)
		arrbItem := int32(arrbItemTemp)
		arrb = append(arrb, arrbItem)
	}

	arrcTemp := strings.Split(readLine(reader), " ")

	var arrc []int32

	for i := 0; i < int(lenc); i++ {
		arrcItemTemp, err := strconv.ParseInt(arrcTemp[i], 10, 64)
		checkError(err)
		arrcItem := int32(arrcItemTemp)
		arrc = append(arrc, arrcItem)
	}

	ans := triplets(arra, arrb, arrc)

	fmt.Fprintf(writer, "%d\n", ans)

	writer.Flush()
}
