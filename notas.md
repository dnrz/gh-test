package main

/*
	go build main.go -> Un archivo ejecutable
	go run main.go -> Ejecutarlo sin compilar un archivo ejecutable
*/

/*const (
	Domingo int = iota + 1000 // 0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)*/

// func division(dividendo, divisor int) (int, error) {

// 	if divisor == 0 {
// 		return 0, errors.New("No es posible dividir por cero")
// 	} else {
// 		return dividendo / divisor, nil
// 	}
// }

func main() {
	/*fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)*/

	// var curso string = "La draca es una perra" //1
	// var curso = "La draca es una perra" //2
	/*curso := "La draca"

	longitud := len(curso) // retorna int

	fmt.Println(longitud)

	a := curso[1] //codigo ASCII
	fmt.Printl"n(a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Printf("%c\n", a)*/
	/*
		var nombre string
		var edad int
		var altura int

		fmt.Print("Ingresa tu nombre: ")
		fmt.Scanf("%s", &nombre)
		fmt.Print("Ingresa tu EDAD: ")
		fmt.Scanf("%d", &edad)
		fmt.Print("Ingresa tu altura en cm: ")
		fmt.Scanf("%d", &altura)

		fmt.Printf("Hola %s con una edad %d y una altura %d cm \n", nombre, edad, altura)*/

	// slice := make([]int, 3, 3)

	// fmt.Println(slice)
	//panic
	/*
		var dividnedo, divisor int

		fmt.Print("INgresa un valor para el dividendo: ")
		fmt.Scanf("%d", &dividnedo)

		fmt.Print("INgresa un valor para el divisor: ")
		fmt.Scanf("%d", &divisor)

		if divisor == 0 {
			panic("Que onda no se puede")
		}
		resultado := dividnedo / divisor

		fmt.Println(resultado)
	*/
	// funcion_anonima := func() {
	// 	fmt.Println("Funcion anonima")
	// }
	// funcion_anonima()

	//Manejo de errores
	// if resultado, err := division(10, 0); err != nil {
	// 	// fmt.Println(err)
	// 	panic(err)
	// } else {
	// 	fmt.Println("El resultado es:", resultado)
	// }

}
