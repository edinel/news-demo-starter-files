/*
  Web client
 This sketch connects to a website (http://www.google.com)
 using the WiFi module.
 This example is written for a network using WPA encryption. For
 WEP or WPA, change the Wifi.begin() call accordingly.


 Circuit:
 * Board with NINA module (Arduino MKR WiFi 1010, MKR VIDOR 4000 and UNO WiFi Rev.2)
 created 13 July 2010
 by dlf (Metodo2 srl)
 modified 31 May 2012
 by Tom Igoe
 */

#include <SPI.h>
#include <WiFiNINA.h>
#include <ArduinoJson.h>
#include <NTPClient.h>
#include <WiFiUdp.h>


//#include "arduino_secrets.h"
///////please enter your sensitive data in the Secret tab/arduino_secrets.h

//int keyIndex = 0;            // your network key Index number (needed only for WEP)

int status = WL_IDLE_STATUS;

char server[] = "api.tidesandcurrents.noaa.gov";
String path = "/api/prod/datagetter?begin_date=20230328&end_date=20230404&station=9414358&product=predictions&datum=MLLW&time_zone=lst_ldt&interval=hilo&units=english&application=DataAPI_Sample&format=json";
String query = "GET "+ path + " HTTP/1.1";
WiFiClient client;
WiFiUDP ntpUDP;
NTPClient timeClient(ntpUDP, "2.north-america.pool.ntp.org", 3600, 60000);

void setup() {
  //Initialize serial and wait for port to open:
  Serial.begin(9600);
  
  while (!Serial) {
    ; // wait for serial port to connect. Needed for native USB port only
  }

  // Check for the WiFi module:
  if (WiFi.status() == WL_NO_MODULE) {
    Serial.println("Communication with WiFi module failed!");
    // don't continue
    while (true);
  }

// Make sure we're running the latest firmware
  String fv = WiFi.firmwareVersion();
  if (fv < WIFI_FIRMWARE_LATEST_VERSION) {
    Serial.println("Please upgrade the firmware");
  }

  // Connect to Wifi network:
  while (status != WL_CONNECTED) {
    Serial.print("Attempting to connect to WiFi: ");
    Serial.print("SSID is->>");
    Serial.print(ssid);
    Serial.println("<--");
    // Connect to WPA/WPA2 network. Change this line if using open or WEP network:
    status = WiFi.begin(ssid, pass);
    // wait 10 seconds for connection:
    delay(10000);
  }
  Serial.println("Connected to wifi");
  printWifiStatus();

  Serial.println("Starting a time Client");
  timeClient.begin();
}

void loop() {
  timeClient.update();
  Serial.println(timeClient.getFormattedTime());
  Serial.println("\nStarting connection to server...");
  // if you get a connection, report back via serial:
  if (client.connectSSL(server, 443)) {
    Serial.println("connected to server");
    // Make a HTTP request:
    client.println(query);
    client.println("Host: api.tidesandcurrents.noaa.gov");
    client.println("Connection: close");
    client.println();
    Serial.println(query);
    Serial.println("Host: api.tidesandcurrents.noaa.gov");
    Serial.println("Connection: close");
    Serial.println();
  }else{
    Serial.println ("AUGH FAIIIIL");  
    Serial.print ("Server was:");
    Serial.println (server);
    Serial.print ("query was:");
    Serial.println (query);
    Serial.println ("Host: api.tidesandcurrents.noaa.gov");
  }

  while (client.available()) {
    char c = client.read();
    Serial.write(c);
  }

  // if the server's disconnected, stop the client:

  if (!client.connected()) {
    Serial.println();
    Serial.println("disconnecting from server.");
    client.stop();
    // do nothing forevermore:
    while (true);
  }
}

void printWifiStatus() {
  // print the SSID of the network you're attached to:
  Serial.print("SSID: ");
  Serial.println(WiFi.SSID());
  // print your board's IP address:
  IPAddress ip = WiFi.localIP();
  Serial.print("IP Address: ");
  Serial.println(ip);
  // print the received signal strength:
  long rssi = WiFi.RSSI();
  Serial.print("signal strength (RSSI):");
  Serial.print(rssi);
  Serial.println(" dBm");
}
