# Example using goroutines for Concurrent API Calls

This Go program demonstrates the power of goroutines by making 36 concurrent third-party API calls. The primary objective is to showcase the efficiency gains achieved through parallelism compared to a sequential execution. Each API call fetches age-related data based on predefined names and countries.

## Overview

The program utilizes goroutines to concurrently fetch data from a third-party API (https://api.agify.io) for different combinations of names and countries. By making 36 parallel API calls, it provides a tangible comparison with the sequential execution, highlighting the potential speedup when leveraging concurrent programming.

## Execution

1. **Names and Countries:** The program is designed to fetch age-related data for a set of names (e.g., John, Jane, Chuck) across various countries (e.g., CA, US, GB, AU).

2. **Parallel Execution:** Goroutines are employed to initiate concurrent API calls for each combination of names and countries, significantly reducing the overall execution time.

3. **Sequential Comparison:** The program also measures the time it would take for the same API calls if executed sequentially, allowing for a clear comparison (just the sum of the individual times).

4. **Results Display:** The results, including names, countries, ages, and execution times, are presented in a formatted table to visualize the performance improvement achieved through goroutines.

Feel free to customize this description according to your specific needs. If you have any further requests or questions, let me know!

```markdown
+--------------------+---------+-----+------------+
| NAME               | COUNTRY | AGE | TIME TAKEN |
|:-------------------|:--------|:----|:-----------|
| Alice              | AU      |  40 | 325ms      |
| Alice              | CA      |  71 | 281ms      |
| Alice              | GB      |  53 | 281ms      |
| Alice              | US      |  64 | 326ms      |
| Ashley             | AU      |  46 | 326ms      |
| Ashley             | CA      |  40 | 325ms      |
| Ashley             | GB      |  39 | 325ms      |
| Ashley             | US      |  37 | 325ms      |
| Bob                | AU      |  73 | 325ms      |
| Bob                | CA      |  71 | 276ms      |
| Bob                | GB      |  69 | 325ms      |
| Bob                | US      |  70 | 319ms      |
| Chuck              | AU      |  54 | 319ms      |
| Chuck              | CA      |  73 | 325ms      |
| Chuck              | GB      |  74 | 325ms      |
| Chuck              | US      |  72 | 325ms      |
| Emma               | AU      |  42 | 325ms      |
| Emma               | CA      |  26 | 325ms      |
| Emma               | GB      |  39 | 325ms      |
| Emma               | US      |  37 | 319ms      |
| Jane               | AU      |  53 | 325ms      |
| Jane               | CA      |  61 | 326ms      |
| Jane               | GB      |  63 | 320ms      |
| Jane               | US      |  70 | 326ms      |
| Jessica            | AU      |  36 | 325ms      |
| Jessica            | CA      |  46 | 276ms      |
| Jessica            | GB      |  31 | 325ms      |
| Jessica            | US      |  39 | 325ms      |
| John               | AU      |  69 | 320ms      |
| John               | CA      |  70 | 281ms      |
| John               | GB      |  70 | 277ms      |
| John               | US      |  66 | 283ms      |
| Sue                | AU      |  76 | 319ms      |
| Sue                | CA      |  65 | 320ms      |
| Sue                | GB      |  67 | 295ms      |
| Sue                | US      |  72 | 295ms      |
| Time if Sequential |         |     | 11.285s    |
| Time when Parallel |         |     | 327ms      |
+--------------------+---------+-----+------------+
