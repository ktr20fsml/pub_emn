package repository

import (
	domainGeneral "api/domain/model/general"
	"context"
)

type GeneralRepository interface {
	CreateTableInformation(context.Context, *domainGeneral.TableInformation) error
	CreateTableInformations(context.Context, []*domainGeneral.TableInformation) error
	UpdateTableInformation(context.Context, *domainGeneral.TableInformation) error
	UpdateTableInformations(context.Context, []*domainGeneral.TableInformation) error
}
