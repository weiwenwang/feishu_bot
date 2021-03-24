package feishu_bot

import (
	"encoding/json"
)

type Chat_obj struct {
	Avatar  string `json:"avatar"`
	Chat_id string `json:"chat_id"`
	Name    string `json:"name"`
}

type Dt struct {
	Groups []Chat_obj `json:"groups"`
}

type Chat_resp struct {
	Code int    `json:"code"`
	Data Dt     `json:"data"`
	Msg  string `json:"msg"`
}

func (*FeiShu_Bot) ChatList() ([]Chat_obj, error) {
	fb_byte, _ := json.Marshal(Feishu_bot)
	resp, err := Post_r(Get_Chat_List, fb_byte)
	if err != nil {
		return nil, err
	}
	var chat_resp Chat_resp

	err = json.Unmarshal(resp, &chat_resp)
	if err != nil {
		return nil, err
	}
	return chat_resp.Data.Groups, nil
}

type zh_ch struct {
	Title   string                `json:"title"`
	Content [][]map[string]string `json:"content"`
}

type post struct {
	Zh_ch zh_ch `json:"zh_cn"`
}
type Content struct {
	Post post `json:"post"`
}
type Msg struct {
	Chat_id  string  `json:"chat_id"`
	Msg_type string  `json:"msg_type"`
	Content  Content `json:"content"`
}

func (*FeiShu_Bot) SendMessage(chat_id string, title string, content []string) ([]byte, error) {
	var msg Msg
	msg.Chat_id = chat_id
	msg.Msg_type = "post"

	ctt2 := make([][]map[string]string, 0)
	for _, v := range content {
		ctt2 = append(ctt2, getc(v, "text"))
	}

	msg.Content.Post.Zh_ch.Title = title
	msg.Content.Post.Zh_ch.Content = ctt2
	content_byte, _ := json.Marshal(msg)
	return Post_r(Get_Send, content_byte)
}

func (*FeiShu_Bot) SendMessageRichText(chat_id string, title string, content []string, at []string) ([]byte, error) {
	var msg Msg
	msg.Chat_id = chat_id
	msg.Msg_type = "post"

	ctt2 := make([][]map[string]string, 0)
	for _, v := range at {
		ctt2 = append(ctt2, getc(v, "at"))
	}

	for _, v := range content {
		ctt2 = append(ctt2, getc(v, "text"))
	}

	msg.Content.Post.Zh_ch.Title = title
	msg.Content.Post.Zh_ch.Content = ctt2
	content_byte, _ := json.Marshal(msg)
	return Post_r(Get_Send, content_byte)
}

func getc(raw string, tag string) []map[string]string {
	text := "text"
	if tag == "at" {
		text = "user_id"
	}
	c := map[string]string{
		"tag": tag,
		text:  raw,
	}
	ctt1 := make([]map[string]string, 0)
	ctt1 = append(ctt1, c)
	return ctt1
}
