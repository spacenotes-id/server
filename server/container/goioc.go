package container

import (
	"context"
	"reflect"

	"github.com/gofiber/fiber/v2/log"
	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/controller"
	postgresDB "github.com/tfkhdyt/SpaceNotes/server/database/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/repository/postgres"
	"github.com/tfkhdyt/SpaceNotes/server/route"
	"github.com/tfkhdyt/SpaceNotes/server/service"
	"github.com/tfkhdyt/SpaceNotes/server/usecase"
)

type bean struct {
	beanID   string
	beanType reflect.Type
}

func registerBeans(beans ...bean) {
	for _, bean := range beans {
		overwritten, err := di.RegisterBean(bean.beanID, bean.beanType)
		if err != nil {
			log.Fatalf("Failed to register %v. Cause: %v \n",
				bean.beanID, err,
			)
		}

		if overwritten {
			log.Infof("%s is overwritten \n", bean.beanID)
		}
	}
}

func InitDi() {
	ctx := context.Background()

	registerBeans(
		bean{
			beanID:   "userRepo",
			beanType: reflect.TypeOf((*postgres.UserRepoPostgres)(nil)),
		},
		bean{
			beanID:   "refreshTokenRepo",
			beanType: reflect.TypeOf((*postgres.RefreshTokenRepoPostgres)(nil)),
		},
		bean{
			beanID:   "bcryptService",
			beanType: reflect.TypeOf((*service.BcryptService)(nil)),
		},
		bean{
			beanID:   "jwtService",
			beanType: reflect.TypeOf((*service.JwtService)(nil)),
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

	db, err := postgresDB.GetPostgresSQLCQuerier(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := di.RegisterBeanInstance("querier", db); err != nil {
		log.Fatal(err)
	}

	if err := di.InitializeContainer(); err != nil {
		log.Fatal(err)
	}
}
