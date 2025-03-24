package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

// Map con los dias de la semana en español
var diaSemana = map[time.Weekday]string{
	time.Sunday:    "Domingo",
	time.Monday:    "Lunes",
	time.Tuesday:   "Martes",
	time.Wednesday: "Miércoles",
	time.Thursday:  "Jueves",
	time.Friday:    "Viernes",
	time.Saturday:  "Sábado",
}

// Función que genera una lista de fechas entre dos fechas dadas
func generarDías(FechaInicio time.Time, FechaFin time.Time) []string {
	var fechas []string // Array para almacenar las fechas

	for fecha := FechaInicio; !fecha.After(FechaFin); fecha = fecha.AddDate(0, 0, 1) {
		if fecha.Weekday() != time.Saturday && fecha.Weekday() != time.Sunday { // Evitar Sábados y Domingos
			diaSemana := diaSemana[fecha.Weekday()]
			fechas = append(fechas, fmt.Sprintf("%s - %s", fecha.Format("2006-01-02"), diaSemana)) // Formato de la fecha
		}
	}
	return fechas
}

// Nueva función que cuenta el número de veces que aparece cada día de la semana en un rango de fechas
func contarDiasSemana(FechaInicio time.Time, FechaFin time.Time) map[string]int {
	conteo := make(map[string]int)

	for fecha := FechaInicio; !fecha.After(FechaFin); fecha = fecha.AddDate(0, 0, 1) {
		dia := diaSemana[fecha.Weekday()]
		conteo[dia]++
	}
	return conteo
}

func main() {
	// Definir el rango de fechas desde el 1 de junio hasta el 31 de agosto de 2025
	FechaInicio := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	FechaFin := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	// Generar la lista de fechas dentro de este rango
	fechas := generarDías(FechaInicio, FechaFin)

	//Abrir archivo CSV
	file, err := os.Open("./fichero.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo", err)
		return
	}
	defer file.Close()

	//Leer archivo
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo", err)
		return
	}
	//Crear un nuevo archivo excel
	f := excelize.NewFile()

	//Escribir los encabezados
	f.SetCellValue("Sheet1", "A1", "Nombre")
	f.SetCellValue("Sheet1", "B1", "Teléfono")

	//Escribimos las fechas como encabezado
	for j, fecha := range fechas {
		colName, _ := excelize.ColumnNumberToName(j + 3)
		f.SetCellValue("Sheet1", fmt.Sprintf("%s1", colName), fecha)
	}

	//Pasar datos de csv a excel

	for i, record := range records {
		f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), record[0])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), record[1])

		//Escribir las fechas del calendario
		for j := 0; j < len(fechas) && j < len(record)-2; j++ {
			colName, _ := excelize.ColumnNumberToName(j + 3)
			f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", colName, i+2), record[j+2])
		}
	}

	//Guardar el archivo Excel
	err = f.SaveAs("./excel.xlsx")
	if err != nil {
		fmt.Println("Error al guardar el archivo Excel.", err)
		return
	}
	fmt.Println("Archivo Excel generado")
}
