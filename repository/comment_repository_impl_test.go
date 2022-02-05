package repository

import (
	"context"
	"fmt"
	"github/andiahmads/go-learnsql/connection"
	"github/andiahmads/go-learnsql/entity"
	"testing"
)

func Test_commentRepositoryImpl_Insert(t *testing.T) {
	commentRepository := NewCommentRepository(connection.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "testing2@gmail.com",
		Comment: "from unit testing",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)

}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(connection.GetConnection())
	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)

}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(connection.GetConnection())
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	for _, comment := range comments {
		fmt.Println(comment)
	}

}
