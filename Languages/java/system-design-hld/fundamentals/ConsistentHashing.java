import java.nio.charset.StandardCharsets;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;
import java.util.*;

/**
 * Consistent Hashing Ring Implementation
 *
 * Consistent hashing distributes keys across nodes in a way that minimizes
 * remapping when nodes are added or removed. Each node gets multiple "virtual nodes"
 * on the ring for more even distribution.
 *
 * Key concepts:
 * - Hash ring: A circular space [0, 2^32) where both nodes and keys are mapped
 * - Virtual nodes: Each physical node maps to multiple positions on the ring
 * - Minimal disruption: Adding/removing a node only affects keys in its immediate range
 *
 * Run: javac ConsistentHashing.java && java -ea ConsistentHashing
 */
public class ConsistentHashing {

    private final TreeMap<Long, String> ring = new TreeMap<>();
    private final Map<String, Set<Long>> nodePositions = new HashMap<>();
    private final int virtualNodes;

    /**
     * @param virtualNodes Number of virtual nodes per physical node.
     *                     Higher values give better distribution but use more memory.
     */
    public ConsistentHashing(int virtualNodes) {
        this.virtualNodes = virtualNodes;
    }

    /**
     * Adds a node to the hash ring with its virtual nodes.
     */
    public void addNode(String node) {
        Set<Long> positions = new HashSet<>();
        for (int i = 0; i < virtualNodes; i++) {
            long hash = hash(node + "#VN" + i);
            ring.put(hash, node);
            positions.add(hash);
        }
        nodePositions.put(node, positions);
    }

    /**
     * Removes a node and all its virtual nodes from the ring.
     */
    public void removeNode(String node) {
        Set<Long> positions = nodePositions.remove(node);
        if (positions != null) {
            for (Long pos : positions) {
                ring.remove(pos);
            }
        }
    }

    /**
     * Gets the node responsible for a given key.
     * Walks clockwise on the ring from the key's hash position.
     */
    public String getNode(String key) {
        if (ring.isEmpty()) {
            return null;
        }
        long hash = hash(key);
        // Find the first node position >= key hash (clockwise walk)
        Map.Entry<Long, String> entry = ring.ceilingEntry(hash);
        // Wrap around to the first entry if we've gone past the end
        if (entry == null) {
            entry = ring.firstEntry();
        }
        return entry.getValue();
    }

    /**
     * Returns the number of physical nodes on the ring.
     */
    public int getNodeCount() {
        return nodePositions.size();
    }

    /**
     * SHA-256 based hash function mapped to a long value.
     * SHA-256 provides excellent distribution properties.
     */
    private long hash(String key) {
        try {
            MessageDigest md = MessageDigest.getInstance("SHA-256");
            byte[] digest = md.digest(key.getBytes(StandardCharsets.UTF_8));
            // Use first 8 bytes to construct a long (unsigned via mask)
            long h = 0;
            for (int i = 0; i < 8; i++) {
                h = (h << 8) | (digest[i] & 0xFF);
            }
            return h;
        } catch (NoSuchAlgorithmException e) {
            throw new RuntimeException("SHA-256 not available", e);
        }
    }

    // ──────────────────────────────────────────────────────────────
    // Tests
    // ──────────────────────────────────────────────────────────────

    public static void main(String[] args) {
        testEvenDistribution();
        testMinimalRemapping();
        testEdgeCases();
        System.out.println("\n✓ All consistent hashing tests passed!");
    }

    /**
     * Demonstrates that virtual nodes provide reasonably even distribution.
     */
    private static void testEvenDistribution() {
        System.out.println("=== Test: Even Distribution ===");
        ConsistentHashing ch = new ConsistentHashing(150);

        String[] nodes = {"NodeA", "NodeB", "NodeC", "NodeD"};
        for (String node : nodes) {
            ch.addNode(node);
        }

        // Distribute 10000 keys and count per node
        Map<String, Integer> distribution = new HashMap<>();
        int totalKeys = 10000;
        for (int i = 0; i < totalKeys; i++) {
            String assigned = ch.getNode("key-" + i);
            distribution.merge(assigned, 1, Integer::sum);
        }

        System.out.println("Distribution of " + totalKeys + " keys across " + nodes.length + " nodes:");
        double idealPercentage = 100.0 / nodes.length; // 25% each
        for (String node : nodes) {
            int count = distribution.getOrDefault(node, 0);
            double pct = (count * 100.0) / totalKeys;
            System.out.printf("  %s: %d keys (%.1f%%)%n", node, count, pct);
            // With 150 virtual nodes, each should get roughly 15-35% (generous bounds)
            assert count > 1000 : node + " has too few keys: " + count;
            assert count < 4000 : node + " has too many keys: " + count;
        }
        System.out.println("  ✓ Distribution within acceptable bounds (ideal: " + String.format("%.0f%%", idealPercentage) + " each)\n");
    }

