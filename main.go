package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var reader *bufio.Reader

type User struct {
	id       int
	username string
	email    string
	age      int
}

var users map[int]User

var id int

func crearUsuario() {
	clearConsole()
	fmt.Print("Ingresa un valor para username: ")
	username := readLine()
	fmt.Print("Ingresa un valor para email: ")
	email := readLine()
	fmt.Print("Ingresa un valor para edad: ")
	age, err := strconv.Atoi(readLine())

	if err != nil {
		panic("No se puede man")
	}

	id++
	user := User{id, username, email, age}
	users[id] = user
	fmt.Println("Usuario creado exitosamente")
}

func listarUsuario() {
	clearConsole()

	for id, user := range users {
		fmt.Println(id, "-", user.username)
	}

}
func actualizarUsuario() {
	fmt.Print("Ingresa ID del usuario a actualizar: ")

	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("NO se puede man")
	}

	if _, ok := users[id]; ok {
		fmt.Print("Ingresa un valor para username: ")
		username := readLine()
		fmt.Print("Ingresa un valor para email: ")
		email := readLine()
		fmt.Print("Ingresa un valor para edad: ")
		age, err := strconv.Atoi(readLine())
		if err != nil {
			panic("No se puede man")
		}
		users[id] = User{id, username, email, age}
	}
}

func eliminarUsuario() {
	clearConsole()
	fmt.Print("Ingresa ID del usuario a eliminar: ")
	id, err := strconv.Atoi(readLine())
	if err != nil {
		panic("NO se puede man")
	}

	delete(users, id)

	fmt.Println("Eliminado")
}

func readLine() string {
	if option, err := reader.ReadString('\n'); err != nil {
		panic("No es posible obtener el valor")
	} else {
		return strings.TrimSuffix(option, "\n")
	}
}

func clearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	users = make(map[int]User)

	var option string

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("A- Crear")
		fmt.Println("B- Listar")
		fmt.Println("C- Actualizar")
		fmt.Println("D- ELminar")

		fmt.Print("Ingresa una opcion: ")
		option = readLine()
		if option == "quit" || option == "q" {
			break
		}

		switch option {
		case "a", "crear":
			crearUsuario()
		case "b", "listar":
			listarUsuario()
		case "c", "actualizar":
			actualizarUsuario()
		case "d", "eliminar":
			eliminarUsuario()
		default:
			fmt.Println("Opcion incorrecta")
		}
	}

	fmt.Println("nv")
}
