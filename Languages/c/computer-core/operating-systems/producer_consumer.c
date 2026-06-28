/**
 * Producer-Consumer Problem using POSIX threads
 * Demonstrates: mutex, condition variables, bounded buffer.
 */

#include <stdio.h>
#include <stdlib.h>
#include <pthread.h>
#include <assert.h>

#define BUFFER_SIZE 5
#define NUM_ITEMS 10

typedef struct {
    int buffer[BUFFER_SIZE];
    int count;
    int in;   // write index
    int out;  // read index
    pthread_mutex_t mutex;
    pthread_cond_t not_full;
    pthread_cond_t not_empty;
} BoundedBuffer;

BoundedBuffer bb;
int produced_count = 0;
int consumed_count = 0;
int consumed_items[NUM_ITEMS];

void buffer_init(BoundedBuffer* b) {
    b->count = 0;
    b->in = 0;
    b->out = 0;
    pthread_mutex_init(&b->mutex, NULL);
    pthread_cond_init(&b->not_full, NULL);
    pthread_cond_init(&b->not_empty, NULL);
}

void buffer_destroy(BoundedBuffer* b) {
    pthread_mutex_destroy(&b->mutex);
    pthread_cond_destroy(&b->not_full);
    pthread_cond_destroy(&b->not_empty);
}

void buffer_put(BoundedBuffer* b, int item) {
    pthread_mutex_lock(&b->mutex);
    while (b->count == BUFFER_SIZE) {
        pthread_cond_wait(&b->not_full, &b->mutex);
    }
    b->buffer[b->in] = item;
    b->in = (b->in + 1) % BUFFER_SIZE;
    b->count++;
    pthread_cond_signal(&b->not_empty);
    pthread_mutex_unlock(&b->mutex);
}

int buffer_get(BoundedBuffer* b) {
    pthread_mutex_lock(&b->mutex);
    while (b->count == 0) {
        pthread_cond_wait(&b->not_empty, &b->mutex);
    }
    int item = b->buffer[b->out];
    b->out = (b->out + 1) % BUFFER_SIZE;
    b->count--;
    pthread_cond_signal(&b->not_full);
    pthread_mutex_unlock(&b->mutex);
    return item;
}

void* producer(void* arg) {
    (void)arg;
    for (int i = 0; i < NUM_ITEMS; i++) {
        buffer_put(&bb, i);
        __sync_fetch_and_add(&produced_count, 1);
    }
    return NULL;
}

void* consumer(void* arg) {
    (void)arg;
    for (int i = 0; i < NUM_ITEMS; i++) {
        int item = buffer_get(&bb);
        consumed_items[__sync_fetch_and_add(&consumed_count, 1)] = item;
    }
    return NULL;
}

int main(void) {
    buffer_init(&bb);

    pthread_t prod_thread, cons_thread;
    pthread_create(&prod_thread, NULL, producer, NULL);
    pthread_create(&cons_thread, NULL, consumer, NULL);

    pthread_join(prod_thread, NULL);
    pthread_join(cons_thread, NULL);

    // Verify all items were produced and consumed
    assert(produced_count == NUM_ITEMS);
    assert(consumed_count == NUM_ITEMS);

    // Verify all items 0..NUM_ITEMS-1 were consumed (order may vary with multiple threads)
    int found[NUM_ITEMS] = {0};
    for (int i = 0; i < NUM_ITEMS; i++) {
        assert(consumed_items[i] >= 0 && consumed_items[i] < NUM_ITEMS);
        found[consumed_items[i]] = 1;
    }
    for (int i = 0; i < NUM_ITEMS; i++) {
        assert(found[i] == 1);
    }

    buffer_destroy(&bb);
    printf("All tests passed!\n");
    return 0;
}
