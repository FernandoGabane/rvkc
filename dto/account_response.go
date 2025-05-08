package dto

import (
	"rvkc/models"
)


type AccountResponse struct {
	ID       string   `json:"id,omitempty"`
	Document string   `json:"document,omitempty"`
	Name     string   `json:"name,omitempty"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Roles    []string `json:"roles,omitempty"`
}



func ToAccountResponse(account *models.Account) AccountResponse {
	roles := make([]string, len(account.Roles))
	for i, role := range account.Roles {
		roles[i] = role.Name
	}

	return AccountResponse{
		ID:       account.ID,
		Document: account.Document,
		Name:     account.Name,
		Phone:    account.Phone,
		Email:    account.Email,
		Roles:    roles,
	}
}


func ToAccountSimpleResponse(account *models.Account) AccountResponse {
	accountResponse := ToAccountResponse(account)

	accountResponse.Document = ""
	accountResponse.Phone = ""
	accountResponse.Email = ""

	return accountResponse
}


func ToAccountSimpleResponseList(account []*models.Account) []AccountResponse {
	accountResponseList := ToAccountResponseList(account)
	var accountResponseSimpleList []AccountResponse

	for _, ac := range accountResponseList {
		ac.Document = ""
		ac.Phone = ""
		ac.Email = ""
		accountResponseSimpleList = append(accountResponseSimpleList, ac)
	}

	return accountResponseSimpleList
}


func ToAccountResponseList(account []*models.Account) []AccountResponse {
	var accountResponseList []AccountResponse = make([]AccountResponse, 0)
	
	for _, account 	:= range account {
		accountResponse    := ToAccountResponse(account)
		accountResponseList = append(accountResponseList, accountResponse)
	}


	return accountResponseList
}
