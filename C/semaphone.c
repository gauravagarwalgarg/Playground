#include <stdio.h>
#include <pthread.h>
#include <semaphore.h>
#include <unistd.h>

sem_t semaphore;  // Declare semaphore

void* producer(void* arg) {
    while (1) {
        sleep(1);  // Simulate work
        printf("Producer: Producing item...\n");
        sem_post(&semaphore);  // Signal consumer
    }
    return NULL;
}

void* consumer(void* arg) {
    while (1) {
        sem_wait(&semaphore);  // Wait for producer signal
        printf("Consumer: Consuming item...\n");
    }
    return NULL;
}

int main() {
    pthread_t prod_thread, cons_thread;

    sem_init(&semaphore, 0, 0);  // Initialize semaphore with value 0

    pthread_create(&prod_thread, NULL, producer, NULL);
    pthread_create(&cons_thread, NULL, consumer, NULL);

    pthread_join(prod_thread, NULL);
    pthread_join(cons_thread, NULL);

    sem_destroy(&semaphore);  // Destroy semaphore
    return 0;
}

