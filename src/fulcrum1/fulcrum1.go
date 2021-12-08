package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type server struct{}

var planetVectors map[string][3]int32

func main() {
	folder := "out"
	planetVectors = make(map[string][3]int32)

	RemoveContents(folder)

}

//Añade una ciudad
func addCity(planet string, city string, value int32) bool {
	success := true
	filename := "Registro_" + planet + ".txt"
	folder := "out/"
	f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false
	}

	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false
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
			return false
		}

		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return false
		}
		updateVector(planet)
		updateLog("AddCity " + planet + " " + city + " " + fmt.Sprint(value))
	}
	return success

}

//Actualiza el nombre de una ciudad determinada
func updateName(planet string, city string, new_value string) bool {
	success := false
	folder := "out/"
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false
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
		updateVector(planet)
		updateLog("UpdateName " + planet + " " + city + " " + new_value)
	}
	return success

}

//Actualiza el número de rebeldes en una ciudad determinada
func updateNumber(planet string, city string, new_value int32) bool {
	success := false
	folder := "out/"
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false
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
		updateVector(planet)
		updateLog("UpdateNumber " + planet + " " + city + " " + fmt.Sprint(new_value))
	}
	return success

}

//Borra una ciudad determinada
func deleteCity(planet string, city string) bool {
	success := false
	folder := "out/"
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		return false
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
		updateVector(planet)
		updateLog("DeleteCity " + planet + " " + city)
	}

	return success
}

//Obtiene la cantidad de rebeldes en una ciudad determinada
func getRebels(planet string, city string) int32 {
	folder := "out/"
	input, err := ioutil.ReadFile(folder + "Registro_" + planet + ".txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	rebels := int32(-1)

	for i, line := range lines {
		if strings.Contains(line, city) {
			aux, _ := strconv.ParseInt(strings.Split(lines[i], " ")[2], 10, 32)
			rebels = int32(aux)
		}
	}
	return rebels

}

//Actualiza el Log
func updateLog(op string) {
	filename := "Logs.txt"
	folder := "out/"
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
func getLogs() []string {
	logs := []string{}
	folder := "out/"
	input, err := ioutil.ReadFile(folder + "Logs.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i := range lines {
		if len(strings.Split(lines[i], " ")) > 1 {
			city := strings.Split(lines[i], " ")[1]
			log.Println(city)
			logs = append(logs, lines[i]+"\n"+fmt.Sprint(planetVectors[city][0])+", "+fmt.Sprint(planetVectors[city][1])+", "+fmt.Sprint(planetVectors[city][2])+"\n")
		}

	}
	deleteLogs()
	return logs

}
func deleteLogs() {
	folder := "out/"
	e := os.Remove(folder + "Logs.txt")
	if e != nil {
		log.Fatal(e)

	}
}
func deletePlanet(planet string) bool {
	folder := "out/"
	e := os.Remove(folder + "Registro_" + planet + ".txt")
	if e != nil {
		return false
	}
	return true
}

func merge(files []string) {
	folder := "out/"
	for file_index := range files {
		lines := strings.Split(files[file_index], "\n")
		planet := lines[0]
		vector := strings.Split(lines[1], ", ")
		x, _ := strconv.ParseInt(vector[0], 10, 32)
		y, _ := strconv.ParseInt(vector[1], 10, 32)
		z, _ := strconv.ParseInt(vector[2], 10, 32)
		planetVectors[planet] = [3]int32{int32(x), int32(y), int32(z)}
		deletePlanet(planet)
		filename := "Registro_" + planet + ".txt"
		for i := 2; i < len(lines)-1; i++ {
			f, err := os.OpenFile(folder+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprintln(f, lines[i])
			if err != nil {
				fmt.Println(err)
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
			}

		}

	}

}

//Actualiza el vector de cambios
func updateVector(planet string) {
	value, check_variable_name := planetVectors[planet]
	if check_variable_name {
		value[0]++
		planetVectors[planet] = value
	} else {
		planetVectors[planet] = [3]int32{1, 0, 0}
	}
}

//Obtiene el vector de un planeta determinado
func getVector(planet string) [3]int32 {
	value, check_variable_name := planetVectors[planet]
	if check_variable_name {
		return value
	} else {
		return [3]int32{0, 0, 0}
	}
}

// Borra los archivos creados durante la ejecución
func RemoveContents(dir string) error {
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
