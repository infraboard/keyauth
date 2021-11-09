package all

import (
	// 加载服务模块
	_ "github.com/infraboard/keyauth/pkg/application/http"
	_ "github.com/infraboard/keyauth/pkg/department/http"
	_ "github.com/infraboard/keyauth/pkg/domain/http"
	_ "github.com/infraboard/keyauth/pkg/endpoint/http"
	_ "github.com/infraboard/keyauth/pkg/ip2region/http"
	_ "github.com/infraboard/keyauth/pkg/mconf/http"
	_ "github.com/infraboard/keyauth/pkg/micro/http"

	_ "github.com/infraboard/keyauth/pkg/namespace/http"
	_ "github.com/infraboard/keyauth/pkg/permission/http"
	_ "github.com/infraboard/keyauth/pkg/policy/http"
	_ "github.com/infraboard/keyauth/pkg/provider/http"

	_ "github.com/infraboard/keyauth/pkg/role/http"
	_ "github.com/infraboard/keyauth/pkg/session/http"
	_ "github.com/infraboard/keyauth/pkg/system/http"
	_ "github.com/infraboard/keyauth/pkg/tag/http"
	_ "github.com/infraboard/keyauth/pkg/token/http"
	_ "github.com/infraboard/keyauth/pkg/user/http"
	_ "github.com/infraboard/keyauth/pkg/verifycode/http"
)
