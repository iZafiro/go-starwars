package concerns

import (
	"fmt"
	"log"

	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Añade una ciudad
func CAddCity(planet string, city string, value int32, planetVectors map[string][3]int32, folder string, node int32, write bool) (bool, map[string][3]int32) {
	success := true
	filename := "Registro_" + planet + ".txt"
	// Se crea el archivo de registro del planeta si es que no existe
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, planetVectors
	}
	// Se lee el archivo de registro del planeta
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, planetVectors
	}

	lines := strings.Split(string(input), "\n")
	// Si la ciudad ya se encuentra en el registro del planeta, se retorna false
	for _, line := range lines {
		if strings.Contains(line, city) {
			success = false
		}
	}
	// Si la ciudad aún no ha sido creada se agrega
	if success {
		// Se actualizan el vector de cambios y los logs de registro
		if write {
			planetVectors = updateVector(planet, planetVectors, node)
			updateLog("AddCity "+planet+" "+city+" "+fmt.Sprint(value), folder)

		}
		// Se añade la ciudad al registro del planeta
		fmt.Fprintln(f, planet+" "+city+" "+fmt.Sprint(value))
		if err != nil {
			fmt.Println(err)
			return false, planetVectors
		}
		// Se cierra el archivo
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return false, planetVectors
		}

	}
	// Se retorna un booleano que indica si la operación fue exitosa y los vectores actualizados
	return success, planetVectors

}

// Actualiza el nombre de una ciudad determinada
func CUpdateName(planet string, city string, new_value string, planetVectors map[string][3]int32, folder string, node int32, write bool) (bool, map[string][3]int32) {
	success := false
	// Se abre el archivo de registro del planeta, si no existe se retorna false y no se generan cambios
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, planetVectors
	}

	lines := strings.Split(string(input), "\n")
	// Si se encuentra la línea que contiene el nombre de la ciudad, esta se reemplaza por el nuevo nombre
	for i, line := range lines {
		if strings.Contains(line, city) {
			lines[i] = strings.Replace(lines[i], city, new_value, 1)
			success = true
		}
	}
	// Se vuelve a armar el archivo y se guardan los cambios
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		return false, planetVectors
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("UpdateName "+planet+" "+city+" "+new_value, folder)
	}
	// Se retorna un booleano que indica si la operación fue exitosa y los vectores actualizados
	return success, planetVectors

}

// Actualiza el número de rebeldes en una ciudad determinada
func CUpdateNumber(planet string, city string, new_value int32, planetVectors map[string][3]int32, folder string, node int32, write bool) (bool, map[string][3]int32) {
	success := false
	// Se abre el archivo de registro del planeta, si no existe se retorna false y no se generan cambios
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, planetVectors
	}

	lines := strings.Split(string(input), "\n")
	// Si se encuentra la línea que contiene el nombre de la ciudad, se reemplaza el valor numérico por el valor nuevo
	for i, line := range lines {
		if strings.Contains(line, city) {
			lines[i] = planet + " " + city + " " + fmt.Sprint(new_value)
			success = true
		}
	}
	// Se vuelve a armar el archivo y se guardan los cambios
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		return false, planetVectors
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("UpdateNumber "+planet+" "+city+" "+fmt.Sprint(new_value), folder)
	}
	// Se retorna un booleano que indica si la operación fue exitosa y los vectores actualizados
	return success, planetVectors

}

// Borra una ciudad determinada
func CDeleteCity(planet string, city string, planetVectors map[string][3]int32, folder string, node int32, write bool) (bool, map[string][3]int32) {
	success := false
	// Se abre el archivo de registro del planeta, si no existe se retorna false y no se generan cambios
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, planetVectors
	}

	lines := strings.Split(string(input), "\n")
	new_lines := []string{}
	// Si se encuentra la línea que contiene el nombre de la ciudad, se omite y se copia el resto del archivo
	for i, line := range lines {
		if !strings.Contains(line, city) {
			new_lines = append(new_lines, lines[i])
		} else {
			success = true
		}
	}
	// Se vuelve a armar el archivo y se guardan los cambios
	output := strings.Join(new_lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		return false, planetVectors
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("DeleteCity "+planet+" "+city, folder)
	}
	// Se retorna un booleano que indica si la operación fue exitosa y los vectores actualizados
	return success, planetVectors
}

// Obtiene la cantidad de rebeldes en una ciudad determinada
func CGetRebels(planet string, city string, folder string) (bool, int32) {
	rebels := int32(-1)
	// Se abre el archivo de registro del planeta, si no existe se retorna false y el valor por defecto -1 que indica error
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, rebels
	}

	lines := strings.Split(string(input), "\n")
	// Se busca la línea que contenga la ciudad y se obtiene el número de rebeldes asociado
	for i, line := range lines {
		if strings.Contains(line, city) {
			aux, _ := strconv.ParseInt(strings.Split(lines[i], " ")[2], 10, 32)
			rebels = int32(aux)
		}
	}
	// Se retorna un booleano que indica si la operación fue exitosa y la cantidad de rebeldes
	return true, rebels

}

