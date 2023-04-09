package db_resources

func WhereMap(req *ListRequests) {
	where := make(map[string]any)
	if req.Ids != nil {
		where["id in (?)"] = req.Ids
	}
	if req.Name != "" {
		where["name like ?"] = "%" + req.Name + "%"
	}
	if req.Phone != "" {
		where["phone = ?"] = req.Phone
	}
	if req.Title != "" {
		where["title like ?"] = "%" + req.Title + "%"
	}
	if req.UserId != 0 {
		where["user_id = ?"] = req.UserId
	}
}

type ListRequests struct {
	Ids    []int32 //in
	Name   string  //like
	Phone  string  //=
	Title  string  //like
	UserId int     //=
}
