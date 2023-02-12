package api

import "encoding/json"

func makeOkResult() ([]byte, error) {
	data := make(map[string]string)
	data["result"] = "OK"
	dataToRet, err := json.Marshal(data)
	return dataToRet, err
}

func makeBadDataResult() ([]byte, error) {
	data := make(map[string]string)
	data["result"] = "FAIL READ DATA"
	dataToRet, err := json.Marshal(data)
	return dataToRet, err
}

func makeFailSplitResult() ([]byte, error) {
	data := make(map[string]string)
	data["result"] = "FAIL SPLIT DATA"
	dataToRet, err := json.Marshal(data)
	return dataToRet, err
}

func makeFailInsertResult() ([]byte, error) {
	data := make(map[string]string)
	data["result"] = "FAIL INSERT DATA"
	dataToRet, err := json.Marshal(data)
	return dataToRet, err
}
