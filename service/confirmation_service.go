package service

import (
	"time"
	"net/http"
	
	"rvkc/dto"
	"rvkc/util"
	"rvkc/models"
	"rvkc/converter"
	"rvkc/middleware"
	"rvkc/context_error"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ConfirmationService struct {
	serviceConfirmation GenericService[models.Confirmation]
	serviceClub         ClubService
	serviceAccount      AccountService
	log                 *logrus.Logger
}


func NewConfirmationService(
	serviceConfirmation GenericService[models.Confirmation],
	serviceClub ClubService,
	serviceAccount AccountService,
) *ConfirmationService {

	return &ConfirmationService{
		serviceConfirmation: serviceConfirmation,
		serviceClub:         serviceClub,
		serviceAccount:      serviceAccount,
		log:                 util.GetLogger(),
	}
}


func (c *ConfirmationService) CreateConfirmation(ctx *gin.Context) {
	var request dto.ConfirmationRequestList
	if err := middleware.ValidateJSONAndStruct(ctx, &request); err != nil {
		return
	}

	// // check accounts exist
	for _, confirmation := range request.Confirmations {
		if _, err := c.serviceAccount.GetById(ctx, *confirmation.AccountId); err != nil {
			return
		}
	}

	// check clubs exist
	clubsList := make([]*models.Club, 0)
	for _, confirmation := range request.Confirmations {
		club, err := c.serviceClub.GetById(*confirmation.ClubId)
		if err != nil {
			context_error.ClubNotFoundError(ctx)
			return
		}

		clubsList = append(clubsList, club)
	}

	// current date time is greater than start date club
	var currentTime = time.Now()
	for _, club := range clubsList {
		if currentTime.After(club.StartAt) {
			context_error.ClubInProgressOrFinishedError(ctx)
			return
		}
	}

	newConfirmation := converter.ToConfirmationEntityList(request.Confirmations, currentTime)

	// check if accounts already confirmed
	for _, nc := range newConfirmation {
		confirmation, _ := c.serviceConfirmation.FindBy("account_id = ? AND club_id = ? ORDER BY action_at DESC LIMIT 1", nc.AccountId, nc.ClubId)
		if confirmation != nil && nc.Status == confirmation.Status {
			context_error.ConfirmationStatusAlreadyRegisteredError(ctx, confirmation.Status)
			return
		}
	}

	// TODO check count slots

	if err := c.serviceConfirmation.CreateBatch(&newConfirmation); err != nil {
		context_error.ConfirmationPersistError(ctx)
		return
	}

	var confirmationPointers []*models.Confirmation
	for i := range newConfirmation {
		confirmationPointers = append(confirmationPointers, &newConfirmation[i])
	}

	ctx.JSON(http.StatusCreated, dto.ToConfirmationResponseList(confirmationPointers))
}


func (c *ConfirmationService) GetConfirmationByQueryParam(ctx *gin.Context) {
	clubID 	  := ctx.Query("club_id")
	accountID := ctx.Query("account_id")

	switch {	
		case clubID != "":
			c.getConfirmationsByClubId(ctx, clubID)
			return
		case accountID != "":
			c.GetConfirmationsByAccountId(ctx, accountID)
			return
		default:
			c.GetConfirmations(ctx)
			return
	}
}


func (c *ConfirmationService) getConfirmationsByClubId(ctx *gin.Context, clubID string) {
	subQuery := c.serviceConfirmation.SubQuery(
		"account_id, club_id, MAX(DATE_TRUNC('milliseconds', action_at)) AS last_action",
		"account_id, club_id",
	)

	joins := []interface{}{
		[]interface{}{
			`JOIN (?) AS latest 
			 ON confirmation.account_id = latest.account_id 
			 AND confirmation.club_id = latest.club_id 
			 AND DATE_TRUNC('milliseconds', confirmation.action_at) = latest.last_action`,
			subQuery,
		},
	}

	confirmations, err := c.serviceConfirmation.FindWithJoinsSubqueryAndPreloads(
		joins,
		[]string{"Account", "Club"},
		"confirmation.club_id = ? AND confirmation.status = ?",
		clubID, "CONFIRMED",
	)

	if err != nil {
		context_error.ConfirmationSearchError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToConfirmationResponseList(confirmations))
}


func (c *ConfirmationService) GetConfirmationsByAccountId(ctx *gin.Context, accountID string) {
	subQuery := c.serviceConfirmation.SubQuery(
		"account_id, club_id, MAX(DATE_TRUNC('milliseconds', action_at)) AS last_action",
		"account_id, club_id",
	)

	joins := []interface{}{
		[]interface{}{
			`JOIN (?) AS latest 
			 ON confirmation.account_id = latest.account_id 
			 AND confirmation.club_id = latest.club_id 
			 AND DATE_TRUNC('milliseconds', confirmation.action_at) = latest.last_action`,
			subQuery,
		},
	}

	confirmations, err := c.serviceConfirmation.FindWithJoinsSubqueryAndPreloads(
		joins,
		[]string{"Account", "Club"},
		"confirmation.account_id = ? AND confirmation.status = ?",
		accountID, "CONFIRMED",
	)


	if err != nil {
		context_error.ConfirmationSearchError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToConfirmationResponseList(confirmations))
}


func (c *ConfirmationService) GetConfirmation(ctx *gin.Context) {
	confirmationParam := ctx.Param("id")
	confirmation, err := c.getById(confirmationParam)
	if err != nil {
		context_error.ClubNotFoundError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToConfirmationResponse(confirmation))
}


func (c *ConfirmationService) GetConfirmations(ctx *gin.Context) {
	confirmations, err := c.serviceConfirmation.GetAll()
	if err != nil {
		context_error.ConfirmationSearchError(ctx)
		return
	}

	ctx.JSON(http.StatusOK, dto.ToConfirmationResponseList(confirmations))
}


func (c *ConfirmationService) getById(id string) (*models.Confirmation, error) {
	return c.serviceConfirmation.GetBy("id = ?", id)
}