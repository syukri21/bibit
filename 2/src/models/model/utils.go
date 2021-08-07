package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type ConvertibleBoolean bool

func (bit *ConvertibleBoolean) UnmarshalJSON(data []byte) error {
	asString := strings.ToLower(strings.ReplaceAll(string(data), "\"", ""))
	if asString == "1" || asString == "true" {
		*bit = true
	} else if asString == "0" || asString == "false" {
		*bit = false
	} else {
		return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
	}
	return nil
}

type ConvertibleInt int64

func (bit ConvertibleInt) UnmarshalJSON(data []byte) error {
	asString := strings.ToLower(strings.ReplaceAll(string(data), "\"", ""))
	result, err := strconv.ParseInt(asString, 10, 64)

	if err != nil {
		return err
	}

	bit = ConvertibleInt(result)
	return nil
}
