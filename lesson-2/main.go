package main

import "github.com/MaxFando/gb-advance/lesson-2/config"

func main() {
	conf := config.Configuration{}
	if err := conf.NewConnection(); err != nil {
		panic("something wrong happened")
	}

	defer func() {
		conf.CloseConnection()
	}()
}
