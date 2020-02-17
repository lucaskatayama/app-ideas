package bin2dec

import "fmt"

func Bin2Dec(bin string) (r string, err error) {
	n := len(bin) - 1
	result := 0
	for i, val := range bin {
		// rune 1 = 49
		if val == 49 {
			if n - i == 0 {
				result += 1
			} else {
				result += 2<<(n - i)
			}
		}
	}
	return fmt.Sprintf("%d", result), err
}
