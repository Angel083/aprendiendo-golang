package main

import "fmt"

func main() {
	// Punteros
	// Variables que asignan el valor de memoria de una variable
	/** Para obtener el valor en memoria de una variable se tiene que
	que declarar de la siguiente manera:
	var p *string
	Y para asignar el valor de memoria se hace uso del amperson (&)
	p = &variable
	*/
	// Desreferenciación
	/** Se puede obtener el valor original de una variable con solo tener
	el valor de memoria del apuntador, a este procesos se le llama
	desreferenciar y solo se tiene que agregar un asterisco a la variable
	que tenga la dirección de memoria (a titulo personal, se puede
	interpretar como si el apuntador se le quitará el asterisco poniendo
	otro asterisco).
	pSinApunta := *p
	*/
	// Array
	/** Arreglos de información de un solo tipo de dato
	var food [3] string
	No se puede cambiar el tamaño de un arreglo después de ser declardo
	Una forma más rápida de declarar un array es de la siguiente manera
	food :=[3]string{"sopa","crema","pizza"}
	*/
	// Slice
	/** Son la forma de trabajar con arreglos de forma dinámica, en sí, los
	slice son apuntadores a arrays. Otra forma de definirlos sería: extrar
	la información de un array o segmentar un array en otro de acuerdo a
	los índices establecidos.
	*/
	set := [7]string{"Gato", "Perro", "Cuyo", "Pavo real", "Rinoceronte", "León", "Lagarto"}
	animals := set[0:5]
	fmt.Println(animals)
	/**
	El primer valor, es de donde empieza y el otro es en donde va a termina
	cabe destacar que es excluyente, es decir, el último valor no será incluido
	dentro de este nuevo arreglo.

	Como son apuntadores, si cambias el valor original de un elemento
	afectará en todas las otras posibles segmentaciones donde se haga referencia,
	en otras palabras si cambio la palabra "Pavo real" a "Aguila" cambiará en
	todos los demás slices
	*/
	aves := set[3:4]
	fmt.Println(aves)
	aves[0] = "Aguila"
	fmt.Println(animals)
	fmt.Println(aves)
	// No es necesario especificar el primer índice si se desea tomar desde
	// el primer valor
	domesticos := set[:3]
	fmt.Println(domesticos)
	zoologico := set[3:]
	fmt.Println(zoologico)
	// No es necesario especificar el último índice si se desea tomar hasta
	// el último valor
	// Se pueden usar todos los elementos solo dejando vacío el arreglo


}