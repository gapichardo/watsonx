package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dslipak/pdf"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type Parametros struct {
	Decoding_method    string `json:"decoding_method"`
	Max_new_tokens     int    `json:"max_new_tokens"`
	Min_new_tokens     int    `json:"min_new_tokens"`
	Stop_sequences     []rune `json:"stop_sequences"`
	Repetition_penalty int    `json:"repetition_penalty"`
}
type OrdenCompra struct {
	Model_id   string     `json:"model_id"`
	Input_m    string     `json:"input"`
	Parameters Parametros `json:"parameters"`
	Project_id string     `json:"project_id"`
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", fileHandler)

	fmt.Println("************************ " + "Iniciando el servidor Web en el puerto: " + CONN_PORT + " ************************+")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error al iniciar el servidor http: ", err)
		return
	}
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	file, header, err := r.FormFile("archivo")
	if err != nil {
		log.Println("error al obtener el archivo para el formulario proporcionada: ", err)
		return
	}
	defer file.Close()

	out, pathError := os.Create("templates/" + header.Filename)
	if pathError != nil {
		log.Println("error al crear un archivo para escribir: ", pathError)
		return
	}
	defer out.Close()

	_, copyFileError := io.Copy(out, file)
	if copyFileError != nil {
		log.Println("se produjo un error al copiar el archivo: ", copyFileError)
	}

	//Extracción de texto del pdf
	content, err := readPdf("templates/" + header.Filename)

	// Creación del JSON de carga

	var oc OrdenCompra

	p := Parametros{
		Decoding_method:    "greedy",
		Max_new_tokens:     20,
		Min_new_tokens:     0,
		Stop_sequences:     []rune{},
		Repetition_penalty: 1,
	}

	oc.Model_id = "google/flan-ul2"
	oc.Input_m = content
	oc.Parameters = p
	oc.Project_id = "6dc32fb8-db62-4530-b4bd-cb97ab03a43c"

	newJson, err := json.MarshalIndent(oc, "", "     ")
	if err != nil {
		log.Println("error marshalling", err)
	}

	outJson, pathError := os.Create("templates/ordenCompra.json")
	if pathError != nil {
		log.Println("error al crear el JSON ", pathError)
		return
	}
	outJson.Write(newJson)
	outJson.Close()
	fmt.Println(string(newJson))

	//  LLamado al API
	resAPI, err := callMyApi(newJson)
	if err != nil {
		log.Println("error llamar el API de Watsonx ", pathError)
		return
	}

	fmt.Fprint(w, "<H2>Archivo cargado exitosamente: "+header.Filename+"</H2>"+"<p>"+string(newJson)+"</p>"+"<H2>Respuesta del API de Watsonx</H2>"+"<p>"+resAPI+"</p>")
}

func index(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/loadFile.gohtml")
	parsedTemplate.Execute(w, nil)
}

func readPdf(path string) (string, error) {
	f, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := f.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func callMyApi(j []byte) (string, error) {
	fmt.Println("Calling API...")
	client := &http.Client{}
	body := bytes.NewReader(j)

	req, err := http.NewRequest("POST", "https://us-south.ml.cloud.ibm.com/ml/v1-beta/generation/text?version=2023-05-29", body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer wYiN6DQkzzlnG4ENml5nuNgZ5QOziJMqzl2UyGkFsKSe")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyText))

	return string(bodyText), nil
}
