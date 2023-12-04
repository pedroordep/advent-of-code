import java.io.File

fun part1(lines: List<String>): Int {
    var sum = 0
    lines.forEach {
        val winningNumbers = it.split(":")[1].split("|")[0].trim().replace("  ", " ").split(" ").map { it.toInt() }
        val ourNumbers = it.split(":")[1].split("|")[1].trim().replace("  ", " ").split(" ").map { it.toInt() }

        var cardSum = 0
        ourNumbers.forEach {
            if (winningNumbers.contains(it)) {
                if (cardSum == 0) {
                    cardSum = 1
                } else {
                    cardSum *= 2
                }
            }
        }
        sum += cardSum
    }
    return sum
}

var lines = File("input_example.txt").readLines()
assert(part1(lines) == 13)

lines = File("input.txt").readLines()
println("part 1: ${part1(lines)}")

fun part2(lines: List<String>): Int {
    val cards = mutableMapOf(0 to 1)
    lines.forEachIndexed { index, _ -> cards[index] = 1 }
    lines.forEachIndexed { index, line ->
        val winningNumbers = line.split(":")[1].split("|")[0].trim().replace("  ", " ").split(" ").map { it -> it.toInt() }
        val ourNumbers = line.split(":")[1].split("|")[1].trim().replace("  ", " ").split(" ").map { it -> it.toInt() }

        var cardSum = 0
        ourNumbers.forEach {
            if (winningNumbers.contains(it)) {
                cardSum++
            }
        }
        (index..<index + cardSum).forEach {
            val cur = cards[it + 1]!!
            cards[it + 1] = cur + cards.getValue(index)
        }
    }
    return cards.values.sum()
}

lines = File("input.txt").readLines()
println("part 2: ${part2(lines)}")