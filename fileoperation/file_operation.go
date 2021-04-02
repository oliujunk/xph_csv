package fileoperation

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"whxph.com/xph_csv/communication"
)

var (
	csvFile *os.File
)

// Start 文件操作
func Start() {
	for !communication.SyncRTC {
	}
	var err error
	csvFile, err = os.OpenFile("/mnt/usb/data.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		csvFile, err = os.OpenFile("data.csv", os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatalf("can not create file, err is %+v", err)
		}
	}
	_, _ = communication.Job.AddFunc("*/1 * * * * *", writeFile)
}

func writeFile() {
	_, _ = csvFile.Seek(0, io.SeekEnd)

	w := csv.NewWriter(csvFile)
	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	row := []string{time.Now().Format("2006-01-02 15:04:05"),
		fmt.Sprintf("%.2f", float32(communication.CurrentData.E1)/100.0),
		fmt.Sprintf("%d", communication.CurrentData.E2),
		fmt.Sprintf("%.1f", float32(communication.CurrentData.E3)/10.0),
		fmt.Sprintf("%.1f", float32(communication.CurrentData.E4)/10.0),
		fmt.Sprintf("%.1f", float32(communication.CurrentData.E5)/10.0),
	}
	err := w.Write(row)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}
	w.Flush()
}
