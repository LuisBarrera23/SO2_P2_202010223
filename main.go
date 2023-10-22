package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"time"

	"github.com/eiannone/keyboard"
)

var bancoDePalabras = []string{
	"manzana", "naranja", "pera", "uva", "kiwi", "platano",
	"perro", "gato", "elefante", "leon", "jirafa", "tigre",
	"sol", "luna", "estrella", "planeta", "galaxia", "nebulosa",
	"flor", "arbol", "pasto", "montana", "oceano", "rio",
	"coche", "bicicleta", "avion", "tren", "barco", "camion",
	"casa", "apartamento", "villa", "castillo", "cabana", "iglu",
	"computadora", "telefono", "television", "tableta", "raton", "teclado",
	"libro", "revista", "periodico", "cuento", "novela", "poesia",
	"pelota", "raqueta", "balon", "bate", "gorra", "uniforme",
	"piano", "violin", "flauta", "guitarra", "bateria", "trompeta",
	"familia", "amigo", "vecino", "colega", "companero", "profesor",
	"arte", "musica", "danza", "pintura", "escultura", "teatro",
	"amor", "felicidad", "tristeza", "enfado", "paz", "esperanza",
	"viaje", "vacaciones", "aventura", "descanso", "turismo", "excursion",
	"mente", "cuerpo", "alma", "corazon", "espiritu", "mente",
	"comida", "bebida", "postre", "ensalada", "sopa", "hamburguesa",
	"ropa", "zapatos", "sombrero", "abrigo", "vestido", "camiseta",
	"lapiz", "pluma", "cuaderno", "libreta", "papel", "borrador",
	"solucion", "problema", "pregunta", "respuesta", "duda", "certeza",
	"arte", "ciencia", "tecnologia", "cultura", "historia", "educacion",
	"fiesta", "celebracion", "aniversario", "cumpleanos", "boda", "graduacion",
	"aventura", "exploracion", "descubrimiento", "misterio", "leyenda", "mito",
	"moda", "tendencia", "estilo", "diseno", "traje", "modelo",
	"ocasion", "reunion", "encuentro", "evento", "conferencia", "congreso",
	"medicina", "enfermedad", "hospital", "medico", "enfermera", "paciente",
	"deporte", "ejercicio", "competicion", "equipo", "jugador", "entrenador",
	"energía", "electricidad", "combustible", "gasolina", "petróleo", "viento",
	"proyecto", "empresa", "negocio", "inversion", "cliente", "mercado",
	"cine", "pelicula", "director", "actor", "actriz", "escenario",
	"naturaleza", "paisaje", "ecosistema", "clima", "tierra", "cielo",
	"historia", "relato", "leyenda", "mito", "cuento", "novela",
	"futuro", "destino", "oportunidad", "sueno", "meta", "esperanza",
	"paz", "guerra", "conflicto", "armas", "ejercito", "batalla",
	"agricultura", "cultivo", "cosecha", "granja", "trabajo", "campo",
	"ruido", "silencio", "sonido", "musica", "voz", "grito",
	"amistad", "compania", "relacion", "conexion", "vinculo", "apego",
	"avion", "vuelo", "aeropuerto", "piloto", "pasajero", "destino",
	"arte", "creatividad", "imaginacion", "inspiracion", "genio", "talento",
	"colores", "rojo", "azul", "verde", "amarillo", "naranja", "morado",
	"numeros", "uno", "dos", "tres", "cuatro", "cinco", "seis",
	"direccion", "camino", "mapa", "norte", "sur", "este", "oeste",
	"emocion", "alegria", "sorpresa", "tristeza", "enojo", "miedo",
	"alimento", "nutricion", "vitamina", "mineral", "dieta", "sabor",
	"vehiculo", "coche", "motocicleta", "bicicleta", "autobus", "tren",
	"estrategia", "plan", "programa", "estrategia", "tactica", "plan",
	"pasaje", "boleto", "boleto", "viaje", "billete", "ticket",
	"instrumento", "musical", "guitarra", "piano", "violoncello", "clarinete",
	"comida", "rapida", "restaurante", "comida", "bebida", "hamburguesa",
	"pelota", "futbol", "baloncesto", "tenis", "voleibol", "béisbol",
	"televisión", "programa", "noticiero", "canales", "serie", "peliculas",
	"carro", "viaje", "estacionamiento", "mecánico", "conductor", "gasolina",
	"producto", "tecnología", "comercio", "venta", "producto", "consumidor",
	"servicio", "cliente", "empresa", "internet", "computadora", "cliente",
	"educación", "clase", "escuela", "profesor", "estudiante", "aprendizaje",
	"cultura", "arte", "literatura", "música", "historia", "religión",
	"fiesta", "celebración", "aniversario", "boda", "fiesta", "cumpleaños",
	"articulo", "noticia", "informacion", "investigacion", "artículo", "novedad",
	"transporte", "vehículo", "coche", "tren", "autobús", "subway",
	"naturaleza", "medio ambiente", "ecosistema", "biodiversidad", "fauna", "flora",
	"humano", "persona", "gente", "hombre", "mujer", "niño",
	"lugar", "sitio", "espacio", "territorio", "localidad", "comunidad",
	"religión", "fe", "creencia", "espiritualidad", "dios", "iglesia",
	"culto", "ritual", "sacrificio", "oración", "festividad", "adoración",
	"industria", "producción", "fabricación", "empresa", "inversión", "comercio",
	"tecnología", "avance", "innovación", "investigación", "desarrollo", "ciencia",
	"concepto", "idea", "conocimiento", "cognición", "entendimiento", "sabiduría",
	"interacción", "relación", "conexión", "contacto", "comunicación", "vinculación",
	"identidad", "personalidad", "individualidad", "carácter", "naturaleza", "yo",
	"compañía", "negocio", "sociedad", "organización", "corporación", "entidad",
	"materia", "sustancia", "sólido", "líquido", "gas", "forma",
	"función", "tarea", "trabajo", "responsabilidad", "ocupación", "puesto",
	"cambio", "modificación", "transformación", "variación", "evolución", "desarrollo",
	"decisión", "elección", "opción", "alternativa", "determinación", "resolución",
	"planeta", "tierra", "mundo", "planeta", "globo", "orbe",
	"descubrimiento", "hallazgo", "revelación", "encuentro", "hallazgo", "revelación",
	"alegría", "felicidad", "entusiasmo", "regocijo", "júbilo", "diversión",
	"preocupación", "inquietud", "ansiedad", "nerviosismo", "tensión", "intriga",
	"convivencia", "relación", "coexistencia", "interacción", "colaboración", "compartir",
	"despedida", "adiós", "separación", "partida", "hasta luego", "hasta pronto",
	"soledad", "aislamiento", "sensación", "emoción", "sentimiento", "impresión",
	"alimento", "comida", "nutrición", "alimentación", "dieta", "nutriente",
	"deseo", "anhelo", "esperanza", "anhelo", "pretensión", "querer",
	"criatura", "ser", "entidad", "organismo", "sujeto", "individuo",
	"habitación", "cuarto", "espacio", "recinto", "área", "sitio",
	"propósito", "meta", "objetivo", "fin", "finalidad", "razón",
	"ejemplo", "muestra", "modelo", "ilustración", "caso", "demostración",
	"valor", "importancia", "relevancia", "significado", "sentido", "consecuencia",
	"aventura", "viaje", "experiencia", "h"}

