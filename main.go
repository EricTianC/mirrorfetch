// Package main 用于检测各大网站及相应镜像站的脚本
package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"mirrorfetch/list"
	"mirrorfetch/model"
)

func main() {
	fmt.Println("Starting MirrorFetch (http/https)")
	fmt.Println("Sources:")
	CheckNamedRemoteList(model.ToNamedRemoteList(list.MirrorSources))

	// 开始检查镜像站连接情况
	fmt.Println("Mirror Sites:")
	CheckNamedRemoteList(model.ToNamedRemoteList(list.MirrorSites))
}

func CheckNamedRemoteList(remotes []model.NamedRemote) {
	var (
		okStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("#04B575")).MarginRight(1)
		infoStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
		errorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#F92F60")).MarginRight(1)
	)

	counts := len(remotes)
	sourcesResp := make(chan model.HttpTracesResponse, counts)
	for _, source := range remotes {
		go func() {
			sourcesResp <- model.TouchHome(source)
		}()
	}

	for range counts {
		select {
		case response := <-sourcesResp:
			if response.Reachable {
				fmt.Println(okStyle.Render("✓"), response.Name,
					infoStyle.Render(fmt.Sprintf("dns: %v conn: %v tls: %v total: %v",
						response.DNSDuration, response.TCPDuration, response.TLSDuration, response.TotalDuration)))
			} else {
				fmt.Println(errorStyle.Render("✗"), response.Name,
					infoStyle.Render(fmt.Sprintf("error: %v", response.ErrorMessage)))
			}
		}
	}
}
