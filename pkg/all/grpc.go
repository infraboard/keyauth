package all

import (
	_ "github.com/infraboard/keyauth/pkg/application/impl"
	_ "github.com/infraboard/keyauth/pkg/department/impl"
	_ "github.com/infraboard/keyauth/pkg/domain/impl"
	_ "github.com/infraboard/keyauth/pkg/endpoint/impl"
	_ "github.com/infraboard/keyauth/pkg/mconf/impl"
	_ "github.com/infraboard/keyauth/pkg/micro/impl"
	_ "github.com/infraboard/keyauth/pkg/namespace/impl"
	_ "github.com/infraboard/keyauth/pkg/permission/impl"
	_ "github.com/infraboard/keyauth/pkg/policy/impl"
	_ "github.com/infraboard/keyauth/pkg/role/impl"
	_ "github.com/infraboard/keyauth/pkg/session/impl"
	_ "github.com/infraboard/keyauth/pkg/tag/impl"
	_ "github.com/infraboard/keyauth/pkg/token/impl"
	_ "github.com/infraboard/keyauth/pkg/user/impl"
	_ "github.com/infraboard/keyauth/pkg/verifycode/impl"
)
