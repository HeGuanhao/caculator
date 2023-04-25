/**
* Author: HeGuanhao
* Email: 64777121@qq.com
* Date: 2023/4/25 9:51
 */

package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	display    *widget.Label
	num1, num2 float64
	operator   string
)

func main() {
	a := app.New()
	w := a.NewWindow("计算器")
	display = widget.NewLabel("0")

	btns := createButtons()

	w.SetContent(container.NewVBox(
		display,
		container.NewGridWrap(fyne.NewSize(80, 80), btns...),
	))

	w.ShowAndRun()
}

func createButtons() []fyne.CanvasObject {
	btns := make([]fyne.CanvasObject, 0)

	// 数字按钮
	for i := 0; i <= 9; i++ {
		num := i
		btn := widget.NewButton(strconv.Itoa(num), func() {
			updateDisplay(strconv.Itoa(num))
		})
		btns = append(btns, btn)
	}

	// 操作符按钮
	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		btn := widget.NewButton(op, func() {
			operator = op
			num1, _ = strconv.ParseFloat(display.Text, 64)
			updateDisplay("")
		})
		btns = append(btns, btn)
	}

	// 等号按钮
	equalBtn := widget.NewButton("=", func() {
		num2, _ = strconv.ParseFloat(display.Text, 64)
		result := calculateResult()
		updateDisplay(fmt.Sprintf("%.2f", result))
	})
	btns = append(btns, equalBtn)

	// 清空按钮
	clearBtn := widget.NewButton("C", func() {
		updateDisplay("")
	})
	btns = append(btns, clearBtn)

	return btns
}


func updateDisplay(value string) {
	display.SetText(value)
}

func calculateResult() float64 {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			return 0
		}
		return num1 / num2
	default:
		return 0
	}
}
