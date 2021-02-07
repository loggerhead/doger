package api

import (
	"fmt"
	neturl "net/url"
	"strings"

	"encoding/json"

	"github.com/gocolly/colly"
	"github.com/pkg/errors"
)

const (
	HOST = "mini.tuodan.tech"
	PATH = "jstd-doger/app"
)

var (
	API2URL map[string]string = map[string]string{
		"UserDetail":         "/user/v6/info",
		"UserDynamicList":    "/dynamic/v2/list",
		"DynamicList":        "/dynamic/v2/list",
		"DynamicDetail":      "/dynamic/v2/dynamicDetail",
		"DynamicCommentList": "/dynamic/v2/commentList",
		"DynamicLikeList":    "/dynamic/v1/likeList",
	}
	ReqUtil Req
)

type Req struct {
	c   *colly.Request
	url string
}

func NewReq(r *colly.Request) *Req {
	req := &Req{c: r}
	return req
}

func (r *Req) URL() string {
	return r.url
}

func (r *Req) GenURL(api string, qq ...interface{}) {
	r.url = GenURL(GetURL(api), qq...)
}

func (r *Req) Visit() error {
	err := r.c.Visit(r.url)
	if errors.Is(err, colly.ErrAlreadyVisited) {
		return nil
	}
	return err
}

func (r *Req) VisitByPost(qq ...interface{}) error {
	aa := make(map[string]string)
	for i := 0; i < len(qq); i += 2 {
		k := fmt.Sprint(qq[i])
		v := ""
		if i+1 < len(qq) {
			v = fmt.Sprint(qq[i+1])
		}
		aa[k] = v
	}

	payload, err := json.Marshal(aa)
	if err != nil {
		return errors.Wrap(err, "VisitByPost failed")
	}
	r.c.Headers.Set("content-type", "application/json; charset=UTF-8")

	err = r.c.PostRaw(r.url, payload)
	if errors.Is(err, colly.ErrAlreadyVisited) {
		return nil
	}
	return err
}

func GenURL(api string, qq ...interface{}) string {
	url := fmt.Sprintf("https://%s/%s/%s", HOST, PATH, strings.TrimLeft(api, "/"))

	u, err := neturl.Parse(url)
	if err != nil {
		panic(fmt.Sprintf("url.Parse(%q) failed: %v", url, err))
	}

	q := u.Query()
	for i := 0; i < len(qq); i += 2 {
		k := fmt.Sprint(qq[i])
		v := ""
		if i+1 < len(qq) {
			v = fmt.Sprint(qq[i+1])
		}
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

func GetURL(api string) string {
	url, ok := API2URL[api]
	if !ok {
		panic(fmt.Sprintf("unknow api: %s", api))
	}
	return url
}
