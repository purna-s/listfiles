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
	err = filepath.Walk(loc, func(pathX string, infoX os.FileInfo, errX error) error {

		if errX != nil {
			fmt.Printf("error 「%v」 at a path 「%q」\n", errX, pathX)
			return errX
		}

		if infoX.IsDir() {
			fmt.Println("\n'", pathX, "'", " is a directory.\n")
		} else if subs == "Y" {
				fmt.Println("FileName", infoX.Name())
				fmt.Println("Directory", filepath.Dir(pathX))
				fmt.Println("Extension", filepath.Ext(pathX))
				fmt.Println("Size", infoX.Size())
				fmt.Println("ModTime", infoX.ModTime())
				fmt.Println("\n")
			} else if pathX == loc {
					fmt.Println("FileName", infoX.Name())
					fmt.Println("Directory", filepath.Dir(pathX))
					fmt.Println("Extension", filepath.Ext(pathX))
					fmt.Println("Size", infoX.Size())
					fmt.Println("ModTime", infoX.ModTime())
					fmt.Println("\n")
				}
	return nil
   })

	if err != nil {
		fmt.Println("error walking the path : \n", loc, err)
	}

	activityLog.Debugf("Activity has listed out the files Successfully")
	fmt.Println("Activity has listed out the files Successfully")
	
	return true, nil
}

