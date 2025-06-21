package controller

import (
	"net/http"
	"strconv"

	"github.com/nitin-kukreti/GoChat/internal/domain"
	"github.com/nitin-kukreti/GoChat/internal/usecase"
	"github.com/nitin-kukreti/GoChat/internal/utils"
)

type UserHandler struct {
	usecase *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{usecase: uc}
}

type createUserRequest struct {
	Name string `json:"name"`
}



func (u *UserHandler) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	var user createUserRequest

	user, err := utils.BodyParser[createUserRequest](req.Body)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, utils.MsgInvalidJSON)
		return
	}
	resUser, err := u.usecase.CreateUser(user.Name)
	if err != nil {
		utils.WriteError(res,http.StatusInternalServerError,utils.MsgUserCreationFailed)
		return
	}

	utils.WriteJSON(res,http.StatusOK,utils.MsgUserCreated,resUser);
}


func (u *UserHandler) GetUserById(res http.ResponseWriter, req *http.Request) {
	idStr := req.PathValue("id")

	userId, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, "Invalid user ID")
		return
	}

	user, err := u.usecase.GetUserByID(userId)
	if err != nil {
		if err == domain.ErrUserNotFound {
			utils.WriteError(res, http.StatusNotFound, utils.MSGUserNotFound)
			return;
		}
		utils.WriteError(res, http.StatusInternalServerError, utils.MSGUserGetFailed)
		return
	}

	utils.WriteJSON(res, http.StatusOK, utils.MSGUserGetSuccess, user)
}
