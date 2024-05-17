package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	`github.com/rhmdnrhuda/simple-rest-api/models`
	`github.com/rhmdnrhuda/simple-rest-api/repositories`

	"go.mongodb.org/mongo-driver/mongo"
)

type HotelHandler struct {
	repo *repositories.HotelRepository
}

func NewHotelHandler(db *mongo.Database) *HotelHandler {
	return &HotelHandler{
		repo: repositories.NewHotelRepository(db),
	}
}

func (h *HotelHandler) GetHotels(c echo.Context) error {
	hotels, err := h.repo.GetHotels()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Data    interface{}
		Message string
	}{
		Data:    hotels,
		Message: "Hotels found",
	})
}

func (h *HotelHandler) CreateHotel(c echo.Context) error {
	var hotel models.Hotel
	if err := c.Bind(&hotel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	newHotel, err := h.repo.CreateHotel(hotel)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, newHotel)
}

func (h *HotelHandler) GetHotelById(c echo.Context) error {
	id := c.Param("id")

	hotel, err := h.repo.GetHotelById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	if hotel == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Hotel not found"})
	}

	return c.JSON(http.StatusOK, hotel)
}

func (h *HotelHandler) UpdateHotel(c echo.Context) error {
	id := c.Param("id")
	var hotel models.Hotel
	if err := c.Bind(&hotel); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.repo.UpdateHotel(id, hotel); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Hotel updated successfully"})
}

func (h *HotelHandler) DeleteHotel(c echo.Context) error {
	id := c.Param("id")

	if err := h.repo.DeleteHotel(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.NoContent(http.StatusNoContent)
}
