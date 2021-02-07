package crawler

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"

	"github.com/gocolly/colly"
	"github.com/loggerhead/doger/api"
)

func Init(c *colly.Collector) error {
	if GetToken() == "" {
		return errors.New("not set doger token")
	}

	c.OnRequest(onRequest)
	c.OnResponse(onResponse)
	c.OnError(onError)
	go logMaxDID()

	return c.Visit(api.ReqUtil.GenURLDynamicList(1, api.AllGender, "hot", "DUMMY_RAND"))
}

func onError(r *colly.Response, err error) {
	log.Printf("ERROR %v: [%v] %v %v\n", err, r.Request.URL, r.StatusCode, string(r.Body))
}

func onRequest(r *colly.Request) {
	r.Headers.Set("token", fmt.Sprintf("dogerProd %s", GetToken()))
	r.Headers.Set("content-type", "application/json; charset=UTF-8")
	r.Headers.Set("accept-language", "zh-cn")
	r.Headers.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 14_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.2(0x1800022c) NetType/4G Language/zh_CN")
	r.Headers.Set("referer", "https://servicewechat.com/wxdbf2b50ca37c30c0/259/page-frame.html")
	r.Headers.Set("accept-encoding", "gzip, deflate, br")

	log.Printf("visiting %s\n", r.URL)
}

func onResponse(r *colly.Response) {
	d, err := api.ParseResp(r)
	if err != nil {
		log.Printf("ERROR onResponse (%v): %v %v\n", err, r.StatusCode, string(r.Body))
		return
	}

	baseResp := reflect.Indirect(reflect.ValueOf(d)).Field(0).Interface().(api.BaseResp)
	if baseResp.Code == 431 {
		err = r.Request.Retry()
		log.Printf("ERROR visiting too fast [%v]. retry (%v)\n", r.Request.URL, err)
		return
	}
	if baseResp.Code != 1 {
		log.Printf("ERROR BaseResp not success [%v]. %v %v\n", r.Request.URL, r.StatusCode, string(r.Body))
		return
	}

	switch v := d.(type) {
	case *api.RespDynamicList:
		onRespDynamicList(r, v)
	case *api.RespDynamicDetail:
		onRespDynamicDetail(r, v)
	case *api.RespDynamicCommentList:
		onRespDynamicCommentList(r, v)
	case *api.RespDynamicLikeList:
		onRespDynamicLikeList(r, v)
	case *api.RespUserDetail:
		onRespUserDetail(r, v)
	}
}

var maxDID = struct {
	RecvCh chan int
	SendCh chan int
	DID    int
}{
	RecvCh: make(chan int),
	SendCh: make(chan int, 1),
	DID:    0,
}

func setMaxDID(did int) {
	maxDID.RecvCh <- did
}

func GetMaxDID() chan int {
	return maxDID.SendCh
}

func logMaxDID() {
	for did := range maxDID.RecvCh {
		if did > maxDID.DID {
			maxDID.DID = did
			if len(maxDID.SendCh) == 0 {
				maxDID.SendCh <- did
			}
			log.Printf("MaxDID: %v\n", GetMaxDID())
		}
	}
}

func GetToken() string {
	kk := []struct {
		key    string
		weight int
	}{
		{"DOGER_TOKEN", 6},
		{"DOGER_TOKEN_ME", 4},
	}

	sum := 0
	for _, v := range kk {
		sum += v.weight
	}

	r := rand.Intn(sum)
	key := ""
	for _, v := range kk {
		if r < v.weight {
			key = v.key
			break
		}
		r -= v.weight
	}

	return os.Getenv(key)
}
