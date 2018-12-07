package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test the shit that waits upto thankful eat your cunt", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.Join([]string{"I have a very", "large penius"}, "-"))
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	var buf bytes.Buffer
	buf.Write([]byte("test"))
	fmt.Println(buf)
	file, err := os.Open("stringTests.go")
	if err != nil {
		fmt.Println("File does not exist")
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return
	}
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
	bs2, err := ioutil.ReadFile("stringTests.go")
	if err != nil {
		return
	}
	str2 := string(bs2)
	fmt.Println(str2)

}
