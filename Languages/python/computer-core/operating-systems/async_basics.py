"""Asyncio fundamentals: coroutines, gather, create_task, and semaphore for rate limiting.

Demonstrates cooperative multitasking with Python's async/await syntax.
"""

import asyncio
import time

# --- Basic Coroutines ---

async def fetch_data(url: str, delay: float) -> str:
    """Simulate an async HTTP fetch."""
    print(f"  Fetching {url}...")
    await asyncio.sleep(delay)
    return f"Data from {url}"

# --- Gather: run coroutines concurrently ---

async def gather_demo() -> None:
    """Run multiple coroutines concurrently with gather."""
    urls = ["api/users", "api/posts", "api/comments"]
    start = time.perf_counter()

    results = await asyncio.gather(
        fetch_data(urls[0], 1.0),
        fetch_data(urls[1], 0.5),
        fetch_data(urls[2], 0.8),
    )

    elapsed = time.perf_counter() - start
    print(f"  Results: {results}")
    print(f"  Total time: {elapsed:.2f}s (parallel, not 2.3s sequential)\n")

# --- create_task: fire-and-forget with control ---

async def background_task(name: str) -> None:
    """A background task that runs independently."""
    print(f"  [{name}] Started")
    await asyncio.sleep(0.5)
    print(f"  [{name}] Completed")

async def task_demo() -> None:
    """Demonstrate create_task for concurrent execution."""
    task1 = asyncio.create_task(background_task("Task-A"))
    task2 = asyncio.create_task(background_task("Task-B"))

    print("  Main coroutine continues while tasks run...")
    await asyncio.sleep(0.2)
    print("  Main did some work")

    await task1
    await task2
    print()

# --- Semaphore: rate limiting concurrent operations ---

async def limited_fetch(sem: asyncio.Semaphore, url: str) -> str:
    """Fetch with semaphore-based rate limiting."""
    async with sem:
        print(f"  [SEM] Fetching {url} (slot acquired)")
        await asyncio.sleep(0.5)
        return f"Done: {url}"

async def semaphore_demo() -> None:
    """Limit concurrency to 3 simultaneous requests."""
    sem = asyncio.Semaphore(3)
    urls = [f"page/{i}" for i in range(8)]

    start = time.perf_counter()
    results = await asyncio.gather(*[limited_fetch(sem, url) for url in urls])
    elapsed = time.perf_counter() - start

    print(f"  Completed {len(results)} requests in {elapsed:.2f}s (max 3 concurrent)\n")

# --- Main ---

async def main() -> None:
    print("=== asyncio.gather ===")
    await gather_demo()

    print("=== asyncio.create_task ===")
    await task_demo()

    print("=== asyncio.Semaphore (rate limiting) ===")
    await semaphore_demo()

if __name__ == "__main__":
    asyncio.run(main())
