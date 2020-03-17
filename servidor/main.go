package main

import (
	_ "github.com/go-sql-driver/mysql" // La librería que nos permite conectar a MySQL
	"database/sql"                     // Interactuar con bases de datos
	"fmt"                              // Imprimir mensajes y esas cosas
	"github.com/rs/cors"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func handler(w http.ResponseWriter, req *http.Request) {
    // ...
	enableCors(&w)
    // ...
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	
}


//----------------------------------------SE OBTIENE ACCESO A LA BASE DE DATOS-----------------------------------------

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := ""
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "pasantia"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}


//Estrucutura consulta video
type Video struct {
	Id string `json:"id,omitempty"`
	Tema string `json:"tema,omitempty"`
	Autor string `json:"autor,omitempty"`
	Fecha string `json:"fecha,omitempty"`
	Link string `json:"link,omitempty"`
  }

//Estrucutura consulta audio
type Audio struct {
	Id string `json:"id,omitempty"`
	Tema string `json:"tema,omitempty"`
	Participantes string `json:"participantes,omitempty"`
	Fecha string `json:"fecha,omitempty"`
	Link string `json:"link,omitempty"`
  }

  //Estrucutura consulta publicidad
type Publicidad struct {
	Id string `json:"id,omitempty"`
	Tema string `json:"tema,omitempty"`
	Tipo string `json:"tipo,omitempty"`
	Participantes string `json:"participantes,omitempty"`
	Fecha string `json:"fecha,omitempty"`
	Link string `json:"link,omitempty"`
  }

    //Estrucutura consulta documentos
type Documento struct {
	Id string `json:"id,omitempty"`
	Tema string `json:"tema,omitempty"`
	Tipo string `json:"tipo,omitempty"`
	Area string `json:"area,omitempty"`
	Autor string `json:"autor,omitempty"`
	Fecha string `json:"fecha,omitempty"`
	Link string `json:"link,omitempty"`
  }

  //----------------------------------------SE REALIZA LA CONSULTA PARA OBTENER VIDEOS ESPECIFICOS-------------------------------
func obtenerVideos(tema string, autor string, fecha string) ([]Video, error) {
	videos := []Video{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	filas, err := db.Query("SELECT id, tema, autor, fecha, link FROM videos where autor LIKE '%"+autor+"%' AND tema LIKE '%"+tema+"%' AND fecha LIKE '%"+fecha+"%'")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Video

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Tema, &c.Autor, &c.Fecha , &c.Link)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		
		// Y si no, entonces agregamos lo leído al arreglo
	
		videos = append(videos, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return videos, nil
}

  //----------------------------------------SE REALIZA LA CONSULTA PARA OBTENER AUDIOS ESPECIFICOS-------------------------------
  func obtenerAudios(tema string, participantes string, fecha string) ([]Audio, error) {
	audios := []Audio{}
	db, err := obtenerBaseDeDatos()
	fmt.Printf(fecha)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	filas, err := db.Query("SELECT id, tema, participantes, fecha, link FROM audios where participantes LIKE '%"+participantes+"%' AND tema LIKE '%"+tema+"%' AND fecha LIKE '%"+fecha+"%'")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Audio

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Tema, &c.Participantes, &c.Fecha , &c.Link)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		
		// Y si no, entonces agregamos lo leído al arreglo
	
		audios = append(audios, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return audios, nil
}

 //----------------------------------------SE REALIZA LA CONSULTA PARA OBTENER PUBLICIDAD ESPECIFICOS-------------------------------
 func obtenerPublicidad(tema string, tipo string, participantes string, fecha string) ([]Publicidad, error) {
	publicidades := []Publicidad{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	filas, err := db.Query("SELECT id, tema, tipo, participantes, fecha, link FROM publicidad where participantes LIKE '%"+participantes+"%' AND tema LIKE '%"+tema+"%' AND tipo LIKE '%"+tipo+"%' AND fecha LIKE '%"+fecha+"%'")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Publicidad

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Tema, &c.Tipo, &c.Participantes, &c.Fecha , &c.Link)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		
		// Y si no, entonces agregamos lo leído al arreglo
	
		publicidades = append(publicidades, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return publicidades, nil
}

 //----------------------------------------SE REALIZA LA CONSULTA PARA OBTENER DOCUMENTOS ESPECIFICOS-------------------------------
 func obtenerDocumento(tema string, tipo string, area string, autor string, fecha string) ([]Documento, error) {
	documentos := []Documento{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	
	filas, err := db.Query("SELECT id, tema, tipo, area, autor, fecha, link FROM documentos where autor LIKE '%"+autor+"%' AND tema LIKE '%"+tema+"%' AND tipo LIKE '%"+tipo+"%' AND area LIKE '%"+area+"%' AND fecha LIKE '%"+fecha+"%'")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Documento

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.Id, &c.Tema, &c.Tipo, &c.Area, &c.Autor, &c.Fecha , &c.Link)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		
		// Y si no, entonces agregamos lo leído al arreglo
	
		documentos = append(documentos, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return documentos, nil
}
//------------------------------  EndPoints ---------------------------------------------------------

// EndPoint Videos
func GetVideoEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
 
  videos, err := obtenerVideos(params["tema"],params["autor"],params["año"])
	if err != nil {
		fmt.Printf("Error obteniendo videos: %v", err)
		return
	}
	var resultado []Video
	for _, video := range videos {
		resultado = append(resultado, Video{Id: video.Id, Tema: video.Tema, Autor: video.Autor, Fecha: video.Fecha, Link: video.Link})
		fmt.Printf("%v\n", video.Tema)
	}
	fmt.Printf("%v\n", resultado)
	json.NewEncoder(w).Encode(resultado)
}

//busqueda de todos los videos
func GetVideosEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
		
	videos, err := obtenerVideos("","","20")
	if err != nil {
		fmt.Printf("Error obteniendo videos: %v", err)
		return
	}
	var resultado []Video
	for _, video := range videos {
		resultado = append(resultado, Video{Id: video.Id, Tema: video.Tema, Autor: video.Autor, Fecha: video.Fecha, Link: video.Link})
		fmt.Printf("%v\n", video.Tema)
	}
	fmt.Printf("%v\n", resultado)
	json.NewEncoder(w).Encode(resultado)
}

// EndPoint Audios
func GetAudioEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
 
  audios, err := obtenerAudios(params["tema"],params["participantes"],params["año"])
	if err != nil {
		fmt.Printf("Error obteniendo audios: %v", err)
		return
	}
	var resultado []Audio
	for _, audio := range audios {
		resultado = append(resultado, Audio{Id: audio.Id, Tema: audio.Tema, Participantes: audio.Participantes, Fecha: audio.Fecha, Link: audio.Link})
		fmt.Printf("%v\n", audio.Tema)
	}
	fmt.Printf("%v\n", resultado)
	json.NewEncoder(w).Encode(resultado)
}

