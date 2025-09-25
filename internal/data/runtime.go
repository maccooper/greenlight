package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Error to be returned if UnmarshalJSON() cannot parse or convert to json-string correctly
var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

// Declare a custom runtime type, has the same underlying type as our movie struct (int32)
type Runtime int32

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	//parse the incoming json to remove double quotes
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	//Split the string to isolate the number
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	// Convert the int32 to a Runtime type and assign it to the receiver.
	*r = Runtime(i)

	return nil

}

// custom marshaljson method on runtime type, should return the json-encoded val for the movie runtime.
// in our case, a string of the format "<runtime> mins"
func (r Runtime) MarshalJSON() ([]byte, error) {

	//int32 to str conversion
	jsonValue := fmt.Sprintf("%d mins", r)

	//wraps string in double quotes, which is a requirement to be a valid json string
	quotedJsonValue := strconv.Quote(jsonValue)

	//convert str to byte slice
	return []byte(quotedJsonValue), nil

}
