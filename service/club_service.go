package service

import (
	"net/http"
	"rvkc/converter"
	"rvkc/dto"
	"rvkc/middleware"
	"rvkc/models"
	"rvkc/util"
	"strconv"
	
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClubService struct {
	service GenericService[models.Club]
	log     *logrus.Logger
}


func NewClubClubService(service GenericService[models.Club]) *ClubService {
	return &ClubService{
		service: service,
		log:     util.GetLogger(),
	}
}


func (c *ClubService) CreateClub(ctx *gin.Context) {
	var request dto.ClubRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

	newClub := converter.ToClubEntity(&request)
	newClub.Higienize()

	_, err := c.service.GetBy(`
			date = ? AND (
				(start_at <= ? AND end_at > ?) OR
				(start_at < ? AND end_at >= ?) OR
				(start_at >= ? AND end_at <= ?)
			)`, newClub.Date, newClub.StartAt, newClub.StartAt, newClub.EndAt, newClub.EndAt, newClub.StartAt, newClub.EndAt,
	)

	if err == nil {
		// there is a club registed conflicting start and end time.
		ctx.JSON(http.StatusConflict, gin.H{"errors": "Já existe um clube com horário conflitante neste dia."})
		return
	}

	if err != gorm.ErrRecordNotFound {
		// database generic error
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao verificar conflitos."})
		return
	}

	if err := c.service.Create(&newClub); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao criar club."})
		return
	}

	ctx.JSON(http.StatusCreated, newClub)
}


func (c *ClubService) GetClub(ctx *gin.Context) {
	clubParam := ctx.Param("id")
	clubId, err := strconv.ParseUint(clubParam, 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Club id inválido."})
	}

	clubs, err := c.service.GetByID(uint(clubId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Club não encontrado."})
		return
	}

	ctx.JSON(http.StatusOK, clubs)
}


func (c *ClubService) GetClubs(ctx *gin.Context) {
	clubs, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar clubs."})
		return
	}

	ctx.JSON(http.StatusOK, clubs)
}


func (c *ClubService) UpdateClub(ctx *gin.Context) {
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Club não encontrado."})
		return
	}

	updateClub := converter.ToClubEntity(&request)
	updateClub.ID = persistedClub.ID
	updateClub.Higienize()
	
    err = c.service.Update(&updateClub)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao atualizar club."})
        return
    }

    ctx.Status(http.StatusAccepted)
}


func (c *ClubService) DeleteClub(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao deletar club."})
		return
	}
	ctx.Status(http.StatusNoContent)
}