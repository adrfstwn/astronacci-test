package dto

type CheckRequest struct {
    FlightNumber string `json:"flightNumber" binding:"required"`
    Date         string `json:"date" binding:"required"`
}

type CheckResponse struct {
    Exists bool `json:"exists"`
}

type GenerateRequest struct {
    Name         string `json:"name" binding:"required"`
    ID           string `json:"id" binding:"required"`
    FlightNumber string `json:"flightNumber" binding:"required"`
    Date         string `json:"date" binding:"required"`
    Aircraft     string `json:"aircraft" binding:"required"`
}

type GenerateResponse struct {
    Success bool     `json:"success"`
    Seats   []string `json:"seats"`
}
