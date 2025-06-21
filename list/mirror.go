// Package list
// 预存储各镜像网站地址
package list

import "mirrorfetch/model"

var MirrorSites = []*model.MirrorSite{
	{
		Name:        "USTC Mirror",
		Description: "中国科学技术大学开源软件镜像 Life, love, linux",
		HomeUrl:     "https://mirrors.ustc.edu.cn/",
	},
	{
		Name:        "STJU Mirror",
		Description: "由 上海交通大学 Linux 用户组 (SJTUG) 维护",
		HomeUrl:     "https://mirror.sjtu.edu.cn/",
	},
	{
		Name:        "TUNA Mirror",
		Description: "本站由清华大学信息化技术中心支持创办，由清华大学 TUNA 协会运行维护。",
		HomeUrl:     "https://mirrors.tuna.tsinghua.edu.cn/",
	},
	{
		Name:        "NJU Mirror",
		Description: "南京大学开源镜像站",
		HomeUrl:     "https://mirrors.nju.edu.cn/",
	},
}

var MirrorSources = []*model.MirrorSource{
	{
		Name:        "Github",
		Description: "Where the world build software",
		HomeUrl:     "https://github.com/",
	},
	{
		Name:        "Pypi",
		Description: "pip packages",
		HomeUrl:     "https://www.pypi.org",
	},
	{
		Name:    "Docker Hub",
		HomeUrl: "https://hub.docker.com/",
	},
}
