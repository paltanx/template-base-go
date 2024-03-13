package handlers

import (
	"encoding/json"
	"get-otp-go/src/repositories"
	"get-otp-go/src/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type IOtpHandler interface {
	Do(w http.ResponseWriter, r *http.Request)
}

type OtpHandler struct {
	logger            utils.ILogger
	otpRepository repositories.IOtpRepository
}

// Container
func NewOtpHandler(
	logger utils.ILogger,
	otpRepository repositories.IOtpRepository,
) *OtpHandler {
	return &OtpHandler{
		logger,
		otpRepository,
	}
}

func (h *OtpHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Calling Get OTP handler")

	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.otpRepository.Get(id)

	if err != nil {
		h.logger.Error(err.Error())
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Failed to encode response", map[string]interface{}{"error": err})
	}
}


