package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"pinkacg/dao/mysql"
	"pinkacg/grpc/user_reco"
	"pinkacg/models"
	"strconv"
	"time"
)

const (
	address     = "hadoop102:9898"
	defaultName = "world"
)

func GetRecommendPost(p *models.PostRecommendList, userId int64) (err error, recoPosts []*models.Post) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return err, nil
	}
	defer conn.Close()
	c := user_reco.NewUserRecommendClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UserRecommend(ctx, &user_reco.User{
		UserId:     strconv.FormatInt(userId, 10),
		CategoryId: int32(p.CategoryId),
		ArticleNum: int32(p.ArticleNum),
		TimeStamp:  p.TimeStamp,
	})
	for _, recommend := range r.Recommends {
		post, _ := mysql.GetPostById(recommend.PostId)
		post.RecoPost = recommend.Params
		post.RecoTimeStamp = r.TimeStamp
		recoPosts = append(recoPosts, post)
	}
	return err, recoPosts
}
