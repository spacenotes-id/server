package container

import (
	"context"
	"log"
	"reflect"

	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/internal/application/usecase"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/database/postgres"
	postgresRepo "github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/repository/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/internal/interface/api/controller"
)

type bean struct {
	beanID   string
	beanType reflect.Type
}

func registerBeans(beans ...bean) {
	for _, bean := range beans {
		overwritten, err := di.RegisterBean(bean.beanID, bean.beanType)
		if err != nil {
			log.Fatalf("ERROR(%s): failed to register bean. Cause: %v \n",
				bean.beanID, err,
			)
		}

		if overwritten {
			log.Printf("INFO: %s is overwritten \n", bean.beanID)
		}
	}
}

func InitDi() {
	registerBeans(
		bean{
			beanID:   "userRepo",
			beanType: reflect.TypeOf((*postgresRepo.UserRepoPostgres)(nil)),
		},
		bean{
			beanID:   "hashingService",
			beanType: nil,
		},
		bean{
			beanID:   "userUsecase",
			beanType: reflect.TypeOf((*usecase.UserUsecase)(nil)),
		},
		bean{
			beanID:   "authController",
			beanType: reflect.TypeOf((*controller.AuthController)(nil)),
		},
		bean{
			beanID:   "userController",
			beanType: reflect.TypeOf((*controller.UserController)(nil)),
		},
	)

	if _, err := di.RegisterBeanInstance(
		"sqlcQuerier",
		postgres.GetPostgresSQLCQuerier(context.Background()),
	); err != nil {
		log.Fatalf(
			"ERROR(sqlcQuerier): %v. %v",
			"Failed to register sqlc querier",
			err,
		)
	}
}
