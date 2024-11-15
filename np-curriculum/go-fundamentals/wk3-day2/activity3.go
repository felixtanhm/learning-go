package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	chnl1 := make(chan string, 1)
	chnl2 := make(chan string, 5)
	chnl3 := make(chan string, 1)

	go write1(chnl1)
	go write2(chnl2)
	go write3(chnl3)
	for {
		select {
		case value1, ok := <-chnl1:
			if !ok {
				chnl1 = nil
			} else {
				fmt.Println("Value 1 is:\n", value1)
				fmt.Println("---------\n")
			}
		case value2, ok := <-chnl2:
			if !ok {
				chnl2 = nil
			} else {
				fmt.Println("Value 2 is:\n", value2)
			}
		case value3, ok := <-chnl3:
			if !ok {
				chnl3 = nil
			} else {
				fmt.Println("Value 3 is:\n", value3)
				fmt.Println("---------\n")
			}
		case <-time.After(30 * time.Second):
			fmt.Println("Timeout: File reads took longer than 30 seconds.")
		}
		if chnl1 == nil && chnl2 == nil && chnl3 == nil {
			break
		}
	}

}

func write1(channel chan string) {
	fileName := "text1.txt"
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("Error reading file:", err)
		return
	}
	channel <- string(fileData)
	close(channel)
}

func write2(channel chan string) {
	fileName := "text2.txt"
	openedFile, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error reading file:", err)
		return
	}
	defer openedFile.Close()
	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		channel <- scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error reading file:", err)
		return
	}
	close(channel)
}

func write3(channel chan string) {
	fileName := "text3.txt"
	list := []string{}
	openedFile, err := os.Open(fileName)
	if err != nil {
		fmt.Errorf("Error reading file:", err)
		return
	}
	defer openedFile.Close()
	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Errorf("Error reading file:", err)
		return
	}
	channel <- strings.Join(list, ", ")
	close(channel)
}
