package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/storage"
	"github.com/revel/revel"
	"github.com/revel/revel/logger"
	"github.com/valyala/fasthttp"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
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
		revel.BeforeAfterFilter,       // Call the before and after filter functions
		revel.ActionInvoker,           // Invoke the action.
	}
	logger.LogFunctionMap["stdoutjson"] =
		func(c *logger.CompositeMultiHandler, options *logger.LogOptions) {
			// Set the json formatter to os.Stdout, replace any existing handlers for the level specified
			c.SetJson(os.Stdout, options)
		}
	revel.AddInitEventHandler(func(event revel.Event, i interface{}) revel.EventResponse {
		switch event {
		case revel.ENGINE_BEFORE_INITIALIZED:

			if revel.RunMode == "dev-fast" {
				revel.AddHTTPMux("/this/is/a/test", fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
					fmt.Fprintln(ctx, "Hi there, it worked", string(ctx.Path()))
					ctx.SetStatusCode(200)
				}))
				revel.AddHTTPMux("/this/is/", fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
					fmt.Fprintln(ctx, "Hi there, shorter prefix", string(ctx.Path()))
					ctx.SetStatusCode(200)
				}))
			} else {
				revel.AddHTTPMux("/this/is/a/test", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hi there, it worked", r.URL.Path)
					w.WriteHeader(200)
				}))
				revel.AddHTTPMux("/this/is/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, "Hi there, shorter prefix", r.URL.Path)
					w.WriteHeader(200)
				}))

			}
		}
		return 0
	})
	explicit("./Firestore_SA.json", "beawebsite-86b5d")
	explicit("./RealtimeDatabase_SA.json", "beawebsite-86b5d")

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
	// revel.OnAppStart(func() {
	// 	var fire = firebase.Config{ProjectID: "beawebsite-86b5d"}
	// 	sa := option.WithCredentialsFile("/Users/r.c.dawson/Desktop/Webdevelopment/beawebsite-86b5d-firebase-adminsdk-zxfqf-d811c41c09.json")
	// 	fs, err := firebase.NewApp(context.Background(), &fire, sa)
	// 	if err != nil {
	// 		log.Fatalf("error initializing app: %v\n", err)
	// 	}
	// 	client, err := fs.Firestore(context.Background())
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// })

}

func explicit(jsonPath, projectID string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(battrs.Name)
	}
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
