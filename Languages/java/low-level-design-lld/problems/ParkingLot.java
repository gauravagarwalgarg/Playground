import java.util.*;

/**
 * Low-Level Design: Parking Lot System
 * 
 * - Multiple floors with multiple spots
 * - Different vehicle types (Motorcycle, Car, Truck)
 * - Assign nearest available spot
 * - Track entry/exit
 */
public class ParkingLot {

    enum VehicleType { MOTORCYCLE, CAR, TRUCK }

    static class Vehicle {
        String licensePlate;
        VehicleType type;

        Vehicle(String licensePlate, VehicleType type) {
            this.licensePlate = licensePlate;
            this.type = type;
        }
    }

    static class ParkingSpot {
        int id;
        VehicleType spotType;
        Vehicle vehicle;

        ParkingSpot(int id, VehicleType spotType) {
            this.id = id;
            this.spotType = spotType;
        }

        boolean isAvailable() { return vehicle == null; }

        boolean park(Vehicle v) {
            if (!isAvailable() || v.type != spotType) return false;
            vehicle = v;
            return true;
        }

        Vehicle remove() {
            Vehicle v = vehicle;
            vehicle = null;
            return v;
        }
    }

    static class ParkingFloor {
        int floorId;
        List<ParkingSpot> spots;

        ParkingFloor(int floorId, List<ParkingSpot> spots) {
            this.floorId = floorId;
            this.spots = spots;
        }

        ParkingSpot findAvailable(VehicleType type) {
            for (ParkingSpot spot : spots) {
                if (spot.isAvailable() && spot.spotType == type) return spot;
            }
            return null;
        }
    }

    private String name;
    private List<ParkingFloor> floors;
    private Map<String, ParkingSpot> activeTickets = new HashMap<>();

    public ParkingLot(String name, List<ParkingFloor> floors) {
        this.name = name;
        this.floors = floors;
    }

    public ParkingSpot parkVehicle(Vehicle vehicle) {
        for (ParkingFloor floor : floors) {
            ParkingSpot spot = floor.findAvailable(vehicle.type);
            if (spot != null) {
                spot.park(vehicle);
                activeTickets.put(vehicle.licensePlate, spot);
                return spot;
            }
        }
        return null;
    }

    public Vehicle unparkVehicle(String licensePlate) {
        ParkingSpot spot = activeTickets.remove(licensePlate);
        if (spot != null) return spot.remove();
        return null;
    }

    public int availableSpots(VehicleType type) {
        int count = 0;
        for (ParkingFloor floor : floors) {
            for (ParkingSpot spot : floor.spots) {
                if (spot.isAvailable() && spot.spotType == type) count++;
            }
        }
        return count;
    }

    public static void main(String[] args) {
        List<ParkingSpot> spots = List.of(
            new ParkingSpot(1, VehicleType.CAR),
            new ParkingSpot(2, VehicleType.CAR),
            new ParkingSpot(3, VehicleType.MOTORCYCLE)
        );
        ParkingFloor floor = new ParkingFloor(1, new ArrayList<>(spots));
        ParkingLot lot = new ParkingLot("Main Lot", List.of(floor));

        // Park cars
        Vehicle car1 = new Vehicle("ABC-123", VehicleType.CAR);
        assert lot.parkVehicle(car1) != null : "Should park car1";
        assert lot.availableSpots(VehicleType.CAR) == 1 : "1 car spot left";

        Vehicle car2 = new Vehicle("DEF-456", VehicleType.CAR);
        assert lot.parkVehicle(car2) != null : "Should park car2";
        assert lot.availableSpots(VehicleType.CAR) == 0 : "No car spots left";

        // Full for cars
        Vehicle car3 = new Vehicle("GHI-789", VehicleType.CAR);
        assert lot.parkVehicle(car3) == null : "Should be full for cars";

        // Motorcycle can still park
        Vehicle bike = new Vehicle("MOTO-1", VehicleType.MOTORCYCLE);
        assert lot.parkVehicle(bike) != null : "Should park motorcycle";

        // Unpark
        Vehicle returned = lot.unparkVehicle("ABC-123");
        assert returned != null && returned.licensePlate.equals("ABC-123") : "Should return car1";
        assert lot.availableSpots(VehicleType.CAR) == 1 : "1 car spot available again";

        System.out.println("All tests passed!");
    }
}
