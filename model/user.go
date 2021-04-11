package model

import (
	"gorm.io/gorm"
)

type Api_user struct {
	gorm.Model
	Name 		string `gorm:"type:varchar(20);not null"`			//用户名
	Telephone 	string `gorm:"type:varchar(11);not null;unique"`	//用户电话
	Password 	string `gorm:"size:255;not null"`					//用户密码
}

type Api_thing struct {
	gorm.Model
	User_id		uint		//用户id
	Table 		string		//事件标题
	Intro		string		//事件内容
	State 		int			//事件状态   0：待办   1：完成
	End			string		//截止时间
}

type Api_alter struct {
	gorm.Model
	Id		int				//事件id
	State	int				//事件状态   0：待办   1：完成
}

type Api_delete struct {
	User_id		uint		//用户id
	Table 		string		//事件标题
	Intro		string		//事件内容
	State 		int			//事件状态   0：待办   1：完成
	End			string		//截止时间
}

type Api_query struct {
	gorm.Model
	Way			int			//查询方式   0: 待办   1: 完成   2:全部
	Page		int 		//查询页面   一页5条记录
	Key			string		//查询关键字
	User_id		uint		//用户id
	Table 		string		//事件标题
	Intro		string		//事件内容
	State 		int			//事件状态   0：待办   1：完成
	End			string		//截止时间
}

