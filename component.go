package validator

import (
	"github.com/Compogo/compogo"
	"github.com/go-playground/validator/v10"
)

// Component — компонент валидатора для Compogo.
// Регистрирует экземпляр validator.Validate в DI-контейнере.
//
// Валидатор используется для проверки структур в HTTP-хендлерах,
// GRPC-сервисах и других компонентах приложения.
//
// Пример использования:
//
//	type CreateUserRequest struct {
//	    Name  string `json:"name" validate:"required,min=3,max=50"`
//	    Email string `json:"email" validate:"required,email"`
//	    Age   int    `json:"age" validate:"gte=18,lte=100"`
//	}
//
//	func (api *UserAPI) Create(w http.ResponseWriter, r *http.Request) {
//	    var req CreateUserRequest
//	    // ... decode json
//
//	    if err := api.validator.Struct(req); err != nil {
//	        http.Error(w, err.Error(), http.StatusBadRequest)
//	        return
//	    }
//	    // ... обработка
//	}
//
// Подключение в приложении:
//
//	app.AddComponents(&validator.Component)
//
// Получение валидатора из контейнера:
//
//	var validate *validator.Validate
//	container.Invoke(func(v *validator.Validate) { validate = v })
var Component = compogo.Component{
	Init: compogo.StepFunc(func(container compogo.Container) error {
		return container.Provide(validator.New)
	}),
}
