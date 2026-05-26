/**
 * Embedded Systems: GPIO Abstraction Layer
 *
 * Demonstrates a hardware abstraction layer (HAL) pattern for GPIO pins.
 * This is a simulation -- on real hardware, register addresses would be memory-mapped.
 */
#include <iostream>
#include <cstdint>
#include <cassert>
using namespace std;

enum class PinMode { INPUT, OUTPUT, ALTERNATE };
enum class PinState { LOW, HIGH };

// Simulated hardware register
static uint32_t GPIO_MODER = 0;
static uint32_t GPIO_ODR = 0;
static uint32_t GPIO_IDR = 0;

class GpioPin {
public:
    GpioPin(int pin) : pin_(pin) {}

    void setMode(PinMode mode) {
        uint32_t bits = static_cast<uint32_t>(mode);
        GPIO_MODER &= ~(0x3 << (pin_ * 2));       // clear
        GPIO_MODER |= (bits << (pin_ * 2));        // set
        mode_ = mode;
    }

    void write(PinState state) {
        if (mode_ != PinMode::OUTPUT) return;
        if (state == PinState::HIGH)
            GPIO_ODR |= (1 << pin_);
        else
            GPIO_ODR &= ~(1 << pin_);
    }

    PinState read() const {
        return (GPIO_IDR & (1 << pin_)) ? PinState::HIGH : PinState::LOW;
    }

    void toggle() {
        if (mode_ != PinMode::OUTPUT) return;
        GPIO_ODR ^= (1 << pin_);
    }

    bool isHigh() const {
        return (GPIO_ODR & (1 << pin_)) != 0;
    }

private:
    int pin_;
    PinMode mode_ = PinMode::INPUT;
};

int main() {
    GpioPin led(5);  // PA5 -- typical onboard LED

    led.setMode(PinMode::OUTPUT);
    assert((GPIO_MODER & (0x3 << 10)) == (1 << 10)); // Output mode = 01

    led.write(PinState::HIGH);
    assert(led.isHigh() == true);

    led.write(PinState::LOW);
    assert(led.isHigh() == false);

    led.toggle();
    assert(led.isHigh() == true);

    led.toggle();
    assert(led.isHigh() == false);

    cout << "All tests passed! GPIO abstraction works." << endl;
    return 0;
}
