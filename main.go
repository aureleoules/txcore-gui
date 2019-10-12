package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/aureleoules/txcore"
)

func main() {
	app := app.New()

	w := app.NewWindow("TXCore GUI")
	w.CenterOnScreen()

	var utxo string
	var index int
	var publicKey string
	var private string
	var compressed bool
	var output string
	var amount int

	utxoEntry := widget.NewEntry()
	utxoEntry.OnChanged = func(content string) {
		utxo = content
	}
	utxoEntry.SetPlaceHolder("TxID")

	outputEntry := widget.NewEntry()
	outputEntry.OnChanged = func(out string) {
		output = out
	}
	outputEntry.SetPlaceHolder("Output Address")

	amountEntry := widget.NewEntry()
	amountEntry.SetPlaceHolder("Amount (Satoshis)")
	amountEntry.OnChanged = func(amnt string) {
		num, err := strconv.Atoi(amnt)
		if err != nil {
			log.Println(err)
			return
		}
		amount = num
	}

	privateEntry := widget.NewEntry()
	privateEntry.SetPlaceHolder("Private key")
	privateEntry.OnChanged = func(str string) {
		private = str
	}

	buildButton := widget.NewButton("Build TX", func() {

		tx := txcore.NewTX()
		tx.AddInput(utxo, publicKey, index, compressed)

		tx.AddOutput(output, amount)

		tx.Build()
		tx.Sign([]string{private})
		log.Println(tx.SignedTXHex)
	})

	publicEntry := widget.NewEntry()
	publicEntry.SetPlaceHolder("Public key")
	publicEntry.OnChanged = func(str string) {
		publicKey = str
	}

	compressedCheck := widget.NewCheck("Compressed", func(checked bool) {
		compressed = checked
	})

	indexEntry := widget.NewEntry()
	indexEntry.SetPlaceHolder("UTXO Index")
	indexEntry.OnChanged = func(str string) {
		i, err := strconv.Atoi(str)
		if err != nil {
			log.Println(err)
			return
		}
		index = i
	}

	utxoEntry.SetText("4a93938ed5ba3157fbf5d565482df19c9fdc9581cdf1508ce1110e1bccc5d82d")
	publicEntry.SetText("mz8NhsSzRXKx66GZRqf2a62iMBN6PqxbwH")
	indexEntry.SetText("0")
	index = 0
	privateEntry.SetText("cQcNmeNmiXysYJT2cGFxYqkh4a3TCniDa25SGvnJJvXmA8DtDJtF")
	compressedCheck.SetChecked(true)
	amountEntry.SetText("7231645")
	outputEntry.SetText("mz8NhsSzRXKx66GZRqf2a62iMBN6PqxbwH")

	inputGroup := widget.NewGroup("Input", utxoEntry, indexEntry, publicEntry, privateEntry, compressedCheck)

	outputGroup := widget.NewGroup("Output", outputEntry, amountEntry)

	w.SetContent(widget.NewVBox(
		widget.NewLabel("TXCore"),
		inputGroup,
		outputGroup,

		buildButton,
	))
	w.ShowAndRun()
}
