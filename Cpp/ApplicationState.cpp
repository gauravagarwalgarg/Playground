#include <iostream>

class ApplicationState {
	private:
		static ApplicationState* instance;
		bool isLoggedIn;

		ApplicationState() : isLoggedIn(false) {}

	public:
		ApplicationState(const ApplicationState&) = delete;
		ApplicationState& operator=(const ApplicationState&) = delete;


		static ApplicationState* getAppState() {
			if(instance == nullptr)
			{
				instance = new ApplicationState();
			}
			return instance;
		}

		bool getLoginState() const {
			return isLoggedIn;
		}

};

int main () {
	ApplicationState* appState1 = ApplicationState::getAppState();

	std::cout << appState1->getLoginState() << std::endl;

}
