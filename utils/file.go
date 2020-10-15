package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// 通过Url网络地址获取
func Url(url string) func() ([]byte, error) {
	return func() ([]byte, error) {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, err
		}
		req.Header["User-Agent"] = []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36 Edg/83.0.478.61"}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		if strings.Contains(resp.Header.Get("Content-Encoding"), "gzip") {
			buffer := bytes.NewBuffer(body)
			r, _ := gzip.NewReader(buffer)
			defer r.Close()
			unCom, err := ioutil.ReadAll(r)
			return unCom, err
		}
		return body, nil
	}
}

// 通过文件获取
func File(file string) func() ([]byte, error) {
	return func() ([]byte, error) {
		return ioutil.ReadFile(file)
	}
}

// 通过Base64字符串获取
func Base64(data string) func() ([]byte, error) {
	return func() ([]byte, error) {
		return base64.StdEncoding.DecodeString(data)
	}
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
