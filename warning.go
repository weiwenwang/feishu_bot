package feishu_bot

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type FeiShu_Bot struct {
	App_id     string `json:"app_id"`
	App_secret string `json:"app_secret"`
	Token      string
	Expire     int64
	Cre_time   int64
}

type Token_resp struct {
	Code                int    `json:"code"`
	Expire              int64  `json:"expire"`
	Msg                 string `json:"msg"`
	Tenant_Access_Token string `json:"tenant_access_token"`
}

const Get_Access_Token = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal/"
const Get_Chat_List = "https://open.feishu.cn/open-apis/chat/v4/list"
const Get_Send = "https://open.feishu.cn/open-apis/message/v4/send/"

var Feishu_bot *FeiShu_Bot
var init_once sync.Once

func Init_Feishu(app_id string, app_secret string) error {
	init_once.Do(func() {
		Feishu_bot = &FeiShu_Bot{
			App_id:     app_id,
			App_secret: app_secret,
			Token:      "",
			Expire:     0,
			Cre_time:   0,
		}
	})
	err := Feishu_bot.UpdateToken()
	if err != nil {
		return err
	}
	return nil

}
func (*FeiShu_Bot) UpdateToken() error {
	fb_byte, _ := json.Marshal(Feishu_bot)
	resp, err := Post_r(Get_Access_Token, fb_byte)
	if err != nil {
		return err
	}
	var token_resp Token_resp
	err = json.Unmarshal(resp, &token_resp)
	if err != nil {
		return err
	}
	Feishu_bot.Token = token_resp.Tenant_Access_Token
	Feishu_bot.Token = token_resp.Tenant_Access_Token
	Feishu_bot.Expire = token_resp.Expire
	Feishu_bot.Cre_time = time.Now().Unix()
	return nil
}

func GetFeiShuObj() (*FeiShu_Bot, error) {
	if Feishu_bot == nil {
		return nil, errors.New("feishu bot doesn't init?")
	}
	if time.Now().Unix() > Feishu_bot.Cre_time+Feishu_bot.Expire-60*20 { // 过期前20分钟update
		err := Feishu_bot.UpdateToken()
		if err != nil {
			return nil, err
		}
	}
	return Feishu_bot, nil
}

func Post_r(api_url string, fb_byte []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(fb_byte)
	request, err := http.NewRequest("POST", api_url, buffer)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8") //添加请求头
	request.Header.Set("Authorization", "Bearer "+Feishu_bot.Token)      //添加请求头
	client := http.Client{}                                              //创建客户端
	resp, err := client.Do(request.WithContext(context.TODO()))          //发送请求
	if err != nil {
		return nil, err
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBytes, nil
}
