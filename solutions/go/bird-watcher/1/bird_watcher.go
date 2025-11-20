package birdwatcher

import "fmt"
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    var total int = 0 
    for i := 0; i < len(birdsPerDay); i++ {
        total = birdsPerDay[i] + total
    }
    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    var total int = 0
    var weekCount int = 0
    weekCount = (week - 1) * 7
    fmt.Println(weekCount)
    for i := weekCount; i < (weekCount + 7); i++ {
        fmt.Printf("count = %d", birdsPerDay[i])
        total = birdsPerDay[i] + total
    }
    return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] = birdsPerDay[i] + 1
        fmt.Printf("count = %d", birdsPerDay[i])
    }
    return birdsPerDay
}
