# Aho-Corasick Benchmark

This benchmark was created to better compare the performance of the available Aho-Corasick implementations in Go.
The benchmark is based on https://github.com/BobuSumisu/aho-corasick-benchmark.

## Running the benchmark

To run the benchmark simply run the following command:
```bash
./run.sh
```

## Results

```
          name    patterns    input len       build     search    matches      alloc
    anknown         128985            6    169.11ms     0.00ms          7    0.21GiB
    bobusumisu      128985            6    168.80ms     0.00ms          7    0.58GiB
    cloudflare      128985            6    709.20ms     0.01ms          7    3.90GiB
    iohub           128985            6     61.97ms     0.02ms          7    0.07GiB
    tmikus          128985            6     81.57ms     0.00ms          2    0.01GiB

    anknown         128985           19    179.74ms     0.01ms         24    0.21GiB
    bobusumisu      128985           19    180.55ms     0.01ms         24    0.58GiB
    cloudflare      128985           19    703.03ms     0.01ms         23    3.90GiB
    iohub           128985           19     61.81ms     0.01ms         17    0.07GiB
    tmikus          128985           19     81.47ms     0.01ms          7    0.01GiB

    anknown         128985           41    182.01ms     0.01ms         49    0.21GiB
    bobusumisu      128985           41    183.87ms     0.02ms         49    0.58GiB
    cloudflare      128985           41    704.64ms     0.02ms         47    3.90GiB
    iohub           128985           41     71.14ms     0.02ms         45    0.07GiB
    tmikus          128985           41     78.97ms     0.01ms         11    0.01GiB

    anknown         128985           73    171.91ms     0.02ms         76    0.21GiB
    bobusumisu      128985           73    175.75ms     0.02ms         76    0.58GiB
    cloudflare      128985           73    740.54ms     0.02ms         67    3.90GiB
    iohub           128985           73     67.37ms     0.02ms         73    0.07GiB
    tmikus          128985           73     84.02ms     0.01ms         22    0.01GiB

    anknown         128985          146    172.23ms     0.03ms        153    0.21GiB
    bobusumisu      128985          146    176.35ms     0.03ms        153    0.58GiB
    cloudflare      128985          146    769.02ms     0.03ms        136    3.90GiB
    iohub           128985          146     64.03ms     0.02ms        146    0.07GiB
    tmikus          128985          146     80.28ms     0.01ms         43    0.01GiB

    anknown         128985          279    175.82ms     0.04ms        287    0.21GiB
    bobusumisu      128985          279    179.27ms     0.05ms        287    0.58GiB
    cloudflare      128985          279    742.38ms     0.04ms        234    3.90GiB
    iohub           128985          279     62.02ms     0.04ms        261    0.07GiB
    tmikus          128985          279     77.15ms     0.01ms         78    0.01GiB

    anknown         128985          534    182.78ms     0.08ms        495    0.21GiB
    bobusumisu      128985          534    184.74ms     0.09ms        495    0.58GiB
    cloudflare      128985          534    797.57ms     0.07ms        357    3.90GiB
    iohub           128985          534     64.96ms     0.07ms        445    0.07GiB
    tmikus          128985          534     79.72ms     0.02ms        137    0.01GiB

    anknown         128985         1118    172.93ms     0.13ms       1066    0.21GiB
    bobusumisu      128985         1118    183.38ms     0.13ms       1066    0.58GiB
    cloudflare      128985         1118    727.14ms     0.10ms        665    3.90GiB
    iohub           128985         1118     63.59ms     0.11ms        941    0.07GiB
    tmikus          128985         1118     83.46ms     0.02ms        299    0.01GiB

    anknown         128985         2233    174.15ms     0.28ms       2169    0.21GiB
    bobusumisu      128985         2233    181.40ms     0.27ms       2169    0.58GiB
    cloudflare      128985         2233    709.97ms     0.19ms       1225    3.90GiB
    iohub           128985         2233     63.11ms     0.21ms       1949    0.07GiB
    tmikus          128985         2233     83.00ms     0.04ms        617    0.01GiB

    anknown         128985         4428    175.41ms     0.44ms       4301    0.21GiB
    bobusumisu      128985         4428    179.99ms     0.48ms       4301    0.58GiB
    cloudflare      128985         4428    751.05ms     0.36ms       2055    3.90GiB
    iohub           128985         4428     63.95ms     0.40ms       3861    0.07GiB
    tmikus          128985         4428     80.45ms     0.06ms       1242    0.01GiB

    anknown         128985         8975    179.30ms     0.97ms       8718    0.21GiB
    bobusumisu      128985         8975    178.74ms     0.83ms       8718    0.58GiB
    cloudflare      128985         8975    780.95ms     0.67ms       3556    3.90GiB
    iohub           128985         8975     63.44ms     0.89ms       7784    0.07GiB
    tmikus          128985         8975     81.15ms     0.11ms       2518    0.01GiB

    anknown         128985        17837    180.66ms     1.78ms      17214    0.21GiB
    bobusumisu      128985        17837    178.35ms     1.47ms      17214    0.58GiB
    cloudflare      128985        17837    828.66ms     1.57ms       6141    3.90GiB
    iohub           128985        17837     65.05ms     1.88ms      15430    0.07GiB
    tmikus          128985        17837     83.62ms     0.20ms       4937    0.01GiB

    anknown         128985        35720    172.68ms     3.55ms      34232    0.22GiB
    bobusumisu      128985        35720    168.47ms     3.21ms      34232    0.58GiB
    cloudflare      128985        35720    719.86ms     2.34ms      10486    3.90GiB
    iohub           128985        35720     61.04ms     3.38ms      30701    0.08GiB
    tmikus          128985        35720     79.16ms     0.36ms       9877    0.01GiB

    anknown         128985        71332    173.67ms     7.08ms      68256    0.22GiB
    bobusumisu      128985        71332    175.75ms     6.70ms      68256    0.58GiB
    cloudflare      128985        71332    827.07ms     5.35ms      17431    3.90GiB
    iohub           128985        71332     61.02ms     6.79ms      61258    0.08GiB
    tmikus          128985        71332     77.40ms     0.69ms      19673    0.01GiB

    anknown         128985       142563    176.35ms    13.30ms     136922    0.22GiB
    bobusumisu      128985       142563    178.92ms    12.84ms     136922    0.59GiB
    cloudflare      128985       142563    745.33ms     8.48ms      28924    3.90GiB
    iohub           128985       142563     64.71ms    12.01ms     122868    0.09GiB
    tmikus          128985       142563     83.89ms     1.71ms      39564    0.01GiB

    anknown         128985       285779    181.23ms    25.57ms     274116    0.23GiB
    bobusumisu      128985       285779    173.83ms    26.18ms     274116    0.60GiB
    cloudflare      128985       285779    828.07ms    18.68ms      46529    3.90GiB
    iohub           128985       285779     68.43ms    25.65ms     246073    0.10GiB
    tmikus          128985       285779     81.18ms     3.02ms      79029    0.01GiB

    anknown         128985       571011    173.75ms    45.49ms     545790    0.26GiB
    bobusumisu      128985       571011    176.11ms    48.80ms     545790    0.63GiB
    cloudflare      128985       571011    740.43ms    27.74ms      70251    3.90GiB
    iohub           128985       571011     60.63ms    48.41ms     490157    0.13GiB
    tmikus          128985       571011     83.61ms     6.06ms     157663    0.01GiB
```
