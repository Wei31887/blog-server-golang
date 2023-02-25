package token

import (
	G "blog/server/global"
	"context"
	"fmt"
	"strconv"
	"time"
)

type BlackList struct {}

// GetJWTBlackList
func (maker *JWTMaker) GetBlackList(payload *Payload) string {
	return "black_list:" + payload.Id.String()
}

// JoinBlackList
func (maker *JWTMaker) JoinBlackList(payload *Payload) error {
	joinNow := time.Now()
	duration := payload.ExpiresAt.Sub(joinNow)
	fmt.Println(duration)
	err := G.GLOBAL_REDIS.SetNX(context.Background(), maker.GetBlackList(payload), joinNow.Unix(), duration).Err()
	return err
}

// IsInBlackList
func (maker *JWTMaker) IsInBlackList(payload *Payload) bool {
	joinUnixStr, err := G.GLOBAL_REDIS.Get(context.Background(), maker.GetBlackList(payload)).Result()

	if err != nil || len(joinUnixStr) == 0 {
		return false
	}

	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if err != nil {
		return false
	}

	gracePeriod := time.Now().Unix() - joinUnix
	if gracePeriod > int64(G.GLOBAL_CONFIG.JWT.BlacklistGracePeriod.Minutes()) {
		return false
	}
	return true
}