// Actualiza el Log
func updateLog(op string, folder string) {
	filename := "Logs.txt"
	// Se abre el archivo de Logs si es que existe, sino se genera uno
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Se agrega la línea recibida
	fmt.Fprintln(f, op)

	if err != nil {
		fmt.Println(err)
		return
	}
	// Se cierra el archivo
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Obtiene los logs de registro
func CGetLogs(planetVectors map[string][3]int32, folder string) []string {
	logs := []string{}
	// Se abre el archivo de logs de registro, si no existe retorna un arreglo vacío
	input, err := ioutil.ReadFile(folder + "Logs.txt")
	if err != nil {
		return logs
	}

	lines := strings.Split(string(input), "\n")
	// Lee los logs y los almacena como strings junto al vector de su planeta respectivo separado por un salto de línea
	//Ej: "AddCity Tattooine Mos_Eisley 5\n3,0,0"
	for i := range lines {
		if len(strings.Split(lines[i], " ")) > 1 {
			city := strings.Split(lines[i], " ")[1]
			logs = append(logs, lines[i]+"\n"+fmt.Sprint(planetVectors[city][0])+", "+fmt.Sprint(planetVectors[city][1])+", "+fmt.Sprint(planetVectors[city][2])+"\n")
		}

	}
	// Borra el archivo de logs de registro
	deleteLogs(folder)
	// Retorna el contenido del archivo antes de ser borrado
	return logs

}

// Borra el archivo que contiene los logs de registro
func deleteLogs(folder string) {
	e := os.Remove(folder + "Logs.txt")
	if e != nil {
		log.Fatal(e)

	}
}

// Borra el registro de un planeta
func deletePlanet(planet string, folder string) bool {
	e := os.Remove(folder + "Registro_" + planet + ".txt")
	// Retorna un booleano indicando si la operación fue o no fue exitosa
	return e == nil
}

// Función que se encarga de hacer merge
func CMerge(files []string, planetVectors map[string][3]int32, folder string) (bool, map[string][3]int32) {
	// Crea un respaldo de planetVectors
	backup_planetVectors := make(map[string][3]int32)

	// Copia el contenido de planetVectors
	for key, value := range planetVectors {
		backup_planetVectors[key] = value
	}
	// Lee el archivo recibido que contiene los nuevos vectores y la información de los nuevos registros
	for file_index := range files {
		lines := strings.Split(files[file_index], "\n")
		// Obtiene el planeta del archivo
		planet := lines[0]
		// Obtiene el vector del planeta y lo guarda
		vector := strings.Split(lines[1], ", ")
		x, _ := strconv.ParseInt(vector[0], 10, 32)
		y, _ := strconv.ParseInt(vector[1], 10, 32)
		z, _ := strconv.ParseInt(vector[2], 10, 32)
		planetVectors[planet] = [3]int32{int32(x), int32(y), int32(z)}
		// Borra el archivo del planeta...
		deletePlanet(planet, folder)
		filename := "Registro_" + planet + ".txt"
		//...para volver a escribirlo
		for i := 2; i < len(lines); i++ {
			f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return false, backup_planetVectors
			}
			fmt.Fprintln(f, lines[i])
			if err != nil {
				return false, backup_planetVectors
			}
			err = f.Close()
			if err != nil {
				return false, backup_planetVectors
			}

		}

	}
	// Se retorna un booleano que indica si la operación fue exitosa y los vectores actualizados
	return true, planetVectors

}

// Actualiza el vector de cambios
func updateVector(planet string, planetVectors map[string][3]int32, node int32) map[string][3]int32 {
	// Comprueba si el planeta está registrado
	value, check_variable_name := planetVectors[planet]
	// Si es que ya existe se añade 1 a los cambios del nodo correspondiente
	if check_variable_name {
		switch node {
		case 1:
			value[0]++
		case 2:
			value[1]++
		case 3:
			value[2]++
		}

		planetVectors[planet] = value
		// Si es que no existe se crea el nuevo vector con el cambio del nodo correspondiente
	} else {
		switch node {
		case 1:
			planetVectors[planet] = [3]int32{1, 0, 0}
		case 2:
			planetVectors[planet] = [3]int32{0, 1, 0}
		case 3:
			planetVectors[planet] = [3]int32{0, 0, 1}
		}

	}
	// Retorna los vectores actualizados
	return planetVectors
}

// Obtiene el vector de un planeta determinado
func CGetVector(planet string, planetVectors map[string][3]int32) [3]int32 {
	value, check_variable_name := planetVectors[planet]
	// Retorna el vector del planeta si es que existe
	if check_variable_name {
		return value
		// Retorna un vector sin cambios si es que no existe
	} else {
		return [3]int32{0, 0, 0}
	}
}

// Borra los archivos creados durante la ejecución
func CRemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}
