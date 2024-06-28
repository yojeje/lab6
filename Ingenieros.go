package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/yojeje/lab6"
	"google.golang.org/grpc"
)

var Ip = map[string]string{
	"Fulcrum1": "dist013",
	"Fulcrum2": "dist014",
	"Fulcrum3": "dist015",
	"Broker": "dist016",
}

func EscribirFichero(operacion string, base string, sector string, valor string) {
	file, err := os.Create(sector+".txt")
	if err != nil {
		fmt.Println("Error al crear eñ archivo", err)
		return
	}

	defer file.Close()

	data := operacion + " " + sector + " " + base

	if (valor != "") {
		data += " " + valor
	}

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(data)

	if err != nil {
		fmt.Println("Error al escribir", err)
		return
	}

	writer.Flush()
}

func AgregarBase(sector string, base string, valor string) {
	EscribirFichero("AgregarBase", base, sector, valor)
}

func RenombrarBase(sector string, base string, nombre string) {
	EscribirFichero("RenombrarBase", base, sector, nombre)
}

func ActualizarValor(sector string, base string, valor string) {
	EscribirFichero("ActualizarValor", sector, base, valor)
}

func BorrarBase(sector string, base string) {
	EscribirFichero("BorrarBase", sector, base, "")
}

func Consultar(m pb.IngenierosClient) {
	cmd := ""
	response, err := m.EnviarBroker(context.Background(), &pb.Comando{Tipo: cmd})

	if (err != nil) {
		log.Fatalf("Error %v", bufio.ErrBadReadCount)
	}

	fmt.Println(response.Dir)

}

func main() {
	// Establecer conexión con el servidor gRPC
	conn, err := grpc.Dial("dist098:50051", grpc.WithInsecure());
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	m := pb.NewIngenierosClient(conn)
	Consultar(m)
	defer conn.Close()
}