func crearArchivo() {
	file, err := os.Create("paginacion.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo", err)
		os.Exit(1)
	}
	file.Close()
}

func borrarArchivo() {
	err := os.Remove("paginacion.txt")
	if err != nil {
		fmt.Println("Error al eliminar el archivo", err)
	}

	fmt.Println("Archivo de paginación borrado.")
}

func procesoNumero(tipo int, proceso int) {
	if tipo == 1 {
		numeroAleatorio := rand.Intn(999) + 1

		file, err := os.OpenFile("paginacion.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%d\n", numeroAleatorio)
		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
		}
		fmt.Printf("Proceso %d escribiendo numero: %d\n", proceso, numeroAleatorio)
	} else if tipo == 0 {
		numeros := []int{}
		archivo, err := os.Open("paginacion.txt")
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return
		}
		defer archivo.Close()

		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			linea := scanner.Text()
			numero, err := strconv.Atoi(linea)
			if err == nil {
				numeros = append(numeros, numero)
			}
		}

		if len(numeros) == 0 {
			fmt.Printf("Proceso %d leyendo numero: Inexistente\n", proceso)
		} else {
			fmt.Printf("Proceso %d leyendo numero: %d\n", proceso, numeros[rand.Intn(len(numeros))])
		}
	}

}

func procesoLetra(tipo int, proceso int) {
	if tipo == 1 {
		palabraAleatoria := bancoDePalabras[rand.Intn(len(bancoDePalabras))]

		file, err := os.OpenFile("paginacion.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return
		}
		defer file.Close()

		_, err = fmt.Fprintf(file, "%s\n", palabraAleatoria)
		if err != nil {
			fmt.Println("Error al escribir en el archivo", err)
		}
		fmt.Printf("Proceso %d escribiendo letra: %s\n", proceso, palabraAleatoria)

	} else if tipo == 0 {
		palabras := []string{}
		archivo, err := os.Open("paginacion.txt")
		if err != nil {
			fmt.Println("Error al abrir el archivo", err)
			return
		}
		defer archivo.Close()

		scanner := bufio.NewScanner(archivo)
		for scanner.Scan() {
			linea := scanner.Text()
			if regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(linea) {
				palabras = append(palabras, linea)
			}
		}

		if len(palabras) == 0 {
			fmt.Printf("Proceso %d leyendo letra: Inexistente\n", proceso)
		} else {
			fmt.Printf("Proceso %d leyendo letra: %s\n", proceso, palabras[rand.Intn(len(palabras))])
		}
	}

}

