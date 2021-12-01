package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/request"
	"time"
)

type Auth struct {
	jwtSession      *JWToken
	staffJwtSession *JWToken
}

func InitAuth() (a *Auth) {
	a = &Auth{
		jwtSession:      initJWT(),
		staffJwtSession: initStaffJWT(),
	}
	return
}

func (a *Auth) AppAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		platform := c.GetHeader("platform")
		switch platform {
		case "ios", "android":

		case "backend":
			c.Next()
			return
		default:
			c.Abort()
			return
		}

		tokenString, err := request.OAuth2Extractor.ExtractToken(c.Request)
		if err != nil {
			c.Abort()
			//response.Error(c, 10002, "token无效", gin.H{})
			return
		}
		claims, e := a.jwtSession.JwtVerify(tokenString)
		if e != nil {
			//.Debug("Expired token:", tokenString)
			//c.Abort()
			//response.Error(c, 10003, "token无效", gin.H{})
			return
		}
		userID, ok := claims["uid"].(float64)
		if !ok {
			c.Abort()
			//response.Error(c, 10004, "未登录", gin.H{})
			return
		}

		c.Set("userID", int(userID))
		c.Next()
	}
}

func (a *Auth) GetUserIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := request.OAuth2Extractor.ExtractToken(c.Request)
		if err == nil {
			claims, e := a.jwtSession.JwtVerify(tokenString)
			if e == nil {
				userID, ok := claims["uid"].(float64)
				if ok {

					c.Set("userID", int(userID))
				}
			}
		}
		c.Next()
	}
}

//func (a *Auth) AppLogin(db *gorm.DB, cache *redis.RedisClient, userID uint, from string, isAdminLogin ...bool) (jwttoken string, user *model.User, err error) {
//	user, err = model.GetUser(db, userID)
//	if !logger.Check(err) {
//		return
//	}
//	jwttoken, err = a.NewAppJwt(user.ID, isAdminLogin...)
//	if !logger.Check(err) {
//		return
//	}
//	valuePrefix := "user|"
//	keyPrefix := fmt.Sprintf("loginToken_%d_%s_", userID, from)
//	if len(isAdminLogin) > 0 && isAdminLogin[0] {
//		// 管理员验证码不走限制
//		valuePrefix = "admin|"
//	} else {
//		switch from {
//		case model.LoginTerminalFromPC:
//			keys, err := cache.GetKeysByPrefix(keyPrefix)
//			if logger.Check(err) && len(keys) >= 1 {
//				for _, v := range keys {
//					redisStr, err := cache.Get(v)
//					if logger.Check(err) {
//						redisSp := strings.Split(redisStr, "|")
//						if len(redisSp) > 1 {
//							switch redisSp[0] {
//							case "user":
//								cache.Del(v)
//							case "admin":
//							}
//						}
//					}
//				}
//			}
//		case model.LoginTerminalFromApp:
//			keys, err := cache.GetKeysByPrefix(keyPrefix)
//			if logger.Check(err) && len(keys) >= 2 {
//				userCount := 0
//				delKey := keys[0]
//				oldTime := time.Now().Add(time.Hour * 24)
//				for _, v := range keys {
//					redisStr, err := cache.Get(v)
//					if logger.Check(err) {
//						redisSp := strings.Split(redisStr, "|")
//						if len(redisSp) > 1 {
//							switch redisSp[0] {
//							case "user":
//								userCount++
//								tokenTime, err := time.ParseInLocation("2006-01-02 15:04:05", redisSp[1], time.Local)
//								if logger.Check(err) {
//									if tokenTime.Before(oldTime) {
//										oldTime = tokenTime
//										delKey = v
//									}
//								}
//							case "admin":
//							}
//						}
//					}
//				}
//				if userCount >= 2 {
//					cache.Del(delKey)
//				}
//			}
//		case model.LoginTerminalFromH5, model.LoginTerminalFromWxa:
//		}
//	}
//
//	key := keyPrefix + util.Md5(jwttoken)
//	logger.Check(cache.Set(key, valuePrefix+time.Now().Format("2006-01-02 15:04:05"), time.Minute*time.Duration(a.jwtSession.timeout)))
//	return
//}

func (a *Auth) GetAppUsedID(c *gin.Context) (userID uint, err error) {
	iVal, ok := c.Get("uid")
	if !ok {
		return 0, errors.New("not available userID")
	}
	userID, ok = iVal.(uint)
	if !ok {
		return 0, errors.New("not available userID")
	}
	return userID, nil
}

func (a *Auth) NewAppJwt(uid uint, isAdminLogin ...bool) (string, error) {
	if uid == 0 {
		return "", errors.New("uid is empty")
	}
	exp := time.Now().Add(time.Second * time.Duration(a.jwtSession.timeout)).Unix()
	if len(isAdminLogin) > 0 && isAdminLogin[0] {
		exp = time.Now().Add(time.Hour * 2).Unix()
	}
	return a.jwtSession.JwtCreate(jwt.MapClaims{
		"uid": uid,
		"exp": exp,
	})
}

func (a *Auth) GetJwtUserInfo(c *gin.Context, key string) (iVal interface{}, err error) {
	//if !StaffJwtKeyMap[key] {
	//	return nil, errors.New("获取jwt员工信息key值异常")
	//}
	tokenString, err := request.OAuth2Extractor.ExtractToken(c.Request)
	if err != nil {
		return nil, errors.New("获取jwt员工信息失败")
	}
	iVal, err = a.jwtSession.Param(tokenString, key)
	if err != nil {
		return nil, errors.New("获取jwt员工信息失败")
	}
	return
}
