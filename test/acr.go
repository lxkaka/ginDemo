package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

const ACR_OPT_REC_AUDIO string = "audio"
const ACR_OPT_REC_HUMMING string = "humming"
const ACR_OPT_REC_BOTH string = "both"

type Recognizer struct {
	Host         string
	AccessKey    string
	AccessSecret string
	RecType      string
	TimeoutS     int
	HttpClient   *http.Client
}

func (r *Recognizer) Post(url string, fieldParams map[string]string, fileParams map[string][]byte, timeoutS int) (string, error) {
	postDataBuffer := bytes.Buffer{}
	mpWriter := multipart.NewWriter(&postDataBuffer)

	for key, val := range fieldParams {
		//fmt.Println(val)
		_ = mpWriter.WriteField(key, val)
	}

	for key, val := range fileParams {
		fw, err := mpWriter.CreateFormFile(key, key)
		if err != nil {
			mpWriter.Close()
			return "", fmt.Errorf("create Form File Error: %v", err)
		}
		fw.Write(val)
	}

	mpWriter.Close()

	//hClient := &http.Client {
	//    Timeout: time.Duration(10 * time.Second),
	//}

	req, err := http.NewRequest("POST", url, &postDataBuffer)
	if err != nil {
		return "", fmt.Errorf("NewRequest Error: %v", err)
	}
	req.Header.Set("Content-Type", mpWriter.FormDataContentType())
	//response, err := hClient.Do(req)
	response, err := r.HttpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http Client Do Error: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("http Response Status Code Is Not 200: %d", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", fmt.Errorf("read From Http Response Error: %v", err)
	}

	return string(body), nil
}

func (r *Recognizer) GetSign(str string, key string) string {
	hmacHandler := hmac.New(sha1.New, []byte(key))
	hmacHandler.Write([]byte(str))
	return base64.StdEncoding.EncodeToString(hmacHandler.Sum(nil))
}

func (r *Recognizer) DoRecognize(audioFp []byte, humFp []byte, userParams map[string]string) (string, error) {
	qurl := "https://" + r.Host + "/v1/identify"
	httpMethod := "POST"
	httpUri := "/v1/identify"
	dataType := "fingerprint"
	signatureVersion := "1"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	stringToSign := httpMethod + "\n" + httpUri + "\n" + r.AccessKey + "\n" + dataType + "\n" + signatureVersion + "\n" + timestamp
	sign := r.GetSign(stringToSign, r.AccessSecret)

	if audioFp == nil && humFp == nil {
		return "", fmt.Errorf("can not Create Fingerprint")
	}

	fieldParams := map[string]string{
		"access_key":        r.AccessKey,
		"timestamp":         timestamp,
		"signature":         sign,
		"data_type":         dataType,
		"signature_version": signatureVersion,
	}

	if userParams != nil {
		for key, val := range userParams {
			fieldParams[key] = val
		}
	}

	fileParams := map[string][]byte{}
	if audioFp != nil && len(audioFp) != 0 {
		fileParams["sample"] = audioFp
		fieldParams["sample_bytes"] = strconv.Itoa(len(audioFp))
	}
	if humFp != nil && len(humFp) != 0 {
		fileParams["sample_hum"] = humFp
		fieldParams["sample_hum_bytes"] = strconv.Itoa(len(humFp))
	}

	result, err := r.Post(qurl, fieldParams, fileParams, r.TimeoutS)
	return result, err
}

func main() {
	fb, _ := ioutil.ReadFile("/Users/lxkaka/Desktop/fb")
	recognizer := Recognizer{
		Host:         "identify-cn-north-1.acrcloud.cn",
		AccessKey:    "30514459464f052a4857b2ec07c33781",
		AccessSecret: "GVyg0f65ZJ1MY1yPal27jwalkHCLj0qPLz11IiFT",
		RecType:      ACR_OPT_REC_AUDIO,
		TimeoutS:     10,
		HttpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	res, err := recognizer.DoRecognize(fb, nil, nil)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	fmt.Printf("%+v", res)
}
