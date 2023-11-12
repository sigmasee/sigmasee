package converters

import (
	"fmt"
	"strings"
	"time"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func PointerStrToStr(ptr *string) string {
	if ptr == nil {
		return ""
	}

	return *ptr
}

func StrToPointerStr(str string) *string {
	if len(str) == 0 {
		return nil
	}

	return &str
}

func PointerBoolToBool(ptr *bool) bool {
	if ptr == nil {
		return false
	}

	return *ptr
}

func BoolToPointerBool(val bool) *bool {
	return &val
}

func PointerIntToInt(ptr *int) int {
	if ptr == nil {
		return 0
	}

	return *ptr
}

func IntToPointerInt(val int) *int {
	return &val
}

func PointerInt8ToInt8(ptr *int8) int8 {
	if ptr == nil {
		return 0
	}

	return *ptr
}

func Int8ToPointerInt8(val int8) *int8 {
	return &val
}

func PointerInt16ToInt16(ptr *int16) int16 {
	if ptr == nil {
		return 0
	}

	return *ptr
}

func Int16ToPointerInt16(val int16) *int16 {
	return &val
}

func PointerInt32ToInt32(ptr *int32) int32 {
	if ptr == nil {
		return 0
	}

	return *ptr
}

func Int32ToPointerInt32(val int32) *int32 {
	return &val
}

func PointerIntToInt32(ptr *int) int32 {
	if ptr == nil {
		return 0
	}

	return int32(*ptr)
}

func Int32ToPointerInt(ptr int32) *int {
	val := int(ptr)

	return &val
}

func TimeToPointerTime(time time.Time) *time.Time {
	if time.IsZero() {
		return nil
	}

	return &time
}

func PointerTimeToPointerProtoTimestamp(time *time.Time) *timestamppb.Timestamp {
	if time != nil && !time.IsZero() {
		return timestamppb.New(*time)
	}

	return nil
}

func PointerProtoTimestampToTime(protoTimestamp *timestamppb.Timestamp) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, protoTimestamp.AsTime().Format(time.RFC3339Nano))
}

func PointerProtoTimestampToPointerTime(protoTimestamp *timestamppb.Timestamp) (*time.Time, error) {
	if protoTimestamp == nil {
		return nil, nil
	}

	convertedTime, err := PointerProtoTimestampToTime(protoTimestamp)
	if err != nil {
		return nil, err
	}

	return TimeToPointerTime(convertedTime), nil
}

func ToRelayID(domainName string, id string) string {
	return fmt.Sprintf("%s%s", domainName, id)
}

func FromRelayID(domainName string, id string) (string, error) {
	if strings.HasPrefix(id, domainName) {
		return strings.TrimPrefix(id, domainName), nil
	}

	return "", fmt.Errorf("id must start with \"%s\"", domainName)
}
