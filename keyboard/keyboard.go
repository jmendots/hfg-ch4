package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const TriangleSides int = 3

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%d \n", TriangleSides)
	return number, nil
}
