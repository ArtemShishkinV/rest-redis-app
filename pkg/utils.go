package pkg

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}

func ComputeHmac512(message string, secret string) string {
	key := []byte(secret)

	h := hmac.New(sha512.New, key)
	h.Write([]byte(message))

	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func GetNumbersFromString(data string) ([]int, error) {
	var numbers []int
	var matchNumbers []string

	for _, s := range strings.Split(data, "\r\n") {
		matchNumbers = append(matchNumbers, strings.Split(s, ",")...)
	}

	for _, element := range matchNumbers {
		if element != "" {
			i, err := strconv.Atoi(element)
			if err != nil {
				return nil, err
			}
			numbers = append(numbers, i)
		}
	}

	return numbers, nil
}
