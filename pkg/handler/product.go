package handler

import (
	"fmt"
	market "market"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHub() *Hub {
	return &Hub{
		ProductsHub: make(map[string]*market.Product),
	}
}

type CreateProductRequest struct {
	Id    string  `json:"id"`
	Image string  `json:"image"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *Handler) CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.hub.ProductsHub[req.Id] = &market.Product{
		Id:    req.Id,
		Image: req.Image,
		Name:  req.Name,
		Price: req.Price,
	}
	c.JSON(http.StatusOK, gin.H{"response": req})

}

type Hub struct {
	ProductsHub map[string]*market.Product `json:"productsHub"`
}

type ProductRes struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *Handler) GetProducts(c *gin.Context) {
	rooms := make([]ProductRes, 0)

	for _, r := range h.hub.ProductsHub {
		rooms = append(rooms, ProductRes{
			Id:    r.Id,
			Price: r.Price,
			Name:  r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *Handler) BuyProduct(c *gin.Context) {
	id := c.Param("prodId")
	count := c.Param("count")
	if _, ok := h.hub.ProductsHub[id]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": "product bought"})
	fmt.Println("got a new order!", "id", id, "items", count)
}
