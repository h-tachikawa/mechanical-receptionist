#ifndef SRF05_H
#define SRF05_H

#include <Arduino.h>

class SRF05 {
public:
    SRF05(int trigPin, int echoPing);
    ~SRF05();

    float distance(void);
private:
    uint8_t _trig;
    uint8_t _echo;

    void emitUltrasonicWaves(void);
    unsigned long waitForUltrasonicWaveReflection(void);
};

#endif // SRF05_H
