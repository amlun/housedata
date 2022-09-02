package configs

var LJ = JobConfig{
	Name:     "链家",
	SavePath: "/tmp/",
	FileName: ".lj.daily.data",
	Regions: map[string]string{
		"chaoyang": "朝阳",
		"daxing":   "大兴",
		"xicheng":  "西城",
		"haidian":  "海淀",
	},
	URL:      "https://bj.lianjia.com/ershoufang/",
	URLExtra: "https://bj.lianjia.com/ershoufang/%s/",
	Selector: ".total span",
}
