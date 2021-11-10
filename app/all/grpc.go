package all

import (
	_ "github.com/infraboard/keyauth/app/application/impl"
	_ "github.com/infraboard/keyauth/app/department/impl"
	_ "github.com/infraboard/keyauth/app/domain/impl"
	_ "github.com/infraboard/keyauth/app/endpoint/impl"
	_ "github.com/infraboard/keyauth/app/mconf/impl"
	_ "github.com/infraboard/keyauth/app/micro/impl"
	_ "github.com/infraboard/keyauth/app/namespace/impl"
	_ "github.com/infraboard/keyauth/app/permission/impl"
	_ "github.com/infraboard/keyauth/app/policy/impl"
	_ "github.com/infraboard/keyauth/app/role/impl"
	_ "github.com/infraboard/keyauth/app/session/impl"
	_ "github.com/infraboard/keyauth/app/tag/impl"
	_ "github.com/infraboard/keyauth/app/token/impl"
	_ "github.com/infraboard/keyauth/app/user/impl"
	_ "github.com/infraboard/keyauth/app/verifycode/impl"
)
