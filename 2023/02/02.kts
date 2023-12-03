import java.io.File

val maxCubesPerColor = mapOf("red" to 12, "green" to 13, "blue" to 14)

fun getId(line: String): Int {
    return line.removePrefix("Game ").split(":")[0].toInt()
}

fun isSetValid(set: String): Boolean {
    set.split(",").forEach {
        val match = Regex(" (\\d+) (\\w+)").find(it)!!
        val (num, color) = match.destructured
        if (num.toInt() > maxCubesPerColor[color]!!) {
            return false
        }
    }
    return true
}

fun isGameValid(game: String): Boolean {
    game.split(":")[1].split(';').forEach {
        if (!isSetValid(it)) {
            return false
        }
    }
    return true
}

var result = File("input_example.txt").useLines { it.sumOf { if (isGameValid(it)) getId(it) else 0 } }
println("test first input: $result")

result = File("input.txt").useLines { it.sumOf { if (isGameValid(it)) getId(it) else 0 } }
println("first input: $result")

fun calcGamePower(game: String): Int {
    val maxCubesPerColor = mutableMapOf("red" to 0, "green" to 0, "blue" to 0)
    game.split(";").forEach {
        it.split(",").forEach {
            val match = Regex(" (\\d+) (\\w+)").find(it)!!
            val (num, color) = match.destructured
            if (num.toInt() > maxCubesPerColor[color]!!) {
                maxCubesPerColor[color] = num.toInt()
            }
        }
    }
    return maxCubesPerColor.toList().fold(1) { acc, (_, num) -> acc * num }
}

File("input_example.txt").useLines {
    result = it.sumOf { calcGamePower(it.split(":")[1]) }
}
println("test second input: $result")

File("input.txt").useLines {
    result = it.sumOf { calcGamePower(it.split(":")[1]) }
}
println("second input: $result")