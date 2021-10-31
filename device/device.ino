#include <SPI.h>
#include <Ethernet2.h>
#define ECHO_PIN 2
#define TRIG_PIN 3
#define MAX_DISTANCE 300

// Arduino Ethernet Shield2 の Mac アドレス
byte mac[] = {
  0xA8, 0x61, 0x0A, 0xAE, 0x69, 0x72
};

IPAddress ip(192, 168, 1, 177);

EthernetClient client;
char server[] = "192.168.1.2";

unsigned int port = 8000;
unsigned int targetMinDistance = 30;

double duration = 0; //受信した間隔
double distance = 0; //距離

boolean notifyToServer() {
  Serial.println("try to connect");

  if (client.connect(server, port)) {
    Serial.println("connected");
    client.println("POST /notify HTTP/1.1");
    client.println("Host: 192.168.1.2:8080");
    client.println("User-Agent: Arduino Post Client");
    client.println("Connection: close");
    client.println();
    client.stop();
    return true;
  }
  else {
    Serial.println("connection failed");
    client.stop();
    return false;
  }
}

double getCurrentDistance() {
  digitalWrite(TRIG_PIN, LOW);
  delayMicroseconds(2);
  digitalWrite(TRIG_PIN, HIGH); //超音波を出力
  delayMicroseconds( 10 ); //
  digitalWrite(TRIG_PIN, LOW); 
  return pulseIn(ECHO_PIN, HIGH); //センサからの入力
}

double calcDistance(double duration) {
  duration = duration / 2; //往復距離を半分にする
  distance = duration * 340 * 100 / 1000000; // 音速を340m/sに設定
  return distance;
}

void setup() {
  Serial.begin(9600);
  pinMode(ECHO_PIN, INPUT);
  pinMode(TRIG_PIN, OUTPUT);
  Ethernet.begin(mac, ip);

  while (!Serial) {
    ; // wait for serial port to connect. Needed for Leonardo only
  }

  delay(1000);
  Serial.println("connecting...");
}

void loop() {
  duration = getCurrentDistance();

  if (duration <= 0) {
    return;
  }

  distance = calcDistance(duration);
  Serial.print("distance:");
  Serial.print(distance);
  Serial.println(" cm");

  boolean isHumanDetected = distance > targetMinDistance;
  
  if (isHumanDetected) {
    delay(3000);
    Serial.println("return");
    return;
  }

  notifyToServer();

  delay(3000);
}
