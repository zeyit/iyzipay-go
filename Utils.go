package iyzipay

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	cnt "github.com/zeyit/iyzipay-go/requestconvert"
)

const client_version = "iyzipay-dotnet-2.1.37" //"iyzipay-golang-1.0"

const authorization = "Authorization"
const random_header_name = "x-iyzi-rnd"
const client_version_header_name = "x-iyzi-client-version"
const iyziws_header_name = "IYZWS "
const colon = ":"

func getHeaders(request cnt.RequestStringConvertible, options *Options) map[string]string {
	t := time.Now()
	//format ddMMyyyyhhmmssffff
	randomString := fmt.Sprintf("%02d-%02d-%d-%02d-%02d-%02d-%04d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second(), t.Second())
	randomString = strings.Replace(randomString, "-", "", 6)
	headers := make(map[string]string)

	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json; charset=utf-16"
	headers[random_header_name] = randomString
	headers[client_version_header_name] = client_version

	headers[authorization] = prepareauthorizationString(request, randomString, options)
	return headers
}

func prepareauthorizationString(request cnt.RequestStringConvertible, randomString string, options *Options) string {
	hash := generateHash(options.APIKey, options.SecretKey, randomString, request)
	return iyziws_header_name + options.APIKey + colon + hash
}

func generateHash(apiKey, secretKey, randomString string, request cnt.RequestStringConvertible) string {

	strReq := request.ToPKIRequestString()

	hashStr := apiKey + randomString + secretKey + strReq
	hasher := sha1.New()
	hasher.Write([]byte(hashStr))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func decodeString(str string) (string, error) {
	if bytes, err := base64.StdEncoding.DecodeString(str); err == nil {
		return string(bytes), nil
	} else {
		return "", err
	}
}
