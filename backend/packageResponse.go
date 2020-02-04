package backend

import "encoding/json"

func packageResponse(response interface{}, e error) []byte {
	ret := make(map[string]interface{})
	if response != nil {
		ret["response"] = response
	} else {
		ret["response"] = ""
	}
	if e != nil {
		ret["error"] = e.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Println("Problem marshaling return data", err)
	}

	return data
}
