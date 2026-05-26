/**
 * LLD Problem: Parking Lot System
 *
 * Design a parking lot with multiple levels, different vehicle sizes,
 * and basic park/unpark operations.
 */
#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>
#include <cassert>
using namespace std;

enum class VehicleSize { SMALL, MEDIUM, LARGE };

class Vehicle {
public:
    Vehicle(string plate, VehicleSize size) : plate_(plate), size_(size) {}
    string plate() const { return plate_; }
    VehicleSize size() const { return size_; }
private:
    string plate_;
    VehicleSize size_;
};

class ParkingSpot {
public:
    ParkingSpot(int id, VehicleSize size) : id_(id), size_(size), vehicle_(nullptr) {}

    bool canFit(const Vehicle& v) const {
        return !occupied() && v.size() <= size_;
    }
    bool occupied() const { return vehicle_ != nullptr; }
    void park(Vehicle* v) { vehicle_ = v; }
    void unpark() { vehicle_ = nullptr; }
    int id() const { return id_; }
private:
    int id_;
    VehicleSize size_;
    Vehicle* vehicle_;
};

class ParkingLot {
public:
    ParkingLot(int small, int medium, int large) {
        int id = 0;
        for (int i = 0; i < small; i++) spots_.emplace_back(id++, VehicleSize::SMALL);
        for (int i = 0; i < medium; i++) spots_.emplace_back(id++, VehicleSize::MEDIUM);
        for (int i = 0; i < large; i++) spots_.emplace_back(id++, VehicleSize::LARGE);
    }

    int park(Vehicle& v) {
        for (auto& spot : spots_) {
            if (spot.canFit(v)) {
                spot.park(&v);
                parked_[v.plate()] = spot.id();
                return spot.id();
            }
        }
        return -1; // lot full
    }

    bool unpark(const string& plate) {
        auto it = parked_.find(plate);
        if (it == parked_.end()) return false;
        spots_[it->second].unpark();
        parked_.erase(it);
        return true;
    }

    int available() const {
        int count = 0;
        for (const auto& s : spots_) if (!s.occupied()) count++;
        return count;
    }

private:
    vector<ParkingSpot> spots_;
    unordered_map<string, int> parked_;
};

int main() {
    ParkingLot lot(2, 3, 1); // 6 total spots
    assert(lot.available() == 6);

    Vehicle car1("MH-01-1234", VehicleSize::MEDIUM);
    Vehicle bike("MH-01-5678", VehicleSize::SMALL);
    Vehicle truck("MH-01-9999", VehicleSize::LARGE);

    int spot1 = lot.park(car1);
    assert(spot1 >= 0);
    assert(lot.available() == 5);

    lot.park(bike);
    lot.park(truck);
    assert(lot.available() == 3);

    assert(lot.unpark("MH-01-1234") == true);
    assert(lot.available() == 4);

    assert(lot.unpark("INVALID") == false);

    cout << "All tests passed!" << endl;
    return 0;
}
