package container

import (
	"log"
	"reflect"

	"github.com/goioc/di"
	"github.com/tfkhdyt/SpaceNotes/server/internal/application/usecase"
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
			beanType: nil,
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
}
