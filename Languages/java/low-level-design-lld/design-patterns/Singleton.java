/**
 * Singleton Design Pattern
 * Thread-safe implementation using double-checked locking.
 */
public class Singleton {

    private static volatile Singleton instance;
    private final String connection;

    private Singleton() {
        this.connection = "PostgreSQL@localhost:5432";
    }

    public static Singleton getInstance() {
        if (instance == null) {
            synchronized (Singleton.class) {
                if (instance == null) {
                    instance = new Singleton();
                }
            }
        }
        return instance;
    }

    public String getConnection() {
        return connection;
    }

    public String query(String sql) {
        return "Executing: " + sql;
    }

    public static void main(String[] args) throws InterruptedException {
        Singleton s1 = Singleton.getInstance();
        Singleton s2 = Singleton.getInstance();

        assert s1 == s2 : "Should be same instance";
        assert s1.getConnection().equals("PostgreSQL@localhost:5432") : "Wrong connection";
        assert s1.query("SELECT 1").equals("Executing: SELECT 1") : "Query failed";

        // Thread safety test
        final Singleton[] instances = new Singleton[10];
        Thread[] threads = new Thread[10];
        for (int i = 0; i < 10; i++) {
            final int idx = i;
            threads[i] = new Thread(() -> instances[idx] = Singleton.getInstance());
            threads[i].start();
        }
        for (Thread t : threads) {
            t.join();
        }
        for (int i = 1; i < 10; i++) {
            assert instances[i] == instances[0] : "Thread safety violated";
        }

        System.out.println("All tests passed!");
    }
}
