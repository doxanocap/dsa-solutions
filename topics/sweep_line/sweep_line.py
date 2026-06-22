from typing import List, Union
import heapq

def count_illuminated_benches(lamps, benches):
    events = []

    for x1, x2 in lamps:
        events.append((x1, 1))
        events.append((x2, -1))

    for b in benches:
        events.append((b, 0))

    # Step 3: Sort events (Primary: x-position, Secondary: start before end before bench)
    events.sort()

    active_lamps = 0
    illuminated_count = 0

    for x, event_type in events:
        if event_type == 1:
            active_lamps += 1
        elif event_type == -1:
            active_lamps -= 1
        else:
            if active_lamps > 0:  # Check if it's illuminated
                illuminated_count += 1

    return illuminated_count
