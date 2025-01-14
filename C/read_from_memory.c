#include <stdint.h>
#include <stddef.h>
#include <stdio.h>

/**
 * @brief Reads data from a specific memory address.
 * 
 * @param address The memory address to read from.
 * @param buffer Pointer to the buffer where the data will be stored.
 * @param length Number of bytes to read.
 * @return int 0 if successful, -1 if an error occurred.
 */
int read_from_memory(uintptr_t address, uint64_t *buffer, size_t length) {
    // Validate inputs
    if (buffer == NULL || length == 0) {
        return -1; // Invalid arguments
    }

    // Ensure the address is within a safe range (example check; adjust as needed)
    if (address < 0x10000000 || address > 0xFFFFFFFF) {
        return -1; // Address out of range
    }

    // Read data from the specified memory address
    for (size_t i = 0; i < length; i++) {
        buffer[i] = *((volatile uint8_t *)(address + i));
    }

    return 0; // Success
}


int main()
{
    uintptr_t target_address = 0xFFFFFFFF;
    uint64_t data[16];
    size_t length = sizeof(data);

    int status = read_from_memory(target_address, data, length);

    if (status == 0) {
        printf("Data read successfully: %ld\n", length);
        for (size_t i = 0; i < length; i++) {
            printf("0x%08X ", data[i]);
        }
        printf("\n");
    } else {
        printf("Failed to read from memory address 0x%08lX\n", target_address);
    }

    return 0;
}
