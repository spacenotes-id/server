package container

import (
	"context"
	"log"
	"reflect"

	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/internal/application/usecase"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/database/postgres"
	postgresRepo "github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/repository/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/internal/infrastructure/security"
	"github.com/tfkhdyt/SpaceNotes/server/internal/interface/api/controller"
	"github.com/tfkhdyt/SpaceNotes/server/internal/interface/api/route"
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
			beanID:   "refreshTokenRepo",
			beanType: reflect.TypeOf((*postgresRepo.RefreshTokenRepoPostgres)(nil)),
		},
		bean{
			beanID:   "hashingService",
			beanType: reflect.TypeOf((*security.BcryptService)(nil)),
		},
		bean{
			beanID:   "authTokenService",
			beanType: reflect.TypeOf((*security.JwtService)(nil)),
		},
		bean{
			beanID:   "userUsecase",
			beanType: reflect.TypeOf((*usecase.UserUsecase)(nil)),
		},
		bean{
			beanID:   "authUsecase",
			beanType: reflect.TypeOf((*usecase.AuthUsecase)(nil)),
		},
		bean{
			beanID:   "authController",
			beanType: reflect.TypeOf((*controller.AuthController)(nil)),
		},
		bean{
			beanID:   "userController",
			beanType: reflect.TypeOf((*controller.UserController)(nil)),
		},
		bean{
			beanID:   "authRoute",
			beanType: reflect.TypeOf((*route.AuthRoute)(nil)),
		},
		bean{
			beanID:   "userRoute",
			beanType: reflect.TypeOf((*route.UserRoute)(nil)),
		},
	)

	if _, err := di.RegisterBeanInstance(
		"querier",
		postgres.GetPostgresSQLCQuerier(context.Background()),
	); err != nil {
		log.Fatalf(
			"ERROR(querier): %v. %v",
			"Failed to register sqlc querier",
			err,
		)
	}

	if err := di.InitializeContainer(); err != nil {
		log.Fatalln("ERROR(InitializeContainer):", err)
	}
}
