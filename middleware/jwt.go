package middleware

import (
	"encoding/json"
	"fmt"
	"jsonrpcmicro/api/response"
	"jsonrpcmicro/global"
	"jsonrpcmicro/internal/auth/model"
	"log"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//parseDuration, _ := time.ParseDuration("24h")
		sessionID := c.Request.Header.Get("sessionID")
		userJson, err := global.REDIS.Get(sessionID).Result()
		if err != nil {
			response.FailWithDetailed(gin.H{"reload": true}, "未登录或者登录已超时,请重新登录", c)
			c.Abort()
			return
		}
		var session model.Session
		err = json.Unmarshal([]byte(userJson), &session)
		if err != nil {
			log.Println("session 获取错误", err.Error())
		}
		opinion := AuthorityOpinion(c.Request.URL.Path, c.Request.Method, session.Authoritys)
		if !opinion {
			response.FailWithDetailed(gin.H{"reload": true}, "权限不足", c)
			c.Abort()
			return
		}
		c.Next()
		fmt.Println("鉴权生效")
	}
}

func AuthorityOpinion(path string, method string, authoritys []model.CasbinInfo) bool {
	for _, v := range authoritys {
		if v.Path == path && v.Method == method {
			return true
		}
	}
	return false
}

// func JWTAuth() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
// 		token := c.Request.Header.Get("x-token")
// 		if token == "" {
// 			response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
// 			c.Abort()
// 			return
// 		}
// 		if service.IsBlacklist(token) {
// 			response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
// 			c.Abort()
// 			return
// 		}
// 		j := NewJWT()
// 		// parseToken 解析token包含的信息
// 		claims, err := j.ParseToken(token)
// 		if err != nil {
// 			if err == TokenExpired {
// 				response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
// 				c.Abort()
// 				return
// 			}
// 			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
// 			c.Abort()
// 			return
// 		}
// 		if err, _ = service.FindUserByUuid(claims.UUID.String()); err != nil {
// 			_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
// 			response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
// 			c.Abort()
// 		}
// 		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
// 			claims.ExpiresAt = time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime
// 			newToken, _ := j.CreateToken(*claims)
// 			newClaims, _ := j.ParseToken(newToken)
// 			c.Header("new-token", newToken)
// 			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
// 			if global.GVA_CONFIG.System.UseMultipoint {
// 				err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
// 				if err != nil {
// 					global.GVA_LOG.Error("get redis jwt failed", zap.Any("err", err))
// 				} else { // 当之前的取成功时才进行拉黑操作
// 					_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
// 				}
// 				// 无论如何都要记录当前的活跃状态
// 				_ = service.SetRedisJWT(newToken, newClaims.Username)
// 			}
// 		}
// 		c.Set("claims", claims)
// 		c.Next()
// 	}
// }

// type JWT struct {
// 	SigningKey []byte
// }

// var (
// 	TokenExpired     = errors.New("Token is expired")
// 	TokenNotValidYet = errors.New("Token not active yet")
// 	TokenMalformed   = errors.New("That's not even a token")
// 	TokenInvalid     = errors.New("Couldn't handle this token:")
// )

// func NewJWT() *JWT {
// 	return &JWT{
// 		[]byte(global.GVA_CONFIG.JWT.SigningKey),
// 	}
// }

// // 创建一个token
// func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(j.SigningKey)
// }

// // 解析 token
// func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		if ve, ok := err.(*jwt.ValidationError); ok {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				return nil, TokenMalformed
// 			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
// 				// Token is expired
// 				return nil, TokenExpired
// 			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
// 				return nil, TokenNotValidYet
// 			} else {
// 				return nil, TokenInvalid
// 			}
// 		}
// 	}
// 	if token != nil {
// 		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
// 			return claims, nil
// 		}
// 		return nil, TokenInvalid

// 	} else {
// 		return nil, TokenInvalid

// 	}

// }

// //更新token
// func (j *JWT) RefreshToken(tokenString string) (string, error) {
// 	jwt.TimeFunc = func() time.Time {
// 		return time.Unix(0, 0)
// 	}
// 	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
// 		jwt.TimeFunc = time.Now
// 		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*60*24*7
// 		return j.CreateToken(*claims)
// 	}
// 	return "", TokenInvalid
// }
