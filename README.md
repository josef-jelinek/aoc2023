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
Part 1 has a couple of similar solutions with different requirements on memory. Since I was guessing this was likely a focus of part 2, I went for a solution which just finds the empty rows/columns (by marking which are not empty) and just creates a list of galaxy co-ordinates. Then the shortest path can be computed simply by using "Manhattan" metric dx+dy and increase by the number of empty rows and empty columns between two galaxies (it just counts it from the arrays of empty rows and columns).

Part 2 is just replacing increment of 1 for each empty row/column by increase of 999999 (as 1 gets replaced by 1000000, it is an increase of 999999). All the work was done in part 1 this time.

### Day 12
Part 1 leads to a pretty straightforward recursive definition of the problem:
 - Terminating conditions: we ran out of ranges and have no `#` slots left - solution found; we ran out of spring slots - solution not found.
 - Reduction step: skip `.`/`?` solve for the rest; skip a given range + 1 character after (range must overlay `?`/`#` only and not be followed by `#`, the character after only if there are any left) and solve for the rest.
 - Synthesis step: accumulate for `?`/`#` alternatives; pass on for `.`.

Part 2 makes the search space for the recursive algorithm significantly bigger and some input lines may be beyond brute-force in a reasonable time. There may be some improvements using cutting the solution tree branches early, but that usually just offsets feasible inputs by constant size increase.

In this case, this problem is very repetitive and e.g. a simple change to a start range  doubles the time, but does not affect the number of solutions for the remaining ranges. Together with the input being of a reasonable size (both springs strings and range sizes), it allows for remembering result for a particular positions in the springs string and range array as a simple 2D array with return numbers. The recursive function has been updated not only to contain the new "memo" array, but also to use the original inputs with indexes instead using "unprocessed" slices of the original input. This change makes all the solutions to be found instantly.

### Day 13
Part 1 is a straightforward problem, where even a direct/naive solution works well. When the grid is recorded as a string array, testing all symmetry lines and compare a pair of lines at a time comes naturally.
Instead of treating row and column versions specially, a simple `transpose` function is used to convert column symmetry into a simpler row symmetry.

Part 2 in this case can be solved by the same naive approach with the only difference that instead of a string comparison a dedicated compare fuction is used, which returns how many characters differ between two strings. It can return 2 as soon as it finds at least 2 differences as there must be at most 1 difference in the entire grid.

### Day 14
Part 1 is straightforward and can be computed without any modification of the input and doing the actual "rolling". Just remember a position where the next stone would be and when found add that position to the sum and move by 1. When a square rock is found move the remembered position after the rock.

Part 2 is a very different problem and actually needs to perform rolling of the stone to the four directions. Fortunately, this produces a repetitive pattern pretty quickly and no more than a few hundreds of operations is needed to detect the pattern automatically. After knowing where the pattern starts and after how many cycles it repeats, the 10^9 cycles needed is reduced to a modulo operation to know how many cycles to do to end on the same place in the loop as we would after 10^9 cycles. This is prone to off-by-1 errors, so care is needed to interpret all the lengths/indexes correctly.

The loop detection is done by computing a simple numeric "hash" (just adding all y * width + x, where the `O` stones are, seems good enough) and remember the hash after every cycle. After each cycle, the closest recent identical hash is looked up until one is found and a potential loop compare is initialized. The loop comparison progresses after each cycle until either a computed hash does not match the expected and the loop compare is canceled, or when the compare reaches the end of the potential loop sucessfully. This makes sure that all hashes in the loop match and that misdetecting the loop is basically impossible.

## Day 15
Part 1 is just a simple sum of results of adding and multiplying a single byte. When byte type is used, no explicit `mod 256` is needed.

Part 2 reinvents a hashtable with open hashing (storing collision as a separate list instead of using only slots of the hashtable). There is no need to use anything other than an array of the lenses for each box and using a linear search plus update/append/delete.

## Day 16
Part 1 can be solved by different path traversing techniques (recursively - depth first, with a queue - breadth first). Solution chosen used a queue (Go slice is used for simplicity although it does not release previously used parts of the queue until finished, but that is not an issue for a one time task at hand). A grid remembering light beams with their directions is used both to get how many tiles are energized and to terminate search if the light with the same direction was already explored. The result is then computed by counting non zero grid cells.

Part 2 is doing the same for the 2 * width + 2 * height individual queue seeds and getting the maximum of the results. Since part 1 was already reasonably efficient, part 2 solution is printed instantly as well.

