"""
Low-Level Design: Parking Lot System

Requirements:
- Multiple floors, each with multiple spots
- Different vehicle types (Motorcycle, Car, Truck)
- Assign nearest available spot by vehicle type
- Track entry/exit
"""

from enum import Enum
from datetime import datetime


class VehicleType(Enum):
    MOTORCYCLE = 1
    CAR = 2
    TRUCK = 3


class Vehicle:
    def __init__(self, license_plate: str, vehicle_type: VehicleType):
        self.license_plate = license_plate
        self.vehicle_type = vehicle_type


class ParkingSpot:
    def __init__(self, spot_id: int, spot_type: VehicleType):
        self.spot_id = spot_id
        self.spot_type = spot_type
        self.vehicle: Vehicle | None = None

    @property
    def is_available(self) -> bool:
        return self.vehicle is None

    def park(self, vehicle: Vehicle) -> bool:
        if not self.is_available:
            return False
        if vehicle.vehicle_type != self.spot_type:
            return False
        self.vehicle = vehicle
        return True

    def remove(self) -> Vehicle | None:
        vehicle = self.vehicle
        self.vehicle = None
        return vehicle


class ParkingFloor:
    def __init__(self, floor_id: int, spots: list[ParkingSpot]):
        self.floor_id = floor_id
        self.spots = spots

    def find_available_spot(self, vehicle_type: VehicleType) -> ParkingSpot | None:
        for spot in self.spots:
            if spot.is_available and spot.spot_type == vehicle_type:
                return spot
        return None


class Ticket:
    def __init__(self, vehicle: Vehicle, spot: ParkingSpot, floor_id: int):
        self.vehicle = vehicle
        self.spot = spot
        self.floor_id = floor_id
        self.entry_time = datetime.now()
        self.exit_time: datetime | None = None


class ParkingLot:
    _instance = None

    def __init__(self, name: str, floors: list[ParkingFloor]):
        self.name = name
        self.floors = floors
        self.active_tickets: dict[str, Ticket] = {}

    def park_vehicle(self, vehicle: Vehicle) -> Ticket | None:
        for floor in self.floors:
            spot = floor.find_available_spot(vehicle.vehicle_type)
            if spot:
                spot.park(vehicle)
                ticket = Ticket(vehicle, spot, floor.floor_id)
                self.active_tickets[vehicle.license_plate] = ticket
                return ticket
        return None  # Lot is full

    def unpark_vehicle(self, license_plate: str) -> Ticket | None:
        ticket = self.active_tickets.pop(license_plate, None)
        if ticket:
            ticket.spot.remove()
            ticket.exit_time = datetime.now()
        return ticket

    def available_spots(self, vehicle_type: VehicleType) -> int:
        count = 0
        for floor in self.floors:
            for spot in floor.spots:
                if spot.is_available and spot.spot_type == vehicle_type:
                    count += 1
        return count


if __name__ == "__main__":
    # Setup: 1 floor, 2 car spots, 1 motorcycle spot
    spots = [
        ParkingSpot(1, VehicleType.CAR),
        ParkingSpot(2, VehicleType.CAR),
        ParkingSpot(3, VehicleType.MOTORCYCLE),
    ]
    floor = ParkingFloor(1, spots)
    lot = ParkingLot("Main Lot", [floor])

    # Test parking
    car1 = Vehicle("ABC-123", VehicleType.CAR)
    ticket1 = lot.park_vehicle(car1)
    assert ticket1 is not None
    assert ticket1.spot.spot_id == 1
    assert lot.available_spots(VehicleType.CAR) == 1

    car2 = Vehicle("DEF-456", VehicleType.CAR)
    ticket2 = lot.park_vehicle(car2)
    assert ticket2 is not None
    assert lot.available_spots(VehicleType.CAR) == 0

    # Lot full for cars
    car3 = Vehicle("GHI-789", VehicleType.CAR)
    ticket3 = lot.park_vehicle(car3)
    assert ticket3 is None

    # Motorcycle can still park
    bike = Vehicle("MOTO-1", VehicleType.MOTORCYCLE)
    ticket_bike = lot.park_vehicle(bike)
    assert ticket_bike is not None

    # Unpark
    returned = lot.unpark_vehicle("ABC-123")
    assert returned is not None
    assert returned.exit_time is not None
    assert lot.available_spots(VehicleType.CAR) == 1

    print("All tests passed!")
