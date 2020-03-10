package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	articleProto "api/proto/article"
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func (r *queryResolver) Article(ctx context.Context, id string) (*model.Article, error) {
	_id, err := graphql.UnmarshalIntID(id)
	if err != nil {
		return nil, err
	}

	articleRes, err := r.ArticleClient.Article(ctx, &articleProto.ArticleRequest{Id: int64(_id)})
	if err != nil {
		return nil, err
	}

	_createAt, err := time.ParseInLocation("2006-01-02 15:04:05", articleRes.CreateAt, time.Local)
	_updateAt, err := time.ParseInLocation("2006-01-02 15:04:05", articleRes.UpdateAt, time.Local)
	if err != nil {
		return nil, err
	}

	categoryMap := make([]*model.Category, len(articleRes.Categorys))
	for categoryKey, categoryValue := range articleRes.Categorys {
		_categoryId, err := graphql.UnmarshalID(categoryValue.Id)
		if err != nil {
			return nil, err
		}
		categoryMap[categoryKey] = &model.Category{
			ID:   _categoryId,
			Name: &categoryValue.Name,
		}
	}

	_userId, err := graphql.UnmarshalString(int(articleRes.UserId))
	User, err := r.Entity().FindUserByID(ctx, _userId)

	return &model.Article{
		ID:        id,
		Content:   &articleRes.Content,
		Title:     &articleRes.Title,
		Excerpt:   &articleRes.Excerpt,
		User:      User,
		CreateAt:  &_createAt,
		UpdateAt:  &_updateAt,
		Categorys: categoryMap,
	}, nil
}

func (r *queryResolver) Articles(ctx context.Context, first *int, after *string, filter *model.ArticlesFilter) (*model.Articles, error) {
	var articlesFilter articleProto.ArticlesFilter
	page, _ := graphql.UnmarshalInt(*after)
	if filter != nil {
		articlesFilter = articleProto.ArticlesFilter{
			Title:      *filter.Title,
			Deleted:    false,
			CategoryId: *filter.Category,
			Asc:        *filter.Asc,
		}
	}
	res, err := r.ArticleClient.Articles(ctx, &articleProto.ArticlesRequest{
		Page:     int64(page),
		PageSize: int64(*first),
		Filter:   &articlesFilter,
	})
	if err != nil {
		return nil, err
	}

	articlesEdges := make([]*model.ArticlesEdges, len(res.Articles))

	for key, value := range res.Articles {
		_id, _ := graphql.UnmarshalID(value.Id)
		_userId, _ := graphql.UnmarshalString(int(value.UserId));
		user, _ := r.Entity().FindUserByID(ctx, _userId)
		_createAt, err := time.ParseInLocation("2006-01-02 15:04:05", value.CreateAt, time.Local)
		_updateAt, err := time.ParseInLocation("2006-01-02 15:04:05", value.UpdateAt, time.Local)
		if err != nil {
			return nil, err
		}
		categoryMap := make([]*model.Category, len(value.Categorys))
		for categoryKey, categoryValue := range value.Categorys {
			_categoryId, err := graphql.UnmarshalID(categoryValue.Id)
			if err != nil {
				return nil, err
			}
			categoryMap[categoryKey] = &model.Category{
				ID:   _categoryId,
				Name: &categoryValue.Name,
			}
		}
		articlesEdges[key] = &model.ArticlesEdges{
			Node: &model.Article{
				ID:        _id,
				Content:   &value.Content,
				Title:     &value.Title,
				Excerpt:   &value.Excerpt,
				User:      user,
				CreateAt:  &_createAt,
				UpdateAt:  &_updateAt,
				Categorys: categoryMap,
			},
		}
	}

	countRes, err := r.ArticleClient.Total(ctx, &articleProto.TotalRequest{})
	_count, err := graphql.UnmarshalInt(countRes.Total)
	if err != nil {
		return nil, err
	}

	startCursor, err := graphql.UnmarshalString(res.Articles[0].Id)
	endCursor, err := graphql.UnmarshalString(res.Articles[len(res.Articles)-1].Id)
	if err != nil {
		return nil, err
	}

	return &model.Articles{
		PageInfo: &model.PageInfo{
			HasNextPage:     res.Articles[len(res.Articles)-1].Id < countRes.Total,
			HasPreviousPage: res.Articles[0].Id > 1,
			StartCursor:     &startCursor,
			EndCursor:       &endCursor,
		},
		Count:    &_count,
		Edges:    articlesEdges,
	}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
