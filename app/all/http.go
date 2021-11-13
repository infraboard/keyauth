package all

import (
	// 加载服务模块
	_ "github.com/infraboard/keyauth/app/application/http"
	_ "github.com/infraboard/keyauth/app/department/http"
	_ "github.com/infraboard/keyauth/app/domain/http"
	_ "github.com/infraboard/keyauth/app/endpoint/http"
	_ "github.com/infraboard/keyauth/app/ip2region/http"
	_ "github.com/infraboard/keyauth/app/mconf/http"
	_ "github.com/infraboard/keyauth/app/micro/http"
	_ "github.com/infraboard/keyauth/app/namespace/http"
	_ "github.com/infraboard/keyauth/app/permission/http"
	_ "github.com/infraboard/keyauth/app/policy/http"
	_ "github.com/infraboard/keyauth/app/provider/http"
	_ "github.com/infraboard/keyauth/app/role/http"
	_ "github.com/infraboard/keyauth/app/session/http"
	_ "github.com/infraboard/keyauth/app/system/http"
	_ "github.com/infraboard/keyauth/app/tag/http"
	_ "github.com/infraboard/keyauth/app/token/http"
	_ "github.com/infraboard/keyauth/app/user/http"
	_ "github.com/infraboard/keyauth/app/verifycode/http"
)
