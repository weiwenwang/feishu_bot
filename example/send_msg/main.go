package main

import (
	"fmt"
	"github.com/weiwenwang/feishu_bot"
)

const APP_ID = "your feishu bot app_id"
const APP_Secret = "your feishu bot app_secret_id"

func main() {
	// 传入配置文件, 初始化一个全局对象*FeiShu_Bot
	err := feishu_bot.Init_Feishu(APP_ID, APP_Secret)
	if err != nil {

	}
	// 获取全局对象*FeiShu_Bot
	bot, _ := feishu_bot.GetFeiShuObj()

	// 查看当前机器人被哪些群添加进去了, 也就是可以把消息发到哪些群里
	chatList, _ := bot.ChatList()
	for _, v := range chatList {
		fmt.Println(v.Chat_id, v.Name)
	}

	// 选择一个群报错
	raw1 := "raw1"
	raw2 := `json: cannot unmarshal string into Go struct field post_data.app_pk_id of type int{"app_id":"wxae70dd50db3ab2d7","app_pk_id":"56","remark":"","sql":"select \n  * \nfrom \n  battle_db.battle_wxae70dd50db3ab2d7 \nwhere \n  STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') = '2020-06-24' \nlimit \n  50","uid":"9"}
	2020/06/24 10:50:44.675 [E] [task.go:140]  json: cannot unmarshal string into Go struct field post_data.app_pk_id of type int{"app_id":"wxae70dd50db3ab2d7","app_pk_id":"56","remark":"","sql":"select \n  * \nfrom \n  battle_db.battle_wxae70dd50db3ab2d7 \nwhere \n  STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') = '2020-06-24' and current_score != '' \nlimit \n  50","uid":"9"}`
	bot.SendMessage(chatList[0].Chat_id, "your title", []string{raw1, raw2})
}
