//
// Code generated by go-jet DO NOT EDIT.
// Generated at Saturday, 14-Mar-20 18:27:21 KST
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type Users struct {
	UUID     string `sql:"primary_key"`
	Email    string
	Name     *string
	SigninDt *int32
}
