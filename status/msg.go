package status

var msg = []string{
	"Success",                          // 成功
	"Request Param Error",              // 请求参数错误
	"Unknown Error",                    // 未知错误
	"Username Has ExistedError ",       // 用户名已存在
	"Generate Token Error",             // 生成token出错
	"Token Is NULL",                    // Token为空
	"Token is Error",                   // Token 解析错误
	"Token Is Expired",                 // Token过期
	"User Not Exist Or Password Wrong", // 用户名不存在或密码错误
	"Load File Error",                  // 加载文件出错
	"Save Up loaded File Error",        // 保存文件出错
	"NoMoreVideos",                     //视频库已经看完
}

func Msg(code int) string {
	if code < 0 || code >= len(msg) {
		return ""
	}
	return msg[code]
}
