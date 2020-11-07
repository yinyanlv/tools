package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"

	myTime "tools/internal/time"

	"github.com/spf13/cobra"
)

var calculateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := myTime.GetNowTime()
		log.Printf("当前时间为：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		var layout = "2006-01-02 15:04:05"

		if calculateTime == "" {
			currentTime = myTime.GetNowTime()
		} else {
			var err error
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}

			if space == 1 {
				layout = "2006-01-02 15:04:05"
			}

			currentTime, err = time.Parse(layout, calculateTime)

			if err != nil {
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}

		t, err := myTime.GetCalculateTime(currentTime, duration)

		if err != nil {
			log.Fatalf("time.GetCalculateTime err: %v", err)
		}

		log.Printf("时间处理结果：%s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculate", "c", "", "需要计算的时间，有效数据为时间戳或已格式化后的时间")
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `时间间隔，有效单位为ns、us、ms、s、m、h`)
}
