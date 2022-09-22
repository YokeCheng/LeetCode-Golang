package exercise

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		data := strings.Split(input.Text(), " ")
		sort.Strings(data)
		fmt.Println(strings.Join(data, " "))
	}
}
