package app

import (
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/blevesearch/bleve"
	"github.com/revel/revel"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	// BleveIndex global reference
	BleveIndex bleve.Index

	// PortalDB application database
	PortalDB *db.DB
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	//open a new index
	// mapping := bleve.NewIndexMapping()
	// bleveIndex, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	revel.ERROR.Println(err)
	// 	return
	// }
	// controllers.BleveIndex = bleveIndex\
	myDBDir := "/home/brow/Documents/PortalDB"
	BleveIndex, _ = bleve.Open("example.bleve")
	PortalDB, _ = db.OpenDB(myDBDir)
	if err := PortalDB.Create("Participants"); err != nil {
		revel.ERROR.Println(err)
	}
	if err := PortalDB.Create("Sponsors"); err != nil {
		revel.ERROR.Println(err)
	}

	sponsors := PortalDB.Use("Sponsors")
	if err := sponsors.Index([]string{"username"}); err != nil {
		revel.ERROR.Println(err)
	}

	// register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
