import Algorithms

struct Day01: AdventDay {
  init(data: String) {
    self.data = data
    setupEntities()
  }
  
  // Save your data in a corresponding text file in the `Data` directory.
  var data: String

  var left: [Int] = []
  var right: [Int] = []
  
  mutating func setupEntities() {
    data.split(separator: "\n").forEach { line in
      let values = line.split(separator: "   ")
      left.append(Int(values[0])!)
      right.append(Int(values[1])!)
    }
    left.sort()
    right.sort()
  }

  // Replace this with your solution for the first part of the day's challenge.
  func part1() -> Any {
    var sum: Int = 0
    for i in 0..<left.count {
      if (right[i] > left[i]) {
        sum += right[i] - left[i]
      } else {
        sum += left[i] - right[i]
      }
    }
    return sum
  }

  // Replace this with your solution for the second part of the day's challenge.
  func part2() -> Any {
    var hits: Dictionary<Int, Int> = [:]
    right.forEach { rightElement in
      hits[rightElement, default: 0] += 1
    }
    
    var sum: Int = 0
    left.forEach { leftElement in
      sum += leftElement * hits[leftElement, default: 0]
    }
    
    return sum
  }
}
