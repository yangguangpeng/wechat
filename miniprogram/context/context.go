package context

import (
	"github.com/yangguangpeng/wechat/v2/credential"
	"github.com/yangguangpeng/wechat/v2/miniprogram/config"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
