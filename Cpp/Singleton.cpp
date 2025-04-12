#include <iostream>
#include <mutex>

class Singleton {
	private:
		Singleton() {
			std::cout << "Singleton instance created\n";
		}

	public:
	Singleton(const Singleton&) = delete;
	Singleton& operator=(const Singleton&) = delete;

	static Singleton& getInstance() {
		static Singleton instance;
		return instance;
	}
};

int main() {
	Singleton& s1 = Singleton::getInstance();
	Singleton& s2 = Singleton::getInstance();

	if(&s1 == &s2) {
		std::cout << "Both instance are same\n" << std::endl;
	}

	return 0;
}

