package datacenter

import (
	"bytes"
	"fmt"
	"gofly/utils/results"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 保存图片-tool
// 视频地址videoPath
// snapshotPath 图片保存路径
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		// log.Fatal("1生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		// log.Fatal("2生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		// log.Fatal("3保存封面图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "./")
	snapshotName = names[len(names)-1] + ".png"
	return
}

// 测试截取视频封面
func TestCover(context *gin.Context) {
	pathurl := "resource/uploads/20230103/test.mp4"
	videopath := fmt.Sprintf("./%s", pathurl)
	pathroot := strings.Split(pathurl, ".")
	imgpath := fmt.Sprintf("./%s", pathroot[0])
	fname, err := GetSnapshot(videopath, imgpath, 1)
	if err != nil {
		results.Failed(context, "截取视频封面失败", err.Error())
	} else {
		results.Success(context, "截取视频封面", fname, nil)
	}
}
