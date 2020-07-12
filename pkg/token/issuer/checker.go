package issuer

import (
	"github.com/infraboard/keyauth/pkg/application"
)

func (i *issuer) CheckClient(clientID, clientSecret string) (*application.Application, error) {
	req := application.NewDescriptApplicationRequest()
	req.ClientID = clientID
	app, err := i.app.DescriptionApplication(req)
	if err != nil {
		return nil, err
	}

	if err := app.CheckClientSecret(clientSecret); err != nil {
		return nil, err
	}

	return app, nil
}
