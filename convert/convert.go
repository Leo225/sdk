package convert

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go/types"
	"strconv"
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprint(v)
}

func ToBool(v interface{}) bool {
	switch b := v.(type) {
	case bool:
		return b
	case types.Nil:
		return false
	case int:
		if v.(int) != 0 {
			return true
		}
		return false
	case string:
		r, err := strconv.ParseBool(ToString(v))
		if err != nil {
			return false
		}
		return r
	default:
		return false
	}
}

func ToInt(v interface{}) int {
	return int(ToInt64(v))
}

func ToInt32(v interface{}) int32 {
	return int32(ToInt64(v))
}

func ToInt64(v interface{}) int64 {
	num, err := strconv.ParseInt(ToString(v), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func ToUint(v interface{}) uint {
	return uint(ToUint64(v))
}

func ToUint32(v interface{}) uint32 {
	return uint32(ToUint64(v))
}

func ToUint64(v interface{}) uint64 {
	num, err := strconv.ParseUint(ToString(v), 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func ToFloat32(v interface{}) float32 {
	return float32(ToFloat64(v))
}

func ToFloat64(v interface{}) float64 {
	num, err := strconv.ParseFloat(ToString(v), 64)
	if err != nil {
		return 0
	}
	return num
}

func BytesToInt(data []byte) int {
	return int(BytesToInt32(data))
}

func BytesToInt32(data []byte) int32 {
	var num int32
	buffer := bytes.NewBuffer(data)
	binary.Read(buffer, binary.BigEndian, &num)
	return num
}

func BytesToInt64(data []byte) int64 {
	var num int64

	buffer := bytes.NewBuffer(data)
	binary.Read(buffer, binary.BigEndian, &num)
	return num
}
