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

double getCurrentDuration() {
  digitalWrite(TRIG_PIN, LOW);
  delayMicroseconds(2);
  digitalWrite(TRIG_PIN, HIGH); //超音波を出力
  delayMicroseconds( 10 ); // 10マイクロ秒待つ
  digitalWrite(TRIG_PIN, LOW);  // 超音波を止める
  return pulseIn(ECHO_PIN, HIGH); //パルスが発生していた時間(μs)を返す
}

double calcDistanceCm(double duration) {
  duration = duration / 2; // duration は往復分の時間になっているので半分にする

  /**
   * 音速 ≒ 340[m/s]
   * 超音波の移動距離(m)=(ECHO の HIGH 時間(μs) × 超音波速度)
   * 超音波の速さ =　音速 ≒ 340[m/s] = 34000[cm/s] = 0.034[cm/μs]
   * 
   * 対象物との距離(cm) = ECHOの HIGH 時間(μs) * 340m * 100(m => cm に変換) / 1000000(s => μs に変換)
   * */
  distance = duration * 340 * 100 / 1000000;
  return distance;
}

void setup() {
  Serial.begin(9600);
  pinMode(ECHO_PIN, INPUT);
  pinMode(TRIG_PIN, OUTPUT);
  Ethernet.begin(mac, ip);

  delay(1000);
  Serial.println("connecting...");
}

void loop() {
  duration = getCurrentDuration();

  if (duration <= 0) {
    return;
  }

  distance = calcDistanceCm(duration);
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
