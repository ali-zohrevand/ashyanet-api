package services

import (
	"gitlab.com/hooshyar/ChiChiNi-API/models"
	"net/http"
)

func Command(command models.Command, User models.UserInDB) (int, []byte) {

	return http.StatusInternalServerError, []byte("")

}
