package organization

// Top level of organization tree view
type Organization struct {
	Batallions []Batallion `json:"batallions"`
}

// Batallion entry
type Batallion struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Platoons []Platoon `json:"platoons"`
}

// Platoon entry
type Platoon struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Squads []Squad `json:"squads"`
}

// Squad entry
type Squad struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Employees []Member `json:"member"`
}

// Member entry
type Member struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	JobFamily    string `json:"job_family"`
	JobTitle     string `json:"job_title"`
	BusinessUnit string `json:"business_unit"`
	StartDate    string `json:"start_date"`
}

// Employee entry, struct from data
type Employee struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	JobFamily    string `json:"job_family"`
	JobTitle     string `json:"job_title"`
	BusinessUnit string `json:"business_unit"`
	Squad        string `json:"squad"`
	Platoon      string `json:"platoon"`
	Battalion    string `json:"battalion"`
	StartDate    string `json:"start_date"`
}
