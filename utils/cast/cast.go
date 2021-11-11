package cast

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"github.com/modern-go/reflect2"
	"github.com/go-ego/gpy"
	"unicode/utf8"
)

func StringSplitQut(strs string) []string {
	if len(strings.TrimSpace(strs)) < 1 {
		return make([]string, 0)
	}

	return strings.Split(strs, ",")
}

func StringsLen(str string) int {
	return utf8.RuneCountInString(str)
}


func ParseInt(s string) int64 {
	if strings.IndexByte(s, '.') != -1 {
		i, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return int64(i)
		}
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	}
	return 0
}

func ParseUint(s string) uint64 {
	if strings.IndexByte(s, '.') != -1 {
		i, err := strconv.ParseFloat(s, 64)
		if err == nil {
			return uint64(i)
		}
	}
	i, err := strconv.ParseUint(s, 10, 64)
	if err == nil {
		return i
	}
	return 0
}

func Int(value interface{}) int {
	return int(Int64(value))
}

func Int8(value interface{}) int8 {
	return int8(Int64(value))
}

func IntArrayToString(v []int) string {
	if len(v) < 1 {
		return ""
	}

	sBuf := strings.Builder{}
	for _, v1 := range v {
		sBuf.WriteString(fmt.Sprintf("%d,", v1))
	}
	str := sBuf.String()
	return str[:len(str)-1]
}

func Int64(value interface{}) int64 {
	switch realValue := value.(type) {
	case int:
		return int64(realValue)
	case int8:
		return int64(realValue)
	case int16:
		return int64(realValue)
	case int32:
		return int64(realValue)
	case int64:
		return int64(realValue)
	case float32:
		return int64(realValue)
	case float64:
		return int64(realValue)
	case bool:
		if realValue {
			return 1
		} else {
			return 0
		}
	case []byte:
		return ParseInt(string(realValue))
	case string:
		return ParseInt(realValue)
	}
	return 0
}

func Uint(value interface{}) uint {
	return uint(Uint64(value))
}
func Uint64(value interface{}) uint64 {
	switch realValue := value.(type) {
	case uint:
		return uint64(realValue)
	case uint8:
		return uint64(realValue)
	case uint16:
		return uint64(realValue)
	case uint32:
		return uint64(realValue)
	case uint64:
		return uint64(realValue)
	case float32:
		return uint64(realValue)
	case float64:
		return uint64(realValue)
	case bool:
		if realValue {
			return 1
		} else {
			return 0
		}
	case []byte:
		return ParseUint(string(realValue))
	case string:
		return ParseUint(realValue)
	}
	return 0
}

func Float(value interface{}) float32 {
	return float32(Float64(value))
}
func Float64(value interface{}) float64 {
	switch realValue := value.(type) {
	case int:
		return float64(realValue)
	case int8:
		return float64(realValue)
	case int16:
		return float64(realValue)
	case int32:
		return float64(realValue)
	case int64:
		return float64(realValue)
	case float32:
		return float64(realValue)
	case float64:
		return float64(realValue)
	case bool:
		if realValue {
			return 1
		} else {
			return 0
		}
	case []byte:
		i, err := strconv.ParseFloat(string(realValue), 10)
		if err == nil {
			return i
		}
	case string:
		i, err := strconv.ParseFloat(realValue, 10)
		if err == nil {
			return i
		}
	}
	return 0
}

func Bytes(value interface{}) []byte {
	return []byte(String(value))
}

func String(value interface{}) string {
	switch realValue := value.(type) {
	case int:
		return strconv.FormatInt(int64(realValue), 10)
	case int8:
		return strconv.FormatInt(int64(realValue), 10)
	case int16:
		return strconv.FormatInt(int64(realValue), 10)
	case int32:
		return strconv.FormatInt(int64(realValue), 10)
	case int64:
		return strconv.FormatInt(realValue, 10)
	case bool:
		if realValue {
			return "true"
		} else {
			return "false"
		}
	case string:
		return realValue
	case []byte:
		return string(realValue)
	}
	t := reflect.TypeOf(value)
	if t != nil {
		if t.Kind() == reflect.Struct || t.Kind() == reflect.Map || t.Kind() == reflect.Slice {
			return Json(value)
		}
	}
	return fmt.Sprint(value)
}

