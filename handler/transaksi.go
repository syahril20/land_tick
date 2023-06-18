package handler

import (
	"fmt"
	resultdto "landtick/dto/result"
	transaksidto "landtick/dto/transaksi"
	"landtick/models"
	"landtick/repositories"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"gopkg.in/gomail.v2"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type HandlerTransaksis struct {
	TransaksiRepositories repositories.TransaksiRepositories
}

func HandlerTransaksi(TransaksiRepositories repositories.TransaksiRepositories) *HandlerTransaksis {
	return &HandlerTransaksis{TransaksiRepositories}
}

func (h *HandlerTransaksis) FindTransaksi(c echo.Context) error {
	transaksi, err := h.TransaksiRepositories.FindTransaksi()

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaksi})

}

func (h *HandlerTransaksis) GetUserId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := h.TransaksiRepositories.GetUserId(id)

	if user.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: user})
}

func (h *HandlerTransaksis) GetTransByUser(c echo.Context) error {
	// id, _ := strconv.Atoi(c.Param("id"))
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)
	transaction, err := h.TransaksiRepositories.GetTransByUser(int(userId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusOK,
			Message: "Waduh"})
	}
	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaction})

}

func (h *HandlerTransaksis) FindTransaksiId(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, _ := h.TransaksiRepositories.FindTransaksiId(id)

	if transaksi.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: "Data Gaada Bos"})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{
		Code: http.StatusOK,
		Data: transaksi})
}

func (h *HandlerTransaksis) DeleteTransaksi(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	transaksi, _ := h.TransaksiRepositories.FindTransaksiId(id)

	if transaksi.Id != id {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: "KOSONG OM"})
	}

	data, err := h.TransaksiRepositories.DeleteTransaksi(id, transaksi)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func (h *HandlerTransaksis) CreateTransaksi(c echo.Context) error {
	request := new(transaksidto.CreateTransaksi)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	request.IdUser = int(userId)
	request.Status = "pending" // next time we will cover this in payment gateway material

	user, _ := h.TransaksiRepositories.GetTransByUser(int(userId))
	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	fmt.Println(user, "MEMEK")

	var transaksiIsMatch = false
	var transaksiId int
	for !transaksiIsMatch {
		transaksiId = int(time.Now().Unix())
		transaksiData, _ := h.TransaksiRepositories.FindTransaksiId(transaksiId)
		if transaksiData.Id == 0 {
			transaksiIsMatch = true
		}
	}
	Tikets, err := h.TransaksiRepositories.GetTiketId(request.IdTiket)
	if err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}

	fmt.Println(userId)
	Transaksi := models.Transaksi{
		TanggalTransaksi: request.TanggalTransaksi,
		QtyDewasa:        request.QtyDewasa,
		QtyAnak:          request.QtyAnak,
		PulangPergi:      request.PulangPergi,
		Total:            request.Total,
		Status:           request.Status,
		IdUser:           request.IdUser,
		IdTiket:          request.IdTiket,
		Tiket:            Tikets,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	data, err := h.TransaksiRepositories.CreateTransaksi(Transaksi)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error()})
	}
	fmt.Println(data, "DATA INI CUKK")
	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)
	// Use to midtrans.Production if you want Production Environment (accept real transaksi).

	// 2. Initiate Snap request param
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(data.Id),
			GrossAmt: int64(data.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: data.User.NamaLengkap,
			Email: data.User.Email,
		},
	}

	// 3. Execute request create Snap transaksi to Midtrans Snap API
	snapResp, _ := s.CreateTransaction(req)

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: snapResp})

	// return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
}

func (h *HandlerTransaksis) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderId := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderId)
	transaction, _ := h.TransaksiRepositories.FindTransaksiId(order_id)

	fmt.Print("ini payloadnya", notificationPayload)

	if transactionStatus == "capture" {
		if fraudStatus == "challenge" {
			// TODO set transaksi status on your database to 'challenge'
			// e.g: 'Payment status challenged. Please take action on your Merchant Administration Portal
			h.TransaksiRepositories.UpdateTransaksi("pending", order_id)
		} else if fraudStatus == "accept" {
			// TODO set transaksi status on your database to 'success'
			SendMail("success", transaction)
			h.TransaksiRepositories.UpdateTransaksi("success", order_id)
		}
	} else if transactionStatus == "settlement" {
		// TODO set transaksi status on your databaase to 'success'
		SendMail("success", transaction)
		h.TransaksiRepositories.UpdateTransaksi("success", order_id)
	} else if transactionStatus == "deny" {
		// TODO you can ignore 'deny', because most of the time it allows payment retries
		// and later can become success
		h.TransaksiRepositories.UpdateTransaksi("failed", order_id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		// TODO set transaksi status on your databaase to 'failure'
		h.TransaksiRepositories.UpdateTransaksi("failed", order_id)
	} else if transactionStatus == "pending" {
		// TODO set transaksi status on your databaase to 'pending' / waiting payment
		h.TransaksiRepositories.UpdateTransaksi("pending", order_id)
	}

	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: notificationPayload})
}

