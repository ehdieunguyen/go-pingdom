package pingdom

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// TransactionCheckService provides an interface to Pingdom transaction checks.
type TransactionCheckService struct {
	client *Client
}

type TransactionCheck interface {
	PutParams() map[string]string
	PostParams() map[string]string
	Valid() error
}

// List returns a list of transaction checks from Pingdom.
// This returns type TransactionCheckResponse rather than TransactionCheck Check since the
// Pingdom API does not return a complete representation of a transaction check.
func (ts *TransactionCheckService) List(params ...map[string]string) ([]TransactionCheckResponse, error) {
	param := map[string]string{}
	if len(params) == 1 {
		param = params[0]
	}
	req, err := ts.client.NewRequest("GET", "/tms/check?extended_tags=true", param)
	if err != nil {
		return nil, err
	}

	resp, err := ts.client.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := validateResponse(resp); err != nil {
		return nil, err
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	bodyString := string(bodyBytes)
	m := &listTransactionJSONResponse{}
	err = json.Unmarshal([]byte(bodyString), &m)

	return m.Checks, err
}

// Read returns detailed information about a pingdom transaction check given its ID.
// This returns type TransactionCheckResponse rather than TransactionCheck Check since the
// pingdom API does not return a complete representation of a transaction check.
func (ts *TransactionCheckService) Read(id int) (*TransactionCheckResponse, error) {
	req, err := ts.client.NewRequest("GET", "/tms/check/"+strconv.Itoa(id)+"?extended_tags=true", nil)
	if err != nil {
		return nil, err
	}

	m := &transactionCheckDetailsJSONResponse{}
	_, err = ts.client.Do(req, m)
	if err != nil {
		return nil, err
	}

	return m.Check, err
}