## Day 17
Part 1 is a non-trivial version of path finding. An attempt to use a simpler depth-first search with a heuristic function (using precomputed grid of smallest cost from each grid cell disregarding 3-in-row-max constraints) works well for smaller examples, but is too slow for the full input. So some sort of Dijkstra algorithm is a good next option. Using a priority queue allows to consider the first solution optimal. It seems like even naive priority queue using plain array and O(N) inserting is sufficiently fast. The only tricky part is then knowing which "nodes" were visited as two nodes are not the same for the same x and y if they are visited from a different directions or different consecutive step count in the same direction. The simplest is probably to just use a struct with x, y, direction, count as a map key to get the smallest found cost for such a node.

Part 2 is the same as part 1 with an additional condition for the end node and slightly different conditions for adding new nodes. Computationally, it is heavier than part 1, but should still finish in 1-2s on a decent laptop. If it is too slow, using e.g. a heap for the priority would be a low hanging fruit. Having more compact node ID would be another (can be 4 bytes instead of 32 bytes on 64bit system).

## Day 18
Part 1 is similar to problem from day 10 in how we detect inside/outside of a polygon. It is solved naively by counting grid cell of the area outline and then counting empty inside grid cells.

Part 2 is a generalization of part 1, where it helps not to think about the grid as an "occupancy" grid but as a area cells and mark vertical cell boundaries as up or down. To make representation compact, a grid is constructed only from unique x and y co-ordinates. To compute volume, not only those delineated rectangular areas are added together, but also the outline needs to be added as it has a thickness, but only half of it is outside areas already counted. The last contribution to the volume is corners, where convex/outside corners contribute by 1/4 and concave by -1/4, resulting in area of +1 total as all other corner contributions cancel each other.

Instead of counting boxes and detecting inside/outside areas of a grid, a polygon area algorithm from geometry can probably be used, but needs to use floating point numbers and rounding. (The algorithm sums up signed areas retrieved from vector cross-product.)

## Day 19
Part 1 is straightforwad FSA with two end states ('A' and 'R') and one start state ('in'). To simplify the logic of running the FSA, the parsed rules contain a predicate boolean function which checks the condition in the rule.

Part 2 is a different problem, where instead of running the FSA, it explores all branches until the final states and collects end intervals which were reduced by the rule conditions. Since each branching splits intervals into non-overlapping segments, we do not need to account for any overlaps when we sum all the posibilities for the resulting intervals. The exploration is done recursively and relies on the fact that the FSA does not contain any cycles.

## Day 20
Part 1 is a form of simulator of a stateful machine. Typically, OOP comes to mind for similar class of problems, but here it is likely not simpler than just storing states of the "modules" separately in an array, which for each module holds either nothing (broadcaster, output), a single flag for flip-flop, or an input flag array for a conjunction module. All three can be represented by `[]byte` which is either nil, one element, or with a length matching the inputs.

The processing is done by dequeueing/enqueueing signals as they are processed/created.

For the small example problems, we can just detect when we get back to the initial state (all values in the state arrays are 0) and then multiply the number of signals detected to get amounts for 1000 button pushes. For the input problem, it will need to go through all the states as it does not get back to the initial state soon enough.

Part 2 requires a change of thinking as there is no way to reach the requested state in a reasonable amount of time just by running the simulation. The solution assumes that the output is fed from a conjunction module with multiple inputs and if we compute when those inputs generate high pulses, we can deduce when the low output pulse is sent.

Another important thing is track signals as they are processed instead of checking states after the processing is complete, since the pulses seem to be reset before the end of each cycle. Many assumptions about periodicity of the state are specific to the given input and obrained by analyzing the processing as it is being run.

After getting a period of high pulses for each end conjunction input, we can compute the least common multiple from the detected cycles. In this case it seems the greatest common divider is 1, so the result is just a multiplication of the end input cycle lengths.

In some aspects this problem resembles day 8, where it was also about detecting cycles and composing them together. In this case the problem is simplified that the cycles start in the beginning, so we do not need to deal with constant offsets of the cycle starts.

## Day 21
Part 1 was a straightforward breadth-first search without keeping track of visited fields. Since the number of steps is not enough to exceed boundaries of the grid, no bound checking / wrapping around is needed.

