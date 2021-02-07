package crawler

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/loggerhead/doger/api"
)

func onRespDynamicList(r *colly.Response, v *api.RespDynamicList) {
	for _, data := range v.Data {
		handleDynamicDetail(r, data.Data)
	}
}

func onRespDynamicDetail(r *colly.Response, v *api.RespDynamicDetail) {
	handleDynamicDetail(r, v.Data)
}

func handleDynamicDetail(r *colly.Response, dynamic *api.DynamicDetail) {
	setMaxDID(dynamic.DynamicID)

	if err := Save(dynamic); err != nil {
		log.Printf("ERROR handleComment Save failed: %v\n", err)
	}

	if err := api.NewReq(r.Request).VisitUserDetail(dynamic.UID); err != nil {
		log.Printf("ERROR handleDynamicDetail VisitUserDetail failed: %v\n", err)
	}

	for _, reply := range dynamic.ReplyInfos {
		if err := api.NewReq(r.Request).VisitUserDetail(reply.FromUserUid); err != nil {
			log.Printf("ERROR handleDynamicDetail VisitUserDetail failed: %v\n", err)
		}
	}

	if err := api.NewReq(r.Request).VisitDynamicCommentList(dynamic.DynamicID); err != nil {
		log.Printf("ERROR handleDynamicDetail VisitDynamicCommentList failed: %v\n", err)
	}

	if err := api.NewReq(r.Request).VisitDynamicLikeList(dynamic.DynamicID); err != nil {
		log.Printf("ERROR handleDynamicDetail VisitDynamicLikeList failed: %v\n", err)
	}
}

func onRespDynamicCommentList(r *colly.Response, v *api.RespDynamicCommentList) {
	for _, comment := range v.Data {
		handleComment(r, comment)
	}
}

func handleComment(r *colly.Response, c *api.DynamicComment) {
	if err := Save(c); err != nil {
		log.Printf("ERROR handleComment Save failed: %v\n", err)
	}

	if err := api.NewReq(r.Request).VisitUserDetail(c.FromUserUID); err != nil {
		log.Printf("ERROR handleComment VisitUserDetail failed: %v\n", err)
	}
	if err := api.NewReq(r.Request).VisitUserDetail(c.ToUserUID); err != nil {
		log.Printf("ERROR handleComment VisitUserDetail failed: %v\n", err)
	}

	for _, comment := range c.Comments {
		handleComment(r, comment)
	}
}

func onRespDynamicLikeList(r *colly.Response, v *api.RespDynamicLikeList) {
	for _, like := range v.Data {
		if err := api.NewReq(r.Request).VisitUserDetail(like.UID); err != nil {
			log.Printf("ERROR onRespDynamicLikeList VisitUserDetail failed: %v\n", err)
		}
	}
}

func onRespUserDetail(r *colly.Response, v *api.RespUserDetail) {
	user := v.Data.UserInfo
	if user == nil || user.ID == 0 || user.Basicinfo.UID == "" {
		log.Printf("onRespUserDetail nil UserInfo: %v\n", string(r.Body))
		return
	}

	if err := Save(user); err != nil {
		log.Printf("ERROR onRespUserDetail Save failed: %v\n", err)
	}

	if err := api.NewReq(r.Request).VisitUserDynamicList(user.Basicinfo.UID); err != nil {
		log.Printf("ERROR onRespUserDetail VisitUserDynamicList failed: %v\n", err)
	}
}
