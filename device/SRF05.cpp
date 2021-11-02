#include "SRF05.h"

#define SPEED_OF_SOUND_METER_PER_SECOND 340 // 音速

// Constructor
SRF05::SRF05(int trigPin, int echoPin) {
    _trig = trigPin;
    _echo = echoPin;

    pinMode(_trig, OUTPUT);
    pinMode(_echo, INPUT);
}

float SRF05::distance(void) {
    this->emitUltrasonicWaves();
    int durationAsMicrosec = this->waitForUltrasonicWaveReflection();

    if (durationAsMicrosec <= 0) {
        return -1;
    }

    float oneWayDurationAsMicrosec = durationAsMicrosec / 2;
    //                                               音速 = 340m/s = 34000cm/s = 0.034cm/μs 
    float durationAsCm = oneWayDurationAsMicrosec * (SPEED_OF_SOUND_METER_PER_SECOND * 100 / 1000000);
    return durationAsCm;
}

void SRF05::emitUltrasonicWaves(void) {
    digitalWrite(_trig, LOW);
    delayMicroseconds(1);
    digitalWrite(_trig, HIGH);
    delayMicroseconds(11);
    digitalWrite(_trig, LOW);
}

unsigned long SRF05::waitForUltrasonicWaveReflection(void) {
    return pulseIn(_echo, HIGH);
}