# aoc2023
https://adventofcode.com/2023

Self-imposed constraints:
 - No regular expressions unless the problem is targetted to need them (unlikely).
 - Just standard built-in packages (including new "generic" slices etc.)
 - No common helper top level functions, each problem is self-contained inside a single function (can use inner function closures). This prevents making solutions seemingly simpler while possibly doing extra work to fit the helper functions (having some map/filter/reduce, []string -> []int etc. may be tempting, but can obscure the real complexity of the solution).
 - No going back and refactoring of the previous solutions based on things learned in the newer solutions. The solutions are a snapshot in time and may be possible to use to track any improvements in the code consistency / style / simplicity.

Solution notes for each day:

### Day 1
Part 1 is straightforward when keeping first digit when found and overriding last digit every time a digit is found. This handles a situation when only one digit is present as the first and last will be set at the same time.

Part 2 adds a need to map words to digits. Instead of matching agains some map keys (and traversing the map each time), a simple switch with checking for a prefix can do.

### Day 2
Part 1 is quite straightforward with multiple nested loops corresponding to the nested problem notation. When a configuration inconsistent with input constrains is found "labeled" continue is used to terminate several loop levels.

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

### Day 6
Part 1 is straightforward as there is a very simple formula for the distance from the time `d = i * (t - i)` which is a nice upside-down parabola and symetric on the range `t`.
So we can go only to the half of the interval and if we find a bound where we beat the record, we know where this ends and do not need to iterate over the whole interval.
One new library function used is `strings.Fields` which handles multiple-whitespace separators simplifying the table parsing.

Part 2 is even simpler when using the observation above. Just running the identical solution on the "new" input gave result instantly. If it was not fast enough, I was considering binary search on the range halve as it is monotonic (upside-down half-parabola), which would probably be fast enough. But it seems like it could have been even faster using Newton's method modified for discrete case, which in theory is similar to binary search with biased "midpoint" computation.

### Day 7
Part 1 boils down to write a function to compare two hands. Then it can be used to sort the hand-bet entries to get the rank and sum all bets multiplied by their rank. Comparing hands has two parts, comparing sorted counts of card kinds in the hand and if the same just comparing the cards with card kinds mapped to values (index in an array in this case). Using `strings.Count(s, s[:1])` and `strings.ReplaceAll(s, s[:1], "")` until the string is empty gives the counts, which need to be sorted and compared.

Part 2 is almost identical with the card kinds updated to move `J` to the end and before the strings is processed to get the kind counts, `J` kind is counted and removed using the mentioned functions and then added to the highest count found. There is one edge case, where there are only `J` cards in a hand as so this returns `[5]` for counts right away.

### Day 8

TBD
