package repository

import (
	"context"
	"database/sql"
	"errors"
	"github/andiahmads/go-learnsql/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{DB: db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlInsert := "INSERT INTO comments(email,comment) VALUES(?,?)"
	result, err := repository.DB.ExecContext(ctx, sqlInsert, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.ID = int64(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int64) (entity.Comment, error) {
	sql := "SELECT * FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, sql, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("id " + strconv.Itoa(int(id)) + "not found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	sql := "SELECT id,email,comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil

}
