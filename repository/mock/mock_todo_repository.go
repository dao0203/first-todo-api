package mock

import (
	"first-todo-api/entity"

	"github.com/stretchr/testify/mock"
)

type MockTodoRepository struct {
	mock.Mock
}

func (_m *MockTodoRepository) Create(todo entity.Todo) error {
	args := _m.Called(todo)
	var r0 error
	if rf, ok := args.Get(0).(func(entity.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (_m *MockTodoRepository) FindAll() ([]entity.Todo, error) {
	args := _m.Called()
	var r0 []entity.Todo
	if rf, ok := args.Get(0).(func() []entity.Todo); ok {
		r0 = rf()
	} else {
		if args.Get(0) != nil {
			r0 = args.Get(0).([]entity.Todo)
		}
	}
	var r1 error
	if rf, ok := args.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (_m *MockTodoRepository) FindByID(id string) (entity.Todo, error) {
	args := _m.Called(id)
	var r0 entity.Todo
	if rf, ok := args.Get(0).(func(string) entity.Todo); ok {
		r0 = rf(id)
	} else {
		r0 = args.Get(0).(entity.Todo)
	}
	var r1 error
	if rf, ok := args.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}

func (_m *MockTodoRepository) Update(todo entity.Todo) error {
	args := _m.Called(todo)
	var r0 error
	if rf, ok := args.Get(0).(func(entity.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (_m *MockTodoRepository) Delete(todo entity.Todo) error {
	args := _m.Called(todo)
	var r0 error
	if rf, ok := args.Get(0).(func(entity.Todo) error); ok {
		r0 = rf(todo)
	} else {
		r0 = args.Error(0)
	}

	return r0
}
