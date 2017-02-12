package commands

import (
	"github.com/jfrogdev/jfrog-cli-go/bintray/utils"
	"github.com/jfrogdev/jfrog-cli-go/utils/io/fileutils"
	"github.com/jfrogdev/jfrog-cli-go/utils/cliutils/log"
)

func DownloadFile(pathDetails *utils.PathDetails, targetPath string, flags *utils.DownloadFlags) (err error) {
	fileutils.CreateTempDirPath()
	defer fileutils.RemoveTempDir()

	if flags.BintrayDetails.User == "" {
		flags.BintrayDetails.User = pathDetails.Subject
	}
	err = utils.DownloadBintrayFile(flags.BintrayDetails, pathDetails, targetPath, flags, "")
	if err != nil {
		return
	}
	log.Info("Downloaded 1 artifact.")
	return
}
