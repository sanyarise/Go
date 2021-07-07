package web

import (
	"fmt"
	"net/http"
	"strconv"
	"github.com/qencept/fibonacci/cache"
	"github.com/qencept/fibonacci/closure"
	"github.com/qencept/fibonacci/recursion"
	"math/rand"
	"time"
)

func Serve() {
	
		http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
			rand.Seed(time.Now().UnixNano())
			fibImplementations := []Fiber{recursion.NewRecursion(), cache.NewCache(), closure.NewClosure()}
			fiber := fibImplementations[rand.Intn(len(fibImplementations))]
			nSlice, ok := request.URL.Query()["n"] //http://localhost:8080/?n=10&n=20

		data := fmt.Sprintf("%s\n", fiber.Name())
		if ok {
			for _, nValue := range nSlice {
				n, err := strconv.Atoi(nValue)
				if err == nil {
					data += fmt.Sprintf("Fib(%d) = %d\n", n, fiber.Fib(n))
				}
			}
		}

		_, err := writer.Write([]byte(data))
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