Part 2 is a very different problem and relies on various assumptions depending on which kind of solution is pursued. After trying to find patterns in the distances for individual fields, it seemed to be better to get patterns for counts for individual steps. The counts are alternating even/odd fields and seem to be too chaotic. However patterns each grid width seem to be predictable.

Very likely tied to how the input is constructed, the low amount of obstacles makes the "area" increase similarly predictably as if there were no obstacles only it is not as easy to derive.

Without any obstacles, the area would form a diamond shape (because of a lack of diagonal moves) and the area increases quadratically with the number of steps (as area of a 2D shape increases quadratically with the axis scale). Fortunately, with this input the formula is the same with some constant and linear components added to the result and all that is needed are three points.

These three points needs to be aligned with the grid multiple for the number of steps in the problem. This is done by sampling at `n%size`, `n + n%size`, and `2n + n%size`, where n is the specified total number of steps and `size` is the size of the grid side. Here we use the fact that the grid is square and the start is right in the middle.

Since this requires searching outside of the input grid, we need to "wrap around" the co-ordinates when testing for the obstacles. In Go, there is no real "modulo" operator, just a "remainder", which means `-1 % size -> -1` instead of `size - 1`, so we simulate modulo by `((x + dx)%size + size)%size` as `x` can be arbitrarily large, positive or negative.

After we get the counts for the selected steps, we need to extrapolate the parabola from the constant offset, first order difference, and second order difference, which correspond to derivatives if we did not deal with discrete input. Each step the value increases by the first order difference and each other step there is second order difference times the step increase. As the second order difference is computed by subtracting two first order diferences, it starts comtributing one step later than the first order difference.

As the first step to computing the field counts uses the same search as part 1 with no optimizations (e.g. taking into account that there are two independent even/odd field sets and that all previously visited even/odd step fields are in the final even/odd set), it takes slightly over 2s to get the solution on a decent laptop.

## Day 22
Part 1 can be done in several steps to find out and count which bricks are not a single supporting brick for other bricks. Most of the work is done by "settling" bricks. First potential candidates with overlapping x and y, which are at lower z are found. Then for each brick which has its supporting bricks already settled its z is adjusted to sit on the highest supporting brick. This is done untill all is settled. Then from the "supporting" brick candidates only those right underneath are kept.

Part 2 reuses most of the code from part 1 except instead of counting "independent" bricks, bricks which cause other bricks to fall are seeds into a kind of breadth-first search to detect which other transitively dependent bricks will fall. The count includes only the dependent bricks, not the "seed" bricks destroyed.

Using straightforward arrays and loops in loops instead of anything more sophisticated seems to be acceptable as the solution is found in ~20ms on a decent laptop.

## Day 23
Part 1 is a straightforward application of a DFS (depth-first-search) using a recursive function. The input seems to be formed in such a way that it forms an DAG (directed-acyclic-graph), which makes the search fast similarly to traversing a tree data structure.

Part 2 relaxes the constraints of intersections to form a DAG and it forms a much less efficiently traversed general graph. It gets too slow to directly traverse a character based grid, but converting to a simplified graph seems to be enough to still get solution in ~1s. The simplified graph just extract intersections (3+ non-blocked neighbors) and go from intersection to intersection adding distance between intersections.

## Day 24
Part 1 is an application of linear algebra / computer graphics for computing intersection of two lines in 2D. It is quite rare to remember the math right away, so some read-up is needed.

Part 2 is a different problem, where we are searching a start condition satisfying some constraints. The most straightforward is to assume that solution for the first few heils will extrapolate for the rest (otherwise there would be no solution).

This leads to a set of linear equations where only rock's position and velocity is unknown and can be solved by solving the set of linear equations. It seems that choosing heils from the beginning and solving for x and y first works well with the input.

## Day 25
Part 1 is a graph problem and there are various ways to solve this task. The problem is typically called st. like "minimum cut", where you limit to the cut having exactly 3 edges. Fortunately the two "halves" are relatively dense enough that simpler intuitive solutions can be used.

If we start with one part containing all nodes and the other none, there are 0 cross-connections between those and we need to move some node to the smaller half. If we choose a node which has the most connections to that second half, it seems that the input is set so that such a node really belongs to the other half. Repeating the same process until there are exactly 3 connections from all nodes of the first half to the second will split the graph to the required parts. Input does not seem to have nodes connected to the bridge edges repeat so as long as we are selecting nodes with 2+ connections to the other half, this will work OK.

Despite how naive the process is, it takes only a fraction of a second to complete.

Part 2 is just... Press the button!
