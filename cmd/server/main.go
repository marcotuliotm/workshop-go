package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"workshop-go/domain/user"
	"workshop-go/internal/server/http"
)

func main() {

	// //user := &domain.User{Name: "Tulio", Age: 33, Phone: []string{"988414831", "3352999999"}, Parentesco: make(map[string]interface{}, 3)
	// user := &domain.User{Name: "Tulio"} // , Age: 33, Phone: []string{"988414831", "3352999999"}, Parentesco: make(map[string]interface{}, 3)
	// // user.AddParentes("pai", "Jose")
	// // user.AddParentes("mae", "Arlete")
	// // user.AddParentes("irmaos", []string{"Romel", "Emmy"})

	// // for parentesco, value := range user.Parentesco {
	// // 	fmt.Printf("Grau de parentesco: [%s]: %v\n", parentesco, value)
	// // }

	// fmt.Printf("User: %s\n", user)
	// fmt.Printf("User: %v\n", user)
	// fmt.Printf("User: %+v\n", user)

	// json, _ := json.Marshal(user)
	// fmt.Println(string(json))

	// fmt.Printf("Refeito: %+v\n", user)

	// ===========================

	userService := user.NewService()

	handler := http.NewHandler(userService)

	server := http.New("8080", handler)
	server.ListenAndServe()
	stopChan := make(chan os.Signal)

	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	<-stopChan
	server.Shutdown()

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func getName(uf string) (name string, err error) {
	name, exists := estados(uf)
	if exists {
		return
	}
	return "", errors.New("Not found.")
}

func estados(uf string) (name string, exists bool) {
	return uf, true
}
