package mongo

import (
	"context"

	"github.com/infraboard/mcube/exception"
	"github.com/infraboard/mcube/http/request"

	"github.com/infraboard/keyauth/pkg/department"
	"github.com/infraboard/keyauth/pkg/policy"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/keyauth/pkg/user"
)

func (s *service) saveAccount(u *user.User) error {
	if _, err := s.col.InsertOne(context.TODO(), u); err != nil {
		return exception.NewInternalServerError("inserted user(%s) document error, %s",
			u.Account, err)
	}

	return nil
}

func (s *service) queryAccount(req *queryUserRequest) (*user.Set, error) {
	userSet := user.NewUserSet(req.PageRequest)

	if !req.SkipItems {
		resp, err := s.col.Find(context.TODO(), req.FindFilter(), req.FindOptions())

		// 查询出该空间下的用户列表
		if req.NamespaceID != "" {
			ps, err := s.queryNamespacePolicy(req.GetToken(), req.NamespaceID)
			if err != nil {
				return nil, err
			}
			req.Accounts = ps.Users()
		}

		if err != nil {
			return nil, exception.NewInternalServerError("find user error, error is %s", err)
		}

		// 循环
		for resp.Next(context.TODO()) {
			u := new(user.User)
			if err := resp.Decode(u); err != nil {
				return nil, exception.NewInternalServerError("decode user error, error is %s", err)
			}

			// 补充用户的部门信息
			if req.WithDepartment && u.DepartmentID != "" {
				depart, err := s.depart.DescribeDepartment(department.NewDescriptDepartmentRequestWithID(u.DepartmentID))
				if err != nil {
					s.log.Errorf("get user department error, %s", err)
				} else {
					u.Department = depart
				}
			}

			u.Desensitize()
			userSet.Add(u)
		}
	}

	// count
	count, err := s.col.CountDocuments(context.TODO(), req.FindFilter())
	if err != nil {
		return nil, exception.NewInternalServerError("get device count error, error is %s", err)
	}
	userSet.Total = count

	return userSet, nil
}

func (s *service) queryNamespacePolicy(tk *token.Token, namespaceID string) (*policy.Set, error) {
	pReq := policy.NewQueryPolicyRequest(request.NewPageRequest(20, 1))
	pReq.NamespaceID = namespaceID
	pReq.WithToken(tk)
	return s.policy.QueryPolicy(pReq)
}
