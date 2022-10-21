package aggregate

//联系人
type Contact struct {
	Id      int64
	Name    string
	GroupId int64
}

//分组，比如 家人、朋友
type Group struct {
	Id   int64
	Name string
}
