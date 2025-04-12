#include <stdio.h>
#include <pthread.h>
#include<unistd.h>

pthread_mutex_t mutex;
int shared_counter = 0;

void* increment(void* arg) {
	for(int i = 0; i < 5; i++) {
		pthread_mutex_lock(&mutex);
		shared_counter++;
		printf("Increment : %d", shared_counter);
		pthread_mutex_unlock(&mutex);
		sleep(1);
	}
	return NULL;
}

void* decrement(void* arg) {
	for(int i = 0; i < 5; i++) {
		pthread_mutex_lock(&mutex);
		shared_counter--;
		printf("Decrement : %d", shared_counter);
		pthread_mutex_unlock(&mutex);
		sleep(1);
	}
	return NULL;
}

int main() {
	pthread_t inc_thread, dec_thread;

	pthread_mutex_init(&mutex, NULL);

	pthread_create(&inc_thread, NULL, increment, NULL);
	pthread_create(&dec_thread, NULL, decrement, NULL);

	pthread_join(inc_thread, NULL);
	pthread_join(dec_thread, NULL);

	pthread_mutex_destroy(&mutex);
	
	return 0;
}
