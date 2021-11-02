#ifndef SRF05_H
#define SRF05_H

#include <Arduino.h>

class SRF05 {
public:
    SRF05(int trigPin, int echoPing);
    ~SRF05();

    auto distance(void) -> float;
private:
    uint8_t _trig;
    uint8_t _echo;

    void emitUltrasonicWaves(void);
    auto waitForUltrasonicWaveReflection(void) -> unsigned long;
};

#endif // SRF05_H
