package data

import (
	"fmt"
	"strconv"
)

// Declare a custom runtime type, has the same underlying type as our movie struct (int32)
type Runtime int32

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
