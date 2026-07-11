package config

type SeatLayout struct {
    Row  int
    Columns []string
}

var SeatLayouts = map[string]SeatLayout{
    "ATR": {
        Row:  18,
        Columns: []string{"A", "C", "D", "F"},
    },
    "Airbus 320": {
        Row:  32,
        Columns: []string{"A", "B", "C", "D", "E", "F"},
    },
    "Boeing 737 Max": {
        Row:  32,
        Columns: []string{"A", "B", "C", "D", "E", "F"},
    },
}