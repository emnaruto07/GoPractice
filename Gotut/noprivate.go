package main

// func idiot(x int) int { // idiot, identity
// 	return x
// }

type DB struct {
	engine string
	url    string
}

func PGConnectDB(connectionUrl string) *DB { // type: func(string) *DB
	url := connectionUrl
	return &DB{"postgresql", url}
}

func MSConnectDB(connectionUrl string) *DB {
	url := connectionUrl
	return &DB{"mysql", url}
}

func LiteConnectDB(connectionUrl string) *DB {
	url := connectionUrl
	return &DB{"sqlite", url}
}

func StoreInDB(input string, connector func(string) *DB) {
	db := connector(userinput())
	db.Store(input)
}

func (r Router) GET(endpoint string, handler func(*Context))

func main() {
	// https://github.com/keogami/dehasher/blob/53089a1e1e3c54486fb37d3e4dd15a13a50a0c3d/crack.go#L20
	var k int = 8
	var myfunc (func(string) *DB) = LiteConnectDB
	StoreInDB("data", MSConnectDB)

	// funcvar := idiot

	// k = funcvar(k)

	// funcvar == idiot

}

//dummy
//security
