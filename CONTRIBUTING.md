# 🛠️ Contributing Guide: Adding a New Shape

Thank you for your interest in contributing to our project! 🎉 To add a new shape to the CLI tool, please follow these steps:

1. **Add the Shape Function in `shapes/shapes.go`** 📄

  - Navigate to the `shapes` directory and open the `shapes.go` fe.
  - Define a new function for your shape. For example, to add a "circle":

    ```go
     // Circle shape
     func Circle() {
         fmt.Println("  ***  ")
         fmt.Println(" *   * ")
         fmt.Println(" *   * ")
         fmt.Println("  ***  ")
     }


2. **Add the Shape Name as a Constant in `cmd/main.go`**

  - Open the `cmd/main.go` file.
  - Declare a new constant for your shape's name:

    ```go
    const (
	    Triangle Shape = iota
        Circle
    )

3. **Update the `ShapeStrings` and `StringShapes` Maps in `cmd/main.go`** 

   - Locate the `ShapeStrings` and `StringShapes` maps in `cmd/main.go`.
   - Add your new shape to both maps:
     ```go
     var ShapeStrings = map[string]string{
         ShapeCircle:   "circle", // Add this line
     }

     var StringShapes = map[string]string{
         "circle":   ShapeCircle, // Add this line
     }

4. **Add a Switch Case for the New Shape in `cmd/main.go`**
   - In the `cmd/main.go` file, locate the switch statement that handles shape printing.
   - Add a new case for your shape:
     ```go
     switch shape {
     case ShapeTriangle:
         shapes.PrintTriangle()
     case ShapeSquare:
         shapes.PrintSquare()
     case ShapeDiamond:
         shapes.PrintDiamond()
     case ShapeStar:
         shapes.PrintStar()
     case ShapeCircle: // Add this case
         shapes.PrintCircle()
     default:
         fmt.Println("Invalid shape name. Run with --help to see available shapes.")
     }

5. **Test Your New Shape**

   - Run the CLI tool to ensure your new shape is printed coectly:

     ```bash
     go run . --shape=circle

6. **Submit a Pull Request**

   - After testing, commit your changes and submit a pull request for review.

Thank you for contributing! Your efforts help make this project better for everyone. 🌟 