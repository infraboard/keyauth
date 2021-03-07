package issuer

import (
	"context"

	"github.com/infraboard/keyauth/pkg/application"
)

func (i *issuer) CheckClient(ctx context.Context, clientID, clientSecret string) (*application.Application, error) {
	req := application.NewDescriptApplicationRequest()
	req.ClientId = clientID
	app, err := i.app.DescribeApplication(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := app.CheckClientSecret(clientSecret); err != nil {
		return nil, err
	}

	return app, nil
}
