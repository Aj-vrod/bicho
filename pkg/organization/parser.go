package organization

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/gosimple/slug"
)

const (
	EMPLOYEES_FILEPATH = "employees.json"
)

// ReadOrgData reads data from a specific file and returns as []Employee
func ReadOrgData(path string) ([]Employee, error) {
	data, err := os.ReadFile(EMPLOYEES_FILEPATH)
	if err != nil {
		return nil, err
	}

	var employees []Employee
	err = json.Unmarshal(data, &employees)
	if err != nil {
		return nil, err
	}

	return employees, nil
}

// ProcessOrgData takes raw data and returns Organization structure
func ProcessOrgData(employees []Employee) (Organization, error) {
	org := buildOrgByName(employees)

	return org, nil

}

// buildOrgByName takes data and translates into tree structure
func buildOrgByName(employees []Employee) Organization {
	org := Organization{}

	for _, e := range employees {
		batallionIdx := slices.IndexFunc(org.Batallions, func(b Batallion) bool { return b.Name == e.Battalion })
		if batallionIdx == -1 {
			org.Batallions = append(org.Batallions, Batallion{
				ID:   fmt.Sprintf("/%s", slug.Make(e.Battalion)),
				Name: e.Battalion,
				Platoons: []Platoon{{
					ID:   fmt.Sprintf("/%s/%s", slug.Make(e.Battalion), slug.Make(e.Platoon)),
					Name: e.Platoon,
					Squads: []Squad{{
						ID:   fmt.Sprintf("/%s/%s/%s", slug.Make(e.Battalion), slug.Make(e.Platoon), slug.Make(e.Squad)),
						Name: e.Squad,
						Employees: []Member{{
							Name:         e.Name,
							Country:      e.Country,
							JobFamily:    e.JobFamily,
							JobTitle:     e.JobTitle,
							BusinessUnit: e.BusinessUnit,
							StartDate:    e.StartDate,
						}},
					}},
				}},
			})
		} else {
			platoonIdx := slices.IndexFunc(org.Batallions[batallionIdx].Platoons, func(p Platoon) bool { return p.Name == e.Platoon })
			if platoonIdx == -1 {
				org.Batallions[batallionIdx].Platoons = append(org.Batallions[batallionIdx].Platoons, Platoon{
					ID:   fmt.Sprintf("/%s/%s", slug.Make(e.Battalion), slug.Make(e.Platoon)),
					Name: e.Platoon,
					Squads: []Squad{{
						ID:   fmt.Sprintf("/%s/%s/%s", slug.Make(e.Battalion), slug.Make(e.Platoon), slug.Make(e.Squad)),
						Name: e.Squad,
						Employees: []Member{{
							Name:         e.Name,
							Country:      e.Country,
							JobFamily:    e.JobFamily,
							JobTitle:     e.JobTitle,
							BusinessUnit: e.BusinessUnit,
							StartDate:    e.StartDate,
						}},
					}},
				})
			} else {
				squadIdx := slices.IndexFunc(org.Batallions[batallionIdx].Platoons[platoonIdx].Squads, func(s Squad) bool { return s.Name == e.Squad })
				if squadIdx == -1 {
					org.Batallions[batallionIdx].Platoons[platoonIdx].Squads = append(org.Batallions[batallionIdx].Platoons[platoonIdx].Squads, Squad{
						ID:   fmt.Sprintf("/%s/%s/%s", slug.Make(e.Battalion), slug.Make(e.Platoon), slug.Make(e.Squad)),
						Name: e.Squad,
						Employees: []Member{{
							Name:         e.Name,
							Country:      e.Country,
							JobFamily:    e.JobFamily,
							JobTitle:     e.JobTitle,
							BusinessUnit: e.BusinessUnit,
							StartDate:    e.StartDate,
						}},
					})
				} else {
					org.Batallions[batallionIdx].Platoons[platoonIdx].Squads[squadIdx].Employees = append(org.Batallions[batallionIdx].Platoons[platoonIdx].Squads[squadIdx].Employees, Member{
						Name:         e.Name,
						Country:      e.Country,
						JobFamily:    e.JobFamily,
						JobTitle:     e.JobTitle,
						BusinessUnit: e.BusinessUnit,
						StartDate:    e.StartDate,
					})
				}
			}
		}
	}

	return org
}

// validateEmployee validates an org entry to ensure it has name, squad, tribe and batallion
func validateEmployee(employee Employee) bool {
	// Must have a name
	if employee.Name == "" {
		return false
	}

	// Must have a squad
	if employee.Squad == "" {
		return false
	}

	// Must have a platoon
	if employee.Platoon == "" {
		return false
	}

	// Must have a batallion
	if employee.Battalion == "" {
		return false
	}

	return true
}
