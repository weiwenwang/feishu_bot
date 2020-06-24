package main

import (
	"fmt"
	"github.com/weiwenwang/feishu_bot"
)

func main() {
	// 传入配置文件, 初始化一个全局对象*FeiShu_Bot
	app_id := "cli_9e38xxx..."
	app_secret := "UJmTdIqncSYNrHxxx..."
	err := feishu_bot.Init_Feishu(app_id, app_secret)
	if err != nil {
		panic(err.Error())
	}
	// 获取全局对象*FeiShu_Bot
	feishu_bot, _ := feishu_bot.GetFeiShuObj()

	// 查看当前机器人被哪些群添加进去了, 也就是可以把消息发到哪些群里
	chatList, _ := feishu_bot.ChatList()
	fmt.Println(chatList[0].Chat_id, chatList[0].Name)

	// 选择一个群报错
	feishu_bot.SendMessage(chatList[0].Chat_id, "线上报错", `json: cannot unmarshal string into Go struct field post_data.app_pk_id of type int{"app_id":"wxae70dd50db3ab2d7","app_pk_id":"56","remark":"","sql":"select \n  * \nfrom \n  battle_db.battle_wxae70dd50db3ab2d7 \nwhere \n  STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') = '2020-06-24' \nlimit \n  50","uid":"9"}
2020/06/24 10:50:44.675 [E] [task.go:140]  json: cannot unmarshal string into Go struct field post_data.app_pk_id of type int{"app_id":"wxae70dd50db3ab2d7","app_pk_id":"56","remark":"","sql":"select \n  * \nfrom \n  battle_db.battle_wxae70dd50db3ab2d7 \nwhere \n  STR_TO_DATE(CONCAT(y, '-', m, '-', d), '%Y-%m-%d') = '2020-06-24' and current_score != '' \nlimit \n  50","uid":"9"}`)
}
