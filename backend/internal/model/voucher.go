package model

type Voucher struct {
    ID            int64  `json:"id"`
    CrewName      string `json:"crewName"`
    CrewID        string `json:"crewId"`
    FlightNumber  string `json:"flightNumber"`
    FlightDate    string `json:"flightDate"`
    AircraftType  string `json:"aircraftType"`
    Seat1         string `json:"seat1"`
    Seat2         string `json:"seat2"`
    Seat3         string `json:"seat3"`
    CreatedAt     string `json:"createdAt"`
}
