# namespace 

[![GoDoc](https://godoc.org/github.com/cogger/namespace?status.png)](http://godoc.org/github.com/cogger/namespace)  
[![Build Status](https://travis-ci.org/cogger/namespace.svg?branch=master)](https://travis-ci.org/cogger/namespace)  
[![Coverage Status](https://coveralls.io/repos/cogger/namespace/badge.svg?branch=master)](https://coveralls.io/r/cogger/namespace?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

namespace addes the ability set a handlers namespace.  This will allow you to set if the function is in some special state such as development, test or production that changes to way certain functions preform.

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"github.com/cogger/cogger"
	"github.com/cogger/namespace"
	"golang.org/x/net/context"
)

func foo(ctx context.Context, w http.ResponseWriter, r *http.Request) int{

	//Do work

	if !namespace.Is(ctx, "dev") {
		//Save to DB
	}

	return http.StatusOK
}




func init() {
	fooHandler := cogger.NewHandler()
	fooHandler.AddContext(namepsace.Add("dev"))

	fooHandler.SetHandler(foo)

  	http.Handle("/foo", fooHandler)
}

~~~

