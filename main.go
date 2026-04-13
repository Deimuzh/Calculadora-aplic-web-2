	package main

	import "fmt"
	func continuar() bool{
		var respuesta string
		fmt.Print("¿Deseas realizar otra operación? (s/n): ")
		fmt.Scan(&respuesta)
		if respuesta == "s" || respuesta == "S" {
			return true // Llamar a la función main para reiniciar la calculadora
		}
		return false // Salir del programa
	}

	func main() {
		var a, b float64
		var operacion string

		fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

		fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
		fmt.Scan(&operacion)

		switch operacion {

		case "+", "-", "*", "/", "^":
			// Operaciones que necesitan 2 números
			fmt.Print("Ingresa el primer número: ")
			fmt.Scan(&a)

			fmt.Print("Ingresa el segundo número: ")
			fmt.Scan(&b)

			switch operacion {
			case "+":
				fmt.Printf("Resultado: %.2f\n", a+b)
			case "-":
				fmt.Printf("Resultado: %.2f\n", a-b)
			case "*":
				fmt.Printf("Resultado: %.2f\n", a*b)
			case "/":
				if b == 0 {
					fmt.Println("Error: no se puede dividir entre cero")
				} else {
					fmt.Printf("Resultado: %.2f\n", a/b)
				}
			case "^":
				resultado := 1.0
				for i := 0; i < int(b); i++ {
					resultado *= a
				}
				fmt.Printf("Resultado: %.2f\n", resultado)
			}

		case "!":
			// Solo necesita 1 número
			fmt.Print("Ingresa el número: ")
			fmt.Scan(&a)

			n := int(a)
			if n < 0 {
				fmt.Println("Error: no existe factorial de negativos")
			} else if n == 0 {
				fmt.Println("Resultado: 1")
			} else {
				resultado := 1
				for i := 1; i <= n; i++ {
					resultado *= i
				}
				fmt.Println("Resultado:", resultado)
			}

		default:
			fmt.Println("Error: operación no reconocida")
		}
		if !continuar() {
			fmt.Println("¡Gracias por usar la calculadora! Hasta luego.")
			return
		}
	}
