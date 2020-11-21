package usersqueries

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/CJN-Team/examanager-server/models"

	//"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//AutomaticCreationUsers Se encarga de crear por medio de una hoja de google sheets los usuarios de manera automatic
func AutomaticCreationUsers(link string, loggedUser string, profile string, loggedInstitution string) (bool, error) {

	response, error := getDocument(link)

	if error != nil {
		return response, error
	}

	error = readDocument(loggedUser, profile, loggedInstitution)

	if error != nil {
		return false, error
	}

	return response, nil

}

func getDocument(link string) (bool, error) {
	client := http.Client{
		Timeout: time.Duration(30 * int64(time.Second)),
	}

	response, error := client.Get(link)

	if error != nil {
		error := errors.New("El archivo no puede ser descargado")
		return false, error
	}

	if response.StatusCode != 200 {
		error := errors.New("El archivo no puede ser descargado debido a que se obtuvo un estatus diferente a 200")
		return false, error
	}

	//fmt.Println(response.Header["Content-Type"][0])
	if response.Header["Content-Type"][0] != "text/csv" {
		error := errors.New("Los datos obtenidos del link no son de tipo CSV, intente con un formato valido")
		return false, error
	}

	documentData, error := ioutil.ReadAll(response.Body)

	if error != nil {
		error := errors.New("No es posible leer el body de la respuesta")
		return false, error
	}

	//fmt.Println(documentData)

	//Esta parte puede ser omitida cuando se tenga la parte de analisis de los datos
	error = ioutil.WriteFile("TemporalData", documentData, 0644)

	if error != nil {
		error := errors.New("Es imposible guardar los datos en el archivo")
		return false, error
	}

	return true, nil
}

func readDocument(loggedUser string, profile string, loggedInstitution string) error {
	dataRaw, error := os.Open("./TemporalData")

	if error != nil {
		return error
	}

	scanner := bufio.NewScanner(dataRaw)

	var position = 0
	for scanner.Scan() {
		if position != 0 {
			register := strings.Split(scanner.Text(), ",")
			if len(register) != 6 {
				error = errors.New("La cantidad de campos pasados no es valida")
				break
			}

			newUser := models.User{
				ID:       register[0],
				IDType:   register[1],
				Profile:  profile,
				Name:     register[2],
				LastName: register[3],
				Email:    register[4],
				Password: register[5],
			}

			AddUser(newUser, loggedUser, loggedInstitution)
		}
		position++
	}

	if error != nil {
		return error
	}

	dataRaw.Close()
	return nil
}
