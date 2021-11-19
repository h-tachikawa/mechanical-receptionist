#include "SRF05.h"

#define SPEED_OF_SOUND_METER_PER_SECOND 340 // 音速(340[m/s])

SRF05::SRF05(int trigPin, int echoPin) {
    _trig = trigPin;
    _echo = echoPin;

    pinMode(_trig, OUTPUT);
    pinMode(_echo, INPUT);
}

// destructor
SRF05::~SRF05(void) {
    // do nothing 
}

auto SRF05::distance(void) -> float {
    this->emitUltrasonicWaves();

    auto pulseDurationAsMicrosec = this->waitForUltrasonicWavesReflection();

    if (pulseDurationAsMicrosec <= 0) {
        return -1;
    }

    auto oneWayDurationAsMicrosec = pulseDurationAsMicrosec / 2;
    //                                               音速 = 340[m/s] = 34000[cm/s] = 0.034[cm/μs] 
    auto durationAsCm = oneWayDurationAsMicrosec * SPEED_OF_SOUND_METER_PER_SECOND * 100 / 1000000;
    return durationAsCm;
}

 auto SRF05::emitUltrasonicWaves(void) -> void {
    digitalWrite(_trig, LOW);
    delayMicroseconds(1);
    digitalWrite(_trig, HIGH);
    delayMicroseconds(11);
    digitalWrite(_trig, LOW);
}

auto SRF05::waitForUltrasonicWavesReflection(void) -> unsigned long {
    return pulseIn(_echo, HIGH); // パルス幅(_echo ピンが HIGH だった時間)をμ秒単位で返す
}