// EndPoint publicidad
func GetPublicidadEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
 
  publicidades, err := obtenerPublicidad(params["tema"],params["tipo"],params["participantes"],params["año"])
	if err != nil {
		fmt.Printf("Error obteniendo publicidad: %v", err)
		return
	}
	var resultado []Publicidad
	for _, publicidad := range publicidades {
		resultado = append(resultado, Publicidad{Id: publicidad.Id, Tema: publicidad.Tema, Tipo: publicidad.Tipo,Participantes: publicidad.Participantes, Fecha: publicidad.Fecha, Link: publicidad.Link})
		fmt.Printf("%v\n", publicidad.Tema)
	}
	fmt.Printf("%v\n", resultado)
	json.NewEncoder(w).Encode(resultado)
}

// EndPoint documentos
func GetDocumentosEndpoint(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
 
  documentos, err := obtenerDocumento(params["tema"],params["tipo"],params["area"],params["autor"],params["año"])
	if err != nil {
		fmt.Printf("Error obteniendo documento: %v", err)
		return
	}
	var resultado []Documento
	for _, documento := range documentos {
		resultado = append(resultado, Documento{Id: documento.Id, Tema: documento.Tema, Tipo: documento.Tipo,Area: documento.Area, Autor: documento.Autor, Fecha: documento.Fecha, Link: documento.Link})
		fmt.Printf("%v\n", documento.Tema)
	}
	fmt.Printf("%v\n", resultado)
	json.NewEncoder(w).Encode(resultado)
}
func main(){
	
	

	
	router := mux.NewRouter()
	
	
	
	router.HandleFunc("/videos", GetVideosEndpoint).Methods("GET")
	router.HandleFunc("/video/{tema},{autor},{año}", GetVideoEndpoint).Methods("GET")
	router.HandleFunc("/audio/{tema},{participantes},{año}", GetAudioEndpoint).Methods("GET")
	router.HandleFunc("/publicidad/{tema},{tipo},{participantes},{año}", GetPublicidadEndpoint).Methods("GET")
	router.HandleFunc("/documento/{tema},{tipo},{area},{autor},{año}", GetDocumentosEndpoint).Methods("GET")
	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))



}