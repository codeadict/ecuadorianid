[![Build Status](https://travis-ci.org/codeadict/ecuadorianid.svg?branch=master)](https://travis-ci.org/codeadict/ecuadorianid)

# What is Ecuadorian ID?
Go library used to validate Ecuadorian identification numbers(Cedula in Spanish).

## Usage.

To use the library from your Golang app just need to import it and use as follow:

  ```
  import (
      "github.com/codeadict/ecuadorianid"
  )

  cedula := string "YOURIDHERE"

  if valid, error := Validate(cedula); error != nil {
    t.Errorf("Error with identification number '%v', response '%+v', error: %+v", cedula, valid, error)
  }
  ```

##Authors:

  * Dairon Medina C. (dairon.medina@gmail.com)
