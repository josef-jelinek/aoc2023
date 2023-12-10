# Advent of Code 2023 - Go Solutions
https://adventofcode.com/2023

Self-imposed constraints:
 - No regular expressions unless the problem is targetted to need them (unlikely).
 - Just standard built-in packages (including new "generic" slices etc.)
 - No common helper top level functions, each problem is self-contained inside a single function (can use inner function closures). This prevents making solutions seemingly simpler while possibly doing extra work to fit the helper functions (having some map/filter/reduce, []string -> []int etc. may be tempting, but can obscure the real complexity of the solution).
 - No going back and refactoring of the previous solutions based on things learned in the newer solutions. The solutions are a snapshot in time and may be possible to use to track any improvements in the code consistency / style / simplicity.
 - No AI/LLM.
 - No hints.

The solutions are tailored for the particular input format and assume its validity, so all input errors are ignored. Also the solution is not meant to handle any edge cases which are not present in the problem input, especially, if it complicates the solution.

Solution notes for each day:

### Day 1
Part 1 is straightforward when keeping first digit when found and overriding last digit every time a digit is found. This handles a situation when only one digit is present, as the first and last ones will be set at the same time.

Part 2 adds a need to map words to digits. Instead of matching agains some map keys (and traversing the map each time), a simple switch with checking for a prefix can do.

### Day 2
Part 1 is quite straightforward with multiple nested loops corresponding to the nested problem notation. When a configuration inconsistent with input constrains is found, "labeled" continue is used to terminate several loop levels.

Part 2 is even more straigtforward as it processes all the input data and does not have additional hardcoded constraints given. It just keeps maximums for colors per line (even though they are called minimum as they are the minimum needed).

### Day 3
Part 1 goes through each line and keeps a flag whether it is inside a number. When it enters a number it checks the three-character column to the left for a symbol and on each digit it keeps checking one character above and one below. When it exits a number it checks a column to the right. To complete the checks for a number even if it is all the way to the right, the column index goes from 0 to line length including the length and all accesses to the characters are guarded from out of bounds.

Each time a word ends, if there was a symbol detected the value is added to the total sum.

Part 2 is a significant change to the problem spec. Instead of detecting a symbol and updating a sum, we detect a gear and update a global map from gear position to a list of found numbers.
We then go over the map and include in the final result only items, which recorded exactly two numbers.

### Day 4
Part 1 is straightforward and just iterates through the numbers we got and checks against the numbers we want and accumulates the resulting sum.

Part 2 sounds more complicated and there may be impulses to overcomplicate the solution, but there is not much more needed compared to part 1 except for a list counting instances of the "won" tickets. Since we are going from top to bottom and never update already processed tickets, we just update the tickets below the current winning one by the current ticket count plus one for the current win. Then the final sum is just summing all the numbers in the list plus one for the original ticket.

### Day 5
Part 1 needs preprocessing of the input to have mapping ranges ready and can be done relatively generally using <src>-to-<dst> pattern for the map name. Then we just start with a seed and "seed" as the first "map" key to resolve and ends when the key becomes "location". Instead of maps, slices are used as the amount of maps is limited to justify not using maps. For each iteration key is resolved using by searching for the right range.

Part 2 is not as straightforward as a naive reuse of the part 1 is sloooow. One observation which helps is that if we resolve a range, the minimum is always the lower bound of the range as all the range values are increasing. So instead of iterating through each seed value one by one, when we resolve a seed to a location, we know that we can safely skip at least as many following seed values as is the shortest remaining range from all the ranges we went through. This optimization is enough to make the solution computation instant.

The implemented solution is only correct, because the initial seeds fall into intervals in such a way that we can skip until the interval ends and do not deal with unmapped IDs. If there were many IDs which do not fall into the explicit ranges, we would have to create "virtual ranges", which map ID->ID between the ranges specified. Fortunately the problem input was well covered by the explicit ranges.

