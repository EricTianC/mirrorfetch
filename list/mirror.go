// Package list
// 预存储各镜像网站地址
package list

import "mirrorfetch/model"

var Mirrors = []model.MirrorSite{USTC, STJU, TUNA}

var USTC = model.MirrorSite{
	Name:        "USTC Mirror",
	Description: "中国科学技术大学开源软件镜像 Life, love, linux",
	HomeUrl:     "https://mirrors.ustc.edu.cn/",
}

var STJU = model.MirrorSite{
	Name:        "STJU Mirror",
	Description: "由 上海交通大学 Linux 用户组 (SJTUG) 维护",
	HomeUrl:     "https://mirror.sjtu.edu.cn/",
}

var TUNA = model.MirrorSite{
	Name:        "TUNA Mirror",
	Description: "本站由清华大学信息化技术中心支持创办，由清华大学 TUNA 协会运行维护。",
	HomeUrl:     "https://mirrors.tuna.tsinghua.edu.cn/",
}
