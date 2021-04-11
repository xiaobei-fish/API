package controller

import (
	"NewTest3/common"
	"NewTest3/model"
	"NewTest3/response"
	"NewTest3/util"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

//注册
func Register(ctx *gin.Context) {
	var requestUser model.Api_user
	ctx.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		//422 Unprocessable Entity 无法处理的请求实体
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		fmt.Println(telephone, len(telephone))
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	// 判断手机号是否存在
	if isTelephoneExist(common.DB, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	// 创建用户
	//返回密码的hash值（对用户密码进行二次处理，防止系统管理人员利用）
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.Api_user{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	common.DB.Create(&newUser) // 新增记录
	// 发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{"token": token,}, "注册成功")
}
//登录
func Login(ctx *gin.Context) {
	var requestUser model.Api_user
	ctx.Bind(&requestUser)
	//name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	// 数据验证
	if len(telephone) != 11 {
		//422 Unprocessable Entity 无法处理的请求实体
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		fmt.Println(telephone, len(telephone))
		return
	}
	if len(password) < 6 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	// 依据手机号，查询用户注册的数据记录
	var user model.Api_user
	common.DB.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	// 判断密码收否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		return
	}
	// 返回结果
	response.Success(ctx, gin.H{"token": token}, "登录成功")

}
//用户信息
func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.Success(ctx, gin.H{
		"user": response.ToUserDTO(user.(model.Api_user))}, "响应成功")
}
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.Api_user
	db.Where("telephone=?", telephone).First(&user)
	//如果没有查询到数据，对于uint数据，默认值为：0
	if user.ID != 0 {
		return true
	}
	return false
}
//增加事件
func Add(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	var requestThing model.Api_thing
	ctx.Bind(&requestThing)
	newThing := model.Api_thing{
		User_id:    user.(model.Api_user).ID,
		Table:      requestThing.Table,
		Intro:      requestThing.Intro,
		State:      requestThing.State,
		End:   		requestThing.End,
	}
	//数据验证
	//.......
	//数据验证报错
	/*
	if .....{
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": "数据输入出错"})
			return
	}
	*/
	common.DB.Create(&newThing)
	// 返回结果
	response.Success(ctx, gin.H{"thing": response.ToThingDTO(requestThing)}, "添加成功")
}
//修改事件
func Alt(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	var requestAlter model.Api_alter
	ctx.Bind(&requestAlter)
	newThing := model.Api_alter{
		Id:         requestAlter.Id,
		State:		requestAlter.State,
	}
	//单改
	if newThing.Id != 0 {
		common.DB.Table("api_things").Where("id=? and user_id=?", newThing.Id, user.(model.Api_user).ID).
			Update("state", newThing.State)
		common.DB.Table("api_things").Where("id=? and user_id=?", newThing.Id, user.(model.Api_user).ID).
			Update("updated_at", time.Now())
		// 返回结果
	} else if newThing.Id == 0{
	//全改
		common.DB.Table("api_things").Where("user_id=? and state!=?", user.(model.Api_user).ID, newThing.State).
			Update("state", newThing.State)
		common.DB.Table("api_things").Where("user_id=? and state!=?", user.(model.Api_user).ID, newThing.State).
			Update("updated_at", time.Now())
	}else {
		//报错
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": "数据输入出错"})
	}
	response.Success(ctx, gin.H{"newState": newThing.State}, "修改成功")
}
//删除事件
func Del(ctx *gin.Context)  {
	user, _ := ctx.Get("user")
	var requestAlter model.Api_alter
	ctx.Bind(&requestAlter)
	newThing := model.Api_alter{
		Id:         requestAlter.Id,
		State:		requestAlter.State,
	}
	//单删
	if newThing.Id != 0 {
		var deleteThings model.Api_delete
		common.DB.Table("api_things").Delete(&deleteThings,"user_id=? and id=?", user.(model.Api_user).ID, newThing.Id)
	}else{
	//全删
		if newThing.State == -1 {
			//全删除
			var deleteThings model.Api_delete
			common.DB.Table("api_things").Delete(&deleteThings,"user_id=?", user.(model.Api_user).ID)
		}else if newThing.State == 0 {
			//删除待办
			var deleteThings model.Api_delete
			common.DB.Table("api_things").Delete(&deleteThings,"user_id=? and state=?", user.(model.Api_user).ID, newThing.State)
		}else if newThing.State == 1 {
			//删除已完成
			var deleteThings model.Api_delete
			common.DB.Table("api_things").Delete(&deleteThings,"user_id=? and state=?", user.(model.Api_user).ID, newThing.State)
		}else{
			//报错
			ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": "数据输入出错"})
		}
	}
	response.Success(ctx, gin.H{"newState": newThing.State}, "删除成功")
}
//查询事件
func Que(ctx *gin.Context)  {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("connect to redis err", err.Error())
		return
	}

	user, _ := ctx.Get("user")
	var requestThing model.Api_query
	ctx.Bind(&requestThing)
	pageIndex := 5
	var queryThings []model.Api_thing
	if requestThing.Way == 0 && requestThing.Key == "" {
		//待办事件
		common.DB.Table("api_things").Where("state=? and user_id=?", requestThing.State, user.(model.Api_user).ID).
			Limit(pageIndex).Offset((requestThing.Page-1)*pageIndex).Find(&queryThings)

		response.Success(ctx, gin.H{"page": requestThing.Page, "thing": queryThings}, "查询成功")
		_, err := c.Do("lpush","history","查询"+strconv.Itoa(requestThing.Page)+"页的待办事件信息")
		if err != nil   {
			fmt.Println("redis lpush failed",err.Error())
		}
	}else if requestThing.Way == 1 && requestThing.Key == "" {
		//完成事件
		common.DB.Table("api_things").Where("state=? and user_id=?", requestThing.State, user.(model.Api_user).ID).
			Limit(pageIndex).Offset((requestThing.Page-1)*pageIndex).Find(&queryThings)

		response.Success(ctx, gin.H{"page": requestThing.Page, "thing": queryThings}, "查询成功")
		_, err := c.Do("lpush","history","查询"+strconv.Itoa(requestThing.Page)+"页的完成事件信息")
		if err != nil   {
			fmt.Println("redis lpush failed",err.Error())
		}
	}else if requestThing.Way == 2 && requestThing.Key == "" {
		//全部事件
		common.DB.Table("api_things").Where("user_id=?", user.(model.Api_user).ID).
			Limit(pageIndex).Offset((requestThing.Page-1)*pageIndex).Find(&queryThings)

		response.Success(ctx, gin.H{"page": requestThing.Page, "thing": queryThings}, "查询成功")
		_, err := c.Do("lpush","history","查询"+strconv.Itoa(requestThing.Page)+"页的全部事件信息")
		if err != nil   {
			fmt.Println("redis lpush failed",err.Error())
		}
	}else if requestThing.Key != "" {
		//关键字查询
		sql := "select * from api_things where user_id=" + strconv.Itoa(int(user.(model.Api_user).ID)) + " and intro like '%" +
			requestThing.Key + "%' limit " + strconv.Itoa(pageIndex*(requestThing.Page-1)) + "," + strconv.Itoa(pageIndex)
		common.DB.Raw(sql).Scan(&queryThings)

		response.Success(ctx, gin.H{"key": requestThing.Key, "thing": queryThings}, "查询成功")
		_, err := c.Do("lpush","history","查询"+strconv.Itoa(requestThing.Page)+"页的含有"+requestThing.Key+"的事件信息")
		if err != nil   {
			fmt.Println("redis lpush failed",err.Error())
		}
	}else {
		//报错
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 400, "msg": "数据输入出错"})
	}
	_, err = c.Do("ltrim", 0, 9)
	if err != nil   {
		fmt.Println("redis ltrim failed",err.Error())
	}

	defer c.Close()
}
