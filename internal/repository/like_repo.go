package repository

import (
	"context"
	"mini_blog/internal/models"
)

func (r *Repository) CreateLike(ctx context.Context, l models.Like) (*models.Like, error) {
	err := r.DB.WithContext(ctx).Create(&l).Error
	if err != nil {
		return nil, err
	}

	return &l, nil
}

func (r *Repository) GetLikeByID(ctx context.Context, likeID int) (*models.Like, error) {
	like := models.Like{}
	err := r.DB.WithContext(ctx).First(&like, likeID).Error
	if err != nil {
		return nil, handleError(err)
	}

	return &like, nil
}

func (r *Repository) GetLikesByPostID(ctx context.Context, postID int) ([]*models.Like, error) {
	likes := []*models.Like{}
	err := r.DB.WithContext(ctx).Where("post_id = ?", postID).Find(&likes).Error
	if err != nil {
		return nil, handleError(err)
	}

	return likes, nil
}

func (r *Repository) CountLikes(ctx context.Context, postID int) (int, error) {
	like := models.Like{}
	var count int64
	err := r.DB.WithContext(ctx).Model(&like).Where("post_id = ?", postID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}