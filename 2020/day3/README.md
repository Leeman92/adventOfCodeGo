# DAY 03
Today this puzzle - at least the way I aproached it - was quite mathematically.

As I am not sure if my code is - without any comments - understandable I will try to explain my logic behind my math
using this readme

Due to the fact that the landscape is unlimited if you go to the right, my first Idea was to just copy it x times into memory.

For the test input this would have been a way to go probably if you wanted to bruteforce it, however looking at the length of the actual input I decided against it.

## Approach for Part 1
----
So I started to think. How to determine how far I have to go each line? Let's take the slope `3 right 1 down` as an example

On line 1 I would not move.

On line 2 I would need to move 3 to the right

On line 3 I would need to move 6 to the right.

Right away it dawned on me that my steps for each line was ```steps = (currentLine - 1) * steps to the right``` (where steps to the right is 3)

However since I start on index 1 I also had to add 1 to the end to get to the right spot.

Since I am working with a language that starts with 0 based indices I was able to ommit both negative values and my formula was

```golang
stepsForLine := currentLine * stepsToTheRight
```

Easy right? Well not really. Let's say we are on the last line (index 10 in the test input). My index would be 30, however the total length of a line was 11 before it began from the beginning.

Here's where a little modulo math comes in handy. Since It is repeating over and over again I can ignore all steps that wrap around. 

30 - 11 = 19

19 - 11 = 8

so we would land on index 8 (since we are 0 based) on position 9 and behold.. thats exactly where we are in the test input. So we now knew what index we would reach with the amount of steps

```golang
endIndex := stepsForLine % len(line)
```

and all that was left to do is check if we landed on a tree `#` or a clearing `.`

## Part 2
----
Part 2 brought a little twist into my logic. We now were able to skip lines. The last slope we took only took every 2nd line into consideration.

Easy right? Let's just skip every x lines based on the input?

Well that was exactly my first thought on this issue. However I soon realised that it wasn't that easy. For what ever reason the test input kept failing for the last slope and instead of 2 trees I kept hitting only 1 tree. What was wrong?

I printed out each line and what index I would land on and compared it the result I would get if I walked down myself and soon I realised I was off. It looked like my solution always walked twice to the right instead of just once when skipping one line. And thats when it dawned on my. When I walk down the slope now I cannot take the lineIndex as an indicator as to how far I already walked and how often I did walk a slope. Because I skipped x amount of lines each time and then I realised I simply had to divide the lineIndex by the amount I skip in order to get back to the right result.

To explain it better with an example. Let's take the example slope of `1 right, 2 down`

When I walked 2 lines down I was now on lineIndex 2. With my old formula that meant I would need to go `lineIndex * steps` to the right so. 2 steps even tho I should only go 1. So by making sure to divide by the amount of lines I skipped I was able to get the correct amount again because on Line 2 I only went once down so far and not 2 already