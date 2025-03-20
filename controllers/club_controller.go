package controllers

import (
	"net/http"
	"rvkc/dto"
	"rvkc/middleware"
	"rvkc/models"
	"rvkc/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ClubController struct {
	service services.GenericService[models.Club]
	log     *logrus.Logger
}


func NewClubController(service services.GenericService[models.Club]) *ClubController {
	return &ClubController{
		service: service,
		log: 	logrus.New(),
	}
}


func (c *ClubController) CreateClub(ctx *gin.Context) {
	var request dto.ClubRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

    newClub := models.Club{
        Name:       *request.Name,
		Recurrence: *request.Recurrence,
		Weekday:    *request.Weekday,
        StartAt:    *request.StartAt,
		EndAt:      *request.EndAt,
    }

	newClub.Higienize()

	// if there is no error it means the club already exists.
	if _, err := c.service.GetBy("name = ?", newClub.Name); err == nil {
		ctx.JSON(http.StatusCreated, gin.H{"errors": "Já existe um club cadastrado com este nome."})
		return
	}
	
	if err := c.service.Create(&newClub); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao criar club"})
		return
	}

	ctx.JSON(http.StatusCreated, newClub)
}


func (c *ClubController) GetClub(ctx *gin.Context) {
	clubParam := ctx.Param("id")
	clubId, err := strconv.ParseUint(clubParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Club id inválido"})
	}

	clubs, err := c.service.GetByID(uint(clubId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar clubs"})
		return
	}

	ctx.JSON(http.StatusOK, clubs)
}


func (c *ClubController) GetClubs(ctx *gin.Context) {
	clubs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar clubs"})
		return
	}

	ctx.JSON(http.StatusOK, clubs)
}


func (c *ClubController) UpdateClub(ctx *gin.Context) {
	var request dto.ClubRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

    clubParam := ctx.Param("id")
	clubId, err := strconv.ParseUint(clubParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Club id inválido"})
	}

	persistedClub, err := c.service.GetByID(uint(clubId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar clubs"})
		return
	}

	updateClub := models.Club{
        ID: 		persistedClub.ID,
        Recurrence: *request.Recurrence,
		Weekday: 	*request.Weekday,
		StartAt: 	*request.StartAt,
		EndAt: 		*request.EndAt,
    }
	
	updateClub.Higienize()
    err = c.service.Update(&updateClub)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao atualizar club"})
        return
    }

    ctx.Status(http.StatusAccepted)
}


func (c *ClubController) DeleteClub(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao deletar club"})
		return
	}
	ctx.Status(http.StatusNoContent)
}