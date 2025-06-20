// ! 用于检测各大网站及相应镜像站的脚本
package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"mirrorfetch/list"
)

type CheckedResponse struct {
	Name string
	Ok   bool
}

func main() {
	mirrorCounts := len(list.Mirrors)

	responses := make(chan CheckedResponse, mirrorCounts)
	for _, mirror := range list.Mirrors {
		go func() {
			responses <- CheckedResponse{Name: mirror.Name, Ok: mirror.TouchHomePage()}
		}()
	}

	for range mirrorCounts {
		select {
		case response := <-responses:
			if response.Ok {
				var style = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575"))
				fmt.Println(response.Name, style.Render(" √"))
			} else {
				fmt.Println("❌", response.Name)
			}
		}
	}
}
