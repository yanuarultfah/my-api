package thirdparty

import (
	"encoding/json"
	"io/ioutil"
	"my-api/utils/errors"
	"net/http"
)

func (listuser *UserListReqRest) Get() *errors.RestErr {
	stmtgetlistuser, err := http.Get("https://reqres.in/api/users?page=2")
	if err != nil {
		return errors.NewBadRequestError(err.Error())
	}
	defer stmtgetlistuser.Body.Close()
	bodylistuser, err := ioutil.ReadAll(stmtgetlistuser.Body)
	// fmt.Println(bodylistuser)
	if err := json.Unmarshal(bodylistuser, &listuser); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	return nil
}
