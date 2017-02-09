#SendSms
A simple web application to send SMS using [burstsms](http://www.burstsms.com.au/) API

### Installation

`go get -u github.com/tabiul/SendSms`

### Requirements
* Bitly - Bitly is used for URL shortening
   * Username and password for [bitly](https://bitly.com) account. In addition you can register a application with bitly and obtain **Client ID** and **Client Secret** so that any API call will be logged under your application
   * BurstSMS - BurstSMS is used to send the SMS message
     * **Client ID** and **Client Secret** from [burstsms](http://www.burstsms.com.au/)
      
### Running
   * Set the following environment variables
     * BITLY_USERNAME  - Username for bitly account (required)
     * BITLY_PASSWORD - Password for bitly account (required)
 	 * BITLY_CLIENTID - Bitly registered Application Client ID (optional)
	 * BITLY_CLIENTSECRET - BITLY registered Application Client Secret (optional)
	 * BURSTSMS_CLIENTID - Burst SMS Client ID (required)
	 * BURSTSMS_CLIENTSECRET - Burst SMS Client Secret (required)
	 
   * Navigate to the `bin` folder after performing `go get`. The `bin` can be found in `$GOPATH` directory
   * Start the application **SendSms** as below

     `./SendSms -webapp ../src/github.com/tabiul/SendSms/webapp` (linux)

     `SendSms.exe -webapp ..\src\github.com\tabiul\SendSms\webapp` (windows)

   * Launch browser and type `localhost:8080`
              

