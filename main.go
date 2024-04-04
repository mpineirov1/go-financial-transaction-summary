package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"go-financial-transaction-summary/models"
	"go-financial-transaction-summary/repository"
	"go-financial-transaction-summary/repository/entity"
	"go-financial-transaction-summary/utils"
	"log"
	"math"
	"os"
	"strconv"
	"text/template"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Calculate total debits and credits of an account
func main() {

	filename := "./txns.csv"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("Error: File '%s' not found.\n", filename)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}
	transactionsByMonth := make(map[int][]models.Transaction)
	var totalBalance float64
	for _, eachrecord := range records[1:] {
		if len(eachrecord) < 3 {
			log.Printf("Record with missing data: %v", eachrecord)
			continue // Skip to next record
		}
		id, _ := strconv.Atoi(eachrecord[0])
		date, _ := time.Parse("1/2", eachrecord[1]) // Formato de fecha MM/DD
		transaction, _ := strconv.ParseFloat(eachrecord[2], 64)
		totalBalance = transaction + totalBalance
		month := date.Format("01")
		monthNumber, _ := strconv.Atoi(month)
		transactionsByMonth[monthNumber] = append(transactionsByMonth[monthNumber], models.Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		})
	}
	summaryData := struct {
		TotalBalance float64
		MonthSummary map[string]models.MonthSummary
	}{
		TotalBalance: totalBalance,
		MonthSummary: make(map[string]models.MonthSummary), // Inicializar el mapa
	}

	for month, transactions := range transactionsByMonth {
		monthName := utils.GetMonthName(month)
		var monthCreditBalance, monthDebitBalance float64
		var monthCreditTotal, monthDebitTotal float64
		for _, t := range transactions {
			fmt.Println("Transaction", t.Transaction)
			if t.Transaction > 0 {
				monthCreditTotal++
				monthCreditBalance += t.Transaction
			} else {
				monthDebitTotal++
				monthDebitBalance += t.Transaction
			}
		}

		monthCreditAvg := monthCreditBalance / monthCreditTotal
		monthDebitAvg := monthDebitBalance / monthDebitTotal
		if math.IsNaN(monthDebitAvg) {
			monthDebitAvg = 0
		}
		if math.IsNaN(monthCreditAvg) {
			monthCreditAvg = 0
		}
		monthSummary := models.MonthSummary{
			TransactionNumber: len(transactions),
			DebitAvg:          monthDebitAvg,
			CreditAvg:         monthCreditAvg,
		}
		summaryData.MonthSummary[monthName] = monthSummary
	}
	var body bytes.Buffer

	tmpl := template.Must(template.ParseFiles("templates/mail.html"))
	tmpl.Execute(&body, summaryData)

	sendMail("mpineirov1@hotmail.com", "Summary", body.String())
}

func insertTransaction(ctx context.Context, repo repository.Transaction, transaction models.Transaction) error {

	transactionData := entity.Transaction{
		ID:          transaction.ID,
		Transaction: transaction.Transaction,
		Date:        transaction.Date,
		CreatedAt:   transaction.CreatedAt,
	}

	err := repo.Insert(ctx, &transactionData)

	return err
}

func sendMail(to, subject, body string) {

	from := goDotEnvVariable("MAIL_FROM_ADDRESS")
	username := goDotEnvVariable("MAIL_USERNAME")
	password := goDotEnvVariable("MAIL_PASSWORD")
	smtpHost := goDotEnvVariable("MAIL_HOST")
	smtpPort, _ := strconv.Atoi(goDotEnvVariable("MAIL_PORT"))

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, username, password)

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
	}
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
