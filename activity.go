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
	
	
		loc := ctx.GetInput("Path").(string)
		subs := ctx.GetInput("SubDirectories[Y/N]").(string)
	
	
	
	// the function that handles each file or dir
		ff = func(pathX string, infoX os.FileInfo, errX error) error {

		// first thing to do, check error. and decide what to do about it
		if errX != nil {
			fmt.Println("error at a path \n", errX, pathX)
			return errX
		}

		// find out if it's a dir or file, if file, print info
		if infoX.IsDir() {
			fmt.Println("\n'", pathX, "'", " is a directory.\n")
		} else if subs == "Y" {
				ctx.SetOutput("FileName", infoX.Name())
				ctx.SetOutput("Directory", filepath.Dir(pathX))
				ctx.SetOutput("Extension", filepath.Ext(pathX))
				ctx.SetOutput("Size", infoX.Size())
				ctx.SetOutput("ModTime", infoX.ModTime())
			} else {
				if filepath.Dir(pathX) == loc {
					ctx.SetOutput("FileName", infoX.Name())
					ctx.SetOutput("Directory", filepath.Dir(pathX))
					ctx.SetOutput("Extension", filepath.Ext(pathX))
					ctx.SetOutput("Size", infoX.Size())
					ctx.SetOutput("ModTime", infoX.ModTime())
					}
				}
		
		return nil
	}

	err = filepath.Walk(loc, ff)

	if err != nil {
		fmt.Println("error walking the path : \n", loc, err)
	}

	activityLog.Debugf("Activity has listed out the files Successfully")
	fmt.Println("Activity has listed out the files Successfully")
	
	return true, nil
}

