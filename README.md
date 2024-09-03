# Hours Day Timewarrior extension

This [Timewarrior extension](https://timewarrior.net/docs/extensions/) calculates the total time registered for each day
in the interval printing:

* The time for each day
    * Whether the day was OVER or UNDER time respect to 8 hours
* The total days
* The total overtime
* The total undertime
* The actual overtime (total overtime - total undertime)

## Installation

1. Download the latest executable for your operating system from
   the [releases page](https://github.com/crossbone-magister/hours-day/releases).
2. Add it to the Timewarrior extension folder as described in the [documentation](https://timewarrior.net/docs/api/).
3. Verify that the extension is active and installed by running `timew extensions`.

## Usage

In a terminal window, run `timew hours-day`. An example output could be:

```bash
2024-01-01 - 8h0m0s
2024-01-02 - 8h20m0s - 20m0s OVER
Total days: 2
Total time: 16h20m0s
Total overtime: 20m0s
Total undertime: 0s
Actual overtime: 20m0s
```
