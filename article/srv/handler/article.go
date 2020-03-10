package handler

import (
	"srv/pkg/model"
	db "srv/pkg/mysql"
	"context"
	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/util/log"

	pb "srv/proto/article"
)

type ArticleSrv struct{}

func (e *ArticleSrv) Total(ctx context.Context, req *pb.TotalRequest, rsp *pb.TotalResponse) error {
	count := 0
	err := db.DB.Table("articles").Where("deleted_at IS NULL").Count(&count).Error
	if err != nil {
		return err
	}

	rsp.Total = int64(count)

	return nil
}

func (e *ArticleSrv) Articles(ctx context.Context, req *pb.ArticlesRequest, rsp *pb.ArticlesResponse) error {
	var articles []model.Article
	var _db = db.DB
	var err error
	var deletedWhere string

	if req.Filter != nil && req.Filter.Deleted == true {
		deletedWhere = "deleted_at IS NOT NULL"
		_db = _db.Unscoped()
	}

	_db = _db.Where(deletedWhere).Where("id > ?", (req.Page - 1) / 10).Limit(req.PageSize)

	err = _db.Preload("Categorys").Preload("UserArticles").Find(&articles).Error
	if err != nil {
		log.Error("查询文章列表错误：", err)
		return nil
	}

	rspMap := make([]*pb.ArticleResponse, len(articles))

	for key, value := range(articles) {
		categoryMap := make([]*pb.CategorysResponse, len(value.Categorys))
		for categoryKey, categoryValue := range value.Categorys {

			categoryMap[categoryKey] = &pb.CategorysResponse{
				Id: int64(categoryValue.ID),
				Name: categoryValue.Name,
			}
		}

		rspMap[key] = &pb.ArticleResponse{
			Id:                   int64(value.ID),
			Content:              value.Content,
			Title:                value.Title,
			Excerpt:              value.Excerpt,
			CreateAt:             value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:             value.UpdatedAt.Format("2006-01-02 15:04:05"),
			Categorys:			  categoryMap,
			UserId:				  int64(value.UserArticles.UserId),
		}
	}

	rsp.Articles = rspMap

	return nil
}

func (e *ArticleSrv) Article(ctx context.Context, req *pb.ArticleRequest, rsp *pb.ArticleResponse) error {
	var article model.Article
	var userArtcle model.UserArticles

	err :=db.DB.Preload("Categorys").Take(&article, req.Id).Error
	if err != nil {
		return err
	}

	db.DB.Model(&model.Article{
		Model:        gorm.Model{
			ID: uint(req.Id),
		},
	}).Related(&userArtcle, "ArticleId")

	categoryMap := make([]*pb.CategorysResponse, len(article.Categorys))
	for categoryKey, categoryValue := range article.Categorys {

		categoryMap[categoryKey] = &pb.CategorysResponse{
			Id: int64(categoryValue.ID),
			Name: categoryValue.Name,
		}
	}

	rsp.Id = int64(article.ID)
	rsp.Title = article.Title
	rsp.Content = article.Content
	rsp.Excerpt = article.Excerpt
	rsp.CreateAt = article.CreatedAt.Format("2006-01-02 15:04:05")
	rsp.UpdateAt = article.UpdatedAt.Format("2006-01-02 15:04:05")
	rsp.UserId = int64(userArtcle.UserId)
	rsp.Categorys = categoryMap

	return nil
}

func (e *ArticleSrv) CreateArticle(ctx context.Context, req *pb.CreateArticleRequest, rsp *pb.ArticleResponse) error {
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *ArticleSrv) Stream(ctx context.Context, req *pb.StreamingRequest, stream pb.ArticleSrv_StreamStream) error {
	log.Logf("Received Article.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&pb.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *ArticleSrv) PingPong(ctx context.Context, stream pb.ArticleSrv_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
