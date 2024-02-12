import re
import matplotlib.pyplot as plt

# matplotlib settings
plt.figure(figsize=(18, 12))

with open('data.txt', 'r') as file:
    data = file.readlines()

# Dictionary to store data for each test function
test_functions = {}

# Regular expression pattern to extract the benchmark number, X and Y values
pattern = r'Benchmark(\w+)/(\w+)_(\d+)-\d+\s+(\d+)\s+(\d+\.\d+)\sns/op'

for line in data:
    matches = re.match(pattern, line)
    if matches:
        benchmark_type = matches.group(1)
        test_function = matches.group(2)
        x_value = int(matches.group(3))
        y_value = float(matches.group(5))

        if test_function not in test_functions:
            test_functions[test_function] = {'x': [], 'y': []}

        test_functions[test_function]['x'].append(x_value)
        test_functions[test_function]['y'].append(y_value)

for test_function, data in test_functions.items():
    plt.plot(data['x'], data['y'], marker='o', label=test_function)

plt.xlabel('Key length')
plt.ylabel('Time ns/op')
plt.title('Benchmark Comparison')
plt.legend()
plt.grid(True)

plt.savefig('graph.png')