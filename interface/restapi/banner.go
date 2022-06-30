package restapi

import (
	"github.com/labstack/echo/v4"
	color "github.com/labstack/gommon/color"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/config"
)

const (
	// https://www.ascii-art-generator.org/ ( lean )
	myBanner = `
Welcome to:

#    # ####### ######  #######        #####  ####### ######  #     # ####### 
#   #  #     # #     # #             #     # #       #     # #     # #       
#  #   #     # #     # #             #       #       #     # #     # #       
###    #     # ######  #####   #####  #####  #####   ######  #     # #####   
#  #   #     # #   #   #                   # #       #   #    #   #  #       
#   #  #     # #    #  #             #     # #       #    #    # #   #       
#    # ####### #     # #######        #####  ####### #     #    #    ####### %s

Machine Learning Model Serving & Monitoring FrameWork
Based on Echo %s (https://echo.labstack.com/)

%s
%s
___________________________________________O/____________
                                    	   O\
`
)

// printSvrHeader print server header - banner
func printSvrHeader(e *echo.Echo, cfg *config.Config) {
	colorer := color.New()
	colorer.SetOutput(e.Logger.Output())
	colorer.Printf(myBanner,
		colorer.Cyan("v"+cfg.Applications.Servers.RestAPI.Version),
		colorer.Red("v"+echo.Version),
		colorer.Cyan(cfg.Applications.Servers.RestAPI.Name),
		colorer.Yellow(cfg.Applications.Servers.RestAPI.Description),
	)
}
