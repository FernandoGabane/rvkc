package service

import (
	"fmt"
	"net/http"
	"rvkc/converter"
	"rvkc/dto"
	"rvkc/middleware"
	"rvkc/models"
	"rvkc/util"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)



type AccountService struct {
	service GenericService[models.Account]
	log     *logrus.Logger
}


func NewAccountService(service GenericService[models.Account]) *AccountService {
	return &AccountService{
		service: service,
		log:     util.GetLogger(),
	}
}


func (c *AccountService) CreatePilot(ctx *gin.Context) {
	var request dto.AccountRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

	newPilot := converter.ToAccountEntity(&request)
	newPilot.Higienize()

	if _, err := c.service.GetBy("document = ?", newPilot.Document); err == nil {
		// if there is no error it means the pilot already exists.
		ctx.JSON(http.StatusCreated, gin.H{"errors": "Já existe uma piloto cadastrado com este documento."})
		return
	}
	
	if err := c.service.Create(&newPilot); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao criar piloto."})
		return
	}

	ctx.JSON(http.StatusCreated, newPilot)
}


func (c *AccountService) GetPilots(ctx *gin.Context) {
	pilots, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar pilotos."})
		return
	}

	ctx.JSON(http.StatusOK, pilots)
}


func (c *AccountService) GetPilotByDocument(ctx *gin.Context) {
	document := ctx.Param("document")
	c.log.Info(fmt.Printf("Searching pilot by document: %v", document))

	pilot, err := c.service.GetBy("document = ?", document)
	if err != nil {
		messageError := "Piloto não encontrado"
		c.log.Warn(fmt.Errorf("%v: %v", messageError, document))
		ctx.JSON(http.StatusNotFound, gin.H{"error": messageError})
		return
	}
	ctx.JSON(http.StatusOK, pilot)
}

func (c *AccountService) UpdatePilot(ctx *gin.Context) {
	var request dto.AccountRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

    persistedPilot, err := c.service.GetBy("document = ?", request.Document)
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"errors": "Piloto não encontrado."})
        return
    }

	updatePilot := converter.ToAccountEntity(&request)
	updatePilot.Higienize()
	updatePilot.ID = persistedPilot.ID


    err = c.service.Update(&updatePilot)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao atualizar piloto."})
        return
    }

    ctx.Status(http.StatusAccepted)
}


func (c *AccountService) DeletePilot(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao deletar piloto."})
		return
	}
	ctx.Status(http.StatusNoContent)
}