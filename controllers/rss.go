package controllers

import (
	"github.com/OhMinsSup/story-server/repository"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/feeds"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetEntireFeed(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	postRepository := repository.NewPostRepository(db)
	posts, code, err := postRepository.FeedPostList()
	if err != nil {
		ctx.AbortWithError(code, err)
		return
	}

	feed := &feeds.Feed{
		Title:       "story",
		Id:          "https://storeis.vercel.app/",
		Link:        &feeds.Link{Href: "https://storeis.vercel.app/"},
		Description: "개발자들을 위한 블로그 서비스. 어디서 글 쓸지 고민하지 말고 벨로그에서 시작하세요.",
		Image:       &feeds.Image{Url: "https://images.velog.io/velog.png"},
		Copyright:   "Copyright (C) 2020. Story. All rights reserved.",
	}

	var feedItems []*feeds.Item
	for _, post := range posts {
		link := "https://storeis.vercel.app/#/story/" + post.ID
		feedItem := feeds.Item{
			Title:       post.Title,
			Id:          link,
			Description: post.Body,
			Link:        &feeds.Link{Href: link},
			Created:     post.CreatedAt,
			Author: &feeds.Author{
				Name:  post.DisplayName,
				Email: post.Email,
			},
		}
		feedItems = append(feedItems, &feedItem)
	}

	feed.Items = feedItems

	rss, err := feed.ToRss()
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.XML(http.StatusOK, rss)
}
