package api

import (
    "encoding/json"
)

func GetQueryParamsFromOptions(data interface{}) map[string]interface{} {
    content, err := json.Marshal(data)
    if err != nil {
        return nil
    }
    var queryParams map[string]interface{}
    err = json.Unmarshal(content, &queryParams)
    if err != nil {
        return nil
    }
    return queryParams
}