func SendMail(status string, transaksi models.Transaksi) {

	if status != transaksi.Status && (status == "success") {
		var CONFIG_SMTP_HOST = "smtp.gmail.com"
		var CONFIG_SMTP_PORT = 587
		var CONFIG_SENDER_NAME = "DumbMerch <demo.dumbways@gmail.com>"
		var CONFIG_AUTH_EMAIL = os.Getenv("EMAIL_SYSTEM")
		var CONFIG_AUTH_PASSWORD = os.Getenv("PASSWORD_SYSTEM")

		var title = transaksi.Tiket.NamaKereta
		var total = strconv.Itoa(transaksi.Total)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", "rodyulo20@gmail.com")
		mailer.SetHeader("Subject", "Transaksi Status")
		mailer.SetBody("text/html", fmt.Sprintf(`<!DOCTYPE html>
	  <html lang="en">
		<head>
		<meta charset="UTF-8" />
		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
		<title>Document</title>
		<style>
		  h1 {
		  color: brown;
		  }
		</style>
		</head>
		<body>
		<h2>Product payment :</h2>
		<ul style="list-style-type:none;">
		  <li>Name : %s</li>
		  <li>Total payment: Rp.%s</li>
		  <li>Status : <b>%s</b></li>
		</ul>
		</body>
	  </html>`, title, total, status))

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Mail sent! to " + "rodyulo20@gmail.com")
	}
}

// func (h *HandlerTransaksis) CreateTransaksi(c echo.Context) error {
// 	request := new(transaksidto.CreateTransaksi)
// 	// Id, _ := strconv.Atoi(c.Param("id"))
// 	// transaksi, _ := h.TransaksiRepositories.FindTransaksiId(Id)

// 	if err := c.Bind(request); err != nil {
// 		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
// 			Code:    http.StatusInternalServerError,
// 			Message: err.Error()})
// 	}
// 	validation := validator.New()
// 	err := validation.Struct(request)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	Tikets, err := h.TransaksiRepositories.GetTiketId(request.IdTiket)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	transaksi := models.Transaksi{
// 		TanggalTransaksi: request.TanggalTransaksi,
// 		QtyDewasa:        request.QtyDewasa,
// 		QtyAnak:          request.QtyAnak,
// 		PulangPergi:      request.PulangPergi,
// 		Total:            request.Total,
// 		Status:           request.Status,
// 		IdUser:           request.IdUser,
// 		IdTiket:          request.IdTiket,
// 		Tiket:            Tikets,

// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}
// 	data, err := h.TransaksiRepositories.CreateTransaksi(transaksi)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
// }

// func (h *HandlerTransaksis) UpdateTransaksi(c echo.Context) error {
// 	request := new(transaksidto.UpdateTransaksi)
// 	if err := c.Bind(&request); err != nil {
// 		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	transaksi, err := h.TransaksiRepositories.FindTransaksiId(id)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	if request.TanggalTransaksi != "" {
// 		transaksi.TanggalTransaksi = request.TanggalTransaksi
// 	}
// 	if request.Qty != 0 {
// 		transaksi.QtyDewasa = request.Qty
// 	}
// 	if request.Total != 0 {
// 		transaksi.Total = request.Total
// 	}
// 	if request.Status != "" {
// 		transaksi.Status = request.Status
// 	}
// 	if request.IdUser != 0 {
// 		transaksi.IdUser = request.IdUser
// 	}
// 	if request.IdTiket != 0 {
// 		transaksi.IdTiket = request.IdTiket
// 	}

// 	transaksi.UpdatedAt = time.Now()

// 	data, err := h.TransaksiRepositories.UpdateTransaksi(id, transaksi)

// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, resultdto.ErrorResult{
// 			Code:    http.StatusBadRequest,
// 			Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, resultdto.SuccessResult{Code: http.StatusOK, Data: convertResponseTransaksi(data)})
// }

func convertResponseTransaksi(Transaksi models.Transaksi) transaksidto.TransaksiResponse {
	return transaksidto.TransaksiResponse{
		Id:               Transaksi.Id,
		TanggalTransaksi: Transaksi.TanggalTransaksi,
		QtyDewasa:        Transaksi.QtyDewasa,
		QtyAnak:          Transaksi.QtyAnak,
		PulangPergi:      Transaksi.PulangPergi,
		Total:            Transaksi.Total,
		Status:           Transaksi.Status,
		IdUser:           Transaksi.IdUser,
		User:             Transaksi.User,
		IdTiket:          Transaksi.IdTiket,
		Tiket:            Transaksi.Tiket,
	}
}
