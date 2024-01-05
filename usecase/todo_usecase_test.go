package usecase_test

import (
	"context"
	"errors"
	"time"

	"first-todo-api/entity"
	mockr "first-todo-api/repository/mock"
	"first-todo-api/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTodoUsecase_Create(t *testing.T) {

	t.Run("success", func(t *testing.T) {
		mockTodo := entity.Todo{
			ID:    "id",
			Title: "title",
		}

		mockTodoRepository := new(mockr.MockTodoRepository)

		mockTodoRepository.On("Create", mockTodo).Return(nil)

		todoUsecase := usecase.NewTodoUsecase(mockTodoRepository, time.Second*2)

		err := todoUsecase.Create(context.Background(), mockTodo)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockTodo := entity.Todo{
			ID:    "id",
			Title: "title",
			Done:  false,
		}

		mockTodoRepository := new(mockr.MockTodoRepository)

		mockTodoRepository.On("Create", mockTodo).Return(errors.New("error"))

		todoUsecase := usecase.NewTodoUsecase(mockTodoRepository, time.Second*2)

		err := todoUsecase.Create(context.Background(), mockTodo)

		assert.Error(t, err)
	})
}
