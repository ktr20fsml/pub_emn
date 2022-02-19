package repository

import (
	domainGeneral "api/domain/model/general"
	"context"
)

type GeneralRepository interface {
	CreateTableInformation(context.Context, *domainGeneral.TableInformation) error
	CreateTableInformations(context.Context, []*domainGeneral.TableInformation) error
	UpdateTableInformation(*domainGeneral.TableInformation) error
	UpdateTableInformations([]*domainGeneral.TableInformation) error
}
