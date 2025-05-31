/* 

* Calculadora del estandar IEEE754 en simple precisión

*/
package  main

import (
"fmt"
)

//Corrobora que el string este coanformados solo de 1 o 0
//Devuelfe Falso en caso de no cumplir
func binario(n string) bool {
  l := len(n);
  var vf bool;
  vf = true;
  for i:= 0; i < l; i++{
    if string(n[i]) != "0" && string(n[i]) != "1" {
      vf = false;
      break;
    }
  }
  return vf;
}

//Calcula las potencias de dos positivas o negativas
func pot_dos(e int) float64 {
  var result float64;
  result = 1;
  if e == 0 {
    result = 1;
  } else if e < 0 {
    for i:= 0; i < e*(-1); i++ {
     result = result / 2; 
    } 
    }else if e > 0 {
      for i:= 0; i < e; i++{
        result = result * 2;
      }
  }
  return result;
}

// Calcula el exponente o la mantisa según corresponda
func cal_bin(n string) float64 {
  var e float64;
  e = 0;
  for i:= 0; i < len(n); i++{
    //Se usa la longitud del substring para saber
    //si se esta calculando la mantisa u el exponente
    if len(n) < 9 {
      if string(n[i]) == "1" {
        //Acá se cacula la exponente, como es entero los indices son positivos
        // Se le resta la longitud del substring 
        // menos 1 a i para no tener que recorrer de forma inversa
        e += pot_dos(len(n)-i-1); 
      }
    } else {
      if string(n[i]) == "1" {
        //Acá se calcula mantisa, como es un numero fracionario los indices
        //son negativos, ademas necesariamente los indices tiene que empezar
        //por |1| por lo que a i se le suma 1 y se lo multiplica por -1
        //para transformar el indice en negativo
        e +=pot_dos((i+1)*(-1)); 
      }
    }
  }
  return e
}


// obtiene la mantisa y el exponente en formatoa decimal
func man_exp(n string, m *float64, e *float64, s *int){
  *m = cal_bin(n[9:32]);
	//No me estaría quedando muy claro como funcionan 
	//los indices en este lenguaje, pero bueno. Funca
  *e = cal_bin(n[1:9]);
	*s = int(cal_bin(string(n[0])));
}

//Demendiendo de los valores de la mantisa y el expontente devuelve el
//valor en su formato decimal o el valor NAN e Infinito
func cal_754(s int,m float64,e float64) {
	var sig float64;
	if(s == 1){
		sig = -1;
	} else {
		sig = 1;
	}
	if(e == 0 && m == 0){
		fmt.Println("-",0);
	} else if(e == 0 && m > 0){
		e -= 126;
    fmt.Println();
    fmt.Println("El valor en decimal es de: ");
		fmt.Println((m*pot_dos(int(e)))*sig);
	} else if(e > 0 && e < 255){
		e -= 127;
		m += 1;
    fmt.Println();
    fmt.Println("El valor en decimal es de: ");
		fmt.Println( m*pot_dos(int(e))*sig);
	} else if(e == 255 && m == 0){
    if(sig < 0){
      fmt.Println("Menos Infinito");
    } else {
      fmt.Println("Mas Infinito");
    }
  } else {
    fmt.Println("NAN: Not an Number");
  }

}

func main(){
  var text string
  var m float64;
  var e float64;
	var s int;
  fmt.Println("Ingrese un valor: ")
  fmt.Scanf("%s",&text)
  // Como no vas a tener un bucle while golang, lctm!!!
  if len(text) < 32 || len(text) > 32 && !binario(text) {  
    for  {
    fmt.Println("El while no existe, son los amigos que hicimos en el camino!")
    fmt.Println("Ingrese un valor: ")
    fmt.Scanf("%s",&text)
    if len(text) == 32 && binario(text){
      break
    }
  }
}
  //Guarda los valores correspondientes en m(mantisa),e(exponente),
  // y S(Signo)
  man_exp(text,&m,&e,&s);
	cal_754(s,m,e)
  //Sigo sin poder creer que no tenga un bucle whilte, ta que te pario Peralta!

}

