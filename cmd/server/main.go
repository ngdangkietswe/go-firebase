/**
 * Author : ngdangkietswe
 * Since  : 10/17/2025
 */

package main

import "go-firebase/internal/app"

// @title Go Firebase API
// @version 1.0
// @description This is a sample server for a Go Firebase application.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email ngdangkietswe@yopmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:300
// @BasePath /api/v1
// @schemes http https
func main() {
	app.Start()
}
