package util

import (
	"github.com/gorilla/websocket"
	"time"
	"encoding/json"
	"net/http"
	"io"
	"crypto/rand"
	"encoding/base64"
	"hatgo/pkg/logs"
)

type ws struct {
	conn *websocket.Conn
}

//消息体
type Message struct {
	Content interface{} `json:"content"`
	Type    string      `json:"type"`
	Time    int64       `json:"time"`
}

//用户体
type User struct {
	Uid  string
	conn *websocket.Conn
}

var (
	member         = make(map[string]*User)
	uidMapWs       = make(map[string]*websocket.Conn)
	groupMapMember = make(map[string][]*User)
)

func CoonId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}

	return Md5(base64.URLEncoding.EncodeToString(b))
}

var Wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//实例化
func NewWs(wss *websocket.Conn) *ws {
	return &ws{conn: wss}
}

//绑定uid和conn
func (this *ws) BindUid(uid string) {
	client := User{Uid: uid, conn: this.conn}
	member[uid] = &client
	uidMapWs[uid] = this.conn
}

//是否在线
func (this *ws) IsOnline(uid string) bool {
	_, exits := member[uid]
	return exits
}

//断开连接
func (this *ws) CloseUid(uid string) () {
	delete(member, uid)
	delete(uidMapWs, uid)
	this.conn.Close()
}

//群发消息
func (this *ws) SendToAll(msg *Message) {
	msg.Time = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

	for k, v := range member {
		if v.conn != this.conn {
			if err := v.conn.WriteMessage(1, sendMess); err != nil {
				logs.SysErr(err)
				delete(member, k)
				delete(uidMapWs, k)
				continue
			}
		}
	}
}

//获取当前组人数
func (this *ws) GetClientCountByGroup(groupName string) int {
	return len(groupMapMember[groupName])
}

//获取某个群的详细信息
func (this *ws) GetClientByGroup(groupName string) []*User {
	return groupMapMember[groupName]
}

//加入某个群
func (this *ws) JoinGroup(groupName, uid string) []*User {
	groupMapMember[groupName] = append(groupMapMember[groupName], member[uid])
	return groupMapMember[groupName]
}

//给指定群发消息
func (this *ws) SendToGroup(groupName string, msg *Message) {
	msg.Time = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

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
func (this *ws) LeaveGroup(groupName, uid string) {
	for k, v := range groupMapMember[groupName] {
		if v.Uid == uid {
			kk := k + 1
			groupMapMember[groupName] = append(groupMapMember[groupName][:k], groupMapMember[groupName][kk:] ...)
			break
		}
	}
}

//发送给指定uid
func (this *ws) SendToUid(uid string, msg *Message) error {
	toWsCoon := uidMapWs[uid]
	msg.Time = time.Now().Unix()
	sendMess, _ := json.Marshal(msg)

	if err := toWsCoon.WriteMessage(1, sendMess); err != nil {
		delete(member, uid)
		return logs.SysErr(err)
	}

	return nil
}

//解析客户端消息
func (this *ws) GetMsg(msg *Message) error {
	var err error
	var reply []byte
	if _, reply, err = this.conn.ReadMessage(); err != nil {
		return logs.SysErr(err)
	}
	if err = json.Unmarshal(reply, msg); err != nil {
		return logs.SysErr(err)
	}

	return nil
}
