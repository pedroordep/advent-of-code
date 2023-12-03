import java.io.File

fun getSum(line: String): Int {
    var first = -1
    var last = -1
    for (i in line.indices) {
        if (first == -1 && line[i].isDigit()) {
            first = line[i].digitToInt()
        }
        if (last == -1 && line[line.length - i - 1].isDigit()) {
            last = line[line.length - i - 1].digitToInt()
        }
        if (first != -1 && last != -1) {
            break
        }
    }
    return first * 10 + last
}

var sum = 0
File("input_example.txt").useLines { lines -> lines.forEach { sum += getSum(it) } }
println("test first input: $sum")

sum = 0
File("input.txt").useLines { lines -> lines.forEach { sum += getSum(it) } }
println("first input: $sum")

val occurrenceHash = mapOf("one" to 1, "two" to 2, "three" to 3, "four" to 4, "five" to 5, "six" to 6, "seven" to 7, "eight" to 8, "nine" to 9, "1" to 1, "2" to 2, "3" to 3, "4" to 4, "5" to 5, "6" to 6, "7" to 7, "8" to 8, "9" to 9)
fun getSumByIndex(line: String): Int {
    var firstIndex = line.length
    var firstValue = -1
    for (entry in occurrenceHash.entries) {
        val index = line.indexOf(entry.key)
        if (index != -1 && index < firstIndex) {
            firstIndex = index
            firstValue = entry.value
        }
    }
    var lastIndex = -1
    var lastValue = -1
    for (entry in occurrenceHash.entries) {
        val index = line.lastIndexOf(entry.key)
        if (index != -1 && index > lastIndex) {
            lastIndex = index
            lastValue = entry.value
        }
    }
    return firstValue * 10 + lastValue
}

sum = 0
File("input_example_2.txt").useLines { lines -> lines.forEach { sum += getSumByIndex(it) } }
println("test second input: $sum")

sum = 0
File("input.txt").useLines { lines -> lines.forEach { sum += getSumByIndex(it) } }
println("second input: $sum")