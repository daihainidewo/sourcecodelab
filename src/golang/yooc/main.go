// Package yooc yooc
// file create by daihao, time is 2018/8/29 10:35
package main

import (
	"net/http"
	"fmt"
	"strings"
	"time"
	"strconv"
	"bytes"
	"encoding/json"
	"io/ioutil"
)

const (
	YoocId         = "435408cf352a14d68ef6861b9d51158c"
	YoocAddSecond  = 1786213
	YoocAddPSecond = 1788095
)

// Cookie
type Cookie struct {
	Lvt        string
	Lpvt       string
	Distinctid string
	Cnzz       string
	Sessionid  string
	Csrftoken  string
}

func main() {
	ih, err := Index()
	if err != nil {
		fmt.Println(err)
		return
	}
	lh, err := Login(ih)
	if err != nil {
		fmt.Println(err)
		return
	}
	cookie, _ := json.Marshal(lh)
	fmt.Println(string(cookie))
	resp, err := Get(lh, "https://www.yooc.me/group/4219/homework/2848/grade")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(html))
}

func Get(c *Cookie, url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Cookie", "csrftoken="+c.Csrftoken+"; UM_distinctid="+c.Distinctid+"; Hm_lvt_435408cf352a14d68ef6861b9d51158c="+c.Lvt+"; CNZZDATA1254048558="+c.Cnzz+"; sessionid="+c.Sessionid+"; Hm_lpvt_435408cf352a14d68ef6861b9d51158c="+c.Lpvt)
	req.Header.Set("Host", "www.yooc.me")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	return resp, err
}

// Login Login
func Login(indexHeader http.Header) (*Cookie, error) {
	// TODO
	// 获取csrftoken
	req, err := http.NewRequest("GET", "https://www.yooc.me/login", nil)
	if err != nil {
		return nil, err
	}

	sessionid, _ := findCookie(indexHeader, "sessionid")
	nt := time.Now()
	ts := int(nt.Unix())
	cnzz := "1786339939-" + strconv.Itoa(ts) + "-https%253A%252F%252Fwww.yooc.me%252F%7C" + strconv.Itoa(ts)

	distinctid := fmt.Sprintf("%x%x-", ts*1000+666, 666) + "0c3040154a7ab7" + "-9393265-1fa400-" + fmt.Sprintf("%x%x", ts*1000+888, 888)

	req.Header.Set("Host", "www.yooc.me")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.AddCookie(&http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
	})

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lvt_" + YoocId,
		Value: strconv.Itoa(ts) + "," + strconv.Itoa(ts+YoocAddSecond),
	})
	req.AddCookie(&http.Cookie{
		Name:  "Hm_lpvt_" + YoocId,
		Value: strconv.Itoa(ts + YoocAddPSecond),
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// 注册登录信息，获取新的sessionid
	body := bytes.NewBufferString("email=15003134964&password=753951%2B%2B%2B&remember=true")
	req, err = http.NewRequest("POST", "https://www.yooc.me/yiban_account/login_ajax", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "www.yooc.me")
	req.Header.Set("Content-Length", strconv.Itoa(body.Len()))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.AddCookie(&http.Cookie{
		Name:  "sessionid",
		Value: sessionid,
	})

	req.AddCookie(&http.Cookie{
		Name:  "Hm_lvt_" + YoocId,
		Value: strconv.Itoa(ts + YoocAddSecond),
	})
	req.AddCookie(&http.Cookie{
		Name:  "Hm_lpvt_" + YoocId,
		Value: strconv.Itoa(ts + YoocAddPSecond),
	})
	req.AddCookie(&http.Cookie{
		Name:  "CNZZDATA1254048558",
		Value: cnzz,
	})
	req.AddCookie(&http.Cookie{
		Name:  "UM_distinctid",
		Value: distinctid,
	})

	csrftoken, _ := findCookie(resp.Header, "csrftoken")
	req.AddCookie(&http.Cookie{
		Name:  "csrftoken",
		Value: csrftoken,
	})

	req.Header.Set("X-CSRFToken", csrftoken)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	sessionid, ok := findCookie(resp.Header, "sessionid")
	if !ok {
		fmt.Println("not found sessionid")
		return nil, err
	}

	// 填充结构体，返回
	ret := new(Cookie)
	ret.Cnzz = cnzz
	ret.Csrftoken = csrftoken
	ret.Distinctid = distinctid
	ret.Sessionid = sessionid
	ret.Lvt = strconv.Itoa(ts) + "," + strconv.Itoa(ts+YoocAddSecond)
	ret.Lpvt = strconv.Itoa(ts + YoocAddPSecond)

	return ret, nil
}

// Index Index
func Index() (http.Header, error) {
	// TODO

	req, err := http.NewRequest("GET", "https://www.yooc.me/", nil)
	if err != nil {
		return nil, err
	}

	client := http.DefaultClient

	req.Header.Set("Host", "www.yooc.me")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Header, nil
}

// findCookie findCookie
func findCookie(header http.Header, s string) (string, bool) {
	// TODO
	for _, d := range header["Set-Cookie"] {
		spl := strings.Split(d, "; ")
		for _, d := range spl {
			if strings.HasPrefix(d, s) {
				tmp := strings.Split(d, "=")
				return tmp[1], true
			}
		}
	}
	return "", false
}

/*

sessionid=82c780299ed18605c3bfc96692f08ad8; Domain=.yooc.me; expires=Wed, 12-Sep-2018 02:54:28 GMT; httponly; Max-Age=1209600; Path=/

sessionid=82c780299ed18605c3bfc96692f08ad8; Hm_lvt_435408cf352a14d68ef6861b9d51158c=1533723170,1535509383; Hm_lpvt_435408cf352a14d68ef6861b9d51158c=1535511265


sessionid=82c780299ed18605c3bfc96692f08ad8;
Hm_lvt_435408cf352a14d68ef6861b9d51158c=1533723170,1535509383;
csrftoken=g1Ix8Cxeoejd1EmVUhdI371BbQVCXVdB;
Hm_lpvt_435408cf352a14d68ef6861b9d51158c=1535511279;


CNZZDATA1254048558=1586372269-1535506578-https%253A%252F%252Fwww.yooc.me%252F%7C1535506578

UM_distinctid=165839ab501f9b-0c3040154a7ab7-9393265-1fa400-165839ab503cc9;
UM_distinctid=16583c2af361de-0f76ea2c45c4ff-9393265-1fa400-16583c2af37100;

sessionid=14e12e82c8028895a637dc7362a86740;
Hm_lvt_435408cf352a14d68ef6861b9d51158c=1533723170,1535509383;
csrftoken=B6IIPGD3W9Wpb7lFS2EZkk3kX7aWBLEU;
Hm_lpvt_435408cf352a14d68ef6861b9d51158c=1535513899;



CNZZDATA1254048558=1786339939-1535512001-https%253A%252F%252Fwww.yooc.me%252F%7C1535512001


sessionid=14e12e82c8028895a637dc7362a86740;
Hm_lvt_435408cf352a14d68ef6861b9d51158c=1533723170,1535509383;
csrftoken=B6IIPGD3W9Wpb7lFS2EZkk3kX7aWBLEU;
Hm_lpvt_435408cf352a14d68ef6861b9d51158c=1535513899;
UM_distinctid=16583c2af361de-0f76ea2c45c4ff-9393265-1fa400-16583c2af37100;
CNZZDATA1254048558=1786339939-1535512001-https%253A%252F%252Fwww.yooc.me%252F%7C1535512001
*/
