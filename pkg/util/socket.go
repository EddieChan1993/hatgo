package util

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"net/http"
	"io"
	"crypto/rand"
	"encoding/base64"
	"hatgo/pkg/logs"
	"github.com/gin-gonic/gin"
	"fmt"
)

const secWebsocketProtocol = "Sec-WebSocket-Protocol"

type Ws struct {
	conn *websocket.Conn
}

//消息体
type Message struct {
	Content interface{} `json:"content"`
	Type    string      `json:"type"`
}

//用户体
type User struct {
	Bid  string
	conn *websocket.Conn
}

var (
	member         = make(map[string]*User)
	uidMapWs       = make(map[string]*websocket.Conn)
	wsMapUid       = make(map[*websocket.Conn]string)
	groupMapMember = make(map[string][]*User)
)

func CoonId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return Md5(base64.URLEncoding.EncodeToString(b))
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//实例化
func NewWs(c *gin.Context) (wss *Ws, token string, err error) {
	resHeader := http.Header{}
	token = c.GetHeader(secWebsocketProtocol)
	resHeader.Add(secWebsocketProtocol, token)
	conn, err := wsupgrader.Upgrade(c.Writer, c.Request, resHeader)
	if err != nil {
		return nil, "", logs.SysErr(err)
	}
	return &Ws{conn: conn}, token, nil
}

//绑定uid和conn
func (this *Ws) BindCoon(bid string) {
	client := User{Bid: bid, conn: this.conn}
	member[bid] = &client
	uidMapWs[bid] = this.conn
	wsMapUid[this.conn] = bid
}

//是否在线
func (this *Ws) IsOnline(bid string) bool {
	_, exits := member[bid]
	return exits
}

//通过Id断开连接
func (this *Ws) CloseBindId(bid string) () {
	delete(member, bid)
	delete(uidMapWs, bid)
	delete(wsMapUid, this.conn)
	this.conn.Close()
}

//通过conn断开直接关闭连接
func (this *Ws) CloseCoon() () {
	this.conn.Close()
}

//群发消息
func (this *Ws) SendToAll(content interface{},t string) {
	m := new(Message)
	m.Content = content
	m.Type = t
	sendMess, _ := json.Marshal(m)
	for k, v := range member {
		if v.conn != this.conn {
			if err := v.conn.WriteMessage(1, sendMess); err != nil {
				logs.SysErr(err)
				delete(member, k)
				delete(uidMapWs, k)
				delete(wsMapUid, v.conn)
				continue
			}
		}
	}
}

//获取当前组人数
func (this *Ws) GetClientCountByGroup(groupName string) int {
	return len(groupMapMember[groupName])
}

//获取某个群的详细信息
func (this *Ws) GetClientByGroup(groupName string) []*User {
	return groupMapMember[groupName]
}

//加入某个群
func (this *Ws) JoinGroup(groupName, bid string) []*User {
	groupMapMember[groupName] = append(groupMapMember[groupName], member[bid])
	return groupMapMember[groupName]
}

//给指定群发消息
func (this *Ws) SendToGroup(groupName string, content interface{}, t string) {
	m := new(Message)
	m.Content = content
	m.Type = t
	sendMess, _ := json.Marshal(content)

	for k, v := range groupMapMember[groupName] {
		if v.conn != this.conn {
			if err := v.conn.WriteMessage(1, sendMess); err != nil {
				logs.SysErr(err)
				//如果发送断裂，则该socket掉线
				//删除当前组下面的切面中的元素即成员
				kk := k + 1
				groupMapMember[groupName] = append(groupMapMember[groupName][:k], groupMapMember[groupName][kk:]...)
				continue
			}
		}
	}
}

//离开某个群
func (this *Ws) LeaveGroup(groupName, bid string) {
	for k, v := range groupMapMember[groupName] {
		if v.Bid == bid {
			kk := k + 1
			groupMapMember[groupName] = append(groupMapMember[groupName][:k], groupMapMember[groupName][kk:] ...)
			break
		}
	}
}

//发送给指定uid
func (this *Ws) SendToUid(bid string, content interface{}, t string) error {
	toWsCoon := uidMapWs[bid]
	m := new(Message)
	m.Content = content
	m.Type = t
	sendMess, _ := json.Marshal(m)
	if err := toWsCoon.WriteMessage(1, sendMess); err != nil {
		delete(member, bid)
		return logs.SysErr(err)
	}

	return nil
}

//给当前连接发消息
func (this *Ws) SendSelf(content interface{}, t string) error {
	m := new(Message)
	m.Content = content
	m.Type = t
	sendMess, _ := json.Marshal(m)
	if err := this.conn.WriteMessage(1, sendMess); err != nil {
		return logs.SysErr(err)
	}
	return nil
}

//解析客户端消息
func (this *Ws) GetMsg(msg *Message) error {
	var err error
	var reply []byte
	if _, reply, err = this.conn.ReadMessage(); err != nil {
		return logs.SysErr(err)
	}
	this.logReqData(reply)
	if err = json.Unmarshal(reply, msg); err != nil {
		return logs.SysErr(err)
	}
	return nil
}

//记录请求数据
func (this *Ws) logReqData(reply []byte) {
	var logStr string
	bid := wsMapUid[this.conn]
	logStr += fmt.Sprintf("\nBid:%s\n", bid)
	logStr += fmt.Sprintf("%s\n", string(reply))
	logStr += fmt.Sprintf("%s", "------------------------------------")
	logs.LogsWs.Info("%s", logStr)
}
