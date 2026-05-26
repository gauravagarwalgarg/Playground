/**
 * GPIO Abstraction Layer
 * Demonstrates hardware abstraction for embedded systems.
 * Simulates GPIO pin control without actual hardware.
 */

#include <stdio.h>
#include <assert.h>
#include <stdint.h>
#include <string.h>

/* GPIO Direction */
typedef enum {
    GPIO_INPUT = 0,
    GPIO_OUTPUT = 1
} gpio_direction_t;

/* GPIO State */
typedef enum {
    GPIO_LOW = 0,
    GPIO_HIGH = 1
} gpio_state_t;

/* GPIO Pin Configuration */
typedef struct {
    uint8_t pin_number;
    gpio_direction_t direction;
    gpio_state_t state;
    int initialized;
} gpio_pin_t;

/* Simulated hardware register (32 pins max) */
static uint32_t gpio_register = 0;
static uint32_t gpio_direction_register = 0;

int gpio_init(gpio_pin_t* pin, uint8_t pin_number, gpio_direction_t direction) {
    if (pin_number > 31) return -1;

    pin->pin_number = pin_number;
    pin->direction = direction;
    pin->state = GPIO_LOW;
    pin->initialized = 1;

    if (direction == GPIO_OUTPUT) {
        gpio_direction_register |= (1U << pin_number);
    } else {
        gpio_direction_register &= ~(1U << pin_number);
    }
    return 0;
}

int gpio_write(gpio_pin_t* pin, gpio_state_t state) {
    if (!pin->initialized || pin->direction != GPIO_OUTPUT) return -1;

    pin->state = state;
    if (state == GPIO_HIGH) {
        gpio_register |= (1U << pin->pin_number);
    } else {
        gpio_register &= ~(1U << pin->pin_number);
    }
    return 0;
}

gpio_state_t gpio_read(gpio_pin_t* pin) {
    if (!pin->initialized) return GPIO_LOW;
    return (gpio_register >> pin->pin_number) & 1U ? GPIO_HIGH : GPIO_LOW;
}

int gpio_toggle(gpio_pin_t* pin) {
    if (!pin->initialized || pin->direction != GPIO_OUTPUT) return -1;
    gpio_state_t new_state = (pin->state == GPIO_HIGH) ? GPIO_LOW : GPIO_HIGH;
    return gpio_write(pin, new_state);
}

int main(void) {
    gpio_pin_t led_pin;
    gpio_pin_t button_pin;

    // Initialize LED as output
    assert(gpio_init(&led_pin, 13, GPIO_OUTPUT) == 0);
    assert(led_pin.direction == GPIO_OUTPUT);
    assert(led_pin.state == GPIO_LOW);

    // Write HIGH
    assert(gpio_write(&led_pin, GPIO_HIGH) == 0);
    assert(gpio_read(&led_pin) == GPIO_HIGH);

    // Toggle
    assert(gpio_toggle(&led_pin) == 0);
    assert(gpio_read(&led_pin) == GPIO_LOW);

    // Initialize button as input
    assert(gpio_init(&button_pin, 7, GPIO_INPUT) == 0);

    // Cannot write to input pin
    assert(gpio_write(&button_pin, GPIO_HIGH) == -1);

    // Invalid pin number
    gpio_pin_t invalid;
    assert(gpio_init(&invalid, 32, GPIO_OUTPUT) == -1);

    printf("All tests passed!\n");
    return 0;
}
