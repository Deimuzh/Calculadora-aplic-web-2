package main

import "fmt"

func main() {
	var a, b float64
	var operacion string

	historial := ""
	contador := 0

	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	for {
		fmt.Print("\nIngresa la operación (+, -, *, /, ^, !): ")
		fmt.Scan(&operacion)

		switch operacion {

		case "+", "-", "*", "/", "^":
			fmt.Print("Ingresa el primer número: ")
			fmt.Scan(&a)

			fmt.Print("Ingresa el segundo número: ")
			fmt.Scan(&b)

			var resultado float64

			switch operacion {
			case "+":
				resultado = a + b
			case "-":
				resultado = a - b
			case "*":
				resultado = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error: no se puede dividir entre cero")
					continue
				}
				resultado = a / b
			case "^":
				resultado = 1
				for i := 0; i < int(b); i++ {
					resultado *= a
				}
			}

			fmt.Printf("Resultado: %.2f %s %.2f = %.2f\n", a, operacion, b, resultado)

			contador++
			historial += fmt.Sprintf("%d) %.2f %s %.2f = %.2f\n", contador, a, operacion, b, resultado)

		case "!":
			fmt.Print("Ingresa el número: ")
			fmt.Scan(&a)

			n := int(a)

			if n < 0 {
				fmt.Println("Error: no existe factorial de negativos")
				continue
			}

			resultado := 1
			for i := 1; i <= n; i++ {
				resultado *= i
			}

			fmt.Printf("Resultado: %d! = %d\n", n, resultado)

			contador++
			historial += fmt.Sprintf("%d) %d! = %d\n", contador, n, resultado)

		default:
			fmt.Println("Error: operación no reconocida")
			continue
		}

		// Preguntar si continúa
		var respuesta string
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scan(&respuesta)

		if respuesta != "s" && respuesta != "S" {
			break
		}
	}

	// Mostrar historial al final
	fmt.Println("\n==== HISTORIAL DE OPERACIONES ====")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)

	fmt.Println("¡Hasta luego!")
}