package service

import (
    "errors"
    "fmt"
    "math/rand"
    "strings"
    "time"

    "backend/internal/config"
    "backend/internal/dto"
    "backend/internal/model"
    "backend/internal/repository"
)

var ErrVoucherExists = errors.New("vouchers already generated for this flight and date")
var ErrInvalidAircraft = errors.New("invalid aircraft type")

type VoucherService struct {
    repo *repository.VoucherRepository
}

func NewVoucherService(repo *repository.VoucherRepository) *VoucherService {
    return &VoucherService{repo: repo}
}

func (s *VoucherService) CheckExists(flightNumber, date string) (bool, error) {
    return s.repo.Exists(flightNumber, date)
}

func (s *VoucherService) GenerateVoucher(req dto.GenerateRequest) ([]string, error) {
    exists, err := s.repo.Exists(req.FlightNumber, req.Date)
    if err != nil {
        return nil, err
    }
    if exists {
        return nil, ErrVoucherExists
    }

    seats, err := generateSeats(req.Aircraft)
    if err != nil {
        return nil, err
    }

    voucher := model.Voucher{
        CrewName:     req.Name,
        CrewID:       req.ID,
        FlightNumber: req.FlightNumber,
        FlightDate:   req.Date,
        AircraftType: req.Aircraft,
        Seat1:        seats[0],
        Seat2:        seats[1],
        Seat3:        seats[2],
        CreatedAt:    time.Now().Format(time.RFC3339),
    }

    if err := s.repo.Insert(voucher); err != nil {
        if strings.Contains(err.Error(), "UNIQUE constraint failed") {
            return nil, ErrVoucherExists
        }
        return nil, err
    }

    return seats, nil
}

func generateSeats(aircraft string) ([]string, error) {
    layout, ok := config.SeatLayouts[aircraft]
    if !ok {
        return nil, ErrInvalidAircraft
    }

    var allSeats []string
    for row := 1; row <= layout.Row; row++ {
        for _, col := range layout.Columns {
            allSeats = append(allSeats, fmt.Sprintf("%d%s", row, col))
        }
    }

    rand.Shuffle(len(allSeats), func(i, j int) {
        allSeats[i], allSeats[j] = allSeats[j], allSeats[i]
    })

    return allSeats[:3], nil
}