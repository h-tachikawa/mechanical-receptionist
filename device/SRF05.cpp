#include "SRF05.h"

#define SPEED_OF_SOUND_METER_PER_SECOND 340 // 音速

// Constructor
SRF05::SRF05(int trigPin, int echoPin) {
    _trig = trigPin;
    _echo = echoPin;
}

void SRF05::begin() {
    pinMode(_trig, OUTPUT);
    pinMode(_echo, INPUT);
}

float SRF05::distance(void) {
    int duration;
    float distance;

    this->trigger();
    duration = this->waitForPulseIn();

    if (duration <= 0) {
        return -1;
    }

    distance = duration / 2;
    distance = distance * SPEED_OF_SOUND_METER_PER_SECOND * 100 / 1000000; // 340m/s = 34000cm/s = 0.034cm/μs

    return distance;
}

void SRF05::trigger(void) {
    digitalWrite(_trig, LOW);
    delayMicroseconds(1);
    digitalWrite(_trig, HIGH);
    delayMicroseconds(11);
    digitalWrite(_trig, LOW);
}

unsigned long SRF05::waitForPulseIn(void) {
    return pulseIn(_echo, HIGH);
}