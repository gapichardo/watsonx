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
	MI_TOKEN  = "Bearer "

	TRAINING = "Input:\\n\\nOrden de Compra\\nNo. 1068647333\\nFecha de Orden de Compra : 29-DIC-2022\\nProveedor:\\nIBM DE MEXICO COMERCIALIZACION Y\\nSERVICIOS S DE RL DE CV\\nALFONSO NAPOLES GANDARA\\n3111,PARQUE CORPORATIVO PENA\\nBLANCA,ALVARO OBREGON\\nMEXICO DF 01210\\nMEXICO\\nTelefono\\nFax\\nNom. de Contacto\\nTel del Contacto\\nTrms de Pagos\\nTransportador\\nFOB\\nFletes\\nCliente #\\n52 5527134678\\n18136049110\\nIMMEDIATE\\nEnviar a:\\nBANCO FELIZ, S.A.\\nFRACCIONAMIENTO JGOLONDRINAS 32\\nNAUCALPAN, CIUDAD DE MEXICO DISTRITO FEDERAL MEXICO\\nFacturar a:\\nBANCO FELIZ, S.A.\\nISABEL LA GUERRERA 65 CENTRO DE LA CIUDAD DE MEXICO AREA 1\\nCIUDAD DE MEXICO DISTRITO FEDERAL\\n06000\\nMEXICO\\nAttn.:  LUIS GONZALEZ\\nMail Stop: Jardines\\nEmail: LUIS.GONZALEZ@BANCOFELIZ.COM Tel: +52 56 3876 2000\\nCódigo de Localización : MX 19530\\nExpedir toda factura a la direccion adiunta aqui.\\n! Nota al Proveedor:\\nEsta Orden de Compra (El \"PO\") constituye la oferta del Comprador. El cumplimiento de esta Orden por el Proveedor significa la aceptación de los Términos y Condiciones aplicables, que fueron establecidos en un acuerdo contraído por las partes previo a la\\nOrden.\\nLos Estándares de Citi para Proveedores (http://www.citigroup.com/citi/suppliers/supplierstandards.htm), han sido establecidos\\nbara facilitar el cumplimiento del broveedor con los terminos contraciales v otras bolticas esenciales de citi que son ablicables a\\ntodos los proveedores que proveen productos y/o servicios a Citi.\\nEste numero de referencia es el Il sumero de la LINEA DE EMBARQUE. Favor indicarle apropiadamente en sus facturas.\\nRef\\nNo\\nLínea de Descripción\\nDescripción: IBM Datapower 6 equipos Gateway X2 y 2 equipos\\nGateway X3\\nNota Especial:\\n!La presente Orden de Compra se rige exclusivamente por los términos y\\nes comencoscind\\npropuesta NO. 00636000001EROA33 de techa del\\n26 de diciembre de 2022 y por el Addendum Local México de fecha 29 de agosto de 2017 suscrito entre Banco Feliz, S.A. \\n e IBM de México, Comercialización y Servicios, S. de R.L.\\nde CV\\nal IBM International Customer Agreement (||CA), de fecha 30 de septiembre de 2009, que para todos los efectos legales pertinentes se dan aqui por reproducidos y se consideran parte integrante de este documento\\nEl número del PO debe aparecer en todas las facturas para facilitar el pago.\\nCantidad\\nPrecio Uni\\nUSD\\n!\\nTotal de Lín\\n(USD\\n967.647.00\\nSe necesita.\\n30-DIC-2022\\nciti\\nOrden de Compra\\nNo. 1068647154\\nFecha de Orden de Compra : 29-DIC-2022\\nRef\\nNo.\\nLínea de Descripción\\nCantidad\\npara:\\n30-DIC-2022\\nPrecio Uni\\n(USD)\\nTotal de Lín\\n(USD)\\nNota Especial :\\n!La presente Orden de Compra se rige exclusivamente por los términos y condiciones contenidos en la propuesta No. 0063h0000011EROAAZ de fecha del\\n26 de diciembre de 2022 y por el Addendum Local México de fecha 29 de agosto de 2017 suscrito entre Banco Feliz S.A. e IBM de México, Comercialización y Servicios, S. de R.L. de C.V. al IBM International Customer Agreement (IICA), de fecha 30 de septiembre de 2009, que para todos los efectos legales pertinentes se dan aquí por reproducidos y se consideran parte integrante de este documento\\nConfirmar con Jesús Ponciano,  como se realizara\\nla distribucion previamente a la entrega\\nEC# 09516-18675586\\nTotal (Exclusivo de Deberes e Impuestos USD)\\n1.333.717.00\\nEsta Orden de Compra esta sujeta a los Terminos y Condiciones adjuntas.\\nNombre del Contacto de Compras : ESC Mexico Service Desk / ESCMexicoServicedesk@citi.com\\nTeléfono del Contacto de Compras : (800)-767-9000 Nombre de Comprador : PROCUREMENT,\\nBANCOFELIZ-BUYING DESK\\nNote: Esta es una copia electrónica y no requiere ninguna firma.\\nReconocimiento del Proveedor\\nFirma Autorizada\\nImprima Nombre:\\nFecha:\\n\\nOutput:\\nLa leyenda se incluye completa\\n\\nInput:\\nPurchase Order: 2000121305\\nThis purchase order was delivered by Ariba Network.\\nFrom:\\nBANCO FELIZ S.A.\\nGOLONDRINAS 32 LOMAS DE CAPISTRANO, \\n52988 CDMX\\nCMX\\nMexico\\nContract Number\\nC3571-V4\\nOther Information\\nCompany Code: 0077 Purchasing Unit Name: MEX1\\nSHIP ALL ITEMS TO\\n0077_748302 DIRECCION EJECUTIVA TECNOLOG\\nGOLONDRIAS 32 LOMAS DE CAPISTRANO, NAUCALPAN, 52988 CDMX\\nCMX\\nMexico\\nShip To Code: A623\\nEmail: AGUZMAN@bancofeliz.com.mx\\nBILL TO\\nBANCO FELIZ S.A. GOLONDRINAS 32 LOMAS DE CAPISTRANO, NAUCALPAN 52988\\n\\nCMX\\nMexico\\nDELIVER TO\\nFor more information about Ariba and Ariba Network, visit https://www.ariba.com.\\nTo:\\nIBM DE MEXICO COMERCIALIZACION Y SERVICIOS S DE RL DE CV ALFONSO NAPOLES GANDARA N 3111 COLONIA PARQUE CORPORATIVO DE PEÑA BLANCA ALVARO OBREGON\\n01210 CIUDAD DE MEXICO\\nCiudad de México\\nMexico\\nPhone: +52 (Sales) 5574749656 Fax:\\nEmail: ncolin@mx1.ibm.com\\nPurchase Order\\n(New)\\n2000121302 Amount: $1,509,319.95 USD Version: 1\\njulio  Ramirez Bernal\\n0077_748302 DIRECCION EJECUTIVA TECNOLOG\\n \\n                                                                       Line Items\\n     \\n     Line #\\n1\\nNot Available Service 1 (EA) 2 Jun 2023 $1,509,319.95 USD $1,509,319.95 USD CLOUD TRUSTEER FRAUD PROTECTION, LICENCIAMIENTO 2023\\nNo. Schedule Lines Part # / Description\\nType Return Qty (Unit) Need By Price Subtotal\\n                STATUS\\n1 Unconfirmed\\nContract Number\\nC3571-V4\\n        Service Period\\n    \\nService Start Date: 1 Jun 2023 Service End Date: 1 Jul 2024\\n Other Information\\nExpected Value for Unplanned Spend:\\nReq. Line No.: Requester:\\nPR No.:\\nContract ID: Service Start Date: Classification Domain: Classification Code:\\n$1,509,319.95 USD\\n1\\nALICIA ROLDAN  BERNAL PR208311\\nC3571-V4\\njue, 1 jun, 2023\\ncustom\\nJ020203\\n    \\nService Sheet Required.\\nPDF generated by FABIOLA RUIZ on Wednesday 31 May 2023 2:36 PM GMT-05:00\\n\\nOutput:\\nNo se encuentra la leyenda\\n\\nInput:\\nCuliacán, Sinaloa, Junio 30 del 2022.\\nPARA: IBM de México, Comercialización y Servicios, S. de R.L. de C.V. CON CARGO A: COSITAS S.A. DE C.V.\\nCoppel, expide el siguiente pedido al proveedor: IBM de México, Comercialización y Servicios, S. de R.L. de C.V., autorizado en la solicitud de compra #84671 con pedido #79041 en el sistema Coupa, de acuerdo a la cotización enviada por medio del evento #13366 en el mismo sistema, la cual consiste en lo siguiente:\\nCantidad\\nDescripción\\nCosto Unitario (USD)\\nSubtotal (USD)\\n1\\nIBM Cloud Pak for Data System Fabric Switch Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$9,828.57\\n$9,828.57\\n1\\nIBM Cloud Pak for Data System Fabric Switch Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$724.50\\n$724.50\\n1\\nIBM Cloud Pak for Data System Fabric Expansion Switch Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$9,828.57\\n$9,828.57\\n1\\nIBM Cloud Pak for Data System Fabric Expansion Switch Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$724.50\\n$724.50\\n1\\nIBM Cloud Pak for Data System Management Switch Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$1,997.67\\n$1,997.67\\n \\n       1\\nIBM Cloud Pak for Data System Management Switch Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$156.98\\n$156.98\\n1\\nIBM Cloud Pak for Data System Management Expansion Switch Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$1,997.67\\n$1,997.67\\n1\\nIBM Cloud Pak for Data System Management Expansion Switch Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$156.98\\n$156.98\\n2\\nIBM Cloud Pak for Data System Y1001-SL64 Platform Only ICP4D SS Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$22,169.70\\n$44,339.40\\n2\\nIBM Cloud Pak for Data System Y1001-SL64 Platform Only ICP4D SS Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$4,081.35\\n$8,162.70\\n4\\nIBM Cloud Pak for Data System Y1001-XL64 Platform Only ICP4D SS Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$22,169.70\\n$88,678.80\\n4\\nIBM Cloud Pak for Data System Y1001-XL64 Platform Only ICP4D SS Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$4,081.35\\n$16,325.40\\n6\\nIBM Cloud Pak for Data System PDU 1-Phase 200-240 V AC with NEMA 30A L6-30 line cord Appliance Install Annual Appliance Maintenance + Subscription and Support Renewal 12 Months\\n$147.80\\n$886.80\\n6\\nIBM Cloud Pak for Data System PDU 1-Phase 200-240 V AC with NEMA 30A L6-30 line cord Appliance Install Subsequent Appliance Business Critical Service Upgrade 12 Months\\n$33.81\\n$202.86\\n65\\nIBM - SERVICIO DE SOPORTE DE INFRAESTRUCTURA (Bolsa de unidades de servicio de Lab Services)\\n$1076.9230\\n$70,000.00\\n           \\n \\n 1\\nIBM - EXPERTISE CONNECT 12 MESES\\nSUBTOTAL\\nIVA\\nTOTAL\\n$56,300.00\\n$56,300.00\\n$310,311.40\\n$49,649.82\\n$359,961.22\\n  \\nLa presente Orden de Compra se rige exclusivamente por los términos y condiciones contenidos en la Propuesta No. HTCMX2112029682359 de fecha 29 de Junio del 2022 y la emisión de la misma constituye la expresión del consentimiento del Cliente a dichos términos y condiciones. Cualquier otro término se tendrá por no puesto.\\nFACTURAR A:\\nCOSITAS, S.A. DE C.V.\\nCALLE: GOLONDRINAS 32. COL. CAPISTRANO, EDO MEX, C.P.52988 R.F.C.: COS-920428-Q20\\nCONDICIONES DE PAGO: 30 DÍAS\\nPRECIO EN DÓLARES\\n\\nOutput:\\nLa leyenda se incluye\\n\\nInput:\\n"
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

	t1 := append([]byte(TRAINING), j...)
	t2 := append(t1, []byte("Output:")...)

	body := bytes.NewReader(t2)

	req, err := http.NewRequest("POST", "https://us-south.ml.cloud.ibm.com/ml/v1-beta/generation/text?version=2023-05-29", body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", MI_TOKEN)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	fmt.Println(string(bodyText))

	return string(bodyText), nil
}
