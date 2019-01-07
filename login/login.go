package login

import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"server/mysql"
	"server/proto"
)



var log = golog.New("login")

func LoginRequestCallback(ev cellnet.Event) {
	msg := ev.Message().(*proto.LoginReq)
	p := new(mysql.Player)
	p.Account = msg.UserName
	if p.GetPlayer() {
		if msg.UserPass == p.Password {
			res := proto.LoginRes {
				Status: proto.LoginStatus_LOGIN_SUCESS,       // 状态
				Attr: &proto.PlayerAttr {
					UserId: 123,
					NickName: "zhou",
					Level: 10,
					Exp: 12,
					Gold: 88,
					Diamond: 98,
				},
			}
			ev.Session().Send(&res)
		} else {
			res := proto.LoginRes {
				Status: proto.LoginStatus_LOGIN_ERR_PASS,       // 状态
				Attr: &proto.PlayerAttr {},
			}
			ev.Session().Send(&res)
		}
	} else {
		res := proto.LoginRes {
			Status: proto.LoginStatus_LOGIN_NO_USER,       // 状态
			Attr: &proto.PlayerAttr {},
		}
		ev.Session().Send(&res)
	}
}

func init() {
}

