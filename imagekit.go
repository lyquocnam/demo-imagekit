package imagekit

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	ApiKey    = "sdjVk+jHVPBgoLJCZhs6piAW8zE="
	ApiSecret = "8uQ5b44MH5lMTCYM6CQArvdRLFE="
)

func Signature(filename string, timestamp string) string {
	input := fmt.Sprintf("apiKey=%s&filename=%s&timestamp=%s", ApiKey, filename, timestamp)
	hash := hmac.New(sha1.New, []byte(ApiSecret))
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

func Base64Image(filepath string) string {
	f, _ := os.Open(filepath)
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(data)
}
