package utils

import (
	"strconv"
	"strings"
)

func StrToUint32(idStr string) (uint32, error) {
	idU64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}

	id := uint32(idU64)
	return id, nil
}

func LowerStrArr(arr []string) []string {
	res := make([]string, len(arr))
	for i, name := range arr {
		res[i] = strings.ToLower(name)
	}
	return res
}
