package aggregate

//联系人
type Contact struct {
	Id          int64 `bson:"_id"`
	Name        string
	PhoneNumber string
	GroupId     int64
}
