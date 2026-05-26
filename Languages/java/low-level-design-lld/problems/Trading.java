// A simplified low-latency trading engine in Java
// Showcasing concurrency, networking, and minimal GC pressure

import java.io.IOException;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.SocketException;
import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.atomic.AtomicBoolean;

public class LowLatencyTradingEngine {

    private static final int PORT = 9000;
    private static final AtomicBoolean running = new AtomicBoolean(true);
    private static final BlockingQueue<String> messageQueue = new ArrayBlockingQueue<>(1024);

    public static void main(String[] args) throws InterruptedException {
        Thread networkThread = new Thread(LowLatencyTradingEngine::networkListener);
        Thread tradingThread = new Thread(LowLatencyTradingEngine::tradingLogic);

        networkThread.start();
        tradingThread.start();

        Thread.sleep(30_000); // run for 30 seconds
        running.set(false);

        networkThread.join();
        tradingThread.join();

        System.out.println("Trading engine shutdown.");
    }

    private static void networkListener() {
        try (DatagramSocket socket = new DatagramSocket(PORT)) {
            byte[] buffer = new byte[1024];
            while (running.get()) {
                DatagramPacket packet = new DatagramPacket(buffer, buffer.length);
                socket.receive(packet);
                String message = new String(packet.getData(), 0, packet.getLength());
                messageQueue.offer(message); // lock-free insertion
            }
        } catch (SocketException e) {
            System.err.println("Socket exception: " + e.getMessage());
        } catch (IOException e) {
            System.err.println("IO exception: " + e.getMessage());
        }
    }

    private static void tradingLogic() {
        while (running.get() || !messageQueue.isEmpty()) {
            String message = messageQueue.poll();
            if (message != null) {
                switch (message.trim()) {
                    case "BUY":
                        System.out.println("Executing BUY order");
                        break;
                    case "SELL":
                        System.out.println("Executing SELL order");
                        break;
                }
            }
        }
    }
}
