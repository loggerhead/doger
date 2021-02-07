package api

import "time"

// gender: 1 男生、2 女生. sort: hot、time
func (r *Req) GenURLDynamicList(pageNo, gender int, sort string, dummys ...interface{}) string {
	args := []interface{}{"pageNo", pageNo, "sort", sort, "pageSize", 100, "type", "STAR"}
	if gender != 0 {
		args = append(args, "gender")
		args = append(args, gender)
	}
	if len(dummys) > 0 {
		args = append(args, "a")
		args = append(args, time.Now().UnixNano())
	}

	r.GenURL("DynamicList", args...)
	return r.URL()
}

func (r *Req) GenURLDynamicDetail(did int) string {
	r.GenURL("DynamicDetail", "dynamicId", did)
	return r.URL()
}

func (r *Req) GenURLDynamicCommentList(did int) string {
	r.GenURL("DynamicCommentList", "dynamicId", did, "pageNo", 1, "pageSize", 1000, "order", "asc")
	return r.URL()
}

func (r *Req) VisitDynamicCommentList(did int) error {
	r.GenURLDynamicCommentList(did)
	return r.Visit()
}

func (r *Req) VisitDynamicLikeList(did int) error {
	r.GenURL("DynamicLikeList")
	return r.VisitByPost("dynamicId", did, "pageNo", 1, "pageSize", 1000)
}
