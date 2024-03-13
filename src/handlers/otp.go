package handlers

import (
	"encoding/json"
	"get-otp-go/src/models"
	"get-otp-go/src/repositories"
	"get-otp-go/src/utils"
	"net/http"
)

type OtpHandler struct {
	logger        utils.ILogger
	otpRepository repositories.IOtpRepository
}

func NewOtpHandler(logger utils.ILogger, otpRepository repositories.IOtpRepository) *OtpHandler {
	return &OtpHandler{
		logger:        logger,
		otpRepository: otpRepository,
	}
}

func (h *OtpHandler) Post(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Calling Post OTP handler")

	var otp models.Otp
	if err := json.NewDecoder(r.Body).Decode(&otp); err != nil {
		h.logger.Error("Failed to decode request body", map[string]interface{}{"error": err})
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.otpRepository.Post(&otp)
	if err != nil {
		h.logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("Failed to encode response", map[string]interface{}{"error": err})
	}
}
