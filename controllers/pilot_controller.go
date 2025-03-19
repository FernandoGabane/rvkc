package controllers

import (
	"fmt"
	"net/http"
	"rvkc/dto"
	"rvkc/middleware"
	"rvkc/models"
	"rvkc/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type PilotController struct {
	service services.GenericService[models.Pilot]
}

func NewPilotController(service services.GenericService[models.Pilot]) *PilotController {
	return &PilotController{service: service}
}

func (c *PilotController) CreatePilot(ctx *gin.Context) {
	var request dto.PilotRequest

	if err := middlewares.ValidateJSON(ctx, &request); err != nil {
		return 
	}

	if err := middlewares.ValidateStruct(ctx, &request); err != nil {
		return
	}

    newPilot := models.Pilot{
        Document: *request.Document,
        Name:     *request.Name,
        Phone:    *request.Phone,
        Email:    *request.Email,
    }

	// if there is no error it means the pilot already exists.
	if _, err := c.service.GetBy("document = ?", newPilot.Document); err == nil {
		ctx.JSON(http.StatusCreated, gin.H{"errors": "Já existe uma piloto cadastrado com este documento."})
		return
	}
	
	newPilot.Higienize()
	if err := c.service.Create(&newPilot); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao criar piloto"})
		return
	}

	ctx.JSON(http.StatusCreated, newPilot)
}

func (c *PilotController) GetPilots(ctx *gin.Context) {
	pilots, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao buscar pilotos"})
		return
	}

	ctx.JSON(http.StatusOK, pilots)
}

func (c *PilotController) GetPilotByDocument(ctx *gin.Context) {
	document := ctx.Param("document")
	log.Info(fmt.Printf("Searching pilot by document: %v", document))

	pilot, err := c.service.GetBy("document = ?", document)
	if err != nil {
		messageError := "Piloto não encontrado"
		log.Warn(fmt.Errorf("%v: %v", messageError, document))
		ctx.JSON(http.StatusNotFound, gin.H{"error": messageError})
		return
	}
	ctx.JSON(http.StatusOK, pilot)
}

func (c *PilotController) UpdatePilot(ctx *gin.Context) {
	var request dto.PilotRequest

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

	updatePilot := models.Pilot{
        ID: persistedPilot.ID,
        Document: persistedPilot.Document,
        Name:  *request.Name,
        Phone: *request.Phone,
        Email: *request.Email,   
    }
	
	updatePilot.Higienize()
    err = c.service.Update(&updatePilot)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao atualizar piloto"})
        return
    }

    ctx.Status(http.StatusAccepted)
}


func (c *PilotController) DeletePilot(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.service.Delete(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": "Erro ao deletar piloto"})
		return
	}
	ctx.Status(http.StatusNoContent)
}