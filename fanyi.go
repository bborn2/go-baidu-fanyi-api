package baidufanyiapi

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	apiURL = "http://api.fanyi.baidu.com/api/trans/vip/translate"
)

//Translate For list of language codes, please refer to `https://api.fanyi.baidu.com/doc/21`
func Translate(appid string, appkey string, fr string, to string, query string) (string, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	rand.Seed(int64(time.Now().UnixNano()))
	salt := strconv.Itoa(rand.Intn(32768) + (65536 - 32768))

	sign := makeMd5(appid + query + salt + appkey)

	payload := url.Values{"appid": {appid}, "q": {query}, "from": {fr}, "to": {to}, "salt": {salt}, "sign": {sign}}

	resp, err := client.Post(apiURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(payload.Encode()))

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	var ret translateResult

	err = json.Unmarshal(data, &ret)
	if err != nil {
		// panic(err)
		return "", err
	}

	if len(ret.TransResult) > 0 {
		return ret.TransResult[0].Dst, nil
	}

	return "", errors.New("no reponse data")
}

func makeMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

type translateResult struct {
	From        string `json:"from"`
	To          string `json:"to"`
	TransResult []struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	} `json:"trans_result"`
}