func StringAnd(strs []string) string {
	if len(strs) > 0 {
		return strings.Join(strs, ",")
	}

	return ""
}

func ArrayToString(atrArray []string) string {
	value := ""
	for _, item := range atrArray {
		value += item
	}
	return value
}

func Bool(value interface{}) bool {
	switch realValue := value.(type) {
	case int, int8, int16, int32, int64, float32, float64:
		return realValue != 0
	case bool:
		return realValue
	case []byte:
		switch string(realValue) {
		case "1", "t", "T", "true", "TRUE", "True":
			return true
		}
	case string:
		switch realValue {
		case "1", "t", "T", "true", "TRUE", "True":
			return true
		}
	}
	return false
}

func Ints(value interface{}) []int64 {
	switch realValue := value.(type) {
	case []interface{}:
		result := make([]int64, len(realValue))
		for i, v := range realValue {
			result[i] = Int64(v)
		}
		return result
	default:
		return []int64{Int64(value)}
	}
}

func Floats(value interface{}) []float64 {
	switch realValue := value.(type) {
	case []interface{}:
		result := make([]float64, len(realValue))
		for i, v := range realValue {
			result[i] = Float64(v)
		}
		return result
	default:
		return []float64{Float64(value)}
	}
}

func Strings(value interface{}) []string {
	switch realValue := value.(type) {
	case []interface{}:
		result := make([]string, len(realValue))
		for i, v := range realValue {
			result[i] = String(v)
		}
		return result
	default:
		return []string{String(value)}
	}
}

func StringToIntArray(value string) []int {
	intArray := make([]int, 0)
	if len(value) < 1 {
		return intArray
	}

	ms := StringSplitQut(value)
	for _, v := range ms {
		vInt, _ := strconv.Atoi(v)
		intArray = append(intArray, vInt)
	}

	return intArray
}

func ValueToString(v reflect.Value) string {
	if v.Kind() == reflect.String {
		return v.String()
	} else {
		return fmt.Sprint(v)
	}
}

func GetLowerName(s string) string {
	buf := []byte(s)
	if buf[0] >= 'A' && buf[0] <= 'Z' {
		buf[0] += 32
	}
	return string(buf)
}

func GetUpperName(s string) string {
	buf := []byte(s)
	if buf[0] >= 'a' && buf[0] <= 'z' {
		buf[0] -= 32
	}
	return string(buf)
}

func FixUpperCase(data []byte, excludesKeys []string) {
	n := len(data)
	types := make([]bool, 0)
	keys := make([]string, 0)
	tpos := -1

	for i := 0; i < n-1; i++ {
		if tpos+1 >= len(types) {
			types = append(types, false)
			keys = append(keys, "")
		}

		if data[i] == '{' {
			tpos++
			types[tpos] = true
			keys[tpos] = ""
			//log.Println(" >>>1 ", types, tpos)
		} else if data[i] == '}' {
			tpos--
			//log.Println(" >>>2 ", types, tpos)
		}
		if data[i] == '[' {
			tpos++
			types[tpos] = false
			keys[tpos] = ""
			//log.Println(" >>>3 ", types, tpos)
		} else if data[i] == ']' {
			tpos--
			//log.Println(" >>>4 ", types, tpos)
		}
		if data[i] == '"' {
			keyPos := -1
			if i > 0 && (data[i-1] == '{' || (data[i-1] == ',' && tpos >= 0 && types[tpos])) {
				keyPos = i + 1
			}
			// skip string
			i++
			for ; i < n-1; i++ {
				if data[i] == '\\' {
					i++
					continue
				}
				if data[i] == '"' {
					if keyPos >= 0 && excludesKeys != nil {
						keys[tpos] = string(data[keyPos:i])
					}
					break
				}
			}

			if keyPos >= 0 && (data[keyPos] >= 'A' && data[keyPos] <= 'Z') {
				if excludesKeys != nil {
					// 是否排除
					excluded := false
					for _, ek := range excludesKeys {
						for j := tpos - 1; j >= 0; j-- {
							if strings.Index(keys[j], ek) != -1 {
								excluded = true
								break
							}
						}
						if excluded {
							break
						}
					}
					if !excluded {
						keyStr := keys[tpos]
						hasLower := false
						for c := len(keyStr) - 1; c >= 0; c-- {
							if keyStr[c] >= 'a' && keyStr[c] <= 'z' {
								hasLower = true
								break
							}
						}
						// 不转换全大写的Key
						if hasLower {
							data[keyPos] += 32
						}
					}
				} else {
					// 不进行排除判断
					hasLower := false
					dataLen := len(data)
					for c := keyPos; c < dataLen; c++ {
						if data[c] == '"' {
							break
						}
						if data[c] >= 'a' && data[c] <= 'z' {
							hasLower = true
							break
						}
					}
					// 不转换全大写的Key
					if hasLower {
						data[keyPos] += 32
					}
				}
			}
			continue
		}
	}
}

