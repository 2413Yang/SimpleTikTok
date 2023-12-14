package repository

import (
	"context"
	"strconv"
	"usersvr/log"
	"usersvr/middleware/cache"
)

// CacheSetUser 将用户信息写入缓存
func CacheSetUser(u User) {

	// 用redis hash结构存储用户信息
	userKey := userKeyPrefix + strconv.FormatInt(u.Id, 10)
	if err := cache.GetRedisCli().HSet(context.Background(), userKey, map[string]interface{}{
		"id":               u.Id,
		"user_name":        u.Name,
		"password":         u.Password,
		"follow_count":     u.Follow,
		"follower_count":   u.Follower,
		"avatar":           u.Avatar,
		"background_image": u.BackgroundImage,
		"signature":        u.Signature,
		"total_favorited":  u.TotalFav,
		"favorite_count":   u.FavCount,
	}).Err(); err != nil {
		log.Error("CacheSetUser error:", err.Error())
		return
	}
}

// CacheGetUser 从缓存中获取用户信息
func CacheGetUser(userId int64) (User, error) {
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)
	data, err := cache.GetRedisCli().HGetAll(context.Background(), userKey).Result()
	if err != nil {
		log.Error("CacheGetUser error:", err.Error())
		return User{}, err
	}
	if len(data) == 0 {
		log.Error("CacheGetUser error:", err.Error())
		return User{}, nil
	}

	user := User{}
	user.Id, _ = strconv.ParseInt(data["id"], 10, 64)
	user.Name = data["user_name"]
	user.Password = data["password"]
	user.Follow, _ = strconv.ParseInt(data["follow_count"], 10, 64)
	user.Follower, _ = strconv.ParseInt(data["follower_count"], 10, 64)
	user.Avatar = data["avatar"]
	user.BackgroundImage = data["background_image"]
	user.Signature = data["signature"]
	user.TotalFav, _ = strconv.ParseInt(data["total_favorited"], 10, 64)
	user.FavCount, _ = strconv.ParseInt(data["favorite_count"], 10, 64)
	return user, nil
}

// CacheIncrByUserFavoritedNum 增加用户被点赞总数
func CacheIncrByUserFavoritedNum(userId int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)

	_, err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "total_favorited", num).Result()
	if err != nil {
		log.Error("CacheIncrByUserFavoritedNum error:", err.Error())
		return err
	}
	return nil
}

// CacheIncrByUserFavoriteNum 增加用户点赞总数
func CacheIncrByUserFavoriteNum(userId int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)

	_, err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "favorite_count", num).Result()
	if err != nil {
		log.Error("CacheIncrByUserFavoriteNum error:", err.Error())
		return err
	}
	return nil
}

// CacheIncrByUserFollowNum 增加用户关注总数
func CacheIncrByUserFollowNum(userId int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)

	_, err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "follow_count", num).Result()
	if err != nil {
		log.Error("CacheIncrByUserFollowNum error:", err.Error())
		return err
	}
	return nil
}

// CacheIncrByUserFollowerNum 增加用户粉丝数量
func CacheIncrByUserFollowerNum(userId int64, num int64) error {
	userKey := userKeyPrefix + strconv.FormatInt(userId, 10)

	_, err := cache.GetRedisCli().HIncrBy(context.Background(), userKey, "follower_count", num).Result()
	if err != nil {
		log.Error("CacheIncrByUserFollowerNum error:", err.Error())
		return err
	}
	return nil
}
