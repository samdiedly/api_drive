package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

// Mapa con los días de la semana en español
var diaSemana = map[time.Weekday]string{
	time.Monday:    "Lunes",
	time.Tuesday:   "Martes",
	time.Wednesday: "Miércoles",
	time.Thursday:  "Jueves",
	time.Friday:    "Viernes",
}

// Mapa de meses en español (solo junio, julio y agosto)
var meses = map[time.Month]string{
	time.June:   "Junio",
	time.July:   "Julio",
	time.August: "Agosto",
}

// Función que genera una lista de fechas entre dos fechas dadas, excluyendo sábados, domingos y otros meses
func generarDías(FechaInicio time.Time, FechaFin time.Time) []string {
	var fechas []string

	for fecha := FechaInicio; !fecha.After(FechaFin); fecha = fecha.AddDate(0, 0, 1) {
		// Filtrar solo los meses de junio, julio y agosto y evitar sábados y domingos
		if _, esMesValido := meses[fecha.Month()]; esMesValido && fecha.Weekday() >= time.Monday && fecha.Weekday() <= time.Friday {
			dia := diaSemana[fecha.Weekday()]
			mes := meses[fecha.Month()]
			fechas = append(fechas, fmt.Sprintf("%s %d %s %d", dia, fecha.Day(), mes, fecha.Year()))
		}
	}
	return fechas
}

func main() {
	// Definir el rango de fechas desde el 1 de junio hasta el 31 de agosto de 2025
	FechaInicio := time.Date(2025, 6, 1, 0, 0, 0, 0, time.UTC)
	FechaFin := time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)

	// Generar la lista de fechas dentro de este rango sin sábados, domingos ni otros meses
	fechas := generarDías(FechaInicio, FechaFin)

	// Abrir archivo CSV
	file, err := os.Open("./fichero.csv")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	// Leer archivo CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error al leer el archivo CSV:", err)
		return
	}

	// Crear un nuevo archivo Excel
	f := excelize.NewFile()
	sheetName := "Datos"
	f.SetSheetName("Sheet1", sheetName)

	// Escribir los encabezados
	f.SetCellValue(sheetName, "A1", "Nombre")
	f.SetCellValue(sheetName, "B1", "Teléfono")

	// Escribir las fechas como encabezado desde la columna C en adelante
	for j, fecha := range fechas {
		colName, _ := excelize.ColumnNumberToName(j + 3)
		f.SetCellValue(sheetName, fmt.Sprintf("%s1", colName), fecha)
	}

	// Pasar datos de CSV a Excel sin duplicar las columnas "Nombre" y "Teléfono"
	for i, record := range records {
		if len(record) < 2 {
			continue // Saltar filas que no tengan al menos 2 columnas
		}
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), record[0])
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), record[1])
	}

	// Guardar el archivo Excel
	err = f.SaveAs("./datos.xlsx")
	if err != nil {
		fmt.Println("Error al guardar el archivo Excel:", err)
		return
	}

	fmt.Println("Archivo Excel generado exitosamente.")
}