func abrirNuevaTerminalYEjecutarComando(comando string) error {
	cmd := exec.Command("bash", "-c", comando+" > /tmp/paginacion_output.txt")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("gnome-terminal", "--", "bash", "-c", "cat /tmp/paginacion_output.txt; read -p 'Presiona Enter para continuar...'")

	return cmd.Run()
}

func ciclo() {
	go func() {
		for {
			rand.Seed(time.Now().UnixNano())

			// Genera un número aleatorio del 1 al 6
			numeroAleatorio := rand.Intn(6) + 1
			fmt.Println("*******************************************************************")
			switch numeroAleatorio {
			case 1:
				procesoNumero(1, 1)
				procesoNumero(0, 2)
				procesoNumero(0, 3)
				procesoLetra(0, 4)
				procesoLetra(0, 5)
				procesoLetra(0, 6)
			case 2:
				procesoNumero(0, 1)
				procesoNumero(1, 2)
				procesoNumero(0, 3)
				procesoLetra(0, 4)
				procesoLetra(0, 5)
				procesoLetra(0, 6)
			case 3:
				procesoNumero(0, 1)
				procesoNumero(0, 2)
				procesoNumero(1, 3)
				procesoLetra(0, 4)
				procesoLetra(0, 5)
				procesoLetra(0, 6)
			case 4:
				procesoNumero(0, 1)
				procesoNumero(0, 2)
				procesoNumero(0, 3)
				procesoLetra(1, 4)
				procesoLetra(0, 5)
				procesoLetra(0, 6)
			case 5:
				procesoNumero(0, 1)
				procesoNumero(0, 2)
				procesoNumero(0, 3)
				procesoLetra(0, 4)
				procesoLetra(1, 5)
				procesoLetra(0, 6)
			case 6:
				procesoNumero(0, 1)
				procesoNumero(0, 2)
				procesoNumero(0, 3)
				procesoLetra(0, 4)
				procesoLetra(0, 5)
				procesoLetra(1, 6)
			default:
				fmt.Println("Número no esperado")
			}
			fmt.Println("Opciones:")
			fmt.Println("Preciona Q para salir")
			fmt.Println("Preciona P para ver el archivo de paginacion")
			fmt.Println("*******************************************************************")
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

func main() {
	crearArchivo()

	err := keyboard.Open()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	ciclo()
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				if char, key, _ := keyboard.GetKey(); key == keyboard.KeyEsc || char == 'q' {
					quit <- true
					return
				} else if char, key, _ := keyboard.GetKey(); key == keyboard.KeyEsc || char == 'p' {
					go abrirNuevaTerminalYEjecutarComando("cat paginacion.txt")
				}
			}
		}
	}()

	<-quit

	borrarArchivo()
}
