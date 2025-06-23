// Package data
// 预存储各镜像网站地址
package data

import "github.com/erictianc/mirrorfetch/model"

var MirrorSites = []*model.MirrorSite{
	{
		Name:        "USTC Mirror",
		Description: "中国科学技术大学开源软件镜像 Life, love, linux",
		HomeURL:     "https://mirrors.ustc.edu.cn/",
	},
	{
		Name:        "STJU Mirror",
		Description: "由 上海交通大学 Linux 用户组 (SJTUG) 维护",
		HomeURL:     "https://mirror.sjtu.edu.cn/",
		ServiceProviders: []model.ServiceProvider{
			{
				ID:          "Lean4/Elan",
				Description: "上交 s3 存储的 Elan 镜像",
				TestTarget: model.TestSpec{
					Sort: "s3",
					URL:  "https://mirrors.sjtug.sjtu.edu.cn/elan/?mirror_intel_list",
				},
			},
			{
				ID:          "Lean4/GithubRepo/Mathlib",
				Description: "Mathlib4 仓库镜像",
				TestTarget: model.TestSpec{
					Sort: "git",
					URL:  "https://mirror.sjtu.edu.cn/git/lean4-packages/mathlib4/",
				},
			},
		},
	},
	{
		Name:        "TUNA Mirror",
		Description: "本站由清华大学信息化技术中心支持创办，由清华大学 TUNA 协会运行维护。",
		HomeURL:     "https://mirrors.tuna.tsinghua.edu.cn/",
	},
	{
		Name:        "NJU Mirror",
		Description: "南京大学开源镜像站",
		HomeURL:     "https://mirrors.nju.edu.cn/",
	},
	{
		Name:    "Aliyun",
		HomeURL: "https://mirrors.aliyun.com/",
	},
	{
		Name:        "Tencent Cloud",
		Description: "腾讯云",
		HomeURL:     "https://mirrors.cloud.tencent.com/",
	},
	{
		Name:        "Huawei Clound",
		Description: "华为云",
		HomeURL:     "https://mirrors.huaweicloud.com/home/",
	},
	{
		Name:        "BFSU Mirror",
		Description: "北京外国语大学开源镜像站",
		HomeURL:     "https://mirrors.bfsu.edu.cn/",
	},
	{
		Name:        "SUSTech Mirror",
		Description: "南方科技大学开源镜像站",
		HomeURL:     "https://mirrors.sustech.edu.cn/",
	},
}

var MirrorSources = []*model.MirrorSource{
	{
		Name:        "Github",
		Description: "Where the world build software",
		HomeURL:     "https://github.com/",
	},
	{
		Name:        "Pypi",
		Description: "pip packages",
		HomeURL:     "https://www.pypi.org",
	},
	{
		Name:    "Docker Hub",
		HomeURL: "https://hub.docker.com/",
	},
	{
		Name:        "Go Proxy",
		Description: "Go 模块的依赖缓存服务器",
		HomeURL:     "https://goproxy.io",
	},
}
