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
	"gorm.io/gorm"
)

type AccountService struct {
	accountService GenericService[models.Account]
	roleService    RoleService
}

func NewAccountService(
	accountService GenericService[models.Account],
	roleService RoleService,
) *AccountService {

	return &AccountService{
		accountService: accountService,
		roleService:    roleService,
	}
}

func (c *AccountService) CreateAccount(ctx *gin.Context) {
	var request dto.AccountRequest
	if err := middleware.ValidateJSONAndStruct(ctx, &request); err != nil {
		return
	}

	middleware.LoggerInfo("creating account", middleware.ToJsonString(request))
	roles, err := c.roleService.GetByName("DEFAULT")
	if err != nil {
		context_error.RoleNotFoundError(ctx)
		return
	}

	newAccount := converter.ToAccountEntity(&request)
	newAccount.Higienize()

	middleware.LoggerInfo("checking if account is already created with document", *request.Document)
	if _, err := c.FindAccountByDocument(ctx, newAccount.Document); err == nil {
		context_error.AccountAlreadyRegisteredError(ctx)
		return
	}

	associations := util.MakeAssociationSlice("Roles", []*models.Role{roles})
	accountPersistErro := c.accountService.CreateWithAssociations(&newAccount, associations)

	if accountPersistErro != nil {
		context_error.AccountPersistError(ctx)
		return
	}

	response := dto.ToAccountResponse(&newAccount)
	middleware.LoggerInfo("account created successfully", middleware.ToJsonString(response))

	ctx.JSON(http.StatusCreated, response)
}

func (c *AccountService) GetAccounts(ctx *gin.Context) {
	account, err := c.GetAllAndRoles(ctx)

	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAccountResponseList(account))
}

func (c *AccountService) GetAccountByDocument(ctx *gin.Context) {
	document := ctx.Param("document")
	// c.log.Info(fmt.Printf("Searching account by document: %v", document))

	account, err := c.GetByDocumentAndRoles(ctx, document)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAccountResponse(account))
}

func (c *AccountService) GetAccountSimpleByDocument(ctx *gin.Context) {
	document := ctx.Param("document")
	// c.log.Info(fmt.Printf("Searching account by document: %v", document))

	account, err := c.GetByDocumentAndRoles(ctx, document)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAccountSimpleResponse(account))
}

func (c *AccountService) GetAccountsSimple(ctx *gin.Context) {
	// c.log.Info(fmt.Printf("Searching all accounts simple."))

	account, err := c.GetAllAndRoles(ctx)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, dto.ToAccountSimpleResponseList(account))
}

func (c *AccountService) UpdateAccount(ctx *gin.Context) {
	var request dto.AccountRequest

	if err := middleware.ValidateJSON(ctx, &request); err != nil {
		return
	}

	if err := middleware.ValidateStruct(ctx, &request); err != nil {
		return
	}

	persistedAccount, err := c.FindAccountByDocument(ctx, *request.Document)
	if err != nil {
		context_error.AccountNotFoundError(ctx)
		return
	}

	requestRole := request.Roles

	UpdateAccount := converter.ToAccountEntity(&request)
	UpdateAccount.Higienize()
	UpdateAccount.ID = persistedAccount.ID
	UpdateAccount.Roles = nil

	roles, findRolesErr := c.roleService.GetRolesByNameList(ctx, requestRole)
	if findRolesErr != nil {
		return
	}

	associations := util.MakeAssociationSlice("Roles", roles)

	updateRoleErr := c.accountService.repo.UpdateWithAssociations(&UpdateAccount, associations)
	if updateRoleErr != nil {
		context_error.RolePersistError(ctx)
		return
	}

	err = c.accountService.Update(&UpdateAccount)
	if err != nil {
		context_error.AccountPersistError(ctx)
		return
	}

	ctx.Status(http.StatusAccepted)
}

func (c *AccountService) DeleteAccount(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.accountService.Delete(uint(id))
	if err != nil {
		context_error.AccountPersistError(ctx)
		return
	}
	ctx.Status(http.StatusNoContent)
}

func (c *AccountService) FindAccountByDocument(ctx *gin.Context, document string) (*models.Account, error) {
	return c.accountService.GetBy("document = ?", document)
}

func (c *AccountService) GetByDocument(ctx *gin.Context, document string) (*models.Account, error) {
	account, err := c.accountService.GetBy("document = ?", document)

	if err == gorm.ErrRecordNotFound {
		context_error.AccountNotFoundError(ctx)
		ctx.Abort()
		return nil, err
	}

	if err != nil {
		context_error.AccountSearchError(ctx)
		ctx.Abort()
		return nil, err
	}

	return account, nil
}

func (c *AccountService) GetById(ctx *gin.Context, id string) (*models.Account, error) {
	account, err := c.accountService.GetBy("id = ?", id)

	if err == gorm.ErrRecordNotFound {
		context_error.AccountNotFoundError(ctx)
		ctx.Abort()
		return nil, err
	}

	if err != nil {
		context_error.AccountSearchError(ctx)
		ctx.Abort()
		return nil, err
	}

	return account, nil
}

func (c *AccountService) GetByDocumentAndRoles(ctx *gin.Context, document string) (*models.Account, error) {
	account, accountServiceErr := c.GetByDocument(ctx, document)

	if accountServiceErr != nil {
		return nil, accountServiceErr
	}

	roles, err := c.roleService.GetRolesByAccount(ctx, account.ID)
	if err != nil {
		return nil, err
	}

	account.Roles = roles
	return account, nil
}

func (c *AccountService) GetAllAccounts(ctx *gin.Context) ([]*models.Account, error) {
	account, err := c.accountService.GetAll()

	if err != nil {
		context_error.AccountSearchError(ctx)
		ctx.Abort()
		return nil, err
	}

	return account, nil
}

func (c *AccountService) GetAllAndRoles(ctx *gin.Context) ([]*models.Account, error) {
	account, err := c.GetAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(account); i++ {
		roles, err := c.roleService.GetRolesByAccount(ctx, account[i].ID)

		if err != nil {
			return nil, err
		}

		account[i].Roles = roles
	}

	return account, nil
}
