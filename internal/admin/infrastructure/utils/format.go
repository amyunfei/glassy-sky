package utils

import (
	"strconv"
	"strings"
)

func IntToHexColor[T int | int32 | int64](color T) string {
	return "#" + strconv.FormatInt(int64(color), 16)
}

func HexColorToInt[T int | int32 | int64](colorStr string) (T, error) {
	i, err := strconv.ParseInt(strings.ReplaceAll(colorStr, "#", ""), 16, 64)
	return T(i), err
}
