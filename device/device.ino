#include <SPI.h>
#include <Ethernet2.h>
#include "SRF05.h"

#define ECHO_PIN 2
#define TRIG_PIN 3

// Arduino Ethernet Shield2 の Mac アドレス
byte mac[] = {
        0xA8, 0x61, 0x0A, 0xAE, 0x69, 0x72
};

IPAddress ip(192, 168, 1, 177);

SRF05 srf05(TRIG_PIN, ECHO_PIN);
EthernetClient client;
char server[] = "192.168.1.2";

unsigned int port = 8000;
unsigned int targetMinDistance = 30;

auto notifyToServer() -> boolean {
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

void setup() {
    Serial.begin(9600);
    Ethernet.begin(mac, ip);

    delay(1000);
    Serial.println("connecting...");
}

void loop() {
    int distance = srf05.distance();
    Serial.print("distance:");
    Serial.print(distance);
    Serial.println(" cm");

    boolean isHumanDetected = distance <= targetMinDistance;

    if (isHumanDetected) {
        notifyToServer();
        delay(3000);
        return;
    }

    delay(3000);
    return;
}
