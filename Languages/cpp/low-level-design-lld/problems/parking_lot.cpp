/**
 * LLD Problem: Parking Lot System
 *
 * Design a parking lot with:
 * - Vehicle hierarchy: Motorcycle, Car, Bus (polymorphism)
 * - ParkingSpot types: MOTORCYCLE, COMPACT, LARGE
 * - ParkingLot class that finds available spots by vehicle type
 *
 * Demonstrates: Inheritance, polymorphism, encapsulation, composition.
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <memory>
#include <cassert>
using namespace std;

// --- Vehicle Hierarchy ---
enum class VehicleType { MOTORCYCLE, CAR, BUS };

class Vehicle {
public:
    Vehicle(string plate, VehicleType type) : plate_(plate), type_(type) {}
    virtual ~Vehicle() = default;
    virtual string description() const = 0;
    string plate() const { return plate_; }
    VehicleType type() const { return type_; }
private:
    string plate_;
    VehicleType type_;
};

class Motorcycle : public Vehicle {
public:
    Motorcycle(string plate) : Vehicle(plate, VehicleType::MOTORCYCLE) {}
    string description() const override { return "Motorcycle [" + plate() + "]"; }
};

class Car : public Vehicle {
public:
    Car(string plate) : Vehicle(plate, VehicleType::CAR) {}
    string description() const override { return "Car [" + plate() + "]"; }
};

class Bus : public Vehicle {
public:
    Bus(string plate) : Vehicle(plate, VehicleType::BUS) {}
    string description() const override { return "Bus [" + plate() + "]"; }
};

// --- Parking Spot ---
enum class SpotType { MOTORCYCLE, COMPACT, LARGE };

class ParkingSpot {
public:
    ParkingSpot(int id, SpotType type) : id_(id), type_(type), vehicle_(nullptr) {}

    bool canFit(const Vehicle& v) const {
        if (occupied()) return false;
        switch (v.type()) {
            case VehicleType::MOTORCYCLE: return true;  // fits anywhere
            case VehicleType::CAR: return type_ == SpotType::COMPACT || type_ == SpotType::LARGE;
            case VehicleType::BUS: return type_ == SpotType::LARGE;
        }
        return false;
    }

    bool occupied() const { return vehicle_ != nullptr; }
    void park(Vehicle* v) { vehicle_ = v; }
    void unpark() { vehicle_ = nullptr; }
    int id() const { return id_; }
    SpotType type() const { return type_; }

private:
    int id_;
    SpotType type_;
    Vehicle* vehicle_;
};

// --- Parking Lot ---
class ParkingLot {
public:
    ParkingLot(int motorcycleSpots, int compactSpots, int largeSpots) {
        int id = 0;
        for (int i = 0; i < motorcycleSpots; i++) spots_.emplace_back(id++, SpotType::MOTORCYCLE);
        for (int i = 0; i < compactSpots; i++) spots_.emplace_back(id++, SpotType::COMPACT);
        for (int i = 0; i < largeSpots; i++) spots_.emplace_back(id++, SpotType::LARGE);
    }

    int park(Vehicle& v) {
        for (auto& spot : spots_) {
            if (spot.canFit(v)) {
                spot.park(&v);
                parked_[v.plate()] = spot.id();
                return spot.id();
            }
        }
        return -1; // no available spot
    }

    bool unpark(const string& plate) {
        auto it = parked_.find(plate);
        if (it == parked_.end()) return false;
        spots_[it->second].unpark();
        parked_.erase(it);
        return true;
    }

    int availableSpots(SpotType type) const {
        int count = 0;
        for (const auto& s : spots_)
            if (!s.occupied() && s.type() == type) count++;
        return count;
    }

    int totalAvailable() const {
        int count = 0;
        for (const auto& s : spots_) if (!s.occupied()) count++;
        return count;
    }

private:
    vector<ParkingSpot> spots_;
    unordered_map<string, int> parked_;
};

int main() {
    // 2 motorcycle spots, 3 compact spots, 1 large spot
    ParkingLot lot(2, 3, 1);
    assert(lot.totalAvailable() == 6);

    // Polymorphism: different vehicle types
    Motorcycle bike("MOTO-001");
    Car car1("CAR-001");
    Car car2("CAR-002");
    Bus bus("BUS-001");

    cout << "Parking: " << bike.description() << endl;
    assert(lot.park(bike) >= 0);

    cout << "Parking: " << car1.description() << endl;
    assert(lot.park(car1) >= 0);

    cout << "Parking: " << bus.description() << endl;
    assert(lot.park(bus) >= 0);
    assert(lot.availableSpots(SpotType::LARGE) == 0);

    // Bus can't park - no large spots left
    Bus bus2("BUS-002");
    assert(lot.park(bus2) == -1);

    // Unpark and repark
    assert(lot.unpark("BUS-001") == true);
    assert(lot.availableSpots(SpotType::LARGE) == 1);
    assert(lot.park(bus2) >= 0);

    assert(lot.unpark("INVALID") == false);

    cout << "All tests passed!" << endl;
    return 0;
}
