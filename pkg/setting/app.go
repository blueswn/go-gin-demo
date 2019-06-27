package setting

type App struct {
	/*
|--------------------------------------------------------------------------
| Application Name
|--------------------------------------------------------------------------
|
| This value is the name of your application. This value is used when the
| framework needs to place the application's name in a notification or
| any other location as required by the application or its packages.
*/
	Name string

	/*
|--------------------------------------------------------------------------
| Application Environment
|--------------------------------------------------------------------------
|
| This value determines the "environment" your application is currently
| running in. This may determine how you prefer to configure various
| services your application utilizes. Set this in your ".env" file.
|
*/
	Env string
	/*
|--------------------------------------------------------------------------
| Application Debug Mode
|--------------------------------------------------------------------------
|
| When your application is in debug mode, detailed error messages with
| stack traces will be shown on every error that occurs within your
| application. If disabled, a simple generic error page is shown.
|
*/
	Debug string

	/*
|--------------------------------------------------------------------------
| Application URL
|--------------------------------------------------------------------------
|
| This URL is used by the console to properly generate URLs when using
| the Artisan command line tool. You should set this to the root of
| your application so that it is used when running Artisan tasks.
|
*/
	Url string
	/*
	|--------------------------------------------------------------------------
	| Application Timezone
	|--------------------------------------------------------------------------
	|
	| Here you may specify the default timezone for your application, which
	| will be used by the PHP date and date-time functions. We have gone
	| ahead and set this to a sensible default for you out of the box.
	|
	*/

	Timezone string

	/*
 |--------------------------------------------------------------------------
 | Application Locale Configuration
 |--------------------------------------------------------------------------
 |
 | The application locale determines the default locale that will be used
 | by the translation service provider. You are free to set this value
 | to any of the locales which will be supported by the application.
 |
 */

	Locale string


	/*
	  |--------------------------------------------------------------------------
	  | Encryption Key
	  |--------------------------------------------------------------------------
	  |
	  | This key is used by the Illuminate encrypter service and should be set
	  | to a random, 32 character string, otherwise these encrypted strings
	  | will not be safe. Please do this before deploying an application!
	  |
	  */

	Key string

	Cipher string

	/*
	|--------------------------------------------------------------------------
	| Logging Configuration
	|--------------------------------------------------------------------------
	|
	| Here you may configure the log settings for your application. Out of
	| the box, Laravel uses the Monolog PHP logging library. This gives
	| you a variety of powerful log handlers / formatters to utilize.
	|
	| Available Settings: "single", "daily", "syslog", "errorlog"
	|
	*/

	Log string

	Log_level string
}

func NewDefaultAppSetting() *App {
	return &App{
		Cipher:"AES-256-CBC",
		Log:"daily",
		Log_level:"debug",
	}
}
