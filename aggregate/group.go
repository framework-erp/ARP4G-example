package aggregate

//分组，比如 家人、朋友
type Group struct {
	Id           int64 `bson:"_id"`
	Name         string
	ContactCount int
	//0启用，1删除
	State int
}

func (group *Group) AddTo() {
	group.ContactCount++
}

func (group *Group) RemoveFrom() {
	group.ContactCount--
}

func (group *Group) SetAsRemoved() {
	group.State = 1
}

func (group *Group) IsDead() bool {
	return group.State == 1 && group.ContactCount == 0
}
