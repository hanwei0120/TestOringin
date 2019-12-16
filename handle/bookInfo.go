package handle

import (
	"context"
	"gitTest2/proto"
)

type BookInfoService struct {
}

func (b *BookInfoService) FindInfoByBookNum(ctx context.Context, req *proto.BookRequest) (*proto.BookResponse, error) {
	model := map[uint32]proto.BookResponse{
		1: {
			PageCount: 500,
			BookName:  "哈利波特",
			Author:    "jk罗琳",
			Price:     "66",
		},
	}
	resp := model[req.BookNum]
	return &resp, nil
}
