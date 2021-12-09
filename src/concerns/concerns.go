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

//Añade una ciudad
func CAddCity(planet string, city string, value int32, planetVectors map[string][3]int32, folder string, node int32) (bool, map[string][3]int32) {
	success := true
	filename := "Registro_" + planet + ".txt"
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, map[string][3]int32{}
	}

	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, map[string][3]int32{}
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if strings.Contains(line, city) {
			success = false
		}
	}

	if success {
		fmt.Fprintln(f, planet+" "+city+" "+fmt.Sprint(value))
		if err != nil {
			fmt.Println(err)
			return false, map[string][3]int32{}
		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return false, map[string][3]int32{}
		}
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("AddCity "+planet+" "+city+" "+fmt.Sprint(value), folder)
	}
	return success, planetVectors

}

//Actualiza el nombre de una ciudad determinada
func CUpdateName(planet string, city string, new_value string, planetVectors map[string][3]int32, folder string, node int32) (bool, map[string][3]int32) {
	success := false
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, map[string][3]int32{}
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, city) {
			lines[i] = strings.Replace(lines[i], city, new_value, 1)
			success = true
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("UpdateName "+planet+" "+city+" "+new_value, folder)
	}
	return success, planetVectors

}

//Actualiza el número de rebeldes en una ciudad determinada
func CUpdateNumber(planet string, city string, new_value int32, planetVectors map[string][3]int32, folder string, node int32) (bool, map[string][3]int32) {
	success := false
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, map[string][3]int32{}
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, city) {
			lines[i] = planet + " " + city + " " + fmt.Sprint(new_value)
			success = true
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("UpdateNumber "+planet+" "+city+" "+fmt.Sprint(new_value), folder)
	}
	return success, planetVectors

}

//Borra una ciudad determinada
func CDeleteCity(planet string, city string, planetVectors map[string][3]int32, folder string, node int32) (bool, map[string][3]int32) {
	success := false
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, map[string][3]int32{}
	}

	lines := strings.Split(string(input), "\n")
	new_lines := []string{}

	for i, line := range lines {
		if !strings.Contains(line, city) {
			new_lines = append(new_lines, lines[i])
		} else {
			success = true
		}
	}
	output := strings.Join(new_lines, "\n")
	err = ioutil.WriteFile(folder+"Registro_"+planet+".txt", []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
	if success {
		planetVectors = updateVector(planet, planetVectors, node)
		updateLog("DeleteCity "+planet+" "+city, folder)
	}

	return success, planetVectors
}

//Obtiene la cantidad de rebeldes en una ciudad determinada
func CGetRebels(planet string, city string, folder string) (bool, int32) {
	rebels := int32(-1)
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false, rebels
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, city) {
			aux, _ := strconv.ParseInt(strings.Split(lines[i], " ")[2], 10, 32)
			rebels = int32(aux)
		}
	}
	return true, rebels

}

//Actualiza el Log
func updateLog(op string, folder string) {
	filename := "Logs.txt"
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(f, op)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
func CGetLogs(planetVectors map[string][3]int32, folder string) []string {
	logs := []string{}
	input, err := ioutil.ReadFile(folder + "Logs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		if len(strings.Split(lines[i], " ")) > 1 {
			city := strings.Split(lines[i], " ")[1]
			logs = append(logs, lines[i]+"\n"+fmt.Sprint(planetVectors[city][0])+", "+fmt.Sprint(planetVectors[city][1])+", "+fmt.Sprint(planetVectors[city][2])+"\n")
		}

	}
	deleteLogs(folder)
	return logs

}
func deleteLogs(folder string) {
	e := os.Remove(folder + "Logs.txt")
	if e != nil {
		log.Fatal(e)

	}
}
func deletePlanet(planet string, folder string) bool {
	e := os.Remove(folder + "Registro_" + planet + ".txt")
	return e == nil
}

func CMerge(files []string, planetVectors map[string][3]int32, folder string) (bool, map[string][3]int32) {
	for file_index := range files {
		lines := strings.Split(files[file_index], "\n")
		planet := lines[0]
		vector := strings.Split(lines[1], ", ")
		x, _ := strconv.ParseInt(vector[0], 10, 32)
		y, _ := strconv.ParseInt(vector[1], 10, 32)
		z, _ := strconv.ParseInt(vector[2], 10, 32)
		planetVectors[planet] = [3]int32{int32(x), int32(y), int32(z)}
		deletePlanet(planet, folder)
		filename := "Registro_" + planet + ".txt"
		for i := 2; i < len(lines); i++ {
			f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				return false, map[string][3]int32{}
			}
			fmt.Fprintln(f, lines[i])
			if err != nil {
				return false, map[string][3]int32{}
			}
			err = f.Close()
			if err != nil {
				return false, map[string][3]int32{}
			}

		}

	}
	return true, planetVectors

}

//Actualiza el vector de cambios
func updateVector(planet string, planetVectors map[string][3]int32, node int32) map[string][3]int32 {
	value, check_variable_name := planetVectors[planet]
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
	return planetVectors
}

//Obtiene el vector de un planeta determinado
func CGetVector(planet string, planetVectors map[string][3]int32) [3]int32 {
	value, check_variable_name := planetVectors[planet]
	if check_variable_name {
		return value
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
