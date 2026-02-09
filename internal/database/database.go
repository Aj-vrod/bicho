package database

import (
	"Aj-vrod/bicho/pkg/config"
	"Aj-vrod/bicho/pkg/organization"
	"database/sql"
	"log"
	"reflect"
)

func connectWithDB() (*sql.DB, error) {
	// Connect with DB
	cfg, err := config.LoadFromEnv()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", cfg.DBDNS)
	if err != nil {
		return nil, err
	}

	// Check connection
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Successfully connected to db")

	return db, nil
}

func SyncOrgWithDB(filePath string) error {
	db, err := connectWithDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Get list of employees from file
	employees, err := organization.ReadOrgData(filePath)
	if err != nil {
		return err
	}

	// Insert employees into db
	for _, e := range employees {
		// Prepare SELECT query
		selectQuery, err := db.Prepare("SELECT name, country, job_family, job_title, business_unit, squad, platoon, battalion, start_date FROM employees WHERE name = $1 AND start_date = $2 AND country = $3;")
		if err != nil {
			return err
		}
		defer selectQuery.Close()

		// Check if employee already in db
		var oldE organization.Employee
		err = selectQuery.QueryRow(e.Name, e.StartDate, e.Country).Scan(&oldE)
		if err == sql.ErrNoRows {
			// Prepare INSERT query
			insertQuery, err := db.Prepare("INSERT INTO employees(name, country, job_family, job_title, business_unit, squad, platoon, battalion, start_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);")
			if err != nil {
				return err
			}
			defer insertQuery.Close()

			result, err := insertQuery.Exec(e.Name, e.Country, e.JobFamily, e.JobTitle, e.BusinessUnit, e.Squad, e.Platoon, e.Battalion, e.StartDate)
			if err != nil {
				return err
			}

			_, err = result.RowsAffected()
			if err != nil {
				return err
			}

			log.Printf("Employee %s successfully added.\n", e.Name)
		} else {
			// Check if there are updates to be made
			wasUpdated := reflect.DeepEqual(e, oldE)
			if !wasUpdated {
				log.Printf("Employee %s needed no update.\n", e.Name)
				continue
			}

			// Prepare UPDATE query
			updateQuery, err := db.Prepare("UPDATE employees e SET(name, country, job_family, job_title, business_unit, squad, platoon, battalion, start_date) = ($1,$2,$3,$4,$5,$6,$7,$8,$9)	WHERE	e.name = $10	AND e.start_date = $11	AND e.country = $12;")
			if err != nil {
				return err
			}
			defer updateQuery.Close()

			result, err := updateQuery.Exec(e.Name, e.Country, e.JobFamily, e.JobTitle, e.BusinessUnit, e.Squad, e.Platoon, e.Battalion, e.StartDate, oldE.Name, oldE.StartDate, oldE.Country)
			if err != nil {
				return err
			}

			_, err = result.RowsAffected()
			if err != nil {
				return err
			}

			log.Printf("Employee %s successfully updated.\n", e.Name)
		}

	}
	return nil
}

func GetEmployees() ([]organization.Employee, error) {
	db, err := connectWithDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	query, err := db.Prepare("SELECT name, country, job_family, job_title, business_unit, squad, platoon, battalion, start_date FROM employees;")
	if err != nil {
		return nil, err
	}

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []organization.Employee
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var emp organization.Employee
		err := rows.Scan(
			&emp.Name,
			&emp.Country,
			&emp.JobFamily,
			&emp.JobTitle,
			&emp.BusinessUnit,
			&emp.Squad,
			&emp.Platoon,
			&emp.Battalion,
			&emp.StartDate,
		)
		if err != nil {
			log.Printf("Failed to scan emplyee: %s", emp.Name)
			continue
		}
		employees = append(employees, emp)
	}
	if err := rows.Err(); err != nil {
		return []organization.Employee{}, err
	}

	return employees, nil
}
