package main

import (
"bufio"
"fmt"
"io"
"os"
"strconv"
"strings"
)

// https://www.hackerrank.com/challenges/insert-a-node-at-a-specific-position-in-a-linked-list/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=linked-lists

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode {
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string, writer *bufio.Writer) {
	for node != nil {
		fmt.Fprintf(writer, "%d", node.data)

		node = node.next

		if node != nil {
			fmt.Fprintf(writer, sep)
		}
	}
}

// Complete the insertNodeAtPosition function below.

/*
 * For your reference:
 *
 * SinglyLinkedListNode {
 *     data int32
 *     next *SinglyLinkedListNode
 * }
 *
 */
func insertNodeAtPosition(llist *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {

	if position == 0 {
		return &SinglyLinkedListNode{
			data: data,
			next: llist,
		}
	}

	nxt := llist.next
	for i := 1; i < int(position-1); i++ {
		nxt = nxt.next
	}

	tmp := nxt.next
	nxt.next = &SinglyLinkedListNode{
		data: data,
		next: tmp,
	}

	return llist;
}

func insertNodeAtPositionApp() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	writer := bufio.NewWriterSize(os.Stdout, 1024 * 1024)

	llistCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	llist := SinglyLinkedList{}
	for i := 0; i < int(llistCount); i++ {
		llistItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		llistItem := int32(llistItemTemp)
		llist.insertNodeIntoSinglyLinkedList(llistItem)
	}

	dataTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	data := int32(dataTemp)

	positionTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	position := int32(positionTemp)

	llist_head := insertNodeAtPosition(llist.head, data, position)

	printSinglyLinkedList(llist_head, " ", writer)
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

