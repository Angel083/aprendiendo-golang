// package main ejemplo de documentacion tipo docstring
package main

import "fmt"

func main() {
	// ## Variables
	// Nos permite almacenar una cadena de texto
	var dog string
	// Asignar un valor
	dog = "Perro";
	/** Todas las variables deben ser usadas*/
	fmt.Println(dog);
	// Esta es otra forma de declarar una variable
	var dogo string = "Perrote"
	fmt.Println(dogo);
	// También es posible declarar más de una variable en una línea 
	var dogito, cat string = "perrito", "gato"
	fmt.Println(dogito, cat);
	// En go no es necesario especificar un tipo de dato ya que identifica el tipo 
	// por lo que se puede declarar de la siguiente manera
	var dog2, cat2 = "perrito", "gato"
	fmt.Println(dog2, cat2);
	// Existe un shorthand para declarar de la siguiente manera
	dog3, cat3 := "perrito", "gato"
	fmt.Println(dog3, cat3);
	/* No se puede modificar el tipo de dato una vez especificado*/
	// int := "string"
	// int = 1
	/* De igual manera no se puede declarar dos veces una variable */
	// int := "string"
	// int := "cadena"
	/* Curiosamente si una variable no ha sido declarada y se junta con una que ya fue declarada, no causa error*/
	// hand := "mano"
	// hand , face := "hand" , "face"

	// --------------------------------------
	// ## Constantes
	// Tambien se pueden declarar constantes que no van a cambiar de valor
	const edad = 12;
	// Hay que tener cuidado ya que se tienen que declarar implicitamente el valor const
	// de otra manera, como usar el shorthand, será declarada como var perdiendo la propiedad 
	// de ser constante.
	//// edad := 12

	// -----------------------------------------
	// ## Comentarios
	/* Los comentarios multilínea por formato de la comunidad
		se usan para comentar código que no se usará
	*/
	// Los comentario en línea se usan para hacer documentación tipo
	// docstring, como en otros idiomas, la regla es:

	// NombreDelPaquete Informacion sobre lo que hace el paquete
  var NombreDelPaquete = "texto"
  fmt.Println(NombreDelPaquete)
	// Usando el comando 
	// go doc --all 
	// permite mostrar todos los comentarios que sirven para documentacion
	/* 
		Para hacer uso de los comentario y así generar la documentación completa sobre
		 un proyecto que esta subido a github	se hace de la siguiente manera:
		 godoc.org/rutaRepositorio
	*/
	// --------------------------------------------------------
	// variables en Golang
	// Booleanos
	var a bool = true;
	fmt.Printf("Tipo: %T, Valor: %v",a ,a)
	// El paquete fmt en su método Printf, nos da la posibilidad de agregar
	// verbos que serán representados con el signo de porciento y la inicial
	// del verbo, existe documentación al respecto para ver más a fondo. Estos
	// valores serán tomados de las variables que sean estipuladas después del
	// cierre de las comillas en el orden que se indica.

	// Enteros
	// Existen dos grupos de enteros, los uint y los int, los uint son enteros
	// sin signo y alcanzan el doble de los int, es decir, si un int8 va desde
	// -128 a 127, el uint8 va de 0 255. Como dato a tomar en cuenta, si solo 
	// se escribe unit y la arquitectura de la computadora es de 64bits 
	// (el procesador) compilará la variable uint64.
	// Existen alias para algunos tipos de datos enteros
	// * byte -> uint8
	// * rune -> int32 o representa los valores unicode 
	//  (var a rune = 'a' -> 97)
	// Flotantes
	// Existen solo dos tipos de float: 
	// float32 y float64

	/* Nota: en Golang no se pueden sumar datos de diferentes tipos de datos
		númericos. Para solucionar este tipo de problemas se usa el casting:
		se usa el tipo de dato al que quieras convertir como un método ej:
			- c := uint16(a) + b
	*/
	// Blank
	// Existe un tipo de dato llamado blank que es cuando queremos hacer uso
	// de una variable pero no ocuparla por el momento, este tipo se representa
	// como un guión bajo `_`. A tomar en cuenta: este dato no see declara, solo
	// se puede asignar el valor por lo que el shorthand (:=) no es aplicable.

	// Valor cero
	// Todas las variables declaradas tendrán por defecto lo que es considerado
	// el valor cero:
	// numericos = 0
	// strings = ""
	// boolean = false
	
	// Operadores aritméticos y de asignación
	// Aritmeticos
	// 	agrupación 			()
	// 	multipliación 	*
	// 	división 				/
	// 	módulo 					%
	// 	suma 						+
	// 	resta 					-
	// Asignación
	// 	Igual							=
	// 	Suma igual				*=
	// 	Resta igual				-=
	// 	Multiplica igual	*=
	// 	Divide igual			/=
	// 	Módulo igual			%=
	// Post incremento ++
	// Post decremento --
	// Comparacion -> return boolean
	// 	Mayor que					>
	// 	Menor que					<
	// 	Igual que					==
	// 	Diferente de			!=
	// 	Mayor igual que  	>=
	// 	Menor igual que		<=
	// Lógicos
	// 	Y (and)	&&
	// 	O (or)	||
	// Unario ! (es un inversor booleano)


}