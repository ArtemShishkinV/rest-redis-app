package pkg

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
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

	return string(h.Sum(nil))
}

func GetNumbersFromString(data string) []int {
	var numbers []int

	re := regexp.MustCompile("-?\\d+")
	matchNumbers := re.FindAllString(data, -1)

	for _, element := range matchNumbers {
		i, _ := strconv.Atoi(element)
		numbers = append(numbers, i)
	}

	return numbers
}
