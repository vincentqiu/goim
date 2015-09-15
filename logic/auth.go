package main

import (
	log "code.google.com/p/log4go"
	"github.com/Terry-Mao/goim/define"
	"strconv"
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
	if userId, roomId, err = tokenDecode(token); err != nil {
		log.Error("tokenDecode(\"%s\") error(%s)", token, err)
		return
	}
	if roomId < 1 {
		roomId = define.NoRoom
	}
	return
}
