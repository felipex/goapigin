package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func servidor1(c *gin.Context) {
	c.JSON(200, gin.H{
		"siape": "123456",
		"nome":  "Gasparzinho",
	})
}
func servidor2(c *gin.Context) {
	c.JSON(200, gin.H{
		"siape": "234567",
		"nome":  "Assombroso",
	})
}
func servidor(c *gin.Context) {
	siape := c.Param("siape")
	s := getServidor(siape)
	c.JSON(200, gin.H{
		"siape": siape,
		"nome":  s,
	})
}

func main() {
	fmt.Println(getServidor("123456"))
	r := gin.Default()
	r.GET("/ping", ping)
	r.GET("/servidor/:siape", servidor)
	r.GET("/servidor/1", servidor1)
	r.GET("/servidor/2", servidor2)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func getServidor(siape string) string {
	ctx := context.Background()

	// Replace with your Google API key
	key := os.Getenv("GOOGLE_API_KEY")

	srv, err := sheets.NewService(ctx, option.WithAPIKey(key))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Replace with your spreadsheet ID and sheet name
	spreadsheetID := "1XSOiiRCMi_CfVZ6tjlZx_cTUf6zvjsxtJEM-AcyuRRo"
	//sheetName := "dados"

	// Read the range A1:B
	readRange := "A1:B"

	resp, err := srv.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
		return ""
	}

	s := "TESTE"
	for _, row := range resp.Values {
		if len(row) >= 2 {
			cellA, _ := row[0].(string)
			cellB, _ := row[1].(string)
			if cellA == siape {
				s = cellB
			}
			//fmt.Printf("Cell A: %s, Cell B: %s\n", cellA, cellB)
		}
	}
	return s
}
