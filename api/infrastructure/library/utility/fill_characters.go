package utility

import (
	"fmt"
	"strings"
)

func FillZeroCharacters(length int, target string) (string, error) {
	if length < len(target) {
		return "", fmt.Errorf("\"%s\" is longer than %d characters.", target, length)
	}

	arrTargetStr := strings.Split(target, "")
	arrResultStr := make([]string, length)
	argLength := len(target)
	zeroLength := length - argLength

	for j := 0; j < zeroLength; j++ {
		arrResultStr = append(arrResultStr, "0")
	}
	for _, v := range arrTargetStr {
		arrResultStr = append(arrResultStr, v)
	}

	return strings.Join(arrResultStr, ""), nil
}
