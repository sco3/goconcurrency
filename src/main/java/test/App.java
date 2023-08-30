package test;

import static java.lang.String.format;
import static java.lang.System.currentTimeMillis;
import static java.lang.System.out;

import java.util.Random;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.atomic.AtomicLong;

public class App {

	public int[] generateNumbers(int max) {
		Random r = new Random();
		int[] numbers = new int[max];
		for (int i = 0; i < max; i++) {
			numbers[i] = r.nextInt();
		}
		return numbers;
	}

	public long add(int[] numbers) {
		long sum = 0;
		for (int i : numbers) {
			sum += i;
		}
		return sum;
	}

	public long addConcurrent(int[] numbers, int numOfCores) {
		AtomicLong sum = new AtomicLong();
		int max = numbers.length;
		int sizeOfParts = max / numOfCores;
		CountDownLatch latch = new CountDownLatch(numOfCores);

		for (int i = 0; i < numOfCores; i++) {
			int start = i * sizeOfParts;
			final int end;
			if (i + 1 == numOfCores) {
				end = numbers.length;
			} else {
				end = start + sizeOfParts;
			}
			Runnable runnable = new Runnable() {
				@Override
				public void run() {
					// out.println(
					// "cores " + numOfCores + " start " + start + " end < " + end
					// );
					long partSum = 0;
					for (int i = start; i < end; i++) {
						partSum += numbers[i];
					}
					sum.addAndGet(partSum);
					latch.countDown();
				}
			};
			Thread t = new Thread(runnable);
			t.setDaemon(true);
			t.start();
		}
		try {
			latch.await();
		} catch (InterruptedException e) {
			out.println("Interrupted" + e);
		}

		return sum.get();
	}

	public static void main(String[] args) {
		App app = new App();
		long start = currentTimeMillis();
		int n = (int) 5e8;
		int[] nums = app.generateNumbers(n);
		long millis = currentTimeMillis() - start;
		out.println(format("Generation of %d numbers took %d ms", n, millis));

		for (int i = 1; i <= 4; i++) {
			start = currentTimeMillis();
			long sum = app.addConcurrent(nums, i);
			millis = currentTimeMillis() - start;

			String msg = format( //
					"Threads: %d, sum: %d, millis: %d", i, sum, millis //
			);
			out.println(msg);
		}
	}
}
