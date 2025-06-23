// Package main 用于检测各大网站及相应镜像站的脚本
package main

import (
	"fmt"
	"github.com/erictianc/mirrorfetch/data"
	"github.com/erictianc/mirrorfetch/model"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println("Starting MirrorFetch (http/https)")
	fmt.Println("Sources:")
	CheckNamedRemoteList(model.ToNamedRemoteList(data.MirrorSources))

	// 开始检查镜像站连接情况
	fmt.Println("Mirror Sites:")
	CheckNamedRemoteList(model.ToNamedRemoteList(data.MirrorSites))

	fmt.Println("Services:")

}

func CheckNamedRemoteList(remotes []model.HTTPRemote) {
	var (
		okStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).MarginRight(1)
		infoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F92F60")).MarginRight(1)
	)

	counts := len(remotes)
	sourcesResp := make(chan model.TestResult, counts)
	for _, source := range remotes {
		go func() {
			sourcesResp <- model.TouchHome(source)
		}()
	}

	for range counts {
		response := <-sourcesResp
		if response.Ok() {
			fmt.Println(okStyle.Render("✓"), response.GetName(),
				infoStyle.Render(response.ToInfo()))
		} else {
			fmt.Println(errorStyle.Render("✗"), response.GetName(),
				infoStyle.Render(fmt.Sprintf("error: %v", response.ToInfo())))
		}
	}
}
