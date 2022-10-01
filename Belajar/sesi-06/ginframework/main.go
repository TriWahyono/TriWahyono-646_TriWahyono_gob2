package main

import routers "ginframework/routers"

func main() {

	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
