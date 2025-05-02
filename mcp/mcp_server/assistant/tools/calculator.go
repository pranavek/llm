    package tools

	import (

		"fmt"
		"context"
	
		"github.com/mark3labs/mcp-go/mcp"
		"github.com/mark3labs/mcp-go/server"

	)


	func RegisterCaculatorTool(s *server.MCPServer) {
		// Add a calculator tool
		calculatorTool := mcp.NewTool("calculate",
			mcp.WithDescription("Perform basic arithmetic operations"),
			mcp.WithString("operation",
				mcp.Required(),
				mcp.Description("The operation to perform (add, subtract, multiply, divide)"),
				mcp.Enum("add", "subtract", "multiply", "divide"),
			),
			mcp.WithNumber("x",
				mcp.Required(),
				mcp.Description("First number"),
			),
			mcp.WithNumber("y",
				mcp.Required(),
				mcp.Description("Second number"),
			),
		)

		s.AddTool(calculatorTool, calculate)
}

    // Add the calculator handler
    func calculate(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
        op := request.Params.Arguments["operation"].(string)
        x := request.Params.Arguments["x"].(float64)
        y := request.Params.Arguments["y"].(float64)

        var result float64
        switch op {
        case "add":
            result = x + y
        case "subtract":
            result = x - y
        case "multiply":
            result = x * y
        case "divide":
            if y == 0 {
                return mcp.NewToolResultError("cannot divide by zero"), nil
            }
            result = x / y
        }

        return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
    }