    /**
     * Demonstrates minimal remapping when a node is added or removed.
     * Only keys in the affected range should move.
     */
    private static void testMinimalRemapping() {
        System.out.println("=== Test: Minimal Remapping ===");
        ConsistentHashing ch = new ConsistentHashing(150);

        String[] initialNodes = {"Server1", "Server2", "Server3"};
        for (String node : initialNodes) {
            ch.addNode(node);
        }

        // Record initial assignment
        int totalKeys = 10000;
        Map<String, String> initialAssignment = new HashMap<>();
        for (int i = 0; i < totalKeys; i++) {
            String key = "request-" + i;
            initialAssignment.put(key, ch.getNode(key));
        }

        // Add a new node
        ch.addNode("Server4");
        int remappedOnAdd = 0;
        for (int i = 0; i < totalKeys; i++) {
            String key = "request-" + i;
            if (!initialAssignment.get(key).equals(ch.getNode(key))) {
                remappedOnAdd++;
            }
        }

        double remapPercentAdd = (remappedOnAdd * 100.0) / totalKeys;
        System.out.printf("  After adding Server4: %d/%d keys remapped (%.1f%%)%n",
                remappedOnAdd, totalKeys, remapPercentAdd);
        // Ideal remapping is ~1/N of keys (1/4 = 25%). Allow generous range.
        assert remapPercentAdd < 40 : "Too many keys remapped on add: " + remapPercentAdd + "%";
        assert remapPercentAdd > 10 : "Too few keys remapped on add: " + remapPercentAdd + "%";
        System.out.println("  ✓ Remapping on add is close to ideal (~25%)");

        // Remove a node
        Map<String, String> beforeRemove = new HashMap<>();
        for (int i = 0; i < totalKeys; i++) {
            String key = "request-" + i;
            beforeRemove.put(key, ch.getNode(key));
        }

        ch.removeNode("Server2");
        int remappedOnRemove = 0;
        for (int i = 0; i < totalKeys; i++) {
            String key = "request-" + i;
            if (!beforeRemove.get(key).equals(ch.getNode(key))) {
                remappedOnRemove++;
            }
        }

        double remapPercentRemove = (remappedOnRemove * 100.0) / totalKeys;
        System.out.printf("  After removing Server2: %d/%d keys remapped (%.1f%%)%n",
                remappedOnRemove, totalKeys, remapPercentRemove);
        // Only keys that were on Server2 should move (~25%)
        assert remapPercentRemove < 40 : "Too many keys remapped on remove: " + remapPercentRemove + "%";
        assert remapPercentRemove > 10 : "Too few keys remapped on remove: " + remapPercentRemove + "%";
        System.out.println("  ✓ Remapping on remove is close to ideal (~25%)\n");
    }

    /**
     * Edge cases: empty ring, single node, duplicate add.
     */
    private static void testEdgeCases() {
        System.out.println("=== Test: Edge Cases ===");

        ConsistentHashing ch = new ConsistentHashing(10);

        // Empty ring
        assert ch.getNode("any-key") == null : "Expected null for empty ring";
        System.out.println("  ✓ Empty ring returns null");

        // Single node
        ch.addNode("OnlyNode");
        for (int i = 0; i < 100; i++) {
            assert "OnlyNode".equals(ch.getNode("key-" + i)) : "Single node should handle all keys";
        }
        System.out.println("  ✓ Single node handles all keys");

        // Remove non-existent node (should not crash)
        ch.removeNode("GhostNode");
        assert ch.getNodeCount() == 1 : "Node count should still be 1";
        System.out.println("  ✓ Removing non-existent node is safe");

        // Remove last node
        ch.removeNode("OnlyNode");
        assert ch.getNode("key") == null : "Expected null after removing all nodes";
        System.out.println("  ✓ Removing last node returns to empty state\n");
    }
}
