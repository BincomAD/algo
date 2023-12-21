func maxWidthOfVerticalArea(points [][]int) int {
    xPoints := make([]int, len(points))

    for i := range points {
        xPoints[i] = points[i][0]
    }

    sort.Ints(xPoints)

    var maxArea int
    for i := range xPoints {
        if i == len(xPoints)-1 {
            continue
        }

        if dist := xPoints[i+1] - xPoints[i]; dist > maxArea {
            maxArea = dist
        }
    }

    return maxArea
}