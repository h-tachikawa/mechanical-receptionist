#ifndef SRF05_H
#define SRF05_H

#include <Arduino.h>

class SRF05 {
    public:
        SRF05(int trigPin, int echoPing);
        virtual ~SRF05() {};

        void begin();
        float distance(void);
    private:
        uint8_t _trig;
        uint8_t _echo;

        void trigger(void);
        unsigned long waitForPulseIn(void);
};

#endif // SRF05_H