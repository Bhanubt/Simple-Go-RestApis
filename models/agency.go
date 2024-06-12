package models

import (
	"pvmsproject/db"
)

//Creating Agency Struct
type Agency struct {
	AgencyId string `json:"agencyId,omitempty"`
	Name     string `json:"name,omitempty"`
	State    string
}

//Inserting Agency Into Agency Table
func (agency *Agency) Insert() (*Agency, error) {
	query := `INSERT INTO agencies (agency_id,name,state) VALUES (?,?,?)`

	//Preparing the Insert Query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	//Executing the Insert Query with Values
	_, err = stmt.Exec(agency.AgencyId, agency.Name, agency.State)
	if err != nil {
		return nil, err
	}

	return agency, nil
}

//Updating State Name of an Agency with Agency ID
func Update(agencyId string, agencyStateName string) (*Agency, error) {
	query := "UPDATE agencies SET state=? WHERE agency_id=?"

	//Preparing the Update with above query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	//Executing the UpdateQuery
	_, err = stmt.Exec(agencyStateName, agencyId)
	if err != nil {
		return nil, err
	}

	//Fetching the Agency with Agency ID
	agency, err := Get(&agencyId)
	if err != nil {
		return nil, err
	}
	return agency, nil

}

//Fetching the Agency with Agenct ID
func Get(agencyId *string) (*Agency, error) {
	query := "SELECT * FROM agencies WHERE agency_id=?"

	//Executing the above query with agencyId value
	row := db.DB.QueryRow(query, agencyId)
	var agency Agency

	//Scanning the values into the agency variable
	err := row.Scan(&agency.AgencyId, &agency.Name, &agency.State)
	if err != nil {
		return nil, err
	}

	return &agency, nil
}
