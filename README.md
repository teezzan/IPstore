Package IPStore
=================

Reliably log, store and count ingress IP addresses.
------------

![Logo](./images/IPstore.png)

**IPStore** is a library for keeping track of IP addresses that make request to your service. 


[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/teezzan/cdenv/blob/master/LICENSE.md)


Features
------------

- Light-Weight
- Fast Insert and Fetch Operations
- Validation of inputs
- Modular and extendable for any kind of storage

Example Code
------------

Using this library is as simple as running the following 
```golang
package main

import 	"github.com/teezzan/IPStore"

func main() {
    // Stores the IP Address 192.34.56.321 in the memory storage
	IPStore.RequestHandled("192.34.56.321") 
    
    // Increments the IP Address 192.54.56.333 count by 1.
	IPStore.RequestHandled("192.34.56.321") 

    // Fetch the top 100 frequent IP addresses
    top100 := IPStore.Top100()

    // Removes all stored IP addresses.
    IPStore.Clear()
}

```

Installation
------------

Use go get.

	go get github.com/teezzan/IPStore

Then import  IPstore library into your own code.

	import "github.com/teezzan/IPStore"


Architecture
-----------
![Logo](./images/arch.jpg)

The modular architecture allows for implementation of the same system using another storage.

Contributing
------------
Issues and pull requests are welcome at [IPStore](https://github.com/teezzan/IPStore). This project is intended to be safe, welcoming, and open for collaboration. Users are expected to adhere to the [Contributor Covenant code of conduct](https://www.contributor-covenant.org/version/2/0/code_of_conduct/). We are all human.

## Authors

**[Taiwo Yusuf](https://github.com/teezzan/)**

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

