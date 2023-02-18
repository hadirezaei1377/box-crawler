package main

func main() {
	// initializing database and router
	// register routes
	// initializing cron jobs
	//      get data
	//      process data and clean up data
	//      store data
}

/*


// The initial architecture of the project

func main() {

	// get config , Recall information about the config
	err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server port : ", config.AppConfig.Server.Port)

	// init server or create it
	server := echo.New()

	// server jobs

	   // know your routes
	   // set  middleware



	// routing
	routing.SetRouting(server)

	// start server or implement it
	server.Start(":" + config.AppConfig.Server.Port)
}

*/

/*
// runnig server by echo framework
func main() {
	// create a new server
	e := echo.New()

	// run this function when you recieve GET request
	e.GET("/", func(c echo.Context) error {
		print("Hi There!")
		return nil
	})

	// run server on localhost
	e.Start("localhost:3124")
}

*/

/*

// decoding config to GO code

// we want to put all config to a struct
type Config struct {
	Server struct {
		// `` is using for convert yaml to GO code
		Port string `yaml:"port"`
	} `yaml:"server"`
}

var AppConfig Config

// decode yml
// set config for running server at first of start by this function
func GetConfig() error {
	file, err := os.Open("config.yml")
	if err != nil {
		return err
	}
	defer file.Close()
	// decode yml and put it to our struct
	// it need file (config file)
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		return err
	}
	return nil
}
*/

/*
func SetRouting(e *echo.Echo) error {

	// What address did I get, what should I do?
	e.GET("/users/get/:id", func(c echo.Context) error {
		// send this response to client
		return c.String(http.StatusOK, "/users/get/   with param")
		// with param shows to client the using part
	})

	e.GET("/users/get/:id/avatars", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/[with param]/avatars")
	})

	// for any params that client sends :
	e.GET("/users/get/:id/avatars/*", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/[with param]/avatars/*")
	})

	e.GET("/users/get/15", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get/15")
	})

	// print routs
	for i, route := range e.Routes() {
		fmt.Println(i, route.Method, route.Path)
	}

	return nil

}

*/

/*

better routing

func SetRouting(e *echo.Echo) error {

	g := e.Group("users")

	g.GET("/getList", controller.GetUserList)
	g.POST("/CreateNewUser", controller.CreateNewUser)

	return nil
}

*/
