package response

import (
	"NewTest3/model"
	"time"
)

type UserSendData struct {
	Name 		string `json:"name"`
	Telephone 	string `json:"telephone"`
}
type ThingSendData struct {
	Table		string
	Intro		string
	State		int
	Start_time  time.Time
	End_time    string
}
//送给数据传输对象
func ToUserDTO(user model.Api_user) UserSendData{
	return UserSendData{
		Name: 		user.Name,
		Telephone:  user.Telephone,
	}
}
func ToThingDTO(thing model.Api_thing) ThingSendData{
	return ThingSendData{
		Table:      thing.Table,
		Intro:      thing.Intro,
		State:      thing.State,
		Start_time: thing.CreatedAt,
		End_time:   thing.End,
	}
}
