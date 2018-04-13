package main

import (
        "github.com/TIBCOSoftware/flogo-lib/core/activity"
        //"encoding/json"
        "fmt"
        //"log"
        //"strings"
        //"io/ioutil"
)




// MyActivity is a stub for your Activity implementation
type MyActivity struct {
        metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
        return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
        return a.metadata
}

func main() {
fmt.Println("hello world!")
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

        // do eval



        return true, nil
}
