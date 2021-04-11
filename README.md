# 注册接口
    
##### 简要描述

- 用户注册接口

##### 请求URL
- ` 127.0.0.1:8080/api/register `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|name 	  |否  |string |用户名   |
|password |是  |string | 密码    |
|telephone|是  |string | 电话号码    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTYxODQwMDM2MCwiaWF0IjoxNjE3Nzk1NTYwLCJpc3MiOiJqaWFuZ3pob3UiLCJzdWIiOiJ1c2VyIHRva2VuIn0.tQyMWQHU8raS1Z9j9zoG_9ZK3s4VxUSYWRqbibI99HM"
	},
	"msg": "注册成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |用户token值  |
|msg |string   |具体说明信息  |




# 登录接口

    
##### 简要描述

- 用户登录接口

##### 请求URL
- ` 127.0.0.1:8080/api/login `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|telephone |是  |string |电话号码   |
|password |是  |string | 密码    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTYxODQwMDYxOSwiaWF0IjoxNjE3Nzk1ODE5LCJpc3MiOiJqaWFuZ3pob3UiLCJzdWIiOiJ1c2VyIHRva2VuIn0.L3iVREiGDZycnNqSkI207PLHzE-M40oV09-6scySOGM"
	},
	"msg": "登录成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |用户token值  |
|msg |string   |具体说明信息  |
# 用户信息

    
##### 简要描述

- 用户信息接口

##### 请求URL
- ` 127.0.0.1:8080/api/info `
  
##### 请求方式
- GET 


##### 返回示例 

``` 
{
	"code": 200,
	"data": {
		"user": {
			"name": "gin",
			"telephone": "15060301341"
		}
	},
	"msg": "响应成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |用户具体信息  |
|msg |string   |具体说明信息  |
# 增加接口

    
##### 简要描述

- 增加接口接口

##### 请求URL
- ` 127.0.0.1:8080/api/add `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|table    |是  |string |标题  |
|intro 	  |是  |string |内容    |
|state    |是  |string |状态(0：未完成，1：已经完成)    |
|end_time |是  |string |截止日期    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"thing": {
			"Table": "Test",
			"Intro": "Test",
			"State": 0,
			"Start_time": "0001-01-01T00:00:00Z",
			"End": "2021/04/10"
		}
	},
	"msg": "添加成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |新增事件具体信息  |
|msg |string   |具体说明信息  |
# 修改接口

    
##### 简要描述

- 修改接口接口

##### 请求URL
- ` 127.0.0.1:8080/api/alt `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|id    |是  |int |事件id(取0为全修改) |
|state    |是  |string |状态(0：未完成，1：已经完成)    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"newState": 1
	},
	"msg": "修改成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |具体信息  |
|msg |string   |具体说明信息  |
# 删除接口

    
##### 简要描述

- 删除接口接口

##### 请求URL
- ` 127.0.0.1:8080/api/del `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|id    |是  |int |事件id(取0为全修改) |
|state    |是  |string |状态(0：未完成，1：已经完成，-1：全删除)    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"newState": -1
	},
	"msg": "删除成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |具体信息  |
|msg |string   |具体说明信息  |
# 查询接口

    
##### 简要描述

- 查询接口接口

##### 请求URL
- ` 127.0.0.1:8080/api/que `
  
##### 请求方式
- POST 

##### 参数

|参数名|必选|类型|说明|
|:----    |:---|:----- |-----   |
|way    |是  |int |0：待办，1：完成，2：全部 |
|page   |是  |int |查询的页面page值    |
|key    |否  |string |关键字查询    |

##### 返回示例 

``` 
  {
	"code": 200,
	"data": {
		"key": "Test",
		"thing": null
	},
	"msg": "查询成功"
}
```

##### 返回参数说明 

|参数名|类型|说明|
|:-----  |:-----|-----                           |
|code |int   |200:成功，422:数据错误，500:系统错误  |
|data |json   |具体信息  |
|msg |string   |具体说明信息  |
