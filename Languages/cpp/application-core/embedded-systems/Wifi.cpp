#include <iostream>

class Wifi
{
	protected:
		int data;
		void Initialize();
	public:
		int getData() { return data; }
		Wifi(int val): data(val) {}	
};

void Wifi::Initialize()
{
	std::cout << "Initialization called" << std::endl;
}


int main()
{
	Wifi wifi(10);
	int data = wifi.getData();
	std::cout << "Data is : " << data << std::endl;

	//wifi.Initialize();

	return 0;
}
