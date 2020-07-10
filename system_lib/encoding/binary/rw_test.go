package binary

import (
	"bufio"
	"encoding/binary"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	rd := bufio.NewReader(os.Stdin)
	//rd := strings.NewReader("hello world")
	buffer := make([]byte,255)
	i := 0
	for {
		err := binary.Read(rd, binary.BigEndian, &buffer[i])
		if err != nil {
			t.Error(err, "--- buffer:", string(buffer))
			return
		}
		i++
	}
	t.Log("success", "--- buffer:",string(buffer))
}