package errno

import "net/http"

var (
	// 100x 通用类型
	OK                    = &Errno{Code: http.StatusOK, Message: "成功"}
	ParamsError           = &Errno{Code: http.StatusBadRequest, Message: "参数错误"}
	AuthError             = &Errno{Code: http.StatusUnauthorized, Message: "无权访问"}
	InternalServerError   = &Errno{Code: http.StatusInternalServerError, Message: "服务器错误"}
	DatabaseError         = &Errno{Code: http.StatusInternalServerError, Message: "数据库错误"}
	TooManyRequestError   = &Errno{Code: http.StatusTooManyRequests, Message: "发送了太多请求"}
	SessionError          = &Errno{Code: http.StatusUnprocessableEntity, Message: "您的 Session 已过期"}
	NotFoundError         = &Errno{Code: http.StatusNotFound, Message: "Not Found"}
	ResourceNotFoundError = &Errno{Code: http.StatusNotAcceptable, Message: "资源未找到"}

	// 200x auth 相关
	SocialAuthorizationError = &Errno{Code: http.StatusBadRequest, Message: "第三方登录失败"}
	LoginError               = &Errno{Code: http.StatusBadRequest, Message: "用户名或密码错误"}
	TokenError               = &Errno{Code: http.StatusBadRequest, Message: "token error"}
	TokenExpireError         = &Errno{Code: http.StatusBadRequest, Message: "token 已过期"}
	TokenRefreshError        = &Errno{Code: http.StatusBadRequest, Message: "token 已过期(已过刷新时间)"}
	TokenInBlackListError    = &Errno{Code: http.StatusBadRequest, Message: "token 不可使用(已加入黑名单)"}
	TokenMissingError        = &Errno{Code: http.StatusBadRequest, Message: "token 没有找到"}

	// 300x 存储相关
	UploadError = &Errno{Code: 3000, Message: "上传失败"}

	// 500x 第三方错误
	SmsError = &Errno{Code: http.StatusInternalServerError, Message: "短信发送异常"}
	StatusBadGatewayError  = &Errno{Code: http.StatusBadGateway, Message: "网关错误"}
)
