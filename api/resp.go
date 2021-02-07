package api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ParseResp(r *colly.Response) (interface{}, error) {
	if r == nil || r.Request == nil || r.Request.URL == nil {
		return nil, fmt.Errorf("invalid Response: %+v", r)
	}

	url := r.Request.URL.String()

	for k, v := range API2URL {
		if !strings.Contains(url, v) {
			continue
		}
		switch k {
		case "UserDetail":
			d := &RespUserDetail{}
			err := json.Unmarshal(r.Body, d)
			return d, err
		case "UserDynamicList":
			fallthrough
		case "DynamicList":
			raw := &struct {
				BaseResp
				Data []struct {
					Type int             `json:"type,omitempty"`
					Data json.RawMessage `json:"data,omitempty"`
				} `json:"data,omitempty"`
			}{}

			err := json.Unmarshal(r.Body, raw)
			if err != nil {
				return nil, err
			}

			d := &RespDynamicList{
				BaseResp: raw.BaseResp,
			}
			for _, rawData := range raw.Data {
				if rawData.Type != 1 {
					continue
				}

				detail := &DynamicDetail{}
				err := json.Unmarshal(rawData.Data, detail)
				if err != nil {
					return nil, err
				}

				d.Data = append(d.Data, DynamicListData{
					Type: rawData.Type,
					Data: detail,
				})
			}
			return d, err
		case "DynamicDetail":
			d := &RespDynamicDetail{}
			err := json.Unmarshal(r.Body, d)
			return d, err
		case "DynamicCommentList":
			d := &RespDynamicCommentList{}
			err := json.Unmarshal(r.Body, d)
			return d, err
		case "DynamicLikeList":
			d := &RespDynamicLikeList{}
			err := json.Unmarshal(r.Body, d)
			return d, err
		}
	}

	return nil, fmt.Errorf("unkown URL: %v", url)
}
