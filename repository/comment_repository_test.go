package repository

import (
	"context"
	"fmt"
	db "go-database"
	"testing"

	"go-database/entity"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test Repository",
	}

	result, err := commentRepository.Insert(ctx, comment)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 32)

	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(db.GetConnection())

	comments, err := commentRepository.FindAll(context.Background())

	if err != nil {
		panic(err)
	}

	for _, comment := range comments {
		fmt.Println("Test", comment)
	}

}
