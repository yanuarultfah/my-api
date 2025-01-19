package users

import (
	"fmt"
	usersdb "my-api/datasource/pgsql/users_db"
	cryptoutils "my-api/utils/crypto"
	dateutils "my-api/utils/date_utils"
	"my-api/utils/errors"
)

const (
	queryInsertUser = `insert into userapi(first_name,last_name,email,date_created,status,password) values($1,$2,$3,$4,$5,$6) returning user_id;`
	queryGetUser    = `select user_id, first_name, last_name, email, date_created,status from userapi where user_id =$1;`
	errorNoRows     = "no rows in result set"
	// queryUpdateUser       = "update userapi set first_name=?, last_name=?, email=? where user_id=?;"
	// queryDeleteUser       = "delete from userapi where user_id=?;"
	queryFindUserByStatus = `select user_id,first_name,last_name,email,date_created,status from userapi where status = $1;`
)

func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	getErr := stmt.QueryRow(user.UserId).Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if getErr != nil {
		return errors.NewNotFoundError(errorNoRows)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	id := ""
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = dateutils.GetNowString()
	user.Status = "Active"
	saveErr := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, cryptoutils.GetMd5(user.Password)).Scan(&id)
	if saveErr != nil {
		return errors.NewBadRequestError(saveErr.Error())
	}
	user.UserId = id

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserId, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		results = append(results, user)
	}
	if len(status) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user matching status %s", status))
	}
	return results, nil
}
