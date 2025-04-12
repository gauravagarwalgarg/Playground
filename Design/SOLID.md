# SOLID Princples

- More extensible code
- Easier to maintain code
- Reusable code

# Single Responsibility

- Classes should have a single purpose
- Every class/module should be responsible for one portion of the overall system.
- Avoid "spaghetti" code
- Allows you to implement proper separation of concerns
- Maintainability

``cpp
class CoffeeMachine {
		void pourCoffee() {
				std::cout << "Pouring coffee";
		}

		void sendCoffeeMetrics() {
				std::cout << "Sending metric";

		UrlRequest request;
		request.uri("/metrics");
		request.perform();
		}
}
```

# Open-closed

- Open for extension, closed for modification

# Liskov Substitution

- Behavioral class substitution

# Interface Segregation

- Favor multiple, specific interfaces over a single interface

# Dependency Inversion

- Depend on abstractions not implementations
