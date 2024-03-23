package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ListGenderRequest struct {
	Gender string `json:"gender"`
}

func (server *Server) listGender(ctx *gin.Context) {
	var req ListGenderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	listGender, err := server.store.ListGender(ctx, req.Gender)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, listGender)
}

type CountGender struct {
	Male   int `json:"male"`
	Female int `json:"female"`
}

func (server *Server) countGender(ctx *gin.Context) {
	genders, err := server.store.GetGender(ctx)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, err)
		return
	}

	count := CountGender{}
	for _, gender := range genders {
		if gender == "Male" {
			count.Male++
		} else if gender == "Female" {
			count.Female++
		}
	}
	ctx.JSON(http.StatusOK, count)
}

func (server *Server) countAge(ctx *gin.Context) {
	ages, err := server.store.GetAge(ctx)
	if err != nil {
		ctx.JSON(http.StatusNonAuthoritativeInfo, err)
		return
	}
	ageCount := make(map[int]int)
	for _, age := range ages {
		ageCount[int(age)]++
	}

	ctx.JSON(http.StatusOK, ageCount)
}
