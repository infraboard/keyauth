package impl

import (
	"context"

	"github.com/infraboard/mcube/exception"

	"github.com/infraboard/keyauth/app/application"
)

func (s *service) save(app *application.Application) (
	*application.Application, error) {
	if _, err := s.col.InsertOne(context.TODO(), app); err != nil {
		return nil, exception.NewInternalServerError("inserted application(%s) document error, %s",
			app.Name, err)
	}
	return app, nil
}
