package all

import (
	// 加载服务模块
	_ "github.com/infraboard/keyauth/apps/application/http"
	_ "github.com/infraboard/keyauth/apps/department/http"
	_ "github.com/infraboard/keyauth/apps/domain/http"
	_ "github.com/infraboard/keyauth/apps/endpoint/http"
	_ "github.com/infraboard/keyauth/apps/ip2region/http"
	_ "github.com/infraboard/keyauth/apps/mconf/http"
	_ "github.com/infraboard/keyauth/apps/micro/http"
	_ "github.com/infraboard/keyauth/apps/namespace/http"
	_ "github.com/infraboard/keyauth/apps/permission/http"
	_ "github.com/infraboard/keyauth/apps/policy/http"
	_ "github.com/infraboard/keyauth/apps/provider/http"
	_ "github.com/infraboard/keyauth/apps/role/http"
	_ "github.com/infraboard/keyauth/apps/session/http"
	_ "github.com/infraboard/keyauth/apps/system/http"
	_ "github.com/infraboard/keyauth/apps/tag/http"
	_ "github.com/infraboard/keyauth/apps/token/http"
	_ "github.com/infraboard/keyauth/apps/user/http"
	_ "github.com/infraboard/keyauth/apps/verifycode/http"
)
