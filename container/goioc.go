package container

import (
	"reflect"

	"github.com/gofiber/fiber/v2/log"
	"github.com/goioc/di"
	"github.com/spacenotes-id/server/controller"
	postgresDB "github.com/spacenotes-id/server/database/postgres"
	"github.com/spacenotes-id/server/repository/postgres"
	"github.com/spacenotes-id/server/service"
	"github.com/spacenotes-id/server/usecase"
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

func InitDI() {
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
			beanID:   "spaceRepo",
			beanType: reflect.TypeOf((*postgres.SpaceRepoPostgres)(nil)),
		},
		bean{
			beanID:   "noteRepo",
			beanType: reflect.TypeOf((*postgres.NoteRepoPostgres)(nil)),
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
			beanID:   "spaceUsecase",
			beanType: reflect.TypeOf((*usecase.SpaceUsecase)(nil)),
		},
		bean{
			beanID:   "noteUsecase",
			beanType: reflect.TypeOf((*usecase.NoteUsecase)(nil)),
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
			beanID:   "spaceController",
			beanType: reflect.TypeOf((*controller.SpaceController)(nil)),
		},
		bean{
			beanID:   "noteController",
			beanType: reflect.TypeOf((*controller.NoteController)(nil)),
		},
	)

	db, err := postgresDB.GetPostgresSQLCQuerier()
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