func If(i bool, a, b interface{}) interface{} {
	if i {
		return a
	}
	return b
}

func StringIf(i bool, a, b string) string {
	if i {
		return a
	}
	return b
}

func Switch(i uint, args ...interface{}) interface{} {
	if i < uint(len(args)) {
		return args[i]
	}
	return nil
}

func StringIn(arr []string, s string) bool {
	for _, d := range arr {
		if d == s {
			return true
		}
	}
	return false
}

func In(arr []interface{}, s interface{}) bool {
	for _, d := range arr {
		if d == s {
			return true
		}
	}
	return false
}

func SplitTrim(s, sep string) []string {
	ss := strings.Split(s, sep)
	for i, s1 := range ss {
		ss[i] = strings.TrimSpace(s1)
	}
	return ss
}

func Json(value interface{}) string {
	j, err := json.Marshal(value)
	if err == nil {
		return string(j)
	}
	return fmt.Sprint(value)
}

func JsonP(value interface{}) string {
	j, err := json.MarshalIndent(value, "", "  ")
	if err == nil {
		return string(j)
	}
	return fmt.Sprint(value)
}

func PageSize(pageSize int) int {
	if pageSize < 20 {
		return 20
	}

	if pageSize > 50 {
		return 50
	}

	return pageSize
}

func PageIndex(pageIndex int) int {
	if pageIndex < 1 {
		return 1
	}
	return pageIndex
}

func IsNil(i interface{}) bool {
	return reflect2.IsNil(i)
}

//String To array
func StringToArray(value string) []string {
	if len(value) < 1 {
		return make([]string, 0)
	}
	arrS := strings.Split(value, ",")
	return arrS
}

func StrArrayToString(values []string) string {
	return strings.Join(values, ",")
}

func IsContainsHanzi(param string) bool {
	for _, r := range param {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func GetCityPinYin(param string) string {
	values := gpy.LazyPinyin(param)
	result := ""
	for _, v := range values {
		result = result + v
	}

	return result
}

func GetCityFirstLetter(param string) string {
	values := gpy.LazyPinyin(param)
	result := ""
	for _, v := range values {
		result = result + v
	}

	return strings.ToUpper(string(result[0]))
}

func CastHanziToPy(param string) []string {
	values := make([]string, 0)
	if len(param) < 1 {
		return values
	}

	args := gpy.Pinyin(param, gpy.Args{
		Heteronym: true,
		Separator: "",
		Fallback: func(r rune, a gpy.Args) []string {
			return []string{string(r)}
		},
	})

	readyArg := make([][]string, 0)
	for _, a := range args {
		a = removeRepByLoop(a)
		readyArg = append(readyArg, a)
	}

	for _, a := range readyArg {
		if len(values) < 1 {
			values = append(values, "")
		}

		if len(a) == 1 {
			for i := 0; i < len(values); i++ {
				values[i] = values[i] + a[0]
			}
		} else {
			valsTemp := make([]string, 0)
			for i := 0; i < len(a); i++ {
				for j := 0; j < len(values); j++ {
					valsTemp = append(valsTemp, values[j]+a[i])
				}
			}
			values = valsTemp
		}
	}

	return values
}

func removeRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
