package main

import (
	log "code.google.com/p/log4go"
<<<<<<< HEAD
	_ "github.com/Terry-Mao/goim/define"
=======
	"github.com/Terry-Mao/goim/define"
>>>>>>> master
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
<<<<<<< HEAD
	// token verify, and return uid, roomid;
	if userId, roomId, err = tokenDecode(token); err != nil {
		log.Error("tokenDecode(\"%s\") error(%s)", token, err)
		return
=======
	if userId, roomId, err = tokenDecode(token); err != nil {
		log.Error("tokenDecode(\"%s\") error(%s)", token, err)
		return
	}
>>>>>>> master
	if roomId < 1 {
		roomId = define.NoRoom
	}

	//userId = 0
	//roomId = define.NoRoom
	return
}
