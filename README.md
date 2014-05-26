lightning
=========

Speedy spark core library written in Go (aka Golang)

```Go
c, err := lightning.Collection{"12345mykey12345"}
//Handle err

myCore := c["mycorename"]

variable, err := myCore.Var("variableName")
//Handle err

fnResult, err := myCore.Fn("functionname", "functionParam") //(Any type for param. will be run through ftm.Sprint())
//Handle err

eventChan, err := myCore.Evt("eventName")
//Handle err

for {
  select {
  case val := <- eventChan:
    //Do something
  }
}

```
