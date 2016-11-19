# jsonconfig
A light-weight JSON config file manager implemented by golang.

## Sample
```
package main

import (
	"fmt"
	"github.com/moakap/jsonconfig"
	"time"
)

func main() {
	err := jsonconfig.InitConf("conf.json")
	if err != nil {
		fmt.Println("jsonconfig initialize failed: ", err.Error())
		return
	}

	jsonconfig.Set("param1", 100)
	jsonconfig.Set("param2", "stringValue")
	jsonconfig.Set("param3", time.Now())

	param1 := jsonconfig.Get("param1")
	fmt.Println("param1 = ", param1)

	param2 := jsonconfig.Get("param2")
	fmt.Println("param2 = ", param2)

	param3 := jsonconfig.Get("param3")
	fmt.Println("param3 = ", param3)

	return
}

```
output
> param1 =  100
> param2 =  stringValue
> param3 =  2016-11-19 15:34:32.19989814 +0800 CST
