package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type caso struct {
	Name         string `json:"name"`
	Location     string `json:"location"`
	Age          int    `json:"age"`
	Infectedtype string `json:"infectedtype"`
	State        string `json:"state"`
}

type arregloCasos struct {
	Casos []caso `json:"casos"`
}

func opcionMenu() {
	fmt.Println("1. Indicar URL LoadBalancer")
	fmt.Println("2. Indicar cantidad de gorutinas")
	fmt.Println("3. Cantidad de solicitudes")
	fmt.Println("4. Ruta del archivo")
	fmt.Println("5. Salir")
}

func main() {
	var opcion int
	var br bool = true
	var direccion string
	var cantidadGorutinas int
	var cantidadSolicitudes int
	var direccionArchivo string
	var data []byte
	var err error
	var wg sync.WaitGroup
	res := arregloCasos{}

	fmt.Println("hello world")
	for br {
		fmt.Println("* --- MENU --- *")
		opcionMenu()
		fmt.Print("Seleccione una opción: ")
		fmt.Scanf("%d", &opcion)
		//fmt.Println(strings.Compare(`1`, opcion))
		switch opcion {
		case 1:
			fmt.Print("Ingrese la direccion del LoadBalancer a utilizar: ")
			fmt.Scanf("%s", &direccion)
			fmt.Println(direccion)
		case 2:
			fmt.Print("Cual es la cantidad de Gorutinas a utilizar?: ")
			fmt.Scanf("%d", &cantidadGorutinas)
			if cantidadSolicitudes >= cantidadGorutinas {
				wg.Add(cantidadGorutinas)
				for x := 0; x < cantidadGorutinas; x++ {
					fmt.Println(x)
					go func() {
						mapDatos := map[string]string{"name": res.Casos[x].Name, "location": res.Casos[x].Location,
							"age": strconv.Itoa(res.Casos[x].Age), "infectedtype": res.Casos[x].Infectedtype, "state": res.Casos[x].State}
						mapB, _ := json.Marshal(mapDatos)
						fmt.Println(string(mapB))
						resp, err := http.Post(direccion+"/crear", "application/json", bytes.NewBuffer(mapB))
						if err != nil {
							log.Fatalln(err)
						}

						defer resp.Body.Close()

						body, err := ioutil.ReadAll(resp.Body)
						if err != nil {
							log.Fatalln(err)
						}

						log.Println(string(body))
						wg.Done()
					}()
				}

			}
		case 3:
			fmt.Print("Ingresa la cantidad de solicitudes: ")
			fmt.Scanf("%d", &cantidadSolicitudes)
		case 4:
			fmt.Print("Ingrese la ruta del archivo: ")
			fmt.Scanf("%s", &direccionArchivo)
			data, err = ioutil.ReadFile(direccionArchivo)
			if err != nil {
				fmt.Println("Ocurrio un error al leer el archivo", err)
			}
			contenidoA := string(data)
			json.Unmarshal([]byte(contenidoA), &res)
			fmt.Println(res)
		default:
			fmt.Println("Gracias! vuelva pronto")
			br = false
		}
	}
}
