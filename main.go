package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := initializeModel()
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
	// fmt.Printf("%#x\n", Checksum(CRC16Algorithm[CRC_16_MODBUS], []byte("123456789asdfadsgqewg")))

}
