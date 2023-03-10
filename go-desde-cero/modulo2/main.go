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
}