### Day 6
Part 1 is straightforward as there is a very simple formula for the distance from the time `d = i * (t - i)` which is a nice upside-down parabola and symetric on the range `t`.
So we can go only to the half of the interval and if we find a bound where we beat the record, we know where this ends and do not need to iterate over the whole interval.
One new library function used is `strings.Fields` which handles multiple-whitespace separators simplifying the table parsing.

Part 2 is even simpler when using the observation above. Just running the identical solution on the "new" input gave result instantly. If it was not fast enough, I was considering binary search on the range halve as it is monotonic (upside-down half-parabola), which would probably be fast enough. But it seems like it could have been even faster using Newton's method modified for discrete case, which in theory is similar to binary search with biased "midpoint" computation.

### Day 7
Part 1 boils down to write a function to compare two hands. Then it can be used to sort the hand-bet entries to get the rank and sum all bets multiplied by their rank. Comparing hands has two parts, comparing sorted counts of card kinds in the hand and if the same just comparing the cards with card kinds mapped to values (index in an array in this case). Using `strings.Count(s, s[:1])` and `strings.ReplaceAll(s, s[:1], "")` until the string is empty gives the counts, which need to be sorted and compared.

Part 2 is almost identical with the card kinds updated to move `J` to the end and before the strings is processed to get the kind counts, `J` kind is counted and removed using the mentioned functions and then added to the highest count found. There is one edge case, where there are only `J` cards in a hand as so this returns `[5]` for counts right away.

### Day 8
Part 1 is again straightforward just having a map from a node name to a pair of node names for `L` and `R` input and traversing while cycling through `L`/`R` directions. End traversing on `ZZZ`.

Part 2 was the biggest increase in difficulty so far and no naive solution would seem to cut it. The solution also maps string symbols to integers for potentially faster operations, but the solution implemented does not really need that step. It only marginally affects the end solution speed as the string IDs are short.

The main observations in this problem were that the paths must be traversed many times in cycles until they all end on the wanted end nodes. These cycles can be easily extracted for individual paths corresponding to all start nodes, however, the cycles are not determined only by the node but also by the position in the direction string. The other observation was that computing a common number of steps to the end state for two paths with the computed end state offsets and periods is fast and ends up with a different end state offset and period (the period is the least common multiple of the two path periods amd the end state offset is the common end state found).

Finding the common end state for two paths can be done fast enough by starting with the two individual end states and adding the corresponding loop period to the smaller one until they are equal (I did not bother to come up with any formula, as it was not needed for the solution). The total common end state (number of steps) is just repeating the process for two with the first path from the previous common end state operation.

When retrieving the end state / period info for each start state, there was only a single end state for each case and the solution took advantage of that. If that was not the case, the problem would get more complicated, and fortunately, it was not needed to be addressed to get the right solution.

### Day 9
Part 1 and Part 2 are basically identical and straightforward when following the examples showing the solution. They construct exacly the same rows of differences. The only change is that instead of accumulating last values in the rows together, we are subtracting the previous accumulated numbr from the first element of the current row.

### Day 10
Part 1 is a case of a breadth-first search following two directions from the start and end when they meet. No tricks and in this case skipping a check for the right connectivity and allowing e.g. "77" to be a valid pipeline connection still gave the right result.

Part 2 requires more careful handling of the data and connections as the inside/outside algorithm is sensitive to the right data. Inside/outside can be determined per-line by counting how many pipes going vertically. Each vertical pipe flips between inside and outside. Vertical pipes are "|", "FJ", "L7", "F-J", "L--7", "F--J", ... Connection like "F7", "LJ", "F-7", "L-J", "F--7", ... are not flipping inside/outside.

To have the correct bahavior, we need to know what pipe type "S" replaces and there are 6 cases, which depend from where the surrounding pipes connect.

Only pipes, which were detected as edges, using a modified algoritm from part 1, should be checked, the rest is the same as ".".

### Day 11
TBD
