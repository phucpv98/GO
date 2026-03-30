//go:build wireinject
// +build wireinject

package wire

import (
	"go-ecommerce/internal/controller"
	"go-ecommerce/internal/repo"
	"go-ecommerce/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		controller.NewUserController,
	)
	return new(controller.UserController), nil
}

// Run: wire
// IoC - Inversion of Control : giảm sự phụ thuộc giữa các thành phần trong hệ thống, giúp dễ dàng thay đổi và mở rộng mà không ảnh hưởng đến các phần khác.
// DI - Dependency Injection : một kỹ thuật để thực hiện IoC, trong đó các phụ thuộc được "injected" vào một đối tượng thay vì đối tượng tự tạo ra chúng. Điều này giúp tăng tính linh hoạt và khả năng kiểm thử của hệ thống.
// => IoC là một nguyên tắc thiết kế, trong khi DI là một kỹ thuật cụ thể để thực hiện nguyên tắc đó. DI giúp giảm sự phụ thuộc giữa các thành phần, làm cho hệ thống dễ dàng mở rộng và bảo trì hơn.
// => Thiết kế phần Controls : gom toàn bộ những thứ liên quan phụ thuộc lại -> wire go.
