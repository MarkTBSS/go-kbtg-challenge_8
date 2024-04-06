package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Storer interface {
	Wallets() ([]Wallet, error)
	WalletsByType(walletType string) ([]Wallet, error)
	WalletsByUserID(user_id string) ([]Wallet, error)
	CreateWallet(wallet Wallet) (Wallet, error)
	UpdateWallet(wallet Wallet) (Wallet, error)
	DeleteWallet(id string) (string, error)
}

type Handler struct {
	store Storer
}

func New(database Storer) *Handler {
	return &Handler{store: database}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
// @Summary		Get all wallets
// @Description	Get all wallets
// @Tags		wallet
// @Accept		json
// @Produce		json
// @Success		200	{object}	Wallet
// @Router		/api/v1/wallets [get]
// @Failure		500	{object}	Err
func (handler *Handler) WalletsHandler(context echo.Context) error {
	wallets, err := handler.store.Wallets()
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

func (handler *Handler) WalletsByTypeHandler(context echo.Context) error {
	walletType := context.QueryParam("wallet_type")
	wallets, err := handler.store.WalletsByType(walletType)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

// PathParamHandler
//
//	@Summary		Get all user_id
//	@Description	Get all user_id
//	@Tags			user_id
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/users/:id/wallets [get]
//	@Failure		500	{object}	Err
func (handler *Handler) WalletsByIDHandler(context echo.Context) error {
	id := context.Param("id")
	wallets, err := handler.store.WalletsByUserID(id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, wallets)
}

func (handler *Handler) CreateWalletHandler(context echo.Context) error {
	var wallet Wallet
	err := context.Bind(&wallet)
	if err != nil {
		return context.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	walletResponse, err := handler.store.CreateWallet(wallet)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, walletResponse)
}

func (handler *Handler) UpdateWalletHandler(context echo.Context) error {
	var wallet Wallet
	err := context.Bind(&wallet)
	if err != nil {
		return context.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	walletResponse, err := handler.store.UpdateWallet(wallet)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, walletResponse)
}

func (handler *Handler) DeleteHandler(context echo.Context) error {
	user_id := context.Param("id")
	deletedID, err := handler.store.DeleteWallet(user_id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return context.JSON(http.StatusOK, "ID : "+deletedID+" deleted")
}
