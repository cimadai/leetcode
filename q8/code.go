package q8

import (
	"math"
	"strconv"
	"strings"
	"regexp"
)

func myAtoi(str string) int {
	tmp := strings.TrimSpace(str)

	base := 10
	if strings.HasPrefix(tmp, "0x") {
		rep := regexp.MustCompile(`([+-]?)0x([0-9a-fA-F]*).*`)
		tmp = rep.ReplaceAllString(tmp, "$1$2")
		base = 16
	} else if strings.HasPrefix(tmp, "0b") {
		rep := regexp.MustCompile(`([+-]?)0b([01]*).*`)
		tmp = rep.ReplaceAllString(tmp, "$1$2")
		base = 2
	} else {
		rep := regexp.MustCompile(`([+-]?[0-9]*).*`)
		tmp = rep.ReplaceAllString(tmp, "$1")
	}

	if val, err := strconv.ParseInt(tmp, base, 64); err != nil {
		if strings.HasSuffix(err.Error(), "value out of range") {
			if strings.HasPrefix(tmp, "-") {
				return math.MinInt32
			} else {
				return math.MaxInt32
			}
		} else {
			return 0
		}
	} else {
		ival := int(val)
		if math.MaxInt32 < ival {
			return math.MaxInt32
		} else if ival < math.MinInt32 {
			return math.MinInt32
		} else {
			return ival
		}
	}

}
