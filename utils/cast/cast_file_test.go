package cast

import (
	"fmt"
	"testing"
)

func TestCastString(t *testing.T) {
	s := []string{"a b","b","c"}
	fmt.Println(String(s))
}

func TestCastHanziToPy(t *testing.T) {
	//str := "aaaa中国"
	//var hzRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")
	//fmt.Println(hzRegexp.MatchString(str))
	fmt.Println(CastHanziToPy("afanda番茄中国"))
}

