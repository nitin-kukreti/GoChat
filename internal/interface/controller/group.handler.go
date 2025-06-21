package controller

import (
	"net/http"
	"github.com/nitin-kukreti/GoChat/internal/usecase"
	"github.com/nitin-kukreti/GoChat/internal/utils"
)

type GroupHandler struct {
	usecase usecase.GroupUseCase
}

func NewGroupHandler(usecase *usecase.GroupUseCase) *GroupHandler {
	return &GroupHandler{usecase: *usecase}
}

type createGroupRequestBody struct {
	Name string `json:"name"`
}

type addUserToGroupBody struct {
	GroupId int `json:"groupId"`
	UserId  int `json:"userId"`
}

func (g *GroupHandler) CreateGroupHandler(res http.ResponseWriter, req *http.Request) {
	grp, err := utils.BodyParser[createGroupRequestBody](req.Body)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, utils.MsgInvalidJSON)
		return
	}

	resgrp, err := g.usecase.CreateGroup(grp.Name)
	if err != nil {
		utils.WriteError(res, http.StatusInternalServerError, utils.MsgGroupCreationFailed)
		return
	}

	utils.WriteJSON(res, http.StatusCreated, utils.MsgGroupCreated, resgrp)
}

func (g *GroupHandler) AddUserToGroup(res http.ResponseWriter, req *http.Request) {
	userGroup, err := utils.BodyParser[addUserToGroupBody](req.Body)
	if err != nil {
		utils.WriteError(res, http.StatusBadRequest, utils.MsgInvalidJSON)
		return
	}

	if userGroup.GroupId <= 0 || userGroup.UserId <= 0 {
		utils.WriteError(res, http.StatusBadRequest, "GroupId and UserId must be valid positive integers")
		return
	}

	if err := g.usecase.AddUserToGroup(userGroup.UserId, userGroup.GroupId); err != nil {
		utils.WriteError(res, http.StatusInternalServerError, utils.MSGUserAddedToGroupFailed)
		return
	}

	utils.WriteJSON[any](res, http.StatusCreated, utils.MsgUserAddedToGroup, nil)
}
