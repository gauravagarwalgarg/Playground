#include <set>
#include <queue>
#include <functional>

enum class OrderType { MARKET, LIMIT };
enum class Side { BUY, SELL };

class Order {
  protected:
    int id;
    Side side;
    double price;
    int quantity;
    OrderType type;

  public:
    Order(int id, Side s, double p, int q, OrderType t)
      : id(id), side(s), price(p), quantity(q), type(t) {}

    virtual ~Order() = default;

    int getId() const { return id; }
    Side getSide() const { return side; }
    double getPrice() const { return price; }
    int getQuantity() const { return quantity; }
    OrderType getType() const { return type; }

    virtual void execute() = 0;
};

class MarketOrder : public Order {
  public:
    MarketOrder(int id, Side s, int q) : Order(id, s, 0, q, OrderType::MARKET) {}
    void execute() override { /* Execution logic */ }
};

class LimitOrder : public Order {
  public:
    LimitOrder(int id, Side s, double p, int q) : Order(id, s, p, q, OrderType::LIMIT) {}
    void execute() override { /* Execution logic */ }
};

class OrderFactory {
  public:
    static std::unique_ptr<Order> createOrder(int id, Side side, double price, int quantity, OrderType type) {
      if(type == OrderType::MARKET)
        return std::make_unique<MarketOrder>(id, side, quantity);

      return std::make_unique<LimitOrder>(id, side, price, quantity);
    }
};


