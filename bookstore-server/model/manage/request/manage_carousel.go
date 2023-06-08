package request

import (
	"bookstore/model/common/request"
	"bookstore/model/manage"
)

type MallCarouselSearch struct {
	manage.MallCarousel
	request.PageInfo
}

type MallCarouselAddParam struct {
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank"`
}

type MallCarouselUpdateParam struct {
	CarouselId   int    `json:"carouselId"`
	CarouselUrl  string `json:"carouselUrl"`
	RedirectUrl  string `json:"redirectUrl"`
	CarouselRank string `json:"carouselRank" `
}
