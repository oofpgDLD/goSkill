package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	//rd := strings.NewReader("hello world")
	buffer := make([]byte,255)
	i := 0
	for {
		err := binary.Read(rd, binary.BigEndian, &buffer[i])
		if err != nil {
			fmt.Println(err, "--- buffer:", string(buffer))
			return
		}
		i++
	}
	fmt.Println("success", "--- buffer:",string(buffer))
}
