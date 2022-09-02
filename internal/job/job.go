package job

import (
	"encoding/json"
	"fmt"
	"github.com/amlun/housedata/configs"
	"github.com/amlun/housedata/internal/service"
	"github.com/amlun/housedata/pkg/xtime"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

//const fileName = ".lj.daily.data"

type DateResult map[string]RegionNumMap
type RegionNumMap map[string]int

func NewJob(config configs.JobConfig) *Job {
	log.Info("新任务：", config.Name)
	return &Job{
		config:   config,
		fileName: config.SavePath + config.FileName,
		result:   make(map[string]RegionNumMap, 0),
		logger: log.WithFields(log.Fields{
			"name": config.Name,
		}),
	}
}

type Job struct {
	// result
	result DateResult
	// config
	config configs.JobConfig
	// save result to file
	fileName string
	// logger
	logger *log.Entry
}

func (j *Job) Do() {
	// load data
	j.load()
	// do and update result
	j.do()
	// clear data and save
	j.clear()
}

func (j *Job) load() {
	j.logger.Info("读取数据")
	var err error
	iot, err := ioutil.ReadFile(j.fileName)
	if os.IsNotExist(err) {
		j.logger.Info("文件不存在")
		return
	}
	if err != nil {
		j.logger.Fatal(err)
	}
	_ = json.Unmarshal(iot, &j.result)
}

func (j *Job) do() {
	var (
		err error
		num int
	)
	j.logger.Info("开始执行抓取任务")
	day := xtime.Date{Time: time.Now()}.String()
	if _, ok := j.result[day]; !ok {
		j.result[day] = make(RegionNumMap, 0)
	}
	var wg sync.WaitGroup
	wg.Add(1)
	// 全市
	num, err = service.GetNum(j.config.URL, j.config.Selector)
	if err != nil {
		return
	}
	j.result[day]["all"] = num
	j.logger.WithField("region", "all").WithField("num", num).Info("获取在售房源数量")
	// 区域
	for k, v := range j.config.Regions {
		regionKey := k
		regionName := v
		num, err = service.GetNum(fmt.Sprintf(j.config.URLExtra, regionKey), j.config.Selector)
		if err != nil {
			continue
		}
		j.result[day][regionName] = num
		j.logger.WithField("region", regionName).WithField("num", num).Info("获取在售房源数量")
	}
}

func (j *Job) clear() {
	j.logger.Info("清除3个月以前的数据")
	before := xtime.Date{Time: time.Now().AddDate(0, -3, 0)}.String()
	delete(j.result, before)

	j.logger.Info("保存数据")
	bs, err := json.Marshal(j.result)
	if err != nil {
		j.logger.Error(err)
		return
	}
	if err := ioutil.WriteFile(j.fileName, bs, os.ModePerm); err != nil {
		j.logger.Error(err)
	}
}
