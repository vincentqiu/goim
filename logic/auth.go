package main

import (
	log "code.google.com/p/log4go"
	_ "github.com/Terry-Mao/goim/define"
)

// developer could implement "ThirdAuth" interface for decide how get userId, or roomId
type Auther interface {
	Auth(token string) (userId int64, roomId int32, err error)
}

type DefaultAuther struct {
}

func NewDefaultAuther() *DefaultAuther {
	return &DefaultAuther{}
}

func (a *DefaultAuther) Auth(token string) (userId int64, roomId int32, err error) {
	// token verify, and return uid, roomid;
	if userId, roomId, err = tokenDecode(token); err != nil {
		log.Error("tokenDecode(\"%s\") error(%s)", token, err)
		return
	}

	//userId = 0
	//roomId = define.NoRoom
	return
}
