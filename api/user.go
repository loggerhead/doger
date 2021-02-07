package api

func (r *Req) GenURLUserDetail(uid string) string {
	r.GenURL("UserDetail", "uid", uid)
	return r.URL()
}

func (r *Req) GenURLUserDynamicList(uid string) string {
	r.GenURL("UserDynamicList", "uid", uid, "pageNo", "1", "pageSize", "100", "type", "INDEX")
	return r.URL()
}

func (r *Req) VisitUserDetail(uid string) error {
	r.GenURLUserDetail(uid)
	return r.Visit()
}

func (r *Req) VisitUserDynamicList(uid string) error {
	r.GenURLUserDynamicList(uid)
	return r.Visit()
}
