package service

import (
	"net/http"

	"rvkc/context_error"
	"rvkc/converter"
	"rvkc/dto"
	"rvkc/middleware"
	"rvkc/models"
	"rvkc/util"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type ClubService struct {
	serviceClub			GenericService[models.Club]
	serviceAccount      AccountService
	log     			*logrus.Logger
}


func NewClubService(
	serviceClub GenericService[models.Club],
	serviceAccount      AccountService,
) *ClubService {
	
		return &ClubService{
		serviceClub: 	serviceClub,
		serviceAccount: serviceAccount,
		log:     	    util.GetLogger(),
	}
}


func (c *ClubService) CreateClub(ctx *gin.Context) {
	var request dto.ClubRequest
	if err := middleware.ValidateJSONAndStruct(ctx, &request); err != nil {
		return
	}

	// check account exists
	if _, err := c.serviceAccount.GetById(ctx, *request.AccountId); err != nil {
		return
	}
	
	// check start_at is before current time
	if  !middleware.TimeCompareFutureValidator(request.StartAt.Time) {
		context_error.ClubStartAtError(ctx)
		return
	}

	newClub := converter.ToClubEntity(&request)
	newClub.Higienize()

	c.checkOverrideRegister(ctx, newClub)
	if ctx.IsAborted() {
		return
	}
	
	if err := c.serviceClub.Create(&newClub); err != nil {
		context_error.ClubPersistError(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, dto.ToClubResponse(&newClub))
}


func (c *ClubService) GetClub(ctx *gin.Context) {
	clubParam := ctx.Param("id")

	clubs, err := c.GetById(clubParam)
	if err != nil {
		context_error.ClubNotFoundError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToClubResponse(clubs))
}


func (c *ClubService) GetClubs(ctx *gin.Context) {
	clubs, err := c.serviceClub.GetAll()
	if err != nil {
		context_error.ClubSearchError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToClubResponseList(clubs))
}


func (c *ClubService) UpdateClub(ctx *gin.Context) {
	var request dto.ClubRequest
	if err := middleware.ValidateJSONAndStruct(ctx, &request); err != nil {
		return
	}
	// check start_at is before current time
	if  !middleware.TimeCompareFutureValidator(request.StartAt.Time) {
		context_error.ClubStartAtError(ctx)
		return
	}

	// check account exists
	if _, err := c.serviceAccount.GetById(ctx, *request.AccountId); err != nil {
		return
	}

    clubParam := ctx.Param("id")
	persistedClub, err := c.GetById(clubParam)
	if err != nil {
		context_error.ClubNotFoundError(ctx)
		return
	}

	updateClub    := converter.ToClubEntity(&request)
	updateClub.ID = persistedClub.ID
	updateClub.Higienize()

	c.checkOverrideRegister(ctx, updateClub)
	if ctx.IsAborted() {
		return
	}
	
    err = c.serviceClub.Update(&updateClub)
    if err != nil {
        context_error.ClubPersistError(ctx)
        return
    }

    ctx.Status(http.StatusAccepted)
}


func (c *ClubService) DeleteClub(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.serviceClub.Delete(uint(id))
	if err != nil {
		context_error.ClubPersistError(ctx)
		return
	}
	ctx.Status(http.StatusNoContent)
}


func (c *ClubService) GetById(id string) (*models.Club, error) {
	return c.serviceClub.GetBy("id = ?", id)
}


func (c *ClubService) checkOverrideRegister(ctx *gin.Context, newClub models.Club) {
	register, err := c.serviceClub.GetBy(
		"start_at < ? AND end_at > ?",
		newClub.EndAt, newClub.StartAt,
	)

	if register.ID != "" {
		context_error.ClubAlreadyRegisteredError(ctx)
		ctx.Abort()
		return
	}

	if err != gorm.ErrRecordNotFound {
		context_error.ClubAlreadyRegisteredGenericError(ctx)
		ctx.Abort()
		return
	}
}