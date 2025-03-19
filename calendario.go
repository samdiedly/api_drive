package main

import (
	"fmt"
	"time"
)

// Función que genera una lista de fechas entre dos fechas dadas
func generarDías(FechaInicio time.Time, FechaFin time.Time) []string {
	var fechas []string // Array para almacenar las fechas

	for fecha := FechaInicio; !fecha.After(FechaFin); fecha = fecha.AddDate(0, 0, 1) {
		fechas = append(fechas, fecha.Format("2006-01-02")) //Siempre se usa "2006-01-02" en este orden para que Go entienda cómo formatear la fecha.
	}
	return fechas
}
func main() {
	// Definir el rango de fechas desde el 1 de junio hasta el 31 de agosto de 2025
	FechaInicio := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	FechaFin := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	// Generar la lista de fechas dentro de este rango
	fechas := generarDías(FechaInicio, FechaFin)

	// Imprimir todas las fechas generadas
	fmt.Println("Fechas entre junio y agosto de 2025:")
	for _, fecha := range fechas {
		fmt.Println(fecha)
	}
}
