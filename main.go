package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sensitiveWord/lib/ahoCorasick"
	Config "sensitiveWord/lib/config"
)

type ResponseData struct {
	Data    string `json:"message"`
	IsMatch bool   `json:"isMatch"`
}
type RequestData struct {
	Text string `json:"text"`
}

func main() {
	ac := ahoCorasick.NewAhoCorasick()
	f, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("words.txt文件不存在")
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	patterns := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		patterns = append(patterns, line)
	}
	ac.Run(patterns)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		reqData, err := getRequestParam(r)
		if err != nil {
			fmt.Println(err)
			return
		}
		strRune := []rune(reqData.Text)
		fixStr, isMatch := ac.MatchAndRewrite(strRune)
		data := &ResponseData{
			Data:    fixStr,
			IsMatch: isMatch,
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonData)
		if err != nil {
			return
		}
	})
	confTemp := Config.GetConfig()
	err = http.ListenAndServe(confTemp.Addr, nil)
	if err != nil {
		fmt.Println("服务启动异常,", err)
		return
	}
}

func getRequestParam(r *http.Request) (reqData *RequestData, err error) {
	reqData = &RequestData{}
	err = r.ParseForm()
	if err != nil {
		return
	}
	if r.Method == http.MethodGet {
		reqData.Text = r.URL.Query().Get("text")
	} else if r.Method == http.MethodPost {
		contentType := r.Header.Get("Content-Type")
		if contentType == "application/json" {
			var body []byte
			body, err = io.ReadAll(r.Body)
			if err != nil {
				return
			}
			err = json.Unmarshal(body, &reqData)
			if err != nil {
				return
			}
		} else {
			err = r.ParseForm()
			if err != nil {
				return
			}
			reqData.Text = r.PostForm.Get("text")
		}
	}
	return

}
