package encoding

import (
	b32 "encoding/base32"
	b64 "encoding/base64"
	"fmt"
)

func Parser(encMethod string, decode bool, encode bool, input string) string {
	if encMethod == "base64" {
		return base64Conv(decode, encode, input)
	}

	if encMethod == "base32" {
		return base32Conv(decode, encode, input)
	}

	return "Unsupported encoding method"

}

func base64Conv(decode bool, encode bool, input string) string {
	var result string

	if encode {
		result = b64.StdEncoding.EncodeToString([]byte(input + "\n"))
	} else if decode {
		data, err := b64.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error decoding base64 : ", err)
			return ""
		}
		result = string(data)
	}
	return result
}

func base32Conv(decode bool, encode bool, input string) string {
	var result string

	if encode {
		result = b32.StdEncoding.EncodeToString([]byte(input + "\n"))
	} else if decode {
		data, err := b32.StdEncoding.DecodeString(input)
		if err != nil {
			fmt.Println("Error decoding base32 : ", err)
			return ""
		}
		result = string(data)
	}
	return result
}
