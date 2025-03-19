package main
import (
	"encoding/csv"
	"fmt"
	"os"
)

// Función main para abrir el fichero csv

func main(){
	//Abrir archivo CSV
	file, err := os.Open("./fichero.csv")
	if err != nil {
		fmt.Println("Error al abir el archivo", err)
		return
	}
	defer file.Close() //Cerrar archivo

	//Crear lector CSV
	reader := csv.NewReader(file)

	//Leer el archivo CSV
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo", err)
		return	
	}
	// Procesar y Mostrar datos
	for i, record := range records {
		fmt.Println("línea", i, "es", record)
	}
}