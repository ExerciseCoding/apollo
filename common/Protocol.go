package common

import "encoding/json"

//定时任务
type CronJob struct {
	Name string `json:"name"` //任务名
	Command string `json:"command"` //shell命令
	CronExpr string `json:"cronExpr"` //cron表达式
}

//HTTP接口应答
type Response struct {
	Errno int `json:"errno"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

//应答方法
func BuildResponse(errno int,msg string,data interface{})(resp []byte,err error){
	//1.定义一个Response
	var(
		response Response
	)
	response.Errno = errno
	response.Msg = msg
	response.Data = data

	//序列化
	resp,err = json.Marshal(response)
	return
}