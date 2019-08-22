// Copyright (c) 2019 Sick Yoon
// This file is part of gocelery which is released under MIT license.
// See file LICENSE for full license details.

package main

import (
	"log"
	"math/rand"
	"reflect"
	"time"

	"github.com/gocelery/gocelery"
)

// Run Celery Worker First!
// celery -A worker worker --loglevel=debug --without-heartbeat --without-mingle

func main() {

	// initialize celery client
	client, err := gocelery.NewCeleryClient(
		gocelery.NewRedisCeleryBroker("redis://"),
		gocelery.NewRedisCeleryBackend("redis://"),
		1,
	)
	if err != nil {
		panic(err)
	}

	// run task
	asyncResult, err := client.Delay("worker.add", 10, 20)
	if err != nil {
		panic(err)
	}
	asyncResultKwargs, err := client.DelayKwargs("worker.add_reflect", map[string]interface{}{
		"a": rand.Intn(10),
		"b": rand.Intn(10),
	})
	if err != nil {
		panic(err)
	}

	// get results from backend with timeout
	res, err := asyncResult.Get(10 * time.Second)
	if err != nil {
		panic(err)
	}
	resKwargs, err := asyncResultKwargs.Get(10 * time.Second)
	if err != nil {
		panic(err)
	}

	log.Printf("result: %+v of type %+v", res, reflect.TypeOf(res))
	log.Printf("result: %+v of type %+v", resKwargs, reflect.TypeOf(resKwargs))
}
