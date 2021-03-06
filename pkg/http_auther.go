package pkg

// // GetInternalAdminTokenCtx 内部调用时的模拟token
// func GetInternalAdminTokenCtx(account string) context.Context {
// 	return session.WithTokenContext(context.Background(), &token.Token{
// 		Account:  account,
// 		Domain:   domain.AdminDomainName,
// 		UserType: types.UserType_INTERNAL,
// 	})
// }
