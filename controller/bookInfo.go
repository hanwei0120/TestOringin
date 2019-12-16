package controller

import (
	"context"
	"gitTest2/proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"strconv"

	//"net/http"
)

func GetBookInfo(ctx *gin.Context) {
	conn, err := grpc.Dial("127.0.0.1:9898", grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc.Dial err", err)
	}
	client := proto.NewBookInfoServiceClient(conn)
	bookNum:=ctx.PostForm("BookNum")
	bookNum2, _ := strconv.Atoi(bookNum)
	req := &proto.BookRequest{BookNum: uint32(bookNum2)}


	resp, err := client.FindInfoByBookNum(context.Background(), req)
	if err != nil {
		log.Fatal("调用微服务错误:", err)
	}
	//jResp, err := json.Marshal(resp)
	//if err != nil {
	//	log.Fatal("json序列化错误:", err)
	//}
	//ctx.String(http.StatusOK, string(jResp))
	ctx.JSON(200, gin.H{
		"PageCount":   resp.PageCount,
		"bookName":    resp.BookName,
		"Author11223": resp.Author,
		"price":       resp.Price,
	})
}
