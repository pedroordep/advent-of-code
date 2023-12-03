import java.io.File
import kotlin.math.log

data class Point(val x: Int, val y: Int)

var a = mutableMapOf<Point, Boolean>()
a[Point(0, 0)] = true
println(a[Point(0, 0)])

fun markParts(x: Int, y: Int, map: MutableMap<String, Boolean>) {
    for (i in -1..1) {
        for (j in -1..1) {
            map["${x + i},${y + j}"] = true
        }
    }
}

fun markPartsFromFile(fileName: String, partsMap: MutableMap<String, Boolean>) {
    File(fileName).useLines {
        it.forEachIndexed { y, line ->
            line.forEachIndexed { x, s ->
                if (!s.isDigit() && s != '.') {
                    markParts(x, y, partsMap)
                }
            }
        }
    }
}

fun getSum(fileName: String, partsMap: MutableMap<String, Boolean>): Int {
    File(fileName).useLines {
        var sum = 0
        it.forEachIndexed { y, line ->
            var x = 0
            while (x < line.length) {
                if (line[x].isDigit()) {
                    // println("char is int ${line.get(x)}")
                    var n = x
                    while (n + 1 < line.length && line[n + 1].isDigit()) {
                        n++
                    }
                    for (xTemp in x..n) {
                        // println("partsMap.get(\"$xTemp,$y\") = ${partsMap.get("$xTemp,$y")}")
                        if (partsMap["$xTemp,$y"] == true) {
                            sum += line.substring(x..n).toInt()
                            // println("adding ${line.substring(x..n).toInt()}, with x=$x n=$n")
                            x = n
                            break
                        }
                    }
                }
                x++
            }
        }
        return sum
    }
}

fun printMatrixWithParts(x: Int, y: Int, partsMap: Map<String, Boolean>) {
    for (i in 0..y) {
        for (j in 0..x) {
            if (partsMap["$j,$i"] == true) {
                print("X")
            } else {
                print("_")
            }
        }
        print("\n")
    }
}

var partNumberCells = mutableMapOf<String, Boolean>()
markPartsFromFile("input_example.txt", partNumberCells)
assert(getSum("input_example.txt", partNumberCells) == 4361)

partNumberCells = mutableMapOf()
markPartsFromFile("input.txt", partNumberCells)
println("part 1: ${getSum("input.txt", partNumberCells)}")

fun getSumRatioFromFile(fileName: String, gearsMap: MutableMap<String, MutableList<Int>>): Int {
    var sum = 0
    val lines = File(fileName).readLines()
    lines.forEachIndexed { y, line ->
        var x = 0
        while (x < line.length) {
            if (line[x].isDigit()) {
                var n = x
                while (n + 1 < line.length && line[n + 1].isDigit()) {
                    n++
                }
                val number = line.substring(x..n).toInt()
                for (yTemp in y - 1..y + 1) {
                    if (yTemp < 0 || yTemp > lines.count() - 1) {
                        continue
                    }
                    for (xTemp in x - 1..n + 1) {
                        if (xTemp < 0 || xTemp > lines[yTemp].count() - 1) {
                            continue
                        }
                        if (yTemp == y && xTemp >= x && xTemp <= n) {
                            continue
                        }
                        if (lines[yTemp][xTemp] == '*') {
                            gearsMap.getOrPut("$xTemp,$yTemp") { mutableListOf() }.add(number)
                        }
                    }
                }
                x = n
            }
            x++
        }
    }
    gearsMap.forEach {
        if (it.value.count() == 2) {
            sum += it.value[0] * it.value[1]
        }
    }
    return sum
}

var gearsMap = mutableMapOf<String, MutableList<Int>>()
assert(getSumRatioFromFile("input_example.txt", gearsMap) == 467835)

gearsMap = mutableMapOf<String, MutableList<Int>>()
println("part 2: ${getSumRatioFromFile("input.txt", gearsMap)}")
