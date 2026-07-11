package repository

import (
	"backend/internal/model"
	"database/sql"
)

type VoucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) *VoucherRepository {
	return &VoucherRepository{db: db}
}

func (r *VoucherRepository) Exists(flightNumber, date string) (bool, error) {
    var id int64
    err := r.db.QueryRow(
        `SELECT id FROM vouchers WHERE flight_number = ? AND flight_date = ?`,
        flightNumber, date,
    ).Scan(&id)

    if err == sql.ErrNoRows {
        return false, nil
    }
    if err != nil {
        return false, err
    }
    return true, nil
}

func (r *VoucherRepository) Insert(v model.Voucher) error {
    _, err := r.db.Exec(
        `INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
        v.CrewName, v.CrewID, v.FlightNumber, v.FlightDate, v.AircraftType, v.Seat1, v.Seat2, v.Seat3, v.CreatedAt,
    )
    return err
}
