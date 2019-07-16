package listfiles

import (

	"fmt"
    "os"
    "path/filepath"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// ActivityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-listfiles")

// MyActivity is a stub for your Activity implementation
type listfiles struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &listfiles{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *listfiles) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *listfiles) Eval(ctx activity.Context) (done bool, err error) {
	
	loc := ctx.GetInput("Loc").(string)
	//fmt.Println("Enter Location:")
	//var loc string
	//fmt.Scan(&loc)
    WalkAllFilesInDir(loc)	


func WalkAllFilesInDir(dir string) error {
    return filepath.Walk(dir, func(path string, info os.FileInfo, e error) error {
        if e != nil {
            return e
        }

        // check if it is a regular file (not dir)
        if info.Mode().IsRegular() {
			activityLog.Debugf("Activity has listed the files Successfully")
			fmt.Println("Activity has listed the files Successfully")
			
			ctx.SetOutput("fullName", path)

			ctx.SetOutput("fileName", info.Name())

			ctx.SetOutput("size", info.Size())

			ctx.SetOutput("lastModified", info.ModTime())

        }
        return nil
    })
}
}


