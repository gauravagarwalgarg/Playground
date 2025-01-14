#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

#define BUFFER_SIZE 5U

typedef struct
{
    int buffer[BUFFER_SIZE];
    int head;
    int tail;
    bool full;
} CircularBuffer;

void CircularBufferInit(CircularBuffer* cb)
{
    cb->head = 0;
    cb->tail = 0;
    cb->full = false;
}

bool CircularBufferIsFull(CircularBuffer* cb)
{
    return cb->full;
}

bool CircularBufferIsEmpty(CircularBuffer* cb)
{
    return (!cb->full && (cb->head == cb->tail));
}

bool CircularBufferPush(CircularBuffer* cb, int item)
{
    if(CircularBufferIsFull(cb))
    {
        return false;
    }

    cb->buffer[cb->head] = item;
    cb->head = (cb->head + 1) % BUFFER_SIZE;

    if(cb->head == cb->tail)
    {
        cb->full = true;
    }

    return true;
}

bool CircularBufferPop(CircularBuffer* cb, int* item)
{
    if(CircularBufferIsEmpty(cb))
    {
        return false;
    }

    *item = cb->buffer[cb->tail];
    cb->tail = (cb->tail + 1) % BUFFER_SIZE;

    cb->full = false;
    return true;
}

void CircularBufferPrint(CircularBuffer *cb) {
    printf("Buffer contents: ");
    
    if (CircularBufferIsEmpty(cb)) {
        printf("Buffer is empty.\n");
        return;
    }

    int index = cb->tail;
    printf("index: %d cb->head: %d cb->tail: %d", index, cb->head, cb->tail);
    while (index != cb->head) {
        printf("%d ", cb->buffer[index]);
        index = (index + 1) % BUFFER_SIZE; // Wrap around when reaching the end
    }

    printf("\n");
}


int main()
{
    CircularBuffer cb;
    CircularBufferInit(&cb);

    printf("Adding elements to the buffer...\n");
    for (int i = 1; i <= 6; i++) {
        if (CircularBufferPush(&cb, i)) {
            printf("Added %d to the buffer.\n", i);
        } else {
            printf("Buffer is full. Cannot add %d.\n", i);
        }
    }

    CircularBufferPrint(&cb);

    printf("\nRemoving elements from the buffer...\n");
    int item;
    while (CircularBufferPop(&cb, &item)) {
        printf("Removed %d from the buffer.\n", item);
    }
    printf("Buffer is empty.\n");

    return 0;    
}