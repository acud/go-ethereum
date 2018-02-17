// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package controllers

type controller interface {
	//Handle(w http.ResponseWriter, r *Request) error
	Name() string
	Method() string
	Resource() string
}

type Controller struct {
	name     string
	method   string
	resource string
}

//Gets the controller name
func (controller *Controller) Name() string {
	return controller.name
}

//gets the controller method
func (controller *Controller) Method() string {
	return controller.name
}

//gets the controller resource
func (controller *Controller) Resource() string {
	return controller.name
}
