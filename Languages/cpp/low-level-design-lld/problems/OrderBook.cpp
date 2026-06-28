#include <iostream>
#include <map>
#include <vector>
#include <mutex>

class OrderBook {
    std::map<double, int, std::greater<double>> buyOrders; // Highest price first
    std::map<double, int> sellOrders;  // Lowest price first
    std::mutex bookMutex;

public:
    void addOrder(bool isBuy, double price, int quantity) {
        std::lock_guard<std::mutex> lock(bookMutex);

        if (isBuy) {
            // Try to match with lowest sell order
            auto it = sellOrders.begin();
            while (it != sellOrders.end() && it->first <= price && quantity > 0) {
                int matchQty = std::min(quantity, it->second);
                std::cout << "Matched Buy @ " << it->first << " Qty: " << matchQty << "\n";
                quantity -= matchQty;
                it->second -= matchQty;

                if (it->second == 0) sellOrders.erase(it++);
                else ++it;
            }
            if (quantity > 0) buyOrders[price] += quantity;  // Add to buy book
        } else {
            // Try to match with highest buy order
            auto it = buyOrders.begin();
            while (it != buyOrders.end() && it->first >= price && quantity > 0) {
                int matchQty = std::min(quantity, it->second);
                std::cout << "Matched Sell @ " << it->first << " Qty: " << matchQty << "\n";
                quantity -= matchQty;
                it->second -= matchQty;

                if (it->second == 0) buyOrders.erase(it++);
                else ++it;
            }
            if (quantity > 0) sellOrders[price] += quantity;  // Add to sell book
        }
    }

    void printBook() {
        std::cout << "Order Book:\n";
        for (const auto& [price, qty] : buyOrders) 
            std::cout << "Buy " << qty << " @ " << price << "\n";
        for (const auto& [price, qty] : sellOrders) 
            std::cout << "Sell " << qty << " @ " << price << "\n";
    }
};

int main() {
    OrderBook ob;
    ob.addOrder(true, 100.5, 10); // Buy 10 @ 100.5
    ob.addOrder(false, 100.5, 5); // Sell 5 @ 100.5 (Matches partially)
    ob.addOrder(false, 99.9, 8);  // Sell 8 @ 99.9
    ob.addOrder(true, 100.1, 12); // Buy 12 @ 100.1 (Matches fully with 99.9 sell)

    ob.printBook();
    return 0;
}

