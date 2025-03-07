package utils

import (
	"runtime"
)

func PrintLogo() {
	switch runtime.GOOS {
	case "darwin":
		colorPrint()
	case "windows":
		textPrint()
	case "linux":
		colorPrint()
	}
}

func colorPrint() {
}

func textPrint() {
}
