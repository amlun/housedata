package configs

var IJ = JobConfig{
	Name:     "我爱我家",
	SavePath: "/tmp/",
	FileName: ".ij.daily.data",
	Regions: map[string]string{
		"chaoyangqu": "朝阳",
		"daxingqu":   "大兴",
		"xichengqu":  "西城",
		"haidianqu":  "海淀",
	},
	URL:      "https://bj.5i5j.com/ershoufang/",
	URLExtra: "https://bj.5i5j.com/ershoufang/%s/",
	Selector: ".total-box span",
}
