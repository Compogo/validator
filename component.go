package validator

import (
	"github.com/Compogo/compogo/component"
	"github.com/Compogo/compogo/container"
	"github.com/go-playground/validator/v10"
)

var Component = &component.Component{
	Init: component.StepFunc(func(container container.Container) error {
		return container.Provide(validator.New)
	}),
}
