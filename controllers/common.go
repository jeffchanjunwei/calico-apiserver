package controllers

import (
        "net/http"
        "io/ioutil"
        "encoding/json"
)

// BodyToObject will convert the incoming HTTP request into the
// passed in 'object'
func BodyToObject(r *http.Request, object interface{}) error {
        body, err := ioutil.ReadAll(r.Body)
        if err != nil {
                return err
        }

        err = json.Unmarshal(body, object)
        if err != nil {
                return err
        }
        return